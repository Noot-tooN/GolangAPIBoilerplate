package models

import (
	"golangapi/constants"
	basemodels "golangapi/models/base"
)

type Role struct {
	basemodels.BaseUuidModelSoftDelete
	Name  constants.RoleName `gorm:"unique;index;not null;default:null"`
	Users []UserInfo         `gorm:"many2many:user_roles;"`
}
