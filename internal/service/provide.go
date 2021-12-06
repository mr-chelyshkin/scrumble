package service

import (
	"context"
	"github.com/google/wire"
	"github.com/mr-chelyshkin/scrumble/internal/daemon"
	"github.com/mr-chelyshkin/scrumble/internal/stat"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func ProvideService(cfg Config, log *zap.Logger, app AppService) daemon.Service {


	return Service{
		name: app.Name(),
		cfg:  cfg,
		log:  log,

		runThirdParty: app.ThirdParty,
		runService:    app.Run,
	}
}

// ProvideConfig initialize and return Service config data.
func ProvideConfig() (cfg Config, err error) {
	if err = viper.UnmarshalKey("service", &cfg); err != nil {
		return
	}

	// validate
	// ...

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
	ProvideService,
	ProvideProbe,
	ProvideConfig,
)
