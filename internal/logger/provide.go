package logger

import (
	"fmt"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func ProvideLoggerZap(cfg Config) (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	if cfg.Level != "" {
		if err := config.Level.UnmarshalText([]byte(cfg.Level)); err != nil {
			return nil, fmt.Errorf("failed to unmarshal log level: %w", err)
		}
	}
	return config.Build()
}

func ProvideConfig() (cfg Config, err error) {
	if err = viper.UnmarshalKey("stat", &cfg); err != nil {
		return
	}

	// validate
	// ...

	cfgToStack := cfg
	return cfgToStack, nil
}

var Set = wire.NewSet(
	ProvideLoggerZap,
	ProvideConfig,
)
