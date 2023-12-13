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
	GetUserRoles(userUuid uuid.UUID, gDB *gorm.DB) ([]models.Role, error)
}

type GormUserRoleDatalayer struct{}

func NewGormUserRoleDatalayer() UserRoleDatalayer {
	return GormUserRoleDatalayer{}
}

func (rdl GormUserRoleDatalayer) GetUserRoles(userUuid uuid.UUID, gDB *gorm.DB) ([]models.Role, error) {
	var roles []models.Role

	res := gDB.
		Table("roles").
		Select("roles.name, roles.uuid").
		Joins("inner join user_roles ur on ur.role_uuid = roles.uuid").
		Where("ur.user_uuid = ?", userUuid).
		Find(&roles)

	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("no rows found")
	}

	return roles, res.Error
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
