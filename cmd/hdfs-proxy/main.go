package main

import (
	hdfs_proxy "github.com/mr-chelyshkin/scrumble/hdfs-proxy"
	"github.com/mr-chelyshkin/scrumble/internal/daemon"
	"github.com/mr-chelyshkin/scrumble/internal/sys"
)


func app(path string) (daemon.Daemon, func(), error) {
	cfg := hdfs_proxy.Config{}
	if err := sys.ParseFile(path, cfg); err != nil {
		panic(err)
	}

	return Init(cfg, hdfs_proxy.Route, hdfs_proxy.In)
}

func main() {
	daemon.Run("hdfs-proxy", app)
}
