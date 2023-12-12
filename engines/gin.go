package engines

import (
	"golangapi/controllers"
	"golangapi/routes"

	"github.com/gin-gonic/gin"
)

func NewGinRouter() *gin.Engine {
	// Create new instance
	router := gin.New()

	// Add middlewares
	router.Use(gin.Logger())

	// Add routes
	healthCheckGroup := router.Group("/healthcheck")
	{
		healthCheckGroup.GET(routes.Healthcheck.Postgre, controllers.CheckPostgre)
		healthCheckGroup.GET(routes.Healthcheck.Server, controllers.CheckServer)
	}

	return router
}