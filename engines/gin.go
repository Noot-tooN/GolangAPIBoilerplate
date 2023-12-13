package engines

import (
	"golangapi/constants"
	"golangapi/controllers"
	"golangapi/middlewares"
	"golangapi/routes"

	"github.com/gin-gonic/gin"
)

func NewGinRouter() *gin.Engine {
	// Init middlewares
	userMw := middlewares.NewDefaultUserMiddleware()

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
		userGroup.DELETE(routes.User.SoftDelete, userMw.UserTokenOk(), userController.SoftDeleteUser)
		userGroup.DELETE(routes.User.HardDelete, userMw.UserTokenOk(), userController.HardDeleteUser)
		userGroup.GET(routes.User.GetProfile, userMw.UserTokenOk(), userController.GetUserProfile)
	}

	adminController := controllers.NewDefaultAdminController()

	adminGroup := router.Group("/admin", userMw.AllowedRolesMW(constants.ADMIN)...)
	{
		adminGroup.GET(routes.Admin.GetUserInfo, adminController.GetUserDataById)
		adminGroup.GET(routes.Admin.GetAllUsers, adminController.GetAllUsers)
	}

	return router
}
