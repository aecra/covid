package route

import (
	"github.com/aecra/covid/object"
	"github.com/aecra/covid/report"
	"github.com/gin-gonic/gin"
)

func ReportTest(c *gin.Context) {
	_user, ok := c.Get("user")
	if !ok {
		c.AbortWithStatus(500)
		return
	}
	user := _user.(*object.User)
	report.ReportSignal(user)
}

func GetRecords(c *gin.Context) {
	_user, ok := c.Get("user")
	if !ok {
		c.AbortWithStatus(500)
		return
	}
	user := _user.(*object.User)
	c.JSON(200, object.GetRecords(user))
}