package main

import (
	"fmt"

	"github.com/aecra/covid/object"
	"github.com/aecra/covid/route"
	"github.com/gin-gonic/gin"
)

func main() {
	object.Init()
	corn()
	r := gin.Default()

	r.Static("/", "./web/dist")

	api := r.Group("/api")
	{
		api.Use(OAuth(Auth_Identify_User))
		api.POST("/getUser", route.GetUser)
		api.POST("/updateUser", route.UpdateUser)
	}

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Start server failed.")
	}
}
