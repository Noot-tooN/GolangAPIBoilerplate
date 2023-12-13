package controllers

import (
	"golangapi/services"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type IAdminController interface {
	GetUserDataById(ctx *gin.Context)
	GetAllUsers(ctx *gin.Context)
}

type AdminController struct {
	UserService services.IUserService
}

func NewDefaultAdminController() IAdminController {
	return &AdminController{
		UserService: services.NewDefaultUserService(),
	}
}

func (ac AdminController) GetAllUsers(ctx *gin.Context) {
	userProfiles, err := ac.UserService.GetAllProfiles()

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Couldn't extract the users from the db",
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		userProfiles,
	)
}

func (ac AdminController) GetUserDataById(ctx *gin.Context) {
	userId := ctx.Param("user_id")

	if userId == "" {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Please provide valid url params",
			},
		)
		return
	}

	userUuid, err := uuid.FromString(userId)

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Bad user uuid",
			},
		)
		return
	}

	userProfile, err := ac.UserService.GetProfile(userUuid.String())

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Couldn't find the user with that uuid",
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		userProfile,
	)
}
