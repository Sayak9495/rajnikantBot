package controller

import (
	"fmt"

	"github.com/Strike-official/rajnikantBot/internal/core"
	"github.com/Strike-official/rajnikantBot/internal/model"
	"github.com/gin-gonic/gin"
)

func CreateBot(ctx *gin.Context) {
	var request model.Request_Structure
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Error:", err)
	}
	strikeObj := core.CreateBot(request)
	ctx.JSON(200, strikeObj)
}

func ManageBot(ctx *gin.Context) {
	var request model.Request_Structure
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Error:", err)
	}
	strikeObj := core.ManageBot(request)
	ctx.JSON(200, strikeObj)
}
func AddHandlerToBot(ctx *gin.Context) {
	var request model.Request_Structure
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Error:", err)
	}
	strikeObj := core.AddHandlerToBot(request)
	ctx.JSON(200, strikeObj)
}
