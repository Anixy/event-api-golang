package main

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/Anixy/event-api-golang/helpers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/joho/godotenv"
)

func TestMigrate(t *testing.T) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/event_golang")
	if err != nil {
		log.Panic(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}

func TestGetEnv(t *testing.T) {

	err := godotenv.Load(".env")
	helpers.PanicIfError(err)
	myEnv, err := godotenv.Read()
	helpers.PanicIfError(err)

	fmt.Println(myEnv["db_name"])
	
}