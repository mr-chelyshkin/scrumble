// +build wireinject

package main

import (
	torrent_fetcher "github.com/mr-chelyshkin/scrumble/torrent-fetcher"

	"github.com/mr-chelyshkin/scrumble/internal/service"
	"github.com/mr-chelyshkin/scrumble/internal/daemon"

	"github.com/google/wire"
)

func Init(cfg torrent_fetcher.Config, app service.AppService) (daemon.Daemon, func(), error) {
	wire.Build(
		wire.NewSet(
			daemon.Set,
			service.Set,
		),
	)
	return daemon.Daemon{}, nil, nil
}
