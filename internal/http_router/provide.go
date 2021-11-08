package http_router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mr-chelyshkin/scrumble/internal/daemon"
	"github.com/mr-chelyshkin/scrumble/internal/stat"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// ProvideHttpRouter initialize and return Stat object.
func ProvideHttpRouter(cfg Config, log *zap.Logger) daemon.Service {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.GET("/", hello)

	return Service{
		name: "http_router",
		cfg:  cfg,
		log:  log,
		e:    e,
	}
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
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
			fmt.Println("READNESS")
			return nil
		},
		Liveness: func(ctx context.Context) error {
			fmt.Println("LIVENESS")
			return nil
		},
	}
}

var Set = wire.NewSet(
	ProvideHttpRouter,
	ProvideProbe,
	ProvideConfig,
)
