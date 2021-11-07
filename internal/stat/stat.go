package stat

import (
	"context"
	"fmt"
	"net/http"
	"net/http/pprof"
	"strings"

	"github.com/arl/statsviz"
	"go.uber.org/zap"
)

type Stat struct {
	mux *http.ServeMux
	srv *http.Server

	probe Probe

	log *zap.Logger
	cfg Config
}

func (s *Stat) Start(_ context.Context) error {
	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Stat) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func (s Stat) String() string {
	return "statistic"
}

// Briefly describe exported endpoints for admin or devops
// for show application statistic data which can be managed by stat config.
func (s Stat) registerRoot() {
	s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var b strings.Builder

		b.WriteString("Service is up and running.\n")
		b.WriteString("Available endpoints:\n")
		b.WriteString("\n")
		b.WriteString("/probe/readiness - readiness check\n")
		b.WriteString("/probe/readiness - readiness check\n")
		b.WriteString("\n")

		if !s.cfg.Metrics.SwitchOff {
			b.WriteString("/live - live statistic visualization\n\n")
		}
		if !s.cfg.Profiler.SwitchOff {
			b.WriteString("/pprof - exported pprof\n\n")
		}
		if !s.cfg.Probe.SwitchOff {
			b.WriteString("/probe/readiness - readiness check\n")
			b.WriteString("/probe/liveness  - liveness check\n\n")
		}
		if !s.cfg.Prometheus.SwitchOff {
			b.WriteString("/metrics - prometheus via HTTP")
		}

		_, _ = fmt.Fprintln(w, b.String())
		s.log.Debug("register stat: '/'")
	})
}

// Live visualization application runtime statistics: Heap, Objects, Goroutines, GC, etc.
// for render plots use https://github.com/arl/statsviz project.
func (s Stat) registerMetrics() {
	if s.cfg.Metrics.SwitchOff {
		return
	}

	_ = statsviz.Register(s.mux, statsviz.Root("/live"))
	s.log.Debug("register stat: '/live'")
}

// Pprof tools for analyze application performance.
// by the way, you can use 'go tool pprof' with this data.
// Example: 'go tool pprof http://{ADDR:PORT}/pprof/heap'.
// For read: https://habr.com/ru/company/badoo/blog/324682/
func (s Stat) registerProfiler() {
	if s.cfg.Profiler.SwitchOff {
		return
	}

	s.mux.HandleFunc("/pprof/",        pprof.Index)
	s.mux.HandleFunc("/pprof/trace",   pprof.Trace)
	s.mux.HandleFunc("/pprof/cmdline", pprof.Cmdline)
	s.mux.HandleFunc("/pprof/profile", pprof.Profile)

	s.mux.Handle("/pprof/threadcreate", pprof.Handler("threadcreate"))
	s.mux.Handle("/pprof/goroutine",    pprof.Handler("goroutine"))
	s.mux.Handle("/pprof/block",        pprof.Handler("block"))
	s.mux.Handle("/pprof/heap",         pprof.Handler("heap"))
	s.mux.Handle("/pprof/allocs",       pprof.Handler("allocs"))
	s.mux.Handle("/pprof/mutex",        pprof.Handler("mutex"))

	s.log.Debug("register stat: '/pprof'")
}

// Register app probes for check application state.
// Required for k8s.
func (s *Stat) registerProbes() {
	if s.cfg.Probe.SwitchOff {
		return
	}

	s.mux.HandleFunc("/probe/readiness", s.handleProbe(s.probe.Readness))
	s.mux.HandleFunc("/probe/liveness",  s.handleProbe(s.probe.Liveness))

	s.log.Debug("register stat: '/probe'")
}

//
//
func (s *Stat) registerPrometheus() {
	//s.mux.Handle("/metrics", promhttp.HandlerFor(s.registry, promhttp.HandlerOpts{}))
}