package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"simple-go/database"
)

type Menu struct {
	ID            int    `db:"id"`
	Name          string `db:"name"`
	IsSpicy       bool   `db:"is_spicy"`
	IsRecommended bool   `db:"is_recommended"`
	Price         int    `db:"price"`
	Type          string `db:"type"`
}

type Employee struct{}

type Cashier struct{}

func main() {
	db, err := database.GetPostgresDB("user=postgres password=password01 dbname=example sslmode=disable")
	if err != nil {
		log.Fatalf("error on connecting to database: %v", err)
	}

	_, err = db.ExecContext(context.Background(), `TRUNCATE TABLE menus`)
	if err != nil {
		log.Fatalf("error on truncate to database: %v", err)
	}

	// insert menu
	menus := []Menu{
		{1, "Nasi Goreng", true, true, 25000, "meal"},
		{2, "Rendang", true, true, 35000, "meal"},
		{3, "Sate Ayam", true, true, 15000, "meal"},
		{4, "Nasi Padang", true, true, 30000, "meal"},
		{5, "Gado-Gado", false, true, 20000, "meal"},
		{6, "Soto", false, true, 18000, "meal"},
		{7, "Nasi Kuning", false, true, 22000, "meal"},
		{8, "Mie Goreng", true, true, 25000, "meal"},
		{9, "Ayam Goreng", false, true, 18000, "meal"},
		{10, "Pempek", true, true, 20000, "meal"},
		{11, "Martabak", false, true, 15000, "meal"},
		{12, "Bakso", true, true, 18000, "meal"},
		{13, "Nasi Uduk", false, true, 22000, "meal"},
		{14, "Tahu Isi", false, true, 5000, "meal"},
		{15, "Sambal", true, false, 5000, "meal"},
		{16, "Nasi Campur", true, true, 25000, "meal"},
		{17, "Lontong Sayur", false, true, 20000, "meal"},
		{18, "Rujak", true, true, 10000, "meal"},
		{19, "Sop Buntut", false, true, 35000, "meal"},
		{20, "Serabi", false, true, 5000, "meal"},
	}

	errChan := make(chan error)
	resultChan := make(chan sql.Result)

	for _, menu := range menus {
		query := `INSERT INTO menus(name,is_spicy,is_recommended,price,type) VALUES($1,$2,$3,$4,$5)`

		go func(menu Menu, resultChan chan sql.Result, errChan chan error) {
			result, err := db.ExecContext(context.Background(), query, menu.Name, menu.IsSpicy, menu.IsRecommended, menu.Price, menu.Type)
			if err != nil {
				errChan <- err
				return
			}

			resultChan <- result

		}(menu, resultChan, errChan)
	}

	for i := 0; i < len(menus); i++ {
		select {
		case res := <-resultChan:
			fmt.Printf("insert is ok: %v\n", res)

		case err := <-errChan:
			fmt.Printf("insert is error: %v\n", err)
		}
	}

	// insert employee
	_, err = db.ExecContext(context.Background(), `TRUNCATE TABLE menus`)
	if err != nil {
		log.Fatalf("error on truncate to database: %v", err)
	}

	// insert menu
	menus := []Menu{}

	errChan := make(chan error)
	resultChan := make(chan sql.Result)

	for _, menu := range menus {
		query := `INSERT INTO menus(name,is_spicy,is_recommended,price,type) VALUES($1,$2,$3,$4,$5)`

		go func(menu Menu, resultChan chan sql.Result, errChan chan error) {
			result, err := db.ExecContext(context.Background(), query, menu.Name, menu.IsSpicy, menu.IsRecommended, menu.Price, menu.Type)
			if err != nil {
				errChan <- err
				return
			}

			resultChan <- result

		}(menu, resultChan, errChan)
	}

	for i := 0; i < len(menus); i++ {
		select {
		case res := <-resultChan:
			fmt.Printf("insert is ok: %v\n", res)

		case err := <-errChan:
			fmt.Printf("insert is error: %v\n", err)
		}
	}
}
