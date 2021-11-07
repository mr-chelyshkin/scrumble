package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Path struct {
	Value string
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
