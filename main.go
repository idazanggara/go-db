package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "postgres"
	dbPassword := "postgres"
	dbName := "db_sample"
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	driverName := "postgres"
	// fmt.Println(dataSourceName)

	db, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		panic(err)
	} else {
		log.Println("Connected")
	}

	defer func(conn *sqlx.DB) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	// defer func() {
	// 	err := db.Close()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }()
}
