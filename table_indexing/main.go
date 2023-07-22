package main

import (
	"table-indexing/cases"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func main() {
	cases.CreateManyTelephoneContacts()
}
