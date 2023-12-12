package main

import (
	"golangapi/config"
	"golangapi/databases/postgre_pool"
	"log"
	"os"
)

func main() {
	// =================== LOAD IN THE CONFIG ===================
	err := config.InitConfig(os.Args[1:])

	if err != nil {
		log.Fatalln(err)
	}

	postgre_pool.InitDefaultPool()
}