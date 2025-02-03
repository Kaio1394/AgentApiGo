package helper

import (
	"AgentApiGo/logger"
	"AgentApiGo/model"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/streadway/amqp"
)

var queue_history string = "Job.Schedule.History"

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
	helper := Helper{}

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
		for msg := range msgs {
			msgStr := string(msg.Body)
			log.Printf(msgStr)
			var msgJson map[string]string
			err := json.Unmarshal([]byte(msgStr), &msgJson)
			if err != nil {
				logger.Log.Error("Error in parsing JSON: %s", err)
				return
			}

			job := model.Job{
				Name:        msgJson["name"],
				Server:      msgJson["server"],
				Script:      msgJson["script"],
				Date:        msgJson["date"],
				Description: msgJson["description"],
				CmdExecute:  msgJson["cmdExecute"],
			}

			if job.Server == helper.GetIp() && job.CmdExecute == "true" {
				logger.Log.Info("Date Execution: " + job.Date + ". Sysdate: " + Sysdate.Format(Layout_date))
				date, errDate := helper.ConvertDate(job.Date, Layout_date)
				if errDate != nil {
					logger.Log.Error("Error to convert date.")
					return
				}

				diff := date.Sub(Sysdate)
				if diff >= -1 || diff <= 1 {
					script := job.Script
					logger.Log.Info("Executing script: " + script)
					cmd := exec.Command("cmd", "/C", script)
					err := cmd.Run()
					if err != nil {
						logger.Log.Error("Executing error: %v", err)
						return
					}
					logger.Log.Error("Finish execute.")
					r.SendMessage(job, queue_history, con)
				}
			} else {
				logger.Log.Info("Not match Server: " + job.Server)
			}
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
