package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password01"
	dbname   = "example"
)

func GetDB(dbname string) *sql.DB {
	switch dbname {
	case "mysql":
		return GetDBMySQL()
	case "postgres":
		return GetDBPostgres()
	}

	return nil
}

func GetDBMySQL() *sql.DB {
	db, err := sql.Open(
		"mysql", "root:password01@tcp(127.0.0.1:3306)/example")
	if err != nil {
		log.Fatalf("error on opening database: %v", err)
	}

	pingErr := db.Ping()
	if pingErr == nil {
		log.Println("successfully connected to database")
	}

	return db
}

func GetDBPostgres() *sql.DB {
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		log.Fatalf("error on opening database: %v", err)
	}

	pingErr := db.Ping()
	if pingErr == nil {
		log.Println("successfully connected to database")
	}

	return db
}
