package v1

import (
	"e-voting-mater/internal/app/api/dto"
	"e-voting-mater/internal/app/api/middleware"
	"e-voting-mater/internal/pkg/entity"
	"e-voting-mater/internal/pkg/service/user"
	"e-voting-mater/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
	userService user.IService
}

func NewUserApi() *UserApi {
	return &UserApi{
		userService: user.Instance(),
	}
}

func (a *UserApi) SetupRoute(rg *gin.RouterGroup) {
	rg.POST("/login", a.login)
	rg.GET("/", a.getByUsername)
	rg.GET("/citizen_id", a.getByCitizenID)
	rg.POST("/voting", a.voting)
}

func (a *UserApi) SetupAdminRoute(rg *gin.RouterGroup) {
	rg.Use(middleware.RequireAPIKey)
	rg.PUT("/create", a.createUser)
	rg.POST("/verify", a.verifyUser)
}

func (a *UserApi) createUser(c *gin.Context) {
	ctx := c

	var params dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.RespondError(c, 500, err.Error())
		return
	}

	if err := a.userService.CreateUser(ctx, &entity.User{
		Username:    params.Username,
		CitizenID:   params.CitizenID,
		CitizenName: params.CitizenName,
		Email:       params.Email,
		PublicKey:   params.PublicKey,
		Verified:    true, // TODO: set to false
		IsAdmin:     params.IsAdmin,
	}, params.Password); err != nil {
		response.RespondError(c, 500, err.Error())
		return
	}

	response.RespondSuccess(c, nil)
}

func (a *UserApi) login(c *gin.Context) {
	ctx := c

	var params dto.LoginRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.RespondError(c, 500, err.Error())
		return
	}

	userInfo, token, err := a.userService.LoginUser(ctx, params.Username, params.Password, params.PersonalCode)
	if err != nil {
		response.RespondError(c, 500, err.Error())
		return
	}

	response.RespondSuccess(c, &dto.LoginResponse{
		Token: token,
		User:  userInfo,
	})
}

func (a *UserApi) getByUsername(c *gin.Context) {
	ctx := c

	var params dto.GetUserByUsernameRequest
	if err := ctx.ShouldBindQuery(&params); err != nil {
		response.RespondError(c, 500, err.Error())
		return
	}

	userInfo, err := a.userService.GetUserByUsername(ctx, params.Username)
	if err != nil {
		response.RespondError(c, 500, err.Error())
		return
	}

	response.RespondSuccess(c, userInfo)
}

func (a *UserApi) getByCitizenID(c *gin.Context) {
	ctx := c

	var params dto.GetUserByCitizenIdRequest
	if err := ctx.ShouldBindQuery(&params); err != nil {
		response.RespondError(c, 500, err.Error())
		return
	}

	userInfo, err := a.userService.GetUserByCitizenID(ctx, params.CitizenId)
	if err != nil {
		response.RespondError(c, 500, err.Error())
		return
	}

	response.RespondSuccess(c, userInfo)
}

func (a *UserApi) voting(c *gin.Context) {
	ctx := c

	var params dto.VotingRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.RespondError(c, 500, err.Error())
		return
	}

	err := a.userService.Vote(ctx, params.Username, params.ServerId, params.VotingHex, params.SignatureHex)
	if err != nil {
		response.RespondError(c, 500, err.Error())
		return
	}

	response.RespondSuccess(c, nil)
}

func (a *UserApi) verifyUser(c *gin.Context) {
	ctx := c

	var params dto.VerifyUsers
	if err := ctx.ShouldBindQuery(&params); err != nil {
		response.RespondError(c, 500, err.Error())
		return
	}

	err := a.userService.VerifyUser(ctx, params.Usernames, params.AdminId, params.SignatureHex)
	if err != nil {
		response.RespondError(c, 500, err.Error())
		return
	}

	response.RespondSuccess(c, nil)
}
