package example

import (
	"implements-interface/infra"
)

type Model struct {
	ID string
}

type RepoInterface interface {
	FindByID(ID string) Model
}

type Repo struct {
	jtDB  infra.DBInterface
	invDB infra.DBInterface
}

func NewRepo(jtDB infra.JTDB, invDB infra.INVDB) *Repo {
	return &Repo{
		jtDB:  jtDB,
		invDB: invDB,
	}
}

func (r *Repo) FindByID(ID string) Model {
	return Model{
		ID: ID,
	}
}
