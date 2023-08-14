package itemRepository

import (
	"context"
	"fmt"
	"time"

	"github.com/Rayato159/go-clean-unit-testing/modules/item"
	"github.com/Rayato159/go-clean-unit-testing/pkg/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//*** Assume this layer is always called a database

type (
	ItemRepositoryService interface {
		InsertOneItem(pctx context.Context, req *item.Item) (primitive.ObjectID, error)
	}

	itemRepository struct{}
)

func NewItemRepository() ItemRepositoryService {
	return &itemRepository{}
}

func (r *itemRepository) itemDbConn(ctx context.Context) *mongo.Collection {
	return database.DbConn().Database("whyyoureadthis").Collection("items")
}

func (r *itemRepository) InsertOneItem(pctx context.Context, req *item.Item) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(pctx, time.Second*10)
	defer cancel()

	db := r.itemDbConn(ctx)
	defer db.Database().Client().Disconnect(ctx)

	result, err := db.InsertOne(ctx, req, nil)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("InsertOne error: %v", err)
	}
	return result.InsertedID.(primitive.ObjectID), nil
}
