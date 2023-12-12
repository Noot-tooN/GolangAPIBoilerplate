package main

import (
	"fmt"
	"golangapi/config"
	"golangapi/databases/gorm"
	postgresqlclient "golangapi/databases/postgre_sql_client"
	"log"
	"os"
)

func main() {
	// =================== LOAD IN THE CONFIG ===================
	err := config.InitConfig(os.Args[1:])

	if err != nil {
		log.Fatalln(err)
	}

	// =================== CONNECT TO THE DB ===================
	err = postgresqlclient.InitDefaultPostgreSqlClient()

	if err != nil {
		log.Fatalln(err)
	}

	_, err = gorm.InitDefaultPostgresGorm()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("OK")
}