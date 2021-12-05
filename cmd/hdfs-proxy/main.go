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
	appRoute := hdfs_proxy.HdfsProxy{}

	return Init(cfg, appRoute)
}

func main() {
	daemon.Run("hdfs-proxy", app)
}
