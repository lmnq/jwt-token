package repo

import (
	"context"
	"fmt"

	"github.com/lmnq/jwt-token/internal/entity"
	"github.com/lmnq/jwt-token/internal/errs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	dbName         = "jwt"
	collectionName = "tokens"
)

type TokenRepo struct {
	client *mongo.Client
}

func New(client *mongo.Client) *TokenRepo {
	return &TokenRepo{
		client: client,
	}
}

func (r *TokenRepo) StoreRefreshToken(ctx context.Context, refreshToken *entity.RefreshToken) error {
	collection := r.client.Database(dbName).Collection(collectionName)

	_, err := collection.InsertOne(ctx, refreshToken)
	if err != nil {
		return fmt.Errorf("failed to store refresh token: %w", err)
	}

	return nil
}

func (r *TokenRepo) GetRefreshToken(ctx context.Context, guid string) (*entity.RefreshToken, error) {
	collection := r.client.Database(dbName).Collection(collectionName)

	var refreshToken entity.RefreshToken
	err := collection.FindOne(ctx, bson.M{"guid": guid}).Decode(&refreshToken)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errs.ErrNotFound{Message: "refresh token not found"}
		}
		return nil, fmt.Errorf("failed to get refresh token: %w", err)
	}

	return &refreshToken, nil
}
