package seedfunctions

import (
	"fmt"
	"golangapi/constants"
	"golangapi/datalayers"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedRoles(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		fmt.Println("Seeding Roles")

		roleDl := datalayers.NewGormRoleDatalayer()

		for _, v := range constants.ValidRoleNames {
			txWithClauses := tx.Clauses(clause.OnConflict{DoNothing: true})

			createErr := roleDl.Create(v, txWithClauses)

			if createErr != nil {
				return createErr
			}
		}

		return nil
	})
}
