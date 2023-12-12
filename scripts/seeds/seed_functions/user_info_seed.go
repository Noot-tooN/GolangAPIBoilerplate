package seedfunctions

import (
	"fmt"

	"gorm.io/gorm"
)

func SeedUserInfo(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		fmt.Println("Seeding UserInfo")
		// TODO@Aleksa finish these seeds once we implement user service
		return nil
	})
}
