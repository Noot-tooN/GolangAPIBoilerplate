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

	// GetUserDataById(ctx *gin.Context)
	// GetUserData(ctx *gin.Context)
	// DeleteUser(ctx *gin.Context)
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

	err = uc.UserService.VerifyUser(loginData)

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			constants.UnauthorizedError,
		)

		return
	}

	ctx.JSON(
		http.StatusOK,
		constants.RequestFinished,
	)
}