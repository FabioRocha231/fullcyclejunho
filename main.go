package main

import (
	"database/sql"
	"github.com/FabioRocha231/fullcyclejunho/internal/infra/database"
	"github.com/FabioRocha231/fullcyclejunho/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")

	if err != nil {
		panic(err)
	}

	orderRepository := database.NewOrderRepository(db)

	newCalculateFinalPrice := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput{
		ID:    "123",
		Price: 10,
		Tax:   5,
	}

	orderOutput, err := newCalculateFinalPrice.Execute(input)
	if err != nil {
		panic(err)
	}

	log.Println(orderOutput)
}
