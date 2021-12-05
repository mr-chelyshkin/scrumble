package logger

import (
	"path/filepath"

	"github.com/mr-chelyshkin/scrumble/internal/sys"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func loggerStdout(cfg Config) (*zap.Logger, error) {
	config := zap.NewProductionConfig()

	if cfg.Level != "" {
		if err := config.Level.UnmarshalText([]byte(cfg.Level)); err != nil {
			return nil, errors.Errorf("failed to unmarshal log level: %w", err)
		}
	}

	return config.Build()
}

func loggerFile(cfg Config) (*zap.Logger, error) {
	maxSize := cfg.MaxSize
	if maxSize == 0 {
		maxSize = 100
	}
	maxAge := cfg.MaxAge
	if maxAge == 0 {
		maxAge = 28
	}
	maxBackups := cfg.MaxBackups
	if maxBackups == 0 {
		maxBackups = 3
	}

	ok, err := sys.IsWritable(filepath.Dir(cfg.LogPath))
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.Errorf("%s doesn't have access to write to %s", sys.CurrentUsername(), filepath.Dir(cfg.LogPath))
	}

	logLevel := zap.NewAtomicLevelAt(zap.InfoLevel)
	if err := logLevel.UnmarshalText([]byte(cfg.Level)); err != nil {
		return nil, errors.Errorf("failed to unmarshal log level: %w", err)
	}
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   cfg.LogPath,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
	})
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		logLevel,
	)

	return zap.New(core), nil
}

func ProvideLoggerZap(cfg Config) (*zap.Logger, error) {
	if cfg.LogPath != "" {
		return loggerFile(cfg)
	}
	return loggerStdout(cfg)
}

func ProvideConfig() (cfg Config, err error) {
	if err = viper.UnmarshalKey("logger", &cfg); err != nil {
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
