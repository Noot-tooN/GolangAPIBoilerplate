package main

import (
	"fmt"
	"golangapi/common"
	"golangapi/config"
	"golangapi/databases/gorm"
	postgresqlclient "golangapi/databases/postgre_sql_client"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	// =================== LOAD IN THE CONFIG ===================
	err := config.InitConfig(os.Args[1:])

	if err != nil {
		log.Fatalln(err)
	}

	// =================== CONNECT TO THE DB ===================
	var wg sync.WaitGroup
	
	wg.Add(2)
	
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
	_, err = gorm.InitDefaultPostgresGorm()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("OK")
}