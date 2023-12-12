package main

import (
	"fmt"
	"golangapi/config"
	"log"
	"os"
)

func main() {
	// =================== LOAD IN THE CONFIG ===================
	err := config.InitConfig(os.Args[1:])

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(config.Config)

	// dsnBuilder, err := postgre.NewPgDsnBuilder(
	// 	"localhost", 
	// 	"6660",
	// 	"some_db_user",
	// 	"some_db_pass",
	// 	"some_db_name",
	// )

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// err = postgre.InitPostgrePool(*dsnBuilder)

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// pool := postgre.GetPostgrePool()

	// conn, err := pool.Acquire(context.Background())

	// defer conn.Release()

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// err = conn.Ping(context.Background())

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println("PINGED!!")
}