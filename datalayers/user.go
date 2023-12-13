package datalayers

import (
	"fmt"
	"golangapi/controllers/outputs"
	"golangapi/models"

	"gorm.io/gorm"
)

type UserDatalayer interface {
	FindUserByEmail(email string, gDB *gorm.DB) (*outputs.UserProfile, error)
	FindUserByUuid(uuid string, gDB *gorm.DB) (*outputs.UserProfile, error)
	FindOrCreateUser(userInfo models.UserInfo, gDB *gorm.DB) (*models.UserInfo, bool, error)

	GetHashedUserPassword(email string, gDB *gorm.DB) (*outputs.HashedUserPass, error)

	FindAllUsers(gDB *gorm.DB) ([]outputs.UserProfile, error)

	CreateUser(userInfo models.UserInfo, gDB *gorm.DB) (*models.UserInfo, error)

	SoftDeleteUser(uuid string, gDB *gorm.DB) error
	HardDeleteUser(uuid string, gDB *gorm.DB) error
}

type GormUserDatalayer struct{}

func NewGormUserDatalayer() UserDatalayer {
	return GormUserDatalayer{}
}

func (pus GormUserDatalayer) FindAllUsers(gDB *gorm.DB) ([]outputs.UserProfile, error) {
	var allUsers []outputs.UserProfile

	res := gDB.Model(models.UserInfo{}).Find(&allUsers)

	return allUsers, res.Error
}

func (pus GormUserDatalayer) FindUserByEmail(email string, gDB *gorm.DB) (*outputs.UserProfile, error) {
	userProfile := outputs.UserProfile{}

	res := gDB.Model(models.UserInfo{}).Where("email = ?", email).Limit(1).Find(&userProfile)

	return &userProfile, res.Error
}

func (pus GormUserDatalayer) FindUserByUuid(uuid string, gDB *gorm.DB) (*outputs.UserProfile, error) {
	userProfile := outputs.UserProfile{}

	res := gDB.Model(models.UserInfo{}).Where("uuid = ?", uuid).Limit(1).Find(&userProfile)

	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("couldn't find user")
	}

	return &userProfile, res.Error
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

func (pus GormUserDatalayer) GetHashedUserPassword(email string, gDB *gorm.DB) (*outputs.HashedUserPass, error) {
	userPassProfile := outputs.HashedUserPass{}

	res := gDB.Model(models.UserInfo{}).Where("email = ?", email).Limit(1).Find(&userPassProfile)

	// If not found res.Error = record not found
	return &userPassProfile, res.Error
}
