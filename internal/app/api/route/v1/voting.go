package v1

import (
	"e-voting-mater/internal/app/api/dto"
	"e-voting-mater/internal/app/api/middleware"
	"e-voting-mater/internal/pkg/entity"
	"e-voting-mater/internal/pkg/service/votingserver"
	"e-voting-mater/pkg/response"

	"github.com/gin-gonic/gin"
)

type Voting struct {
	votingService votingserver.IService
}

func NewVoting() *Voting {
	return &Voting{
		votingService: votingserver.Instance(),
	}
}

func (a *Voting) SetupRoute(rg *gin.RouterGroup) {
	rg.GET("/info", a.getById)
	rg.GET("/", a.getAll)
	rg.GET("/publish_vote", a.publishVote)
}

func (a *Voting) SetupAdminRoute(rg *gin.RouterGroup) {
	rg.Use(middleware.RequireAPIKey)
	rg.POST("/open", a.openVote)
	rg.PUT("/create", a.createServer)
	rg.POST("/active", a.activeServer)
}

func (a *Voting) createServer(c *gin.Context) {
	ctx := c

	var params dto.CreateVotingServerRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.RespondError(c, 500, err.Error())
		return
	}

	err := a.votingService.CreateVotingServer(ctx, &entity.VotingServer{
		AdminId:               params.AdminId,
		NumberOfCandidates:    params.NumberOfCandidates,
		MaximumNumberOfVoters: params.MaximumNumberOfVoters,
		ServerName:            params.ServerName,
	}, params.Candidates, params.SignatureHex)
	if err != nil {
		response.RespondError(c, 500, err.Error())
		return
	}

	response.RespondSuccess(c, nil)
}

func (a *Voting) getById(c *gin.Context) {
	ctx := c

	var params dto.GetVotingServerByIdRequest
	if err := ctx.ShouldBindQuery(&params); err != nil {
		response.RespondError(c, 500, err.Error())
		return
	}

	votingServer, err := a.votingService.GetVotingServerByID(ctx, params.ServerId)
	if err != nil {
		response.RespondError(c, 500, err.Error())
		return
	}

	response.RespondSuccess(c, votingServer)
}

func (a *Voting) getAll(c *gin.Context) {
	ctx := c

	votingServers, err := a.votingService.GetAllVotingServers(ctx)
	if err != nil {
		response.RespondError(c, 500, err.Error())
		return
	}

	response.RespondSuccess(c, votingServers)
}

func (a *Voting) openVote(c *gin.Context) {
	ctx := c

	var params dto.OpenVoteRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.RespondError(c, 500, err.Error())
		return
	}

	voteResult, err := a.votingService.EndVote(ctx, params.AdminId, params.ServerId, params.SignatureHex)
	if err != nil {
		response.RespondError(c, 500, err.Error())
		return
	}

	response.RespondSuccess(c, voteResult)
}

func (a *Voting) publishVote(c *gin.Context) {
	ctx := c

	var params dto.PublishVoteRequest
	if err := ctx.ShouldBindQuery(&params); err != nil {
		response.RespondError(c, 500, err.Error())
		return
	}

	voteResult, err := a.votingService.PublishVote(ctx, params.ServerId)
	if err != nil {
		response.RespondError(c, 500, err.Error())
		return
	}

	response.RespondSuccess(c, voteResult)
}

func (a *Voting) activeServer(c *gin.Context) {
	ctx := c

	var params dto.ActiveVotingServerRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.RespondError(c, 500, err.Error())
		return
	}

	err := a.votingService.ActiveVoting(ctx, params.AdminId, params.ServerId, params.SignatureHex)
	if err != nil {
		response.RespondError(c, 500, err.Error())
		return
	}

	response.RespondSuccess(c, nil)
}
