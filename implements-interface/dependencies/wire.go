//go:build wireinject
// +build wireinject

package dependencies

import (
	"implements-interface/example"
	"implements-interface/infra"

	"github.com/google/wire"
)

func InitRepo(jtDB infra.JTDB, invDb infra.INVDB) *example.Repo {
	wire.Build(
		example.NewRepo,
	)

	return &example.Repo{}
}
