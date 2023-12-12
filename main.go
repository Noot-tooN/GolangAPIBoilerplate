package main

import (
	"context"
	"fmt"
	"golangapi/databases/postgre"
	"log"
)

func main() {
	fmt.Println("Hello World")

	dsnBuilder, err := postgre.NewPgDsnBuilder(
		"localhost", 
		"6660",
		"some_db_user",
		"some_db_pass",
		"some_db_name",
	)

	if err != nil {
		log.Fatalln(err)
	}

	err = postgre.InitPostgrePool(*dsnBuilder)

	if err != nil {
		log.Fatalln(err)
	}

	pool := postgre.GetPostgrePool()

	conn, err := pool.Acquire(context.Background())

	defer conn.Release()

	if err != nil {
		log.Fatalln(err)
	}

	err = conn.Ping(context.Background())

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("PINGED!!")
}