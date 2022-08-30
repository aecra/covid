package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"status" format:"string"`
	Message string      `json:"message" format:"string"`
	Data    interface{} `json:"data" format:"json"`
}

func respondWithSuccess(response *Response) gin.H {
	g := gin.H{
		"status":  200,
		"message": response.Message,
	}
	if response.Data != nil {
		g["data"] = response.Data
	}

	return g
}

func respondWithError(response *Response) gin.H {
	g := gin.H{
		"status":  400,
		"message": response.Message,
	}
	if response.Status != 0 {
		g["status"] = response.Status
	}
	if response.Data != nil {
		g["data"] = response.Data
	}

	return g
}

func RespondBadRequest(c *gin.Context, errorMessage string) {
	c.JSON(http.StatusBadRequest, respondWithError(&Response{Message: errorMessage}))
}

func RespondUnauthorized(c *gin.Context, errorMessage string) {
	c.JSON(http.StatusUnauthorized, respondWithError(&Response{Message: errorMessage, Status: http.StatusUnauthorized}))
}

func RespondOk(c *gin.Context, response Response) {
	c.JSON(http.StatusOK, respondWithSuccess(&response))
}
