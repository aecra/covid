package auth

import (
	"github.com/aecra/covid/common"
	"github.com/aecra/covid/object"
	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username             string `json:"username" binding:"required"`
	Email                string `json:"email"  binding:"required"`
	Password             string `json:"password"  binding:"required"`
	ConfirmationPassword string `json:"confirmation_password"  binding:"required"`
}

func Register(c *gin.Context) {

	var registerRequest RegisterRequest

	if err := c.BindJSON(&registerRequest); err != nil {
		common.RespondBadRequest(c, err.Error())
		return
	}

	username := registerRequest.Username
	email := registerRequest.Email
	password := registerRequest.Password
	confirmPassword := registerRequest.ConfirmationPassword

	if !common.VerifyUsername(username) {
		common.RespondBadRequest(c, "username is invalid")
		return
	}

	if password != confirmPassword {
		common.RespondBadRequest(c, "Passwords & Confirm Passwords donot match")
		return
	}

	if !common.VerifyPassword(password) {
		common.RespondBadRequest(c, "password is invalid")
		return
	}

	if !common.VerifyEmail(email) {
		common.RespondBadRequest(c, "Invalid Email")
		return
	}

	if !object.IsUniqueUser(username, email) {
		common.RespondBadRequest(c, "The username or email you are trying to register already exists")
		return
	}

	if err := object.Register(username, password, email); err != nil {
		common.RespondBadRequest(c, err.Error())
	} else {
		common.RespondOk(c, common.Response{Message: "Registered successfully"})
	}
}
