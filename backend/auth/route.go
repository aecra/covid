package auth

import (
	auth_middleware "github.com/aecra/covid/auth/middleware"
	auth "github.com/aecra/covid/auth/service"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(route *gin.RouterGroup) {
	route.POST("/register", auth.Register)
	route.POST("/login", auth.Login)
	route.Use(auth_middleware.AuthJwtHeaderToken).POST("/online", auth.Online)
}
