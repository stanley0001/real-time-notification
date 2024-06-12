package util

import (
	"log"
	"notification-service/services"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ListenForEvents() {
	//TODO: connect to rabbitmq
	var connectionUrl = os.Getenv("RABBIT_MQ_URI")
	conn, err := amqp.Dial(connectionUrl)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	var events = []string{"create-comment", "create-post", "create-reaction", "update-post", "send-message"}
	for _, e := range events {
		go func(event string) {
			q, err := ch.QueueDeclare(event, false, false, false, false, nil)
			failOnError(err, "Failed to declare a queue")
			msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
			failOnError(err, "Failed to register a consumer")
			forever := make(chan bool)
			go func() {
				for d := range msgs {
					log.Printf("Received a message in event %s: %s", event, d.Body)
					//TODO:process recived message
					//Send notification push/email
					services.HandleEvents(event, d.Body)
				}
			}()
			log.Printf(" [*] Initialize event listenert for messages in event %s.", event)
			<-forever
		}(e)
	}
	select {}
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
