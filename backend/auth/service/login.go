package auth

import (
	"net/http"

	"github.com/aecra/covid/common"
	"github.com/aecra/covid/db"
	"github.com/aecra/covid/object"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

func Login(c *gin.Context) {
	var requestBody LoginRequest

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username := requestBody.Username
	password := requestBody.Password

	user, isAuthenticated := Authenticate(username, password)

	if !isAuthenticated {
		common.RespondUnauthorized(c, "Sorry username or password is incorrect")
		return
	}

	token, _ := GenerateJwtToken(user)
	dataMap := make(map[string]string)
	dataMap["token"] = token
	common.RespondOk(c, common.Response{Message: "Authenticated successfully", Data: dataMap})
}

func Authenticate(username string, password string) (*object.User, bool) {
	var users []object.User
	db.GetConnection().Where(&object.User{Username: username}).Find(&users)
	if len(users) == 0 {
		return nil, false
	}
	existingPw := users[0].Password
	if existingPw != password {
		return nil, false
	} else {
		return &users[0], true
	}
}

func Online(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "You are online"})
}
