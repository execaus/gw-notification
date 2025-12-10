package service

import (
	"context"
	"gw-notification/internal/domain"
	"gw-notification/internal/repository"

	"go.uber.org/zap"
)

type Service interface {
	Save(ctx context.Context, exchange domain.Exchange) error
}

type ExchangeService struct {
	r repository.Repository
}

func (s *ExchangeService) Save(ctx context.Context, exchange domain.Exchange) error {
	if _, err := s.r.Save(ctx, exchange); err != nil {
		zap.L().Error(err.Error())
		return err
	}
	return nil
}

func NewExchangeService(r repository.Repository) Service {
	return &ExchangeService{
		r: r,
	}
}
