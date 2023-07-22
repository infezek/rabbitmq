package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/cobra"
)

func publishQueueCmd() *cobra.Command {
	priory := uint8(0)
	id := 0
	reject := false
	message := "Hello World"
	delay := 0

	cmd := cobra.Command{
		Use:   "publish",
		Short: "A brief description of your command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("id: %d, message: %s, priory: %d, reject: %t, delay: %d\n", id, message, priory, reject, delay)
			rabbitMQURL := "amqp://admin:admin@localhost:5673/"
			//queueName := "my_queue"
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

			type Message struct {
				ID      int    `json:"id"`
				Message string `json:"message"`
				Time    string `json:"time"`
				Reject  bool   `json:"Reject"`
			}
			body, err := json.Marshal(Message{
				ID:      id,
				Message: message,
				Time:    time.Now().Format("2006-01-02 15:04:05"),
				Reject:  reject,
			})
			if err != nil {
				fmt.Println(err)
				return
			}
			err = ch.PublishWithContext(
				context.Background(),
				exchangeName, // Nome da Exchange
				routeKey,     // Route Key
				false,        // Mandatório
				true,         // Imediato
				amqp.Publishing{
					ContentType: "application/json",
					Body:        body,
					Priority:    priory,
				}, // Publishing
			)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Mensagem publicada com sucesso!")
		},
	}
	cmd.Flags().Uint8VarP(&priory, "priory", "p", 0, "Priority")
	cmd.Flags().BoolVarP(&reject, "reject", "r", false, "Reject")
	cmd.Flags().IntVarP(&id, "id", "i", 0, "ID")
	cmd.Flags().IntVarP(&delay, "delay", "d", 0, "Delay")
	cmd.Flags().StringVarP(&message, "messsage", "m", "Hello World", "Message")

	return &cmd
}
