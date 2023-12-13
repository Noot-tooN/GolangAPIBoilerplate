package middlewares

import (
	"fmt"
	"golangapi/constants"
	gormdb "golangapi/databases/gorm"
	"golangapi/datalayers"
	"golangapi/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type IUserMiddleware interface {
	UserTokenOk() gin.HandlerFunc
	AllowedRolesMW(allowedRoles ...constants.RoleName) []gin.HandlerFunc
}

type UserMiddleware struct {
	UserRoleDataLayer datalayers.UserRoleDatalayer
	TokenHandler      services.ITokenHandler
}

func NewUserMiddleware(
	tokenHandler services.ITokenHandler,
	useRoleDL datalayers.GormUserRoleDatalayer,
) IUserMiddleware {
	return UserMiddleware{
		TokenHandler:      tokenHandler,
		UserRoleDataLayer: useRoleDL,
	}
}

func NewDefaultUserMiddleware() IUserMiddleware {
	return UserMiddleware{
		UserRoleDataLayer: datalayers.NewGormUserRoleDatalayer(),
		TokenHandler:      services.NewDefaultSymmetricalPasetoTokenHandler(),
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

func (um UserMiddleware) RoleMiddlewareFactory(allowedRoles ...constants.RoleName) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userUuidString, exists := ctx.Get("USER_UUID")

		if !exists {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"message": "Tried using the app without user uuid",
				},
			)

			return
		}

		userUuid, err := uuid.FromString(fmt.Sprintf("%v", userUuidString))

		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"message": "Invalid uuid",
				},
			)

			return
		}

		userRoles, err := um.UserRoleDataLayer.GetUserRoles(userUuid, gormdb.GetDefaultGormClient())

		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{
					"message": "Couldn't fetch users roles",
				},
			)

			return
		}

		matchingRole := false

		for _, allowedRole := range allowedRoles {
			for _, userRole := range userRoles {
				if allowedRole == userRole.Name {
					matchingRole = true
					break
				}
			}

			if matchingRole {
				break
			}
		}

		if !matchingRole {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"message": "Unauthorized",
				},
			)

			return
		}

		// Go forward
		ctx.Next()
	}
}

func (um UserMiddleware) AllowedRolesMW(allowedRoles ...constants.RoleName) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		um.UserTokenOk(),
		um.RoleMiddlewareFactory(allowedRoles...),
	}
}
