package main

import (
	torrent_fetcher "github.com/mr-chelyshkin/scrumble/torrent-fetcher"

	"github.com/mr-chelyshkin/scrumble/internal/daemon"
	"github.com/mr-chelyshkin/scrumble/internal/sys"
)

func app(path string) (daemon.Daemon, func(), error) {
	cfg := torrent_fetcher.Config{}
	if err := sys.ParseFile(path, cfg); err != nil {
		panic(err)
	}
	return Init(cfg, torrent_fetcher.App{})

}

func main() {
	daemon.Run("torrent-fetcher", app)
}
