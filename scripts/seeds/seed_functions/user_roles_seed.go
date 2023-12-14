package seedfunctions

import (
	"fmt"
	"golangapi/constants"
	"golangapi/datalayers"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedUserRoles(db *gorm.DB) error {
	fmt.Println("Seeding User Roles")

	// Get Roles
	roleDl := datalayers.NewGormRoleDatalayer()

	allRoles, err := roleDl.GetRoles(db)

	if err != nil {
		return err
	}

	if len(allRoles) == 0 {
		return fmt.Errorf("no roles")
	}

	roleUuidMap := map[string]uuid.UUID{}

	for _, v := range allRoles {
		roleUuidMap[fmt.Sprintf("%v", v.Name)] = v.Uuid
	}

	// Get Users
	userDl := datalayers.NewGormUserDatalayer()

	allUsers, err := userDl.FindAllUsers(db)

	if err != nil {
		return err
	}

	if len(allUsers) == 0 {
		return fmt.Errorf("no users")
	}

	userUuidMap := map[string]uuid.UUID{}

	for _, v := range allUsers {
		userUuidMap[fmt.Sprintf("%v", v.Email)] = v.Uuid
	}

	userRoleDl := datalayers.NewGormUserRoleDatalayer()

	return db.Transaction(func(tx *gorm.DB) error {
		txWithClauses := tx.Clauses(clause.OnConflict{DoNothing: true})

		adminUuid := userUuidMap[randomUsers[0].Email]

		// Add ADMIN roles
		userRoleDl.AddRoleForUserByUuid(
			roleUuidMap[fmt.Sprintf("%v", constants.ADMIN)],
			adminUuid,
			txWithClauses,
		)

		return nil
	})
}
