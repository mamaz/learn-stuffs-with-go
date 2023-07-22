package cases

import (
	"log"

	"github.com/go-faker/faker/v4"
	"github.com/jmoiron/sqlx"
)

type PhoneBook struct {
	FirstName string `db:"first_name" faker:"first_name"`
	LastName  string `db:"last_name" faker:"last_name"`
	Number    string `db:"number" faker:"e_164_phone_number"`
}

func CreateManyTelephoneContacts() {
	db, err := sqlx.Connect("mysql", "root:password01@(localhost:3306)/example")
	if err != nil {
		log.Fatalf("error on connecting to mysql %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("error on ping %v", err)
	}

	log.Println("successfully connected to db")

	numOfData := 200000
	perInsert := 100

	for i := 0; i < numOfData; i += perInsert {

		phoneBooks := []PhoneBook{}
		for j := 0; j < perInsert; j++ {
			var phoneBook PhoneBook
			faker.FakeData(&phoneBook)
			if err != nil {
				log.Fatalf("error on data generation %v", err)
			}

			phoneBooks = append(phoneBooks, phoneBook)
		}

		_, err := db.NamedExec("INSERT INTO phone_book (first_name, last_name, number) VALUES (:first_name, :last_name, :number)", phoneBooks)
		if err != nil {
			log.Fatalf("error on inserting %v", err)
		}
	}

	log.Printf("successfully inserted %d\n", numOfData)
	if err := db.Close(); err != nil {
		log.Printf("error on closing db %v", err)
	}
}
