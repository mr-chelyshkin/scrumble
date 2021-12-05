// +build wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	hdfs_proxy "github.com/mr-chelyshkin/scrumble/hdfs-proxy"
	"github.com/mr-chelyshkin/scrumble/internal/daemon"
	"github.com/mr-chelyshkin/scrumble/internal/http_router"
)

func Init(cfg hdfs_proxy.Config, route func(e *echo.Echo), in func(ctx context.Context)) (daemon.Daemon, func(), error){
	wire.Build(
		wire.NewSet(
			daemon.Set,
			http_router.Set,
		),
	)
	return daemon.Daemon{}, nil, nil
}
