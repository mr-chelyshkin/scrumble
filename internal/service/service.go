package service

import (
	"context"
	"go.uber.org/zap"
)

type AppService interface {
	Name() string

	Run() error
	ThirdParty(chan <- error)
}

type Service struct {
	name string
	cfg  Config

	log  *zap.Logger

	runThirdParty func(chan <- error)
	runService    func() error
}

func (s Service) Start(ctx context.Context) error {
	s.log.Info("Starting service", zap.String("app", s.name))

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

	if err := s.runService(); err != nil {
		return err
	}
	return nil
}

func (s Service) Shutdown(ctx context.Context) error {
	return nil
}

func (s Service) String() string {
	return s.name
}