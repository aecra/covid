package route

import (
	"github.com/aecra/covid/common"
	"github.com/aecra/covid/object"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		c.AbortWithStatus(500)
		return
	}
	c.JSON(200, user)
}

func UpdateUser(c *gin.Context) {
	_user, ok := c.Get("user")
	if !ok {
		c.AbortWithStatus(500)
		return
	}
	user := _user.(*object.User)

	inUser := object.User{}

	if err := c.ShouldBind(&inUser); err != nil {
		c.AbortWithStatus(403)
		return
	}

	if inUser.Username != user.Username {
		c.AbortWithStatus(401)
		return
	} else if !(inUser.Position == "school" || inUser.Position == "home") {
		c.AbortWithStatus(401)
		return
	} else if !common.VerifyEmail(inUser.Email) {
		c.AbortWithStatus(401)
		return
	}
	user.Position = inUser.Position
	user.Email = inUser.Email
	user.State = inUser.State
	user.Eaisess = inUser.Eaisess
	user.Uukey = inUser.Uukey
	user.Home = inUser.Home
	object.UpdateUser(user)
}
