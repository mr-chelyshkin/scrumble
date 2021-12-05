// +build wireinject

package main

import (
	"github.com/google/wire"
	hdfs_proxy "github.com/mr-chelyshkin/scrumble/hdfs-proxy"
	"github.com/mr-chelyshkin/scrumble/internal/daemon"
	"github.com/mr-chelyshkin/scrumble/internal/http_router"
)

func Init(cfg hdfs_proxy.Config, router http_router.Router) (daemon.Daemon, func(), error){
	wire.Build(
		wire.NewSet(
			daemon.Set,
			http_router.Set,
		),
	)
	return daemon.Daemon{}, nil, nil
}
