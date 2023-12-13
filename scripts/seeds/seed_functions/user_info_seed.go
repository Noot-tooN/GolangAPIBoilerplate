package seedfunctions

import (
	"fmt"
	"golangapi/controllers/inputs"
	"golangapi/datalayers"
	"golangapi/services"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var randomUsers []inputs.Registration = []inputs.Registration{
	{
		Email:    "olora@fi.mv",
		Password: "Test1234",
	},
	{
		Email:    "besuzujo@fecutet.vn",
		Password: "Test1234",
	},
	{
		Email:    "lipore@ih.ca",
		Password: "Test1234",
	},
	{
		Email:    "niw@itragu.mn",
		Password: "Test1234",
	},
	{
		Email:    "kanamtu@womtev.ai",
		Password: "Test1234",
	},
	{
		Email:    "ab@jafed.in",
		Password: "Test1234",
	},
	{
		Email:    "wafsivpu@levun.su",
		Password: "Test1234",
	},
}

func SeedUserInfo(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		fmt.Println("Seeding UserInfo")

		txWithClauses := tx.Clauses(clause.OnConflict{DoNothing: true})

		userService := services.NewUserService(
			datalayers.NewGormUserDatalayer(),
			services.NewDefaultCryptoService(),
			services.NewDefaultSymmetricalPasetoTokenHandler(),
			txWithClauses,
		)

		for _, v := range randomUsers {
			err := userService.CreateUser(v)

			if err != nil {
				return err
			}
		}

		return nil
	})
}
