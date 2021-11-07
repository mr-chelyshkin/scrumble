package daemon

import (
	"context"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"os"
	"os/signal"

	"github.com/mr-chelyshkin/scrumble/internal/logger"
	"github.com/mr-chelyshkin/scrumble/internal/stat"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

// ProvideContext initialize and return signal.NotifyContext for daemon.
func ProvideContext() (ctx context.Context, cancel func()) {
	return signal.NotifyContext(context.Background(), os.Interrupt)
}

// ProvideConfig initialize and return Daemon config data.
func ProvideConfig() (cfg Config, err error) {
	if err = viper.UnmarshalKey("daemon", &cfg); err != nil {
		return
	}

	// validate
	// ...

	return
}

// ProvideDaemon initialize and return Daemon object.
func ProvideDaemon(ctx context.Context, log *zap.Logger, s *stat.Stat, cfg Config, service Service) Daemon {
	wg, ctx := errgroup.WithContext(ctx)

	return Daemon{
		service: service,

		cfg:  cfg,
		log:  log,
		ctx:  ctx,
		wg:   wg,
		stat: s,
	}
}

var Set = wire.NewSet(
	ProvideDaemon,
	ProvideConfig,
	ProvideContext,

	logger.Set,
	stat.Set,
)
