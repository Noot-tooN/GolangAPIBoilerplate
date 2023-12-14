package controllers

import (
	"fmt"
	"golangapi/controllers/inputs"
	"golangapi/services"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type IAdminController interface {
	GetUserDataById(ctx *gin.Context)
	GetAllUsers(ctx *gin.Context)

	AddRoleForUser(ctx *gin.Context)
	RemoveRoleForUser(ctx *gin.Context)
}

type AdminController struct {
	UserService     services.IUserService
	UserRoleService services.IUserRoleService
}

func NewDefaultAdminController() IAdminController {
	return &AdminController{
		UserService:     services.NewDefaultUserService(),
		UserRoleService: services.NewDefaultUserRoleService(),
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

func (ac AdminController) AddRoleForUser(ctx *gin.Context) {
	var userRoleData inputs.RoleForUser

	err := ctx.ShouldBindUri(&userRoleData)

	if err != nil {
		fmt.Println(err)

		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Please provide valid url params",
			},
		)
		return
	}

	// This FromStringOrNil should always pass because validator should catch non uuids
	err = ac.UserRoleService.AddRoleForUser(
		uuid.FromStringOrNil(userRoleData.UserUuid),
		userRoleData.Role,
	)

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Couldn't add the role for the user",
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Finished",
	})
}

func (ac AdminController) RemoveRoleForUser(ctx *gin.Context) {
	var userRoleData inputs.RoleForUser

	err := ctx.ShouldBindUri(&userRoleData)

	if err != nil {
		fmt.Println(err)

		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Please provide valid url params",
			},
		)
		return
	}

	// This FromStringOrNil should always pass because validator should catch non uuids
	err = ac.UserRoleService.RemoveRoleForUser(
		uuid.FromStringOrNil(userRoleData.UserUuid),
		userRoleData.Role,
	)

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Couldn't remove the role for the user",
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Finished",
	})
}
