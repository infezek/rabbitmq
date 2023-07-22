package pkg

import (
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/cobra"
)

var consumerQueueCmd = &cobra.Command{
	Use:   "consumer",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		rabbitMQURL := "amqp://admin:admin@localhost:5673/"
		//queueName := "my_queue"
		//exchangeName := "my_exchange"
		//routeKey := "interna.create"

		conn, err := amqp.Dial(rabbitMQURL)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()

		ch, err := conn.Channel()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer ch.Close()

		type Message struct {
			ID      int    `json:"id"`
			Message string `json:"message"`
			Time    string `json:"time"`
			Reject  bool   `json:"Reject"`
		}

		if err != nil {
			fmt.Println(err)
			return
		}
		msg, err := ch.Consume(
			"my_queue", // Nome da fila
			"",         // Consumer
			false,      // Auto Ack
			false,      // Exclusive
			false,      // No Local
			false,      // No Wait
			nil,        // Args
		)
		if err != nil {
			fmt.Println(err)
			return
		}
		for m := range msg {
			var message Message
			err := json.Unmarshal(m.Body, &message)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(message)
			m.Ack(false)
		}

		fmt.Println("Mensagem publicada com sucesso!")

	},
}
