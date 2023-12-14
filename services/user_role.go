package services

import (
	"golangapi/constants"
	gormdb "golangapi/databases/gorm"
	"golangapi/datalayers"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type IUserRoleService interface {
	AddRoleForUser(userUuid uuid.UUID, role constants.RoleName) error
	RemoveRoleForUser(userUuid uuid.UUID, role constants.RoleName) error
}

type UserRoleService struct {
	UserRoleDataLayer datalayers.UserRoleDatalayer
	GormDb            *gorm.DB
}

func NewUserRoleService(
	userRoleDL datalayers.UserRoleDatalayer,
	gormDb *gorm.DB,
) IUserRoleService {
	return UserRoleService{
		UserRoleDataLayer: userRoleDL,
		GormDb:            gormDb,
	}
}

func NewDefaultUserRoleService() IUserRoleService {
	return UserRoleService{
		UserRoleDataLayer: datalayers.NewGormUserRoleDatalayer(),
		GormDb:            gormdb.GetDefaultGormClient(),
	}
}

func (urs UserRoleService) AddRoleForUser(userUuid uuid.UUID, role constants.RoleName) error {
	return urs.UserRoleDataLayer.UpsertRoleForUser(role, userUuid, gormdb.GetDefaultGormClient())
}

func (urs UserRoleService) RemoveRoleForUser(userUuid uuid.UUID, role constants.RoleName) error {
	return urs.UserRoleDataLayer.RemoveRoleForUser(role, userUuid, gormdb.GetDefaultGormClient())
}
