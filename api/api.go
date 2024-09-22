package api

import (
	"github.com/gin-gonic/gin"
	"merinio/api/routes"
)

func RegisterRoutes(router *gin.Engine) {

	api := router.Group("/api")
	{
		api.GET("/branches", routes.GetListBranches)
		api.POST("/branches", routes.SaveBranch)
	}
}
