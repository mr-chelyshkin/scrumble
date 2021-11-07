package stat

import (
	"context"
	"fmt"
	"net/http"
)

// ProbeFunc probes service, returning error on failed probe.
type ProbeFunc func(ctx context.Context) error

type Probe struct {
	Readness ProbeFunc
	Liveness ProbeFunc
}

// handleProbe execute ProbeFunc(readness / liveness) and write to ResponseWriter with result.
// no-op ProbeFunc that return nil, return OK status.
func (s *Stat) handleProbe(f func(ctx context.Context) error) func(w http.ResponseWriter, r *http.Request) {
	check := func(ctx context.Context) error {
		if f == nil {
			return nil
		}
		return f(ctx)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if err := check(r.Context()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprintln(w, "Probe error:", err)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintln(w, "Probe OK")
	}
}
