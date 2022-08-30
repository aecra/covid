package auth_middleware

import (
	"net/http"

	auth "github.com/aecra/covid/auth/service"
	"github.com/gin-gonic/gin"
)

func AuthJwtHeaderToken(c *gin.Context) {

	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if authHeader[:len("Bearer ")] != "Bearer " {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	jwtToken := authHeader[len("Bearer "):]

	if user, err := auth.ValidateJwtToken(jwtToken); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	} else {
		c.Set("user", user)
		c.Next()
	}
}
