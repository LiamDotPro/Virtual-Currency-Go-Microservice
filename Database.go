package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

var Connection *gorm.DB

// If you want to start a service locally to test..
// docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d postgres
// After which you need to create a database called 'notifications' and add postgres as default admin
// After inspect your container to get the ip it runs on for pgadmin..
// 172.17.0.2
func database() {

	db, err := gorm.Open("postgres", os.Getenv("DBCONNECTIONSTRING"))

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	defer db.Close()

	// Turn logging for the database on.
	// Database logging mode needs to be passed as env
	if os.Getenv("DBDEBUGMODE") != "false" {
		db.LogMode(true)
	}

	// Make Master connection available globally.
	Connection = db

	fmt.Println("Database Successfully Connected!")

	migrations()
	seed()
}

func migrations() {
	Connection.AutoMigrate(&Wallet{})
}

func seed() {

}
