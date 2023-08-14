package server

import (
	"github.com/Rayato159/go-clean-unit-testing/modules/item/itemHandler"
	"github.com/Rayato159/go-clean-unit-testing/modules/item/itemRepository"
	"github.com/labstack/echo/v4"
)

type (
	server struct {
		app *echo.Echo
	}
)

func Start() {
	s := &server{app: echo.New()}

	s.ItemService()

	s.app.Logger.Fatal(s.app.Start(":1323"))
}

func (s *server) ItemService() {
	itemRepository := itemRepository.NewItemRepository()
	itemHandler := itemHandler.NewItemHandler(itemRepository)

	item := s.app.Group("/item")
	_ = item
	_ = itemHandler
}
