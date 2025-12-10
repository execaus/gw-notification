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

func NewRepository(ctx context.Context, cfg config.DatabaseConfig) Repository {
	return NewMongoRepository(ctx, cfg)
}
