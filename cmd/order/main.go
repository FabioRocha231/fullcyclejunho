package main

import (
	"database/sql"
	"encoding/json"
	"github.com/FabioRocha231/fullcyclejunho/internal/infra/database"
	"github.com/FabioRocha231/fullcyclejunho/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
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

func rabbitmqWorker(msgChan chan amqp.Delivery, uc usecase.CalculateFinalPrice) {
	log.Println("Saving rabbitmq")

	for msg := range msgChan {
		var input usecase.OrderInput
		err := json.Unmarshal(msg.Body, &input)
		if err != nil {
			panic(err)
		}

		output, err := uc.Execute(input)
		if err != nil {
			panic(err)
		}

		if err := msg.Ack(false); err != nil {
			panic(err)
		}

		log.Println("Mensagem processada e salva no banco", output)
	}
}
