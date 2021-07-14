package website

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.com/gaming0skar123/go/pingbot/config"
	"gitlab.com/gaming0skar123/go/pingbot/website/routes"
	"gitlab.com/gaming0skar123/go/pingbot/website/routes/api"
)

var router *gin.Engine

func Server() {
	// Disable GIN Debug
	gin.SetMode(gin.ReleaseMode)

	router = gin.New()

	// Allows all origins
	router.Use(cors.Default())

	router.GET("/", routes.Index)

	api.ApplyRoutes(router)

	err := router.Run(config.Port)
	if err != nil {
		panic("Starting GIN Server: " + err.Error())
	}
}