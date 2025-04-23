package votingserver

import (
	"context"
	"encoding/json"

	"e-voting-mater/internal/pkg/client/votingcore"
	"e-voting-mater/internal/pkg/entity"
	"e-voting-mater/internal/pkg/repository"
	"e-voting-mater/internal/pkg/service/auth"
	"e-voting-mater/internal/pkg/service/blockchain"
	"e-voting-mater/pkg/logger"
)

type Service struct {
	votingServerRepo  repository.IVotingServerRepo
	candidateRepo     repository.ICandidateRepo
	userRepo          repository.IUserRepo
	coreClient        votingcore.IClient
	jwtService        auth.IJWTService
	blockchainService blockchain.IService
}

var service *Service

func NewService(
	votingServerRepo repository.IVotingServerRepo,
	candidateRepo repository.ICandidateRepo,
	userRepo repository.IUserRepo,
	coreClient votingcore.IClient,
) *Service {
	if service == nil {
		service = &Service{
			votingServerRepo:  votingServerRepo,
			candidateRepo:     candidateRepo,
			userRepo:          userRepo,
			coreClient:        coreClient,
			jwtService:        auth.Instance(),
			blockchainService: blockchain.Instance(),
		}
	}
	return service
}

func Instance() *Service {
	return service
}

func (s *Service) CreateVotingServer(ctx context.Context, votingServer *entity.VotingServer, candidates []*Candidate, msgHex string) error {
	_, err := s.checkAdminAction(ctx, votingServer.AdminId, votingServer.ServerId, msgHex)
	if err != nil {
		logger.Errorf(ctx, "CreateVotingServer: check admin action error: %v", err)
		return err
	}

	if votingServer.NumberOfCandidates != int64(len(candidates)) {
		logger.Errorf(ctx, "CreateVotingServer: number of candidates not match: %v", err)
		return err
	}

	server, err := s.coreClient.CreateVotingServer(ctx, &votingcore.CreateVotingServerRequest{
		NumberOfCandidates:    votingServer.NumberOfCandidates,
		MaximumNumberOfVoters: votingServer.MaximumNumberOfVoters,
		ServerName:            votingServer.ServerName,
	})
	if err != nil {
		logger.Errorf(ctx, "CreateVotingServer: create voting server error: %v", err)
		return err
	}

	votingServer.ServerId = server.ServerId
	votingServer.OpenedVote = false
	votingServer.Results = ""
	votingServer.Active = false

	if err := s.votingServerRepo.Save(ctx, votingServer); err != nil {
		logger.Errorf(ctx, "CreateVotingServer: save voting server error: %v", err)
		return err
	}

	var candidateNames []string

	for i, c := range candidates {
		candidate := &entity.Candidate{
			CandidateName:  c.CandidateName,
			CitizenID:      c.CitizenID,
			AvatarURL:      c.AvatarURL,
			ServerId:       server.ServerId,
			CandidateIndex: int64(i),
		}
		if err := s.candidateRepo.Save(ctx, candidate); err != nil {
			logger.Errorf(ctx, "CreateVotingServer: save candidates error: %v", err)
			return err
		}
		candidateNames = append(candidateNames, c.CandidateName)
	}

	txHash, err := s.blockchainService.AddCandidates(candidateNames)
	if err != nil {
		logger.Errorf(ctx, "CreateVotingServer: add candidates error: %v", err)
		return err
	}
	logger.Infof(ctx, "CreateVotingServer: txHash: %v", txHash)

	users, err := s.userRepo.GetAll(ctx)
	if err != nil {
		logger.Errorf(ctx, "CreateVotingServer: get all users error: %v", err)
		return err
	}

	for _, user := range users {
		address := user.AptosAddress
		if _, err = s.blockchainService.GiveRightToVote(address); err != nil {
			logger.Errorf(ctx, "CreateVotingServer: give right to vote error: %v", err)
			return err
		}
	}
	logger.Infof(ctx, "CreateVotingServer: successfully give vote to user: %v", txHash)

	if _, err = s.blockchainService.StartVote(); err != nil {
		logger.Errorf(ctx, "CreateVotingServer: start vote error: %v", err)
		return err
	}
	logger.Infof(ctx, "CreateVotingServer: successfully created voting server %s", votingServer.ServerId)

	return nil
}

func (s *Service) GetAllVotingServers(ctx context.Context) ([]*ExtendedVotingServer, error) {
	servers, err := s.votingServerRepo.GetAll(ctx)
	if err != nil {
		logger.Errorf(ctx, "GetAllVotingServers: failed to get all voting servers: %v", err)
		return nil, err
	}

	var extendedServers []*ExtendedVotingServer
	for _, server := range servers {
		candidates, err := s.candidateRepo.GetByServerID(ctx, server.ServerId)
		if err != nil {
			logger.Errorf(ctx, "GetAllVotingServers: failed to get candidates by server ID: %v", err)
			return nil, err
		}
		listCandidates := make([]*Candidate, 0, len(candidates))
		for _, candidate := range candidates {
			listCandidates = append(listCandidates, &Candidate{
				CandidateName: candidate.CandidateName,
				CitizenID:     candidate.CitizenID,
				AvatarURL:     candidate.AvatarURL,
				Index:         candidate.CandidateIndex,
			})
		}
		extendedServers = append(extendedServers, &ExtendedVotingServer{
			VotingServer: &VotingServer{
				NumberOfCandidates:    server.NumberOfCandidates,
				MaximumNumberOfVoters: server.MaximumNumberOfVoters,
				ServerName:            server.ServerName,
				ServerId:              server.ServerId,
				OpenedVote:            server.OpenedVote,
				Results:               server.Results,
				Active:                server.Active,
				ExpTime:               server.ExpTime,
			},
			Candidates: listCandidates,
		})
	}
	return extendedServers, nil

}

func (s *Service) GetVotingServerByID(ctx context.Context, serverId string) (*ExtendedVotingServer, error) {
	server, err := s.votingServerRepo.GetByServerID(ctx, serverId)
	if err != nil {
		logger.Errorf(ctx, "GetVotingServerByID: failed to get voting server by ID: %v", err)
		return nil, err
	}

	candidates, err := s.candidateRepo.GetByServerID(ctx, serverId)
	if err != nil {
		logger.Errorf(ctx, "GetVotingServerByID: failed to get candidates by server ID: %v", err)
		return nil, err
	}
	listCandidates := make([]*Candidate, 0, len(candidates))
	for _, candidate := range candidates {
		listCandidates = append(listCandidates, &Candidate{
			CandidateName: candidate.CandidateName,
			CitizenID:     candidate.CitizenID,
			AvatarURL:     candidate.AvatarURL,
			Index:         candidate.CandidateIndex,
		})
	}
	return &ExtendedVotingServer{
		VotingServer: &VotingServer{
			NumberOfCandidates:    server.NumberOfCandidates,
			MaximumNumberOfVoters: server.MaximumNumberOfVoters,
			ServerName:            server.ServerName,
			ServerId:              server.ServerId,
			OpenedVote:            server.OpenedVote,
			Results:               server.Results,
			Active:                server.Active,
			ExpTime:               server.ExpTime,
		},
		Candidates: listCandidates,
	}, nil
}

func (s *Service) EndVote(ctx context.Context, adminId, serverId string, msgHex string) (*Candidate, map[string]int64, error) {
	_, err := s.checkAdminAction(ctx, adminId, serverId, msgHex)
	if err != nil {
		logger.Errorf(ctx, "EndVote: check admin action error: %v", err)
		return nil, nil, err
	}

	results, err := s.coreClient.EndVote(ctx, &votingcore.EndVoteRequest{
		ServerId: serverId,
	})
	if err != nil {
		logger.Errorf(ctx, "EndVote: open vote error: %v", err)
		return nil, nil, err
	}

	resultBytes, err := json.Marshal(results.Results)
	if err != nil {
		return nil, nil, err
	}

	err = s.votingServerRepo.OpenVote(ctx, serverId, string(resultBytes))
	if err != nil {
		logger.Errorf(ctx, "EndVote: save voting server error: %v", err)
		return nil, nil, err
	}

	mapResult := make(map[string]int64)
	candidates, err := s.candidateRepo.GetByServerID(ctx, serverId)
	if err != nil {
		logger.Errorf(ctx, "EndVote: get candidates error: %v", err)
		return nil, nil, err
	}

	tx, err := s.blockchainService.EndVote()
	if err != nil {
		logger.Errorf(ctx, "EndVote: end vote error: %v", err)
		return nil, nil, err
	}
	logger.Infof(ctx, "EndVote: txHash: %v", tx)

	onChainResults, err := s.blockchainService.GetResults()
	if err != nil {
		logger.Errorf(ctx, "EndVote: get results onchain error: %v", err)
		return nil, nil, err
	}

	for _, candidate := range candidates {
		if onChainResults[candidate.CandidateIndex].VoteCount != results.Results[candidate.CandidateIndex] {
			logger.Errorf(ctx, "EndVote: onchain result not match with core result")
			return nil, nil, err
		}
		mapResult[candidate.CandidateName] = results.Results[candidate.CandidateIndex]
	}

	winningCandidate, err := s.blockchainService.WinningCandidate()
	if err != nil {
		logger.Errorf(ctx, "EndVote: get winning candidate error: %v", err)
		return nil, nil, err
	}

	candidate, err := s.candidateRepo.GetByName(ctx, serverId, winningCandidate)
	if err != nil {
		logger.Errorf(ctx, "EndVote: get winning candidate error: %v", err)
		return nil, nil, err
	}

	logger.Infof(ctx, "EndVote: successfully ended vote for server %s", serverId)
	return &Candidate{
		Index:         candidate.CandidateIndex,
		CandidateName: candidate.CandidateName,
		CitizenID:     candidate.CitizenID,
		AvatarURL:     candidate.AvatarURL,
	}, mapResult, nil
}

func (s *Service) PublishVote(ctx context.Context, serverId string) (map[string]int64, error) {
	server, err := s.votingServerRepo.GetByServerID(ctx, serverId)
	if err != nil {
		logger.Errorf(ctx, "EndVote: get voting server error: %v", err)
		return nil, err
	}

	if !server.OpenedVote {
		logger.Errorf(ctx, "EndVote: voting server already closed")
		return nil, nil
	}

	var results []int64
	if err := json.Unmarshal([]byte(server.Results), &results); err != nil {
		logger.Errorf(ctx, "EndVote: unmarshal results error: %v", err)
		return nil, err
	}

	mapResult := make(map[string]int64)
	candidates, err := s.candidateRepo.GetByServerID(ctx, serverId)
	if err != nil {
		logger.Errorf(ctx, "EndVote: get candidates error: %v", err)
		return nil, err
	}
	for _, candidate := range candidates {
		mapResult[candidate.CandidateName] = results[candidate.CandidateIndex]
	}
	return mapResult, nil
}

func (s *Service) ActiveVoting(ctx context.Context, adminId, serverId string, signatureHex string) error {
	_, err := s.checkAdminAction(ctx, adminId, serverId, signatureHex)
	if err != nil {
		logger.Errorf(ctx, "EndVote: check admin action error: %v", err)
		return err
	}

	err = s.votingServerRepo.ActiveServer(ctx, serverId)
	if err != nil {
		logger.Errorf(ctx, "EndVote: active voting server error: %v", err)
		return err
	}

	return nil
}

func (s *Service) checkAdminAction(ctx context.Context, adminId, serverId string, msgHex string) (*entity.User, error) {
	admin, err := s.userRepo.GetAdminByAdminId(ctx, adminId)
	if err != nil {
		logger.Errorf(ctx, "CreateVotingServer: failed to get admin by adminId: %v", err)
		return nil, err
	}

	//msg, err := signature.BuildAdminVerifyMessage(adminId, serverId)
	//if err != nil {
	//	logger.Errorf(ctx, "CreateVotingServer: failed to build admin verify message: %v", err)
	//	return nil, err
	//}
	//
	//if err := signature.VerifySignature(admin.PublicKey, msgHex, msg); err != nil {
	//	logger.Errorf(ctx, "CreateVotingServer: verify signature error: %v", err)
	//	return nil, err
	//}

	return admin, nil
}
