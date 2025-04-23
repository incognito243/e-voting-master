package route

import (
	v1 "e-voting-mater/internal/app/api/route/v1"
	"e-voting-mater/pkg/route"

	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {
	route.RegisterAPI(engine, NewHealthAPI(), "health")
	route.RegisterAPIGroup(engine, v1.New(), "/api/v1")
	route.RegisterAdminAPIGroup(engine, v1.New(), "/admin/api/v1")
}
