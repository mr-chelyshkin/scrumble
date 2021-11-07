package main

import (
	"github.com/mr-chelyshkin/scrumble/internal/config"
	"github.com/mr-chelyshkin/scrumble/internal/daemon"
)

func qwe(path string) (daemon.Daemon, func(), error) {
	cfgg := config.Config{}
	err := config.FromFile(path, cfgg)

	if err != nil {
		panic(err)
	}
	return Init(cfgg)
}

func main() {
	daemon.Run("qwe", qwe)
}