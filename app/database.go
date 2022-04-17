package app

import (
	"database/sql"
	"log"

	_ "embed"

	"github.com/Anixy/event-api-golang/helpers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/joho/godotenv"
)



func GetDBConnection() *sql.DB {
	err := godotenv.Load(".env")
	helpers.PanicIfError(err)
	testEnv, err := godotenv.Read(".env")
	helpers.PanicIfError(err)
	dbConfig := testEnv["DB_USERNAME"] + ":" + testEnv["DB_PASSWORD"] + "@tcp(" + testEnv["DB_HOST"] + ":" + testEnv["DB_PORT"] + ")/" + testEnv["DB_DATABASE"]+"?parseTime=true"
	helpers.PanicIfError(err)
	db, err := sql.Open(testEnv["DB_CONNECTION"], dbConfig)
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
	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal(err)
	}
	return db
}