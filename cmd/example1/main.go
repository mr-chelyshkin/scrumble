package main

import (
	"github.com/mr-chelyshkin/scrumble/internal/config"
	"github.com/mr-chelyshkin/scrumble/internal/daemon"
)

func app(path string) (daemon.Daemon, func(), error) {
	cfg := config.Config{}
	if err := config.FromFile(path, cfg); err != nil {
		panic(err)
	}

	return Init(cfg)
}

func main() {
	daemon.Run("example1", app)
}
