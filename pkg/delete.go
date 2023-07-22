package pkg

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/cobra"
)

func deleteQueueCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "A brief description of your command",
		Run: func(cmd *cobra.Command, args []string) {
			rabbitMQURL := "amqp://admin:admin@localhost:5673/"
			queueName := "my_queue"
			exchangeName := "my_exchange"
			routeKey := "interna.create"

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
			_, err = ch.QueueDelete(
				queueName, // Nome da Queue
				false,     // Se tiver consumidores
				false,     // Se tiver mensagens
				false,     // Sem-Wait
			)
			if err != nil {
				fmt.Println(err)
				return
			}

			err = ch.ExchangeDelete(
				exchangeName, // Nome da Exchange
				false,        // Se tiver consumidores
				false,        // Sem-Wait
			)

			if err != nil {
				fmt.Println(err)
				return
			}

			err = ch.QueueUnbind(
				queueName,    // Nome da Queue
				routeKey,     // Route Key
				exchangeName, // Nome da Exchange
				nil,
			)
			if err != nil {
				fmt.Println(err)
				return
			}

		},
	}
	return cmd
}
