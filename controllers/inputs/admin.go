package inputs

import (
	"golangapi/constants"
)

type RoleForUser struct {
	UserUuid string             `uri:"user_uuid" binding:"required,uuid"`
	Role     constants.RoleName `uri:"role_name" binding:"required"`
}
