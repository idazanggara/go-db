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

	_, err = db.NamedExec(`INSERT INTO customer (id, name, address, phone, email, saldo) VALUES (:id, :name, :address, :phone, :email, :saldo)`, map[string]interface{}{
		"id":      "C001",
		"name":    "Idaz",
		"address": "Jakarta",
		"phone":   "12321",
		"email":   "Idaz@gmail.com",
		"saldo":   5000000,
	})

	if err != nil {
		log.Fatalln(err)
	} else {
		log.Print("Successfully insert data")
	}

}
