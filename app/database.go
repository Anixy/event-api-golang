package app

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



func GetDBConnection() *sql.DB {
	dbConfig := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_DATABASE")+"?parseTime=true"
	db, err := sql.Open(os.Getenv("DB_CONNECTION"), dbConfig)
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