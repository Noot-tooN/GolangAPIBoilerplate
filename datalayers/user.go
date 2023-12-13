package datalayers

import (
	"fmt"
	"golangapi/models"

	"gorm.io/gorm"
)

type UserDatalayer interface {
	FindUserByEmail(email string, gDB *gorm.DB) (*models.UserInfo, error)
	FindUserByUuid(uuid string, gDB *gorm.DB) (*models.UserInfo, error)
	FindOrCreateUser(userInfo models.UserInfo, gDB *gorm.DB) (*models.UserInfo, bool, error)
	CreateUser(userInfo models.UserInfo, gDB *gorm.DB) (*models.UserInfo, error)
	SoftDeleteUser(uuid string, gDB *gorm.DB) error
	HardDeleteUser(uuid string, gDB *gorm.DB) error
}

type GormUserDatalayer struct {}

func NewGormUserDatalayer() UserDatalayer {
	return GormUserDatalayer{}
}

func (pus GormUserDatalayer) FindUserByEmail(email string, gDB *gorm.DB) (*models.UserInfo, error) {
	userInfo := models.UserInfo{}

	res := gDB.Where("email = ?", email).First(&userInfo)

	// If not found res.Error = record not found
	return &userInfo, res.Error
}

func (pus GormUserDatalayer) FindUserByUuid(uuid string, gDB *gorm.DB) (*models.UserInfo, error) {
	userInfo := models.UserInfo{}

	res := gDB.Where("uuid = ?", uuid).First(&userInfo)

	// If not found res.Error = record not found
	return &userInfo, res.Error
}

func (pus GormUserDatalayer) FindOrCreateUser(userInfo models.UserInfo, gDB *gorm.DB) (*models.UserInfo, bool, error) {
	res := gDB.
		Where("email = ?", userInfo.Email).
		FirstOrCreate(&userInfo)

	created := false

	if res.RowsAffected != 0 {
		created = true
	}

	return &userInfo, created, res.Error
}

func (pus GormUserDatalayer) CreateUser(userInfo models.UserInfo, gDB *gorm.DB) (*models.UserInfo, error) {
	res := gDB.Create(&userInfo)

	return &userInfo, res.Error
}

func (pus GormUserDatalayer) SoftDeleteUser(uuid string, gDB *gorm.DB) error {
	res := gDB.Where("uuid = ?", uuid).Delete(&models.UserInfo{})

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return fmt.Errorf("no rows updated")
	}

	return nil
}

func (pus GormUserDatalayer) HardDeleteUser(uuid string, gDB *gorm.DB) error {
	res := gDB.Unscoped().Where("uuid = ?", uuid).Delete(&models.UserInfo{})

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return fmt.Errorf("no rows updated")
	}

	return nil
}