package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"e-voting-mater/internal/pkg/client/votingcore"
	"e-voting-mater/internal/pkg/entity"
	"e-voting-mater/internal/pkg/repository"
	"e-voting-mater/internal/pkg/service/auth"
	blockchainSrv "e-voting-mater/internal/pkg/service/blockchain"
	"e-voting-mater/internal/pkg/service/pw"
	"e-voting-mater/pkg/blockchain"
	"e-voting-mater/pkg/logger"
	"e-voting-mater/pkg/signature"

	"github.com/gin-gonic/gin"
)

const (
	PrefixBearer = "Bearer "
)

type Service struct {
	userRepo          repository.IUserRepo
	coreClient        votingcore.IClient
	jwtService        auth.IJWTService
	pwService         pw.IService
	candidateRepo     repository.ICandidateRepo
	votingServerRepo  repository.IVotingServerRepo
	trackingRepo      repository.ITrackingRepo
	blockchainService blockchainSrv.IService
}

var service *Service

func NewService(
	coreClient votingcore.IClient,
	userRepo repository.IUserRepo,
	candidateRepo repository.ICandidateRepo,
	votingServerRepo repository.IVotingServerRepo,
	trackingRepo repository.ITrackingRepo,
) *Service {
	if service == nil {
		service = &Service{
			coreClient:        coreClient,
			jwtService:        auth.Instance(),
			pwService:         pw.Instance(),
			userRepo:          userRepo,
			candidateRepo:     candidateRepo,
			votingServerRepo:  votingServerRepo,
			trackingRepo:      trackingRepo,
			blockchainService: blockchainSrv.Instance(),
		}
	}
	return service
}

func Instance() *Service {
	return service
}

func (s *Service) CreateUser(ctx context.Context, user *entity.User, password string) error {
	userCore, err := s.coreClient.CreateUser(ctx, &votingcore.CreateUserRequest{
		Username: user.Username,
	})

	if err != nil {
		logger.Errorf(ctx, "CreateUser: create user error: %v", err)
		return err
	}

	pwPair, err := s.pwService.HashAndEncrypt(password)
	if err != nil {
		logger.Errorf(ctx, "CreateUser: hash and encrypt error: %v", err)
		return err
	}

	aptosAddress, err := blockchain.GenRandomAptosAddress()
	if err != nil {
		return err
	}

	user.EncryptedHash = pwPair.EncryptedHash
	user.Nonce = pwPair.Nonce
	user.UserIdCore = userCore.UserID
	user.AptosAddress = aptosAddress
	err = s.userRepo.Save(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) LoginUser(ctx context.Context, username string, password string, citizenId string) (*InfoUser, string, error) {
	user, err := s.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		logger.Errorf(ctx, "LoginUser: get user by username error: %v", err)
		return nil, "", err
	}
	if user.CitizenID != citizenId {
		logger.Errorf(ctx, "LoginUser: citizen id not match")
		return nil, "", errors.New("citizen id not match")
	}
	if err := s.pwService.Verify(password, &pw.PasswordPair{
		EncryptedHash: user.EncryptedHash,
		Nonce:         user.Nonce,
	}); err != nil {
		logger.Errorf(ctx, "LoginUser: verify password error: %v", err)
		return nil, "", err
	}
	if !user.Verified {
		logger.Errorf(ctx, "LoginUser: user not verified")
		return nil, "", errors.New("user not verified")
	}
	token, err := s.jwtService.GenerateToken(&auth.User{
		Username:    user.Username,
		CitizenID:   user.CitizenID,
		CitizenName: user.CitizenName,
		Verified:    user.Verified,
		Email:       user.Email,
		PublicKey:   user.PublicKey,
	})
	if err != nil {
		logger.Errorf(ctx, "LoginUser: generate token error: %v", err)
		return nil, "", err
	}

	return &InfoUser{
		Username:      user.Username,
		CitizenID:     user.CitizenID,
		CitizenName:   user.CitizenName,
		Verified:      user.Verified,
		Email:         user.Email,
		CompressedKey: fmt.Sprintf("%s...%s", user.PublicKey[:6], user.PublicKey[len(user.PublicKey)-4:]),
		IsAdmin:       user.IsAdmin,
	}, token, nil
}

func (s *Service) GetAllUsers(ctx context.Context) ([]*InfoUser, error) {
	users, err := s.userRepo.GetAll(ctx)
	if err != nil {
		logger.Errorf(ctx, "GetAllUsers: get all users error: %v", err)
		return nil, err
	}
	var result []*InfoUser
	for _, user := range users {
		result = append(result, &InfoUser{
			Username:    user.Username,
			CitizenID:   user.CitizenID,
			CitizenName: user.CitizenName,
			Verified:    user.Verified,
			Email:       user.Email,
			IsAdmin:     user.IsAdmin,
		})
	}
	return result, nil
}

func (s *Service) GetUserByUsername(ctx context.Context, username string) (*InfoUser, error) {
	user, err := s.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		logger.Errorf(ctx, "GetUserByUsername: get user by username error: %v", err)
		return nil, err
	}
	return &InfoUser{
		Username:    user.Username,
		CitizenID:   user.CitizenID,
		CitizenName: user.CitizenName,
		Verified:    user.Verified,
		Email:       user.Email,
		IsAdmin:     user.IsAdmin,
	}, nil
}

func (s *Service) GetUserByCitizenID(ctx context.Context, citizenID string) (*InfoUser, error) {
	user, err := s.userRepo.GetUserByCitizenID(ctx, citizenID)
	if err != nil {
		logger.Errorf(ctx, "GetUserByCitizenID: get user by citizen id error: %v", err)
		return nil, err
	}
	return &InfoUser{
		Username:    user.Username,
		CitizenID:   user.CitizenID,
		CitizenName: user.CitizenName,
		Verified:    user.Verified,
		Email:       user.Email,
		IsAdmin:     user.IsAdmin,
	}, nil
}

func (s *Service) VerifyUser(ctx context.Context, usernames []string, adminId, signatureHex string) error {
	admin, err := s.userRepo.GetAdminByAdminId(ctx, adminId)
	if err != nil {
		logger.Errorf(ctx, "VerifyUser: get admin by admin id error: %v", err)
		return err
	}

	msg, err := signature.BuildApproveUser(adminId)
	if err != nil {
		logger.Errorf(ctx, "VerifyUser: build approve user error: %v", err)
		return err
	}

	if err := s.verifyToken(ctx, admin.Username); err != nil {
		logger.Errorf(ctx, "CreateUser: verify token error: %v", err)
		return err
	}

	if err := signature.VerifySignature(admin.PublicKey, signatureHex, msg); err != nil {
		logger.Errorf(ctx, "VerifyUser: verify signature error: %v", err)
		return err
	}

	err = s.userRepo.VerifyUsers(ctx, usernames)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Vote(ctx context.Context, username string, serverId string, votingHex string, signatureHex string) error {
	voted, err := s.trackingRepo.IsExist(ctx, username, serverId)
	if err != nil {
		logger.Errorf(ctx, "Vote: check if user voted error: %v", err)
		return err
	}

	if voted {
		logger.Errorf(ctx, "Vote: user already voted")
		return errors.New("user already voted")
	}

	if err := s.verifyToken(ctx, username); err != nil {
		logger.Errorf(ctx, "CreateUser: verify token error: %v", err)
		return err
	}

	user, err := s.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		logger.Errorf(ctx, "Vote: get user by username error: %v", err)
		return err
	}

	server, err := s.votingServerRepo.GetByServerID(ctx, serverId)
	if err != nil {
		logger.Errorf(ctx, "Vote: get voting server by server id error: %v", err)
		return err
	}

	msg, err := signature.BuildVerifyMessage(username, serverId)
	if err != nil {
		logger.Errorf(ctx, "Vote: build voting message error: %v", err)
		return err
	}

	if err := signature.VerifySignature(user.PublicKey, signatureHex, msg); err != nil {
		logger.Errorf(ctx, "Vote: verify signature error: %v", err)
		return err
	}

	index, err := signature.FindCandidateIndex(user.PublicKey, votingHex, server.NumberOfCandidates)
	if err != nil {
		logger.Errorf(ctx, "Vote: find candidate index error: %v", err)
		return err
	}

	if err := s.coreClient.Vote(ctx, &votingcore.VoteRequest{
		UserID:      user.UserIdCore,
		ServerID:    serverId,
		CandidateID: index,
	}); err != nil {
		logger.Errorf(ctx, "Vote: vote error: %v", err)
		return err
	}

	if _, err = s.blockchainService.Vote(user.AptosAddress, uint64(index), server.ContractAddress); err != nil {
		logger.Errorf(ctx, "Vote: vote on blockchain error: %v", err)
		return err
	}

	if err := s.trackingRepo.Save(ctx, &entity.Tracking{
		Username: username,
		ServerId: serverId,
	}); err != nil {
		logger.Errorf(ctx, "Vote: save tracking error: %v", err)
		return err
	}

	return nil
}

func (s *Service) IsVoted(ctx context.Context, username, serverId string) (bool, error) {
	voted, err := s.trackingRepo.IsExist(ctx, username, serverId)
	if err != nil {
		logger.Errorf(ctx, "IsVoted: check if user voted error: %v", err)
		return false, err
	}
	return voted, nil
}

func (s *Service) ActiveUser(ctx context.Context, username string) (bool, error) {
	//isExist, err := s.coreClient.CreateUser()
	return false, nil
}

func (s *Service) verifyToken(ctx context.Context, username string) error {
	req, ok := ctx.Value(gin.ContextRequestKey).(*http.Request)
	if !ok {
		logger.Errorf(ctx, "verifyToken: get request from context error")
		return errors.New("get request from context error")
	}
	authHeader := req.Header.Get("Authorization")
	if len(authHeader) == 0 || authHeader[:len(PrefixBearer)] != PrefixBearer {
		logger.Errorf(ctx, "verifyToken: get authorization header error")
		return errors.New("get authorization header error")
	}
	authHeader = authHeader[len(PrefixBearer):]

	if len(authHeader) == 0 {
		logger.Errorf(ctx, "verifyToken: authorization header is empty")
		return errors.New("authorization header is empty")
	}

	user, err := s.jwtService.GetUserFromClaims(authHeader)
	if err != nil {
		logger.Errorf(ctx, "verifyToken: validate token error: %v", err)
		return err
	}

	if user.Username != username {
		logger.Errorf(ctx, "verifyToken: username not match")
		return errors.New("username not match")
	}

	if !user.Verified {
		return errors.New("user not verified")
	}

	return nil
}
