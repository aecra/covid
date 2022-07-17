package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/aecra/covid/object"
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

func OAuth(identity string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, err := auth(c, identity); err == nil && user != nil {
			c.Set("user", user)
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

func auth(c *gin.Context, identity string) (*object.User, error) {
	authorization := c.Request.Header.Get("Authorization")

	if authorization == "" {
		return nil, errors.New(Auth_Unauthorized)
	}

	res := strings.Split(authorization, " ")

	if len(res) != 2 || res[0] != "Bearer" {
		return nil, errors.New(Auth_Unauthorized)
	}

	// read public key from ./cert/key.pub
	file, err := os.Open("./cert/key.pub")
	if err != nil {
		return nil, errors.New(Server_Error)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(content))
	if err != nil {
		return nil, errors.New(Server_Error)
	}
	token, err := jwt.Parse(res[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, errors.New(Auth_InvalidToken)
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return nil, errors.New(Auth_InvalidToken)
	}

	if identity == Auth_Identify_Admin && token.Claims.(jwt.MapClaims)["isAdmin"] == false {
		return nil, errors.New(Insuficient_Scope)
	}

	user := object.GetUserByName(token.Claims.(jwt.MapClaims)["name"].(string))
	if user.Name == "" {
		user.Name = token.Claims.(jwt.MapClaims)["name"].(string)
		user.Email = token.Claims.(jwt.MapClaims)["email"].(string)
		user.State = false
		object.AddUser(&user)
	}

	return &user, nil
}
