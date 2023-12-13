package engines

import (
	"golangapi/controllers"
	"golangapi/middlewares"
	"golangapi/routes"

	"github.com/gin-gonic/gin"
)

func NewGinRouter() *gin.Engine {
	// Init middlewares
	checkTokenMW := middlewares.NewDefaultUserMiddleware()

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

	// User routes
	userController := controllers.NewDefaultUserController()

	userGroup := router.Group("/user")
	{
		userGroup.POST(routes.User.Register, userController.RegisterUser)
		userGroup.POST(routes.User.Login, userController.LoginUser)
		userGroup.DELETE(routes.User.SoftDelete, checkTokenMW.UserTokenOk(), userController.SoftDeleteUser)
		userGroup.DELETE(routes.User.HardDelete, checkTokenMW.UserTokenOk(), userController.HardDeleteUser)
		userGroup.GET(routes.User.GetProfile, checkTokenMW.UserTokenOk(), userController.GetUserProfile)
	}

	return router
}