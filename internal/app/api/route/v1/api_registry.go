package v1

import (
	"e-voting-mater/pkg/route"

	"github.com/gin-gonic/gin"
)

type APIv1 struct{}

func New() *APIv1 {
	return &APIv1{}
}

func (a APIv1) RegisterAPIs(rg *gin.RouterGroup) {
	route.RegisterAPI(rg, NewVoting(), "/voting")
	route.RegisterAPI(rg, NewUserApi(), "/user")
}

func (a APIv1) RegisterAdminAPIs(rg *gin.RouterGroup) {
	route.RegisterAdminAPI(rg, NewVoting(), "/voting")
	route.RegisterAdminAPI(rg, NewUserApi(), "/user")
}
