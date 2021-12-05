package http_router

import (
	"context"
	"github.com/mr-chelyshkin/scrumble/internal/daemon"
	"github.com/mr-chelyshkin/scrumble/internal/http_router/custom_middleware"
	"github.com/mr-chelyshkin/scrumble/internal/stat"
	"go.uber.org/zap/zapcore"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func setEchoLogger(echoLogger echo.Logger, zapLogger *zap.Logger) {
	switch {
	case zapLogger.Core().Enabled(zapcore.DebugLevel):
		echoLogger.SetLevel(1)
	case zapLogger.Core().Enabled(zapcore.InfoLevel):
		echoLogger.SetLevel(2)
	case zapLogger.Core().Enabled(zapcore.WarnLevel):
		echoLogger.SetLevel(3)
	case zapLogger.Core().Enabled(zapcore.ErrorLevel):
		echoLogger.SetLevel(4)
	}
}

// ProvideHttpRouter initialize and return Stat object.
func ProvideHttpRouter(ctx context.Context, cfg Config, log *zap.Logger, appRoute func(e *echo.Echo), In func(ctx context.Context)) daemon.Service {
	go func() {
		In(ctx)
	}()

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	//e.Logger =

	e.Use(middleware.RequestID())
	e.Use(custom_middleware.RequestLogger(log))

	setEchoLogger(e.Logger, log)
	appRoute(e)


	return Service{
		name: "http_router",
		cfg:  cfg,
		log:  log,
		e:    e,
	}
}

// ProvideConfig initialize and return Service config data.
func ProvideConfig() (cfg Config, err error) {
	if err = viper.UnmarshalKey("http_router", &cfg); err != nil {
		return
	}

	// validate
	if cfg.Addr == "" {
		err = errors.New("config must have 'http_router.addr' value.")
		return
	}

	return cfg, err
}

// ProvideProbe initialize probes for http service.
func ProvideProbe() stat.Probe {
	return stat.Probe{
		Readness: func(ctx context.Context) error {
			return nil
		},
		Liveness: func(ctx context.Context) error {
			return nil
		},
	}
}

var Set = wire.NewSet(
	ProvideHttpRouter,
	ProvideProbe,
	ProvideConfig,
)
