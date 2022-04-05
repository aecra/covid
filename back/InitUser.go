package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type User struct {
	Name     string `db:"name" json:"name" binding:"required"`
	Position string `db:"position" json:"position"`
	State    string `db:"state" json:"state"`
	Email    string `db:"email" json:"email"`
	Eaisess  string `db:"eaisess" json:"eaisess"`
	Uukey    string `db:"uukey" json:"uukey"`
	Home     string `db:"home" json:"home"`
}

func InitUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := c.Get("claims")
		if ok {
			if claims, ok := claims.(jwt.MapClaims); ok {
				name := claims["name"].(string)
				email := claims["email"].(string)
				db := c.MustGet("db").(*gorm.DB)
				type ExistUser struct {
					Name string `db:"name"`
				}
				existUsers := []ExistUser{}
				db.Raw("SELECT name FROM users WHERE name=?", name).Take(&existUsers)
				if len(existUsers) == 0 {
					db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", name, email)
				}
				c.Next()
			} else {
				c.AbortWithStatus(401)
			}
		} else {
			c.AbortWithStatus(401)
		}
	}
}
