package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mr-chelyshkin/scrumble/internal/daemon"
	"github.com/mr-chelyshkin/scrumble/internal/http_router"
	"github.com/mr-chelyshkin/scrumble/internal/logger"
	"github.com/mr-chelyshkin/scrumble/internal/stat"

	"github.com/spf13/viper"
)

type Config struct {
	Log        logger.Config      `mapstructure:"logger" json:"log"    yaml:"log"`
	Stat       stat.Config        `mapstructure:"stat"   json:"stat"   yaml:"stat"`
	Daemon     daemon.Config      `mapstructure:"daemon" json:"daemon" yaml:"daemon"`
	HttpRouter http_router.Config `mapstructure:"http"   json:"http"   yaml:"http_router"`
}

func FromFile(path string, cfg interface{}) error {
	file, err := os.OpenFile(filepath.Clean(path), os.O_RDONLY, 0600)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}

	viper.SetConfigType("toml")
	if err := viper.ReadConfig(file); err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	return nil
}
