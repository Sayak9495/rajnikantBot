package main

import (
	"github.com/Strike-official/rajnikantBot/internal/controller"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	cb := router.Group("/rajnikantBot/create")
	{
		cb.POST("/", controller.CreateBot)
	}

	eb := router.Group("/rajnikantBot/edit")
	{
		eb.POST("/", controller.EditBot)
	}

	ab := router.Group("/rajnikantBot/addHandler")
	{
		ab.POST("/", controller.AddHandlerToBot)
	}
}
