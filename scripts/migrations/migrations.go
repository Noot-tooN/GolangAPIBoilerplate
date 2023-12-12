package main

import (
	"fmt"
	"golangapi/common"
	"golangapi/config"
	"golangapi/databases/gorm"
	postgresqlclient "golangapi/databases/postgre_sql_client"
	"golangapi/models"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/gorm/logger"
)

func main() {	
	// =================== LOAD IN THE CONFIG ===================
	err := config.InitConfig(os.Args[1:])

	if err != nil {
		log.Fatalln(err)
	}

	// =================== CONNECT TO THE DB ===================
	var wg sync.WaitGroup
		
	wg.Add(1)

	go common.CheckWg(
		common.EaseOffRetryStrategy(3, time.Second, func() (bool, error) {
			err = postgresqlclient.InitDefaultPostgreSqlClient()

			if err != nil {
				return true, err
			}
			
			return false, err
		}),
		&wg,
	)

	wg.Wait()

	// =================== LINK THE POSTGRE TO GORM ===================
	gDb, err := gorm.InitDefaultPostgresGorm()

	if err != nil {
		log.Fatalln(err)
	}

	gDb.Logger = logger.Default.LogMode(logger.Info)
	
	// // =================== PERFORM THE MIGRATIONS ===================
	fmt.Println("Starting migrations")

	err = gDb.AutoMigrate(
		&models.UserInfo{},
	)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Finished migrating models")
}