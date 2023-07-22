package pkg

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/cobra"
)

var createQueueCmd = &cobra.Command{
	Use:   "create",
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
		_, err = ch.QueueDeclare(
			queueName, // Nome da Queue
			true,      // Durável
			false,     // Não-Auto-Deletável
			false,     // Exclusiva
			false,     // Sem-Wait
			amqp.Table{
				"x-max-priority": 10,
				"x-message-ttl":  10000,
			}, // Argumentos
		)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = ch.ExchangeDeclare(
			exchangeName, // Nome da Exchange
			"direct",     // Tipo
			true,         // Durável
			false,        // Não-Auto-Deletável
			false,        // Interno
			false,        // Sem-Wait
			nil,          // Argumentos
		)

		if err != nil {
			fmt.Println(err)
			return
		}

		err = ch.QueueBind(
			queueName,    // Nome da Queue
			routeKey,     // Chave de roteamento
			exchangeName, // Nome da Exchange (vazio para usar a default)
			false,        // Sem-Wait
			amqp.Table{}, // Argumentos
		)
		if err != nil {
			fmt.Println(err)
			return
		}

	},
}

func init() {
	RootCmd.AddCommand(createQueueCmd)
}
