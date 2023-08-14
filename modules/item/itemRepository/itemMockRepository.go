package itemRepository

import (
	"context"

	"github.com/Rayato159/go-clean-unit-testing/modules/item"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockItemRepository struct {
	mock.Mock
}

func (m *MockItemRepository) InsertOneItem(pctx context.Context, req *item.Item) (primitive.ObjectID, error) {
	args := m.Called(pctx, req)
	return args.Get(0).(primitive.ObjectID), args.Error(1)
}
