package repository

import (
	"context"
	"fmt"
	"gw-notification/config"
	"gw-notification/internal/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

const (
	connectTimeout = 10 * time.Second
)

const (
	ExchangeCollection = "exchange"
)

type MongoRepository struct {
	db *mongo.Database
}

func (r *MongoRepository) Save(ctx context.Context, exchange domain.Exchange) (*primitive.ObjectID, error) {
	result, err := r.db.Collection(ExchangeCollection).InsertOne(ctx, exchange)
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		zap.L().Error(ErrNotObjectID.Error())
		return nil, ErrNotObjectID
	}

	return &oid, nil
}

func NewMongoRepository(ctx context.Context, cfg config.DatabaseConfig) *MongoRepository {
	r := &MongoRepository{}

	ctx, cancel := context.WithTimeout(ctx, connectTimeout)
	defer cancel()

	uri := fmt.Sprintf("mongodb://%s:%v", cfg.Host, cfg.Port)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil
	}

	r.db = client.Database(cfg.Name)

	return r
}
