package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()
	r.Use(Cors())

	db, err := gorm.Open(mysql.Open(DB_DNS), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	r.Use(Oauth(Auth_Identify_User))
	r.Use(InitUser())

	r.POST("/User", func(c *gin.Context) {
		claims, ok := c.Get("claims")
		if !ok {
			c.AbortWithStatus(500)
			return
		}
		db, ok := c.Get("db")
		if !ok {
			c.AbortWithStatus(500)
			return
		}

		name := claims.(jwt.MapClaims)["name"].(string)
		users := []User{}
		db.(*gorm.DB).Where("name = ?", name).Take(&users)
		if len(users) == 1 {
			c.JSON(200, users[0])
		} else {
			c.AbortWithStatus(500)
		}
	})

	r.POST("/UpdateUser", func(c *gin.Context) {
		claims, ok := c.Get("claims")
		if !ok {
			c.AbortWithStatus(500)
			return
		}

		db, ok := c.Get("db")
		if !ok {
			c.AbortWithStatus(500)
			return
		}

		user := User{}

		if err := c.ShouldBind(&user); err != nil {
			c.AbortWithStatus(403)
			return
		}

		if claims, ok := claims.(jwt.MapClaims); ok {
			name := claims["name"].(string)
			if claims["name"].(string) != user.Name {
				c.AbortWithStatus(401)
				return
			} else if !(user.Position == "school" || user.Position == "home") {
				c.AbortWithStatus(401)
				return
			} else if !(user.State == "on" || user.State == "off") {
				c.AbortWithStatus(401)
				return
			} else if VerifyEmailFormat(user.Email) == false {
				c.AbortWithStatus(401)
				return
			}
			db.(*gorm.DB).Exec("UPDATE users SET position = ?, state = ?, email = ?, eaisess = ?, uukey = ?, home = ? WHERE name = ?", user.Position, user.State, user.Email, user.Eaisess, user.Uukey, user.Home, name)

			c.JSON(200, gin.H{
				"message": "success",
			})
		} else {
			c.AbortWithStatus(401)
		}
	})

	r.Run(":8080")
}
