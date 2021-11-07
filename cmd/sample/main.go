package main

import (

	"github.com/mr-chelyshkin/scrumble/internal/config"
	"github.com/mr-chelyshkin/scrumble/internal/daemon"
	"github.com/mr-chelyshkin/scrumble/internal/http"
	"github.com/mr-chelyshkin/scrumble/internal/logger"
	"github.com/mr-chelyshkin/scrumble/internal/stat"

)

type cfg struct {
	LogCfg  logger.Config `mapstructure:"logger"`
	MetrCfg stat.Config   `mapstructure:"stat"`
	DCfg    daemon.Config `mapstructure:"daemon"`
	Http    http.Config   `mapstructure:"http"`
}

func qwe(path config.Path) (daemon.Daemon, func(), error) {
	cfgg := cfg{}
	err := config.FromFile(path.Value, cfgg)

	if err != nil {
		panic(err)
	}
	return Init(cfgg)
}

func main() {
	daemon.Run("qwe", qwe)
}