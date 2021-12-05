package http_router

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Router interface {
	Name() string
	Echo(e *echo.Echo)

	Utils(ctx context.Context) error
	ThirdParty(chan <- error)
}

type Service struct {
	name string
	cfg  Config

	log  *zap.Logger
	e    *echo.Echo

	runUtils      func(ctx context.Context) error
	runThirdParty func(chan <- error)
}

func (s Service) Start(ctx context.Context) error {
	s.log.Info("Starting HTTP server", zap.String("http.addr", s.cfg.Addr), zap.String("app", s.name))

	thirdPartyErr := make(chan error, 1)
	go func() {
		s.runThirdParty(thirdPartyErr)

		for {
			select {
			case err := <-thirdPartyErr:
				s.log.Error("ThirdParty error", zap.Error(err), zap.String("app", s.name))
			default:
			}
		}
	}()

	if err := s.runUtils(ctx); err != nil {
		return err
	}
	if err := s.e.Start(s.cfg.Addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s Service) Shutdown(ctx context.Context) error {
	return s.e.Shutdown(ctx)
}

func (s Service) String() string {
	return s.name
}
