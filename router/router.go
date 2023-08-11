package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lampesm/game-time/controllers"
)

func SetupRoute() *gin.Engine {
	router := gin.New()

	v1 := router.Group("/api/v1")
	{
		v1.POST("game", controllers.CreateGame)
	}

	return router
}
