package util

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishEvent(event string, data interface{}) {
	//TODO: connect to rabbitmq
	var connectionUrl = os.Getenv("RABBIT_MQ_URI")
	conn, err := amqp.Dial(connectionUrl)
	failOnError(err, "Failed to connect to RabbitMQ")
	//TODO: create channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	//TODO: create exchange
	//TODO: create queue
	q, err := ch.QueueDeclare(
		event, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")
	//TODO: bind queue to exchange
	//TODO: publish message to queue
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, _ := json.Marshal(data)
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
	//TODO: close connection
	defer conn.Close()
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
