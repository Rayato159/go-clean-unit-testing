package whydoweneedtest

import (
	"context"
	"net/http"
	"testing"

	"github.com/Rayato159/go-clean-unit-testing/modules/item"
	"github.com/Rayato159/go-clean-unit-testing/modules/item/itemHandler"
	"github.com/Rayato159/go-clean-unit-testing/modules/item/itemRepository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	testCreateItemSuccess struct {
		Input    *item.Item
		Expected int
	}

	testCreateItemErr struct {
		Input    *item.Item
		Expected int
	}
)

func TestCreateItemSuccess(t *testing.T) {
	tests := []testCreateItemSuccess{
		{
			Input: &item.Item{
				Title: "Mock Item",
			},
			Expected: 201,
		},
	}

	for _, test := range tests {
		c := NewEchoContext(http.MethodPost, "/item", test.Input)

		mockItemUsecase := new(itemRepository.MockItemRepository)
		mockItemUsecase.On("InsertOneItem", context.Background(), mock.AnythingOfType("*item.Item")).Return(primitive.NewObjectID(), nil)

		itemHandler := itemHandler.NewItemHandler(mockItemUsecase)

		_ = itemHandler.CreateItem(c)
		assert.Equal(t, test.Expected, c.Response().Status)
	}
}

func TestCreateItemError(t *testing.T) {
	tests := []testCreateItemErr{
		{
			Input: &item.Item{
				Title: "",
			},
			Expected: 400,
		},
	}

	for _, test := range tests {
		c := NewEchoContext(http.MethodPost, "/item", test.Input)

		mockItemUsecase := new(itemRepository.MockItemRepository)
		mockItemUsecase.On("InsertOneItem", context.Background(), mock.AnythingOfType("*item.Item")).Return(primitive.NewObjectID(), nil)

		itemHandler := itemHandler.NewItemHandler(mockItemUsecase)

		_ = itemHandler.CreateItem(c)
		assert.Equal(t, test.Expected, c.Response().Status)
	}
}
