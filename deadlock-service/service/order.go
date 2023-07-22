package service

import (
	"deadlock-service/database"
	"log"
)

func UpdateData(ID int, name string) bool {
	db := database.GetDB("mysql")

	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("error on beginning transaction %+v", err)
	}

	insertQuery := `UPDATE deadlock SET name = ? WHERE id = ?`
	result, err := tx.Exec(insertQuery, name, ID)
	if err != nil {
		log.Fatalf("error on updating query %+v", err)
	}

	affectedNumbers, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("error on getting rows affected: %+v", err)
	}

	tx.Commit()

	return affectedNumbers == 1
}
