package models

import (
	basemodels "golangapi/models/base"

	uuid "github.com/satori/go.uuid"
)

type UserRole struct {
	basemodels.BaseUuidModelSoftDelete
	RoleUuid uuid.UUID `gorm:"not null;default:null;uniqueIndex:idx_user_role"`
	Role     Role      `gorm:"foreignKey:RoleUuid;references:Uuid"`
	UserUuid uuid.UUID `gorm:"not null;default:null;uniqueIndex:idx_user_role"`
	User     UserInfo  `gorm:"foreignKey:UserUuid;references:Uuid"`
}
