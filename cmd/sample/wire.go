// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/mr-chelyshkin/scrumble/internal/config"
	"github.com/mr-chelyshkin/scrumble/internal/daemon"
	"github.com/mr-chelyshkin/scrumble/internal/http"

)

func Init(cfg config.Config) (daemon.Daemon, func(), error){
	wire.Build(
		wire.NewSet(
			daemon.Set,
			http.Set,
		),
	)
	return daemon.Daemon{}, nil, nil
}
