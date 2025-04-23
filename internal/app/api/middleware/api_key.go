package middleware

import (
	"e-voting-mater/configs"
	"e-voting-mater/pkg/response"

	"github.com/gin-gonic/gin"
)

func RequireAPIKey(c *gin.Context) {
	apiKey := c.GetHeader("X-API-Key")

	if apiKey != configs.G.Server.AdminAPIKey {
		message := "Invalid API key"
		response.RespondError(c, 401, message)
		return
	}

	c.Next()
}
