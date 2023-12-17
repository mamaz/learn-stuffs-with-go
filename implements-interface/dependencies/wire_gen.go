// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package dependencies

import (
	"implements-interface/example"
	"implements-interface/infra"
)

// Injectors from wire.go:

func InitRepo(jtDB infra.JTDB, invDb infra.INVDB) *example.Repo {
	repo := example.NewRepo(jtDB, invDb)
	return repo
}
