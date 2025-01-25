package model

import (
	"AgentApiGo/logger"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/streadway/amqp"
)

type IRabbit interface {
	TestConnection() (bool, error)
	HasEmptyParams() bool
	SendMessage(message interface{}, queue string, con *amqp.Connection) bool
	GetStringConnection() string
	Connection() (*amqp.Connection, error)
	Consumer(queue string, con *amqp.Connection)
}

type Rabbit struct {
	User     string
	Password string
	Host     string
	Port     uint32
}

func (r Rabbit) HasEmptyParams() bool {
	return r.User == "" || r.Password == "" || r.Host == "" || r.Port == 0
}

func (r Rabbit) Connection() (*amqp.Connection, error) {
	connectionString := r.GetStringConnection()
	con, err := amqp.Dial(connectionString)
	if err != nil {
		log.Printf("Failed to connect to RabbitMQ: %s", err)
		return nil, err
	}
	log.Println("Successfully connected to RabbitMQ")
	return con, nil
}

func (r Rabbit) TestConnection() (bool, error) {
	_, err := amqp.Dial(r.GetStringConnection())
	if err != nil {
		return false, fmt.Errorf("erro na conexão: %v", err)
	}
	return true, nil
}

func (r Rabbit) SendMessage(message interface{}, queue string, con *amqp.Connection) bool {
	jobJSON, err := json.Marshal(message)
	if err != nil {
		log.Fatalf("Error in converting Job to JSON: %s", err)
	}
	defer con.Close()

	ch, err := con.Channel()
	if err != nil {
		log.Fatalf("Channel error: %s", err)
	}
	defer ch.Close()

	logger.Log.Info("Declare queue: " + queue)
	_, err = ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %s", err)
	}

	err = ch.Publish(
		"",
		queue,
		true,
		false,
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         jobJSON,
			DeliveryMode: amqp.Persistent,
		},
	)
	if err != nil {
		log.Fatalf("Erro ao enviar mensagem: %s", err)
	}
	logger.Log.Info("Send message with successfull!")
	logger.Log.Info("Message:")
	logger.Log.Info(message)

	return true
}

func (r Rabbit) Consumer(queue string, con *amqp.Connection) {
	defer con.Close()

	ch, err := con.Channel()
	if err != nil {
		log.Fatalf("Channel error: %s", err)
	}
	defer ch.Close()

	logger.Log.Info("Declare queue: " + queue)
	_, err = ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %s", err)
	}

	msgs, err := ch.Consume(
		queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("Erro ao enviar mensagem: %s", err)
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("message received: %v", string(d.Body))
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-sigchan

	log.Printf("interrupted, shutting down")
	forever <- true
}

func (r Rabbit) GetStringConnection() string {
	return "amqp://" + r.User + ":" + r.Password + "@" + r.Host + ":" + strconv.Itoa(int(r.Port)) + "/"
}
