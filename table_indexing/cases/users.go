package cases

import (
	"log"

	"github.com/go-faker/faker/v4"
	"github.com/jmoiron/sqlx"
)

type Users struct {
	ID        string `json:"id"`
	Fullname  string `json:"fullname" faker:"last_name"`
	NIK       string `json:"nik" faker:"uuid_digit"`
	Birthdate string `json:"birthdate" faker:"date"`
}

func CreateManyData() {
	db, err := sqlx.Connect("mysql", "root:password01@(localhost:3306)/example")
	if err != nil {
		log.Fatalf("error on connecting to mysql %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("error on ping %v", err)
	}

	log.Println("successfully connected to db")

	numOfData := 50000

	for i := 0; i < numOfData; i++ {
		var user Users
		faker.FakeData(&user)
		if err != nil {
			log.Fatalf("error on data generation %v", err)
		}

		_, err := db.Exec("INSERT INTO users (fullname, nik, birthdate) VALUES (?,?,?)", user.Fullname, user.NIK, user.Birthdate)
		if err != nil {
			log.Fatalf("error on inserting %v", err)
		}
	}

	log.Printf("successfully inserted %d\n", numOfData)
	if err := db.Close(); err != nil {
		log.Printf("error on closing db %v", err)
	}
}
