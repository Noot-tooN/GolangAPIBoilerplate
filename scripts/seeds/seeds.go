package main

import (
	"fmt"
	"golangapi/common"
	"golangapi/config"
	"golangapi/databases/gorm"
	postgresqlclient "golangapi/databases/postgre_sql_client"
	seedfunctions "golangapi/scripts/seeds/seed_functions"
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
	
	// // =================== PERFORM THE SEEDS ===================
	fmt.Println("Starting seeds")

	seedfunctions.SeedUserInfo(gDb)

	fmt.Println("Finished seeding")
	
}