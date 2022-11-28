package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aecra/covid/auth"
	auth_middleware "github.com/aecra/covid/auth/middleware"
	"github.com/aecra/covid/corn"
	"github.com/aecra/covid/migration"
	"github.com/aecra/covid/route"
	"github.com/gin-gonic/gin"
)

var (
	SERVER_PORT = os.Getenv("SERVER_PORT")
)

func init() {
	if SERVER_PORT == "" {
		SERVER_PORT = "8080"
	}
	// Verify that the port is legal, 0-65535
	port, err := strconv.Atoi(SERVER_PORT)
	if err != nil || port < 0 || port > 65535 {
		panic("SERVER_PORT is not set or not supported")
	}
}

func main() {
	migration.AutoMigrate()
	go corn.Start()
	gin.SetMode(gin.ReleaseMode)
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

	err := r.Run(fmt.Sprintf(":%s", SERVER_PORT))
	if err != nil {
		fmt.Println("Start server failed.")
	}
}
