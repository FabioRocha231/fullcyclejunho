package main

import (
	"database/sql"
	"encoding/json"
	"github.com/FabioRocha231/fullcyclejunho/internal/infra/database"
	"github.com/FabioRocha231/fullcyclejunho/internal/usecase"
	"github.com/FabioRocha231/fullcyclejunho/pkg/rabbitmq"
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

	ch, err := rabbitmq.OpenChanel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()

	msgRabbitmqChannel := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgRabbitmqChannel)

	rabbitmqWorker(msgRabbitmqChannel, newCalculateFinalPrice)
}

func rabbitmqWorker(msgChan chan amqp.Delivery, uc *usecase.CalculateFinalPrice) {
	log.Println("Starting rabbitmq")

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
