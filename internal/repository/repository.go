package repository

import (
	"context"
	"gw-notification/config"
	"gw-notification/internal/domain"
)

type Repository interface {
	Save(ctx context.Context, exchange domain.Exchange) (*domain.Exchange, error)
}

func NewRepository(ctx context.Context, cfg config.DatabaseConfig) Repository {
	return NewMongoRepository(ctx, cfg)
}
