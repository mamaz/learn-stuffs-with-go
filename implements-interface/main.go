package main

import (
	"fmt"
	"implements-interface/dependencies"
	"implements-interface/infra"
)

func main() {
	// using normal dependencies
	db := &DB{}
	jtdb := &JTDB{
		JTDBInterface: db,
	}
	fmt.Println(getJTDB(jtdb))

	// using Wire
	repo := dependencies.InitRepo(infra.NewDB(), infra.NewDB())
	fmt.Println(repo.FindByID("id"))
}

type JTDBInterface interface {
	GetMaster() string
	GetSlave() string
}

type JTDB struct {
	JTDBInterface
}

type DB struct {
}

func (db *DB) GetMaster() string {
	return "master"
}

func (db *DB) GetSlave() string {
	return "slave"
}

func getJTDB(db *JTDB) string {
	return db.GetMaster() + db.GetSlave()
}
