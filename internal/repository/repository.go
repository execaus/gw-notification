package repository

import (
	"context"
	"gw-notification/config"
	"gw-notification/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	Save(ctx context.Context, exchange domain.Exchange) (primitive.ObjectID, error)
}

func NewRepository(
	ctx context.Context,
	cfg config.DatabaseConfig,
) (repo Repository, close func(ctx context.Context) error) {
	r := NewMongoRepository(ctx, cfg)
	return r, r.Close
}
