package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func RespondSuccess(c *gin.Context, data interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, Response{
		Code:    ClientErrCodeOK,
		Message: ClientErrMsgOK,
		Data:    data,
	})
}

func RespondError(c *gin.Context, code int, message string) {
	errorResponse := Response{
		Code:    code,
		Message: message,
	}

	c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse)
}
