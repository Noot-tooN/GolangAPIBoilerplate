package datalayers

import (
	"fmt"
	"golangapi/models"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserRoleDatalayer interface {
	Create(userUuid uuid.UUID, roleUuid uuid.UUID, gDB *gorm.DB) error
	Delete(userUuid uuid.UUID, roleUuid uuid.UUID, gDB *gorm.DB) error
}

type GormUserRoleDatalayer struct{}

func NewGormUserRoleDatalayer() UserRoleDatalayer {
	return GormUserRoleDatalayer{}
}

func (rdl GormUserRoleDatalayer) Create(userUuid uuid.UUID, roleUuid uuid.UUID, gDB *gorm.DB) error {
	userRole := models.UserRole{
		UserUuid: userUuid,
		RoleUuid: roleUuid,
	}

	res := gDB.Create(&userRole)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return fmt.Errorf("no rows created")
	}

	return nil
}

func (rdl GormUserRoleDatalayer) Delete(userUuid uuid.UUID, roleUuid uuid.UUID, gDB *gorm.DB) error {
	res := gDB.Delete(&models.Role{}).Where("user_uuid = ?", userUuid).Where("role_uuid = ?", roleUuid)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return fmt.Errorf("no rows deleted")
	}

	return nil
}
