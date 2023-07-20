package main

import (
	"github.com/FabioRocha231/fullcyclejunho/internal/entity"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {

	e := echo.New()

	e.GET("/order", Order)

	e.Logger.Fatal(e.Start(":8080"))

}

func Order(c echo.Context) error {
	order := entity.Order{
		ID:    "1",
		Price: 10.0,
		Tax:   0.1,
	}

	err := order.CalculateFinalPrice()

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, order)
}
