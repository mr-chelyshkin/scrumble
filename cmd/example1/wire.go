// +build wireinject

package main

import (
	"github.com/mr-chelyshkin/scrumble/internal/config"
	"github.com/mr-chelyshkin/scrumble/internal/daemon"
	"github.com/mr-chelyshkin/scrumble/internal/http_router"

	"github.com/google/wire"
)

func Init(cfg config.Config) (daemon.Daemon, func(), error){
	wire.Build(
		wire.NewSet(
			daemon.Set,
			http_router.Set,
		),
	)
	return daemon.Daemon{}, nil, nil
}
