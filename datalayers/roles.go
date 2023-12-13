package datalayers

import (
	"fmt"
	"golangapi/constants"
	"golangapi/models"

	"gorm.io/gorm"
)

type RoleDatalayer interface {
	GetRoles(gDB *gorm.DB) ([]models.Role, error)
	Create(name constants.RoleName, gDB *gorm.DB) error
	Delete(name constants.RoleName, gDB *gorm.DB) error
}

type GormRoleDatalayer struct{}

func NewGormRoleDatalayer() RoleDatalayer {
	return GormRoleDatalayer{}
}

func (rdl GormRoleDatalayer) GetRoles(gDB *gorm.DB) ([]models.Role, error) {
	var allRoles []models.Role

	res := gDB.Find(&allRoles)

	if res.Error != nil {
		return nil, res.Error
	}

	return allRoles, nil
}

func (rdl GormRoleDatalayer) Create(name constants.RoleName, gDB *gorm.DB) error {
	role := models.Role{
		Name: name,
	}

	res := gDB.Create(&role)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return fmt.Errorf("no rows created")
	}

	return nil
}

func (rdl GormRoleDatalayer) Delete(name constants.RoleName, gDB *gorm.DB) error {
	res := gDB.Delete(&models.Role{}).Where("name = ?", name)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return fmt.Errorf("no rows deleted")
	}

	return nil
}
