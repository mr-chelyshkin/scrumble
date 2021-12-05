package daemon

import (
	"context"
	"flag"
	"os"
	"path/filepath"
	"time"

	"github.com/mr-chelyshkin/scrumble/internal/stat"
	"github.com/mr-chelyshkin/scrumble/internal/version"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type Service interface {
	Start(ctx context.Context) error
	Shutdown(ctx context.Context) error
	String() string
}

type Daemon struct {
	service Service

	cfg   Config
	log   *zap.Logger
	stat  *stat.Stat
	wg    *errgroup.Group
	ctx   context.Context
}

const (
	exitCodeOk             = 0
	exitCodeApplicationErr = 1
	exitCodeWatchdog       = 2
)

const (
	shutdownTimeout = time.Second * 5
	watchdogTimeout = shutdownTimeout + time.Second*5
)

func (d *Daemon) startService(ctx context.Context, service Service) {
	log := d.log.Named(service.String())

	d.wg.Go(func() error {
		log.Info("Starting service")
		return service.Start(ctx)
	})
	d.wg.Go(func() error {
		<-ctx.Done()
		log.Info("Shutting down service")

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer shutdownCancel()

		return service.Shutdown(shutdownCtx)
	})
}

func (d *Daemon) Run() {
	log := d.log.Named("app")
	log.Info("Starting", version.LogFields()...)

	if !d.cfg.NoStat {
		d.startService(d.ctx, d.stat)
	} else {
		d.log.Info("Stat switch off by application config")
	}
	d.startService(d.ctx, d.service)

	go func() {
		// Guaranteed way to kill application.
		<-d.ctx.Done()
		// Context is canceled, giving application time to shut down gracefully.
		d.log.Info("Waiting for application shutdown")
		time.Sleep(watchdogTimeout)

		// Probably deadlock, forcing shutdown.
		d.log.Warn("Graceful shutdown watchdog triggered: forcing shutdown")
		os.Exit(exitCodeWatchdog)
	}()

	// Note that we are calling os.Exit() here and no
	if err := d.wg.Wait(); err != nil {
		d.log.Error("Application returned error", zap.Error(err))
		os.Exit(exitCodeApplicationErr)
	}

	d.log.Info("Application stopped")
	os.Exit(exitCodeOk)
}

func Run(name string, f func(p string) (Daemon, func(), error)) {
	cfgPath := flag.String("config", filepath.Join("_config", name, "config.toml"), "config path")
	flag.Parse()

	app, cleanup, err := f(*cfgPath)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	app.Run()
}
