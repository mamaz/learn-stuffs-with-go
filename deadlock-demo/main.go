package main

import (
	"context"
	"deadlock-demo/call"
	"deadlock-demo/database"
	"log"
)

type Deadlock struct {
	ID          int
	Name        string
	Description string
}

func main() {
	db := database.GetDB("mysql")

	query := `
		SELECT * from deadlock where id = 1
	`

	var d Deadlock
	err := db.QueryRowContext(context.Background(), query).Scan(&d.ID, &d.Name, &d.Description)
	if err != nil {
		log.Fatalf("error on querying %v", err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("error on beginning transaction %v", err)
	}

	var firstD Deadlock
	err = tx.QueryRowContext(context.Background(), query).Scan(&firstD.ID, &firstD.Name, &firstD.Description)
	if err != nil {
		log.Fatalf("error on querying %v", err)
	}

	log.Printf("dapet d %+v \n", d)

	var secondD Deadlock
	err = tx.QueryRowContext(context.Background(), query).Scan(&secondD.ID, &secondD.Name, &secondD.Description)
	if err != nil {
		log.Fatalf("error on querying %v", err)
	}

	log.Printf("dapet d %+v \n", secondD)

	res := call.UpdateData(1, "Durmadji")
	log.Printf("updated? %+v \n", res)

	tx.Commit()
}
