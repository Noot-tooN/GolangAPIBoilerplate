package middlewares

import (
	"golangapi/datalayers"
	"golangapi/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type IUserMiddleware interface {
	UserTokenOk() gin.HandlerFunc
}

type UserMiddleware struct {
	UserDatalayer datalayers.UserDatalayer
	TokenHandler services.ITokenHandler
}

func NewUserMiddleware(
	userDL datalayers.UserDatalayer,
	tokenHandler services.ITokenHandler,
) IUserMiddleware {
	return UserMiddleware{
		UserDatalayer: userDL,
		TokenHandler: tokenHandler,
	}
}

func NewDefaultUserMiddleware() IUserMiddleware {
	return UserMiddleware{
		UserDatalayer: datalayers.NewGormUserDatalayer(),
		TokenHandler: services.NewDefaultSymmetricalPasetoTokenHandler(),
	}
}

func (um UserMiddleware) UserTokenOk() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header["Authorization"]

		if len(authHeader) == 0 {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"message": "Expecting an Authorization header",
				},
			)
			return
		}

		bearerToken := authHeader[0]

		splitTokens := strings.Split(bearerToken, " ")

		if splitTokens[0] != "Bearer" {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"message": "Expecting a Bearer token",
				},
			)
			return
		}

		userMap, err := um.TokenHandler.ReadToken(splitTokens[1])

		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"message": "Invalid token",
				},
			)
			return
		}

		validUserUuid, err := uuid.FromString(userMap["uuid"])

		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"message": "Invalid user uuid format",
				},
			)
			return
		}

		ctx.Set("USER_UUID", validUserUuid)
	}
}