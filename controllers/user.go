package controllers

import (
	"fmt"
	"golangapi/constants"
	"golangapi/controllers/inputs"
	"golangapi/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	RegisterUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)

	SoftDeleteUser(ctx *gin.Context)
	HardDeleteUser(ctx *gin.Context)

	GetUserProfile(ctx *gin.Context)
}

type UserController struct {
	UserService services.IUserService
}

func NewDefaultUserController() IUserController {
	return &UserController{
		UserService: services.NewDefaultUserService(),
	}
}

func (uc UserController) RegisterUser(ctx *gin.Context) {
	var registrationData inputs.Registration

	err := ctx.ShouldBindJSON(&registrationData)

	if err != nil {
		fmt.Println(err)

		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			constants.InvalidInputError,
		)

		return
	}

	err = uc.UserService.CreateUser(registrationData)

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			constants.CreateError,
		)

		return
	}

	ctx.JSON(
		http.StatusOK,
		constants.RequestFinished,
	)
}

func (uc UserController) LoginUser(ctx *gin.Context) {
	var loginData inputs.Login

	err := ctx.ShouldBindJSON(&loginData)

	if err != nil {
		fmt.Println(err)

		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			constants.InvalidInputError,
		)

		return
	}

	token, err := uc.UserService.VerifyUser(loginData)

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			constants.UnauthorizedError,
		)

		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"token": token,
		},
	)
}

func (uc UserController) SoftDeleteUser(ctx *gin.Context) {
	userUuid, exists := ctx.Get("USER_UUID")

	if !exists {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Tried using the app without user uuid",
			},
		)

		return
	}

	err := uc.UserService.SoftDeleteUser(fmt.Sprintf("%v", userUuid))

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Couldn' delete the user",
			},
		)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "test",
	})
}

func (uc UserController) HardDeleteUser(ctx *gin.Context) {
	userUuid, exists := ctx.Get("USER_UUID")

	if !exists {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Tried using the app without user uuid",
			},
		)

		return
	}

	err := uc.UserService.HardDeleteUser(fmt.Sprintf("%v", userUuid))

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Couldn' delete the user",
			},
		)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "test",
	})
}

func (uc UserController) GetUserProfile(ctx *gin.Context) {
	userUuid, exists := ctx.Get("USER_UUID")

	if !exists {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Tried using the app without user uuid",
			},
		)

		return
	}

	response, err := uc.UserService.GetProfile(fmt.Sprintf("%v", userUuid))

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Couldn't fetch the profile",
			},
		)

		return
	}

	ctx.JSON(
		http.StatusOK,
		response,
	)
}
