package example1

import (
	"github.com/mr-chelyshkin/scrumble/internal/daemon"
	"github.com/mr-chelyshkin/scrumble/internal/http"
	"github.com/mr-chelyshkin/scrumble/internal/logger"
	"github.com/mr-chelyshkin/scrumble/internal/stat"
)

type cfg struct {
	LogCfg    logger.Config `mapstructure:"logger" json:"log_cfg"  yaml:"log"`
	MetrCfg   stat.Config   `mapstructure:"stat"   json:"metr_cfg" yaml:"metrics"`
	DaemonCfg daemon.Config `mapstructure:"daemon" json:"daemon"   yaml:"daemon" `
	Http      http.Config   `mapstructure:"http"   json:"http"     yaml:"http"`
}

func main() {

}
