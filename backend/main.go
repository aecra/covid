package main

import (
	"fmt"

	"github.com/aecra/covid/auth"
	auth_middleware "github.com/aecra/covid/auth/middleware"
	"github.com/aecra/covid/corn"
	"github.com/aecra/covid/migration"
	"github.com/aecra/covid/route"
	"github.com/gin-gonic/gin"
)

func main() {
	migration.AutoMigrate()
	go corn.Start()
	r := gin.Default()

	r.StaticFile("/", "./web/dist/index.html")
	r.StaticFile("/index.html", "./web/dist/index.html")
	r.NoRoute(func(c *gin.Context) {
		c.File("./web/dist/index.html")
	})
	r.StaticFile("/logo.png", "./web/dist/logo.png")
	r.Static("/assets", "./web/dist/assets")

	api := r.Group("/api")
	{
		userRoute := api.Group("user").Use(auth_middleware.AuthJwtHeaderToken)
		{
			userRoute.GET("/profile", route.GetUser)
			userRoute.PUT("/profile", route.UpdateUser)
			userRoute.POST("/reportTest", route.ReportTest)
			userRoute.GET("/records", route.GetRecords)
		}

		authRoute := api.Group("/auth")
		auth.RegisterAuthRoutes(authRoute)
	}

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Start server failed.")
	}
}
