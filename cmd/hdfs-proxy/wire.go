// +build wireinject

package main

import (
	hdfs_proxy "github.com/mr-chelyshkin/scrumble/hdfs-proxy"

	"github.com/mr-chelyshkin/scrumble/internal/daemon"
	"github.com/mr-chelyshkin/scrumble/internal/http_router"

	"github.com/google/wire"
)

func Init(cfg hdfs_proxy.Config, app http_router.AppHttpRouter) (daemon.Daemon, func(), error) {
	wire.Build(
		wire.NewSet(
			daemon.Set,
			http_router.Set,
		),
	)
	return daemon.Daemon{}, nil, nil
}
