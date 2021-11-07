package stat

import (
	"net/http"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// ProvideStat initialize and return Stat object.
func ProvideStat(cfg Config, log *zap.Logger, probe Probe) *Stat {
	mux := http.NewServeMux()

	s := &Stat{
		mux:   mux,
		log:   log,
		cfg:   cfg,
		probe: probe,

		srv: &http.Server{
			Handler: mux,
			Addr:    cfg.Addr,
		},
	}

	// Register http routes.
	s.registerRoot()
	s.registerProbes()
	s.registerMetrics()
	s.registerProfiler()

	log.Info("Stat initialized", zap.String("http.addr", cfg.Addr))
	return s
}

// ProvideConfig initialize and return Stat config data.
func ProvideConfig() (cfg Config, err error) {
	if err = viper.UnmarshalKey("stat", &cfg); err != nil {
		return
	}

	// validate
	if cfg.Addr == "" {
		err = errors.New("config must have 'stat.addr' value.")
		return
	}

	return cfg, err
}

var Set = wire.NewSet(
	ProvideStat,
	ProvideConfig,
)
