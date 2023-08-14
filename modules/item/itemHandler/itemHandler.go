package itemHandler

import (
	"context"
	"net/http"

	"github.com/Rayato159/go-clean-unit-testing/modules/item"
	"github.com/Rayato159/go-clean-unit-testing/modules/item/itemRepository"
	"github.com/Rayato159/go-clean-unit-testing/modules/request"
	"github.com/labstack/echo/v4"
)

type (
	ItemHandlerService interface {
		CreateItem(c echo.Context) error
	}

	itemHandler struct {
		itemRepository itemRepository.ItemRepositoryService
	}
)

func NewItemHandler(itemRepository itemRepository.ItemRepositoryService) ItemHandlerService {
	return &itemHandler{itemRepository: itemRepository}
}

func (h *itemHandler) CreateItem(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	reqBody := new(item.Item)

	if err := wrapper.Bind(reqBody); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{"message": err.Error()},
		)
	}

	itemId, err := h.itemRepository.InsertOneItem(ctx, reqBody)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{"message": err.Error()},
		)
	}
	reqBody.Id = itemId.Hex()

	return c.JSON(
		http.StatusCreated,
		reqBody,
	)
}
