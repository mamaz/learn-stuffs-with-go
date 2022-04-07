package infrastructure

import (
	"database/sql"
	"log"

	_ "github.com/newrelic/go-agent/v3/integrations/nrpq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewGormDB() *gorm.DB {
	db, err := sql.Open("nrpostgres", "host=localhost port=5432 user=postgres dbname=postgres password=password01 sslmode=disable")
	if err != nil {
		log.Fatalf("cannot connect to postgres")
	}

	config := postgres.Config{
		Conn: db,
	}
	gormdb, err := gorm.Open(postgres.New(config), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "public.", // schema name
			SingularTable: false,
		},
	})
	if err != nil {
		log.Fatalf("cannot open postgres")
	}

	return gormdb
}
