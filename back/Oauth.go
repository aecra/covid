package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var (
	Auth_Unauthorized   = "Unauthorized"
	Auth_InvalidToken   = "invalid_token"
	Auth_Identify_User  = "user"
	Auth_Identify_Admin = "admin"
	Server_Error        = "server_error"
	Insuficient_Scope   = "insuficient_scope"
)

type AuthError struct {
	msg string
}

func (e *AuthError) Error() string {
	return e.msg
}

func Oauth(identity string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if claims, err := auth(c, identity); err == nil {
			c.Set("claims", claims)
			c.Next()
		} else {
			switch err.Error() {
			case Auth_Unauthorized:
				c.AbortWithStatus(401)
			case Auth_InvalidToken:
				c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			case Insuficient_Scope:
				c.AbortWithStatusJSON(403, gin.H{"error": err.Error()})
			case Server_Error:
				c.AbortWithStatus(500)
			default:
				c.AbortWithStatus(500)
			}
		}
	}
}

func auth(c *gin.Context, identity string) (jwt.MapClaims, error) {
	authorization := c.Request.Header.Get("Authorization")

	if authorization == "" {
		return nil, &AuthError{msg: Auth_Unauthorized}
	}

	res := strings.Split(authorization, " ")

	if len(res) != 2 || res[0] != "Bearer" {
		return nil, &AuthError{msg: Auth_Unauthorized}
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(PUB_PEM))
	if err != nil {
		return nil, &AuthError{msg: Server_Error}
	}

	token, err := jwt.Parse(res[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, &AuthError{msg: Auth_InvalidToken}
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return nil, &AuthError{msg: Auth_InvalidToken}
	}

	if identity == Auth_Identify_Admin && token.Claims.(jwt.MapClaims)["isAdmin"] == false {
		return nil, &AuthError{msg: Insuficient_Scope}
	}

	return token.Claims.(jwt.MapClaims), nil
}
