package test

import (
	"database/sql"
	"log"
	"os"

	_ "embed"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func ConnectToDB() *sql.DB {
	dbConfig := os.Getenv("DB_USERNAME_TEST") + ":" + os.Getenv("DB_PASSWORD_TEST") + "@tcp(" + os.Getenv("DB_HOST_TEST") + ":" + os.Getenv("DB_PORT_TEST") + ")/" + os.Getenv("DB_DATABASE_TEST")
	db, err := sql.Open(os.Getenv("DB_CONNECTION_TEST"), dbConfig)
	if err != nil {
		log.Panic(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://../migrations",
		"mysql",
		driver,
	)
	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal(err)
	}
	return db
}