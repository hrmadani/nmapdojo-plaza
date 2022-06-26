package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hrmadani/nmapdojo-plaza/pkg/models"
	"github.com/streadway/amqp"
)

const (
	contentType              = "application/json"
	amqpDSN                  = "amqp://guest:guest@localhost:5672/"
	RabbitMQQueueName        = "user_report"
	RabbitMQDurable          = false
	RabbitMQDeleteWhenUnused = false
	RabbitMQExecutive        = false
	RabbitMQNoWait           = false
	RabbitMQExchange         = ""
	RabbitMQMandatory        = false
	RabbitMQImmediate        = false
)

//FailOnError the Plaza error handler
func FailOnError(err error) {
	if err != nil {
		log.Fatalf("[x] Error Handleing : %s", err)
	}
}

//RabbitMQConnection is the connection to the RabbitMQ server
func RabbitMQMessaging(message string) {
	connection, err := amqp.Dial(amqpDSN)
	FailOnError(err)
	defer connection.Close()

	channel, err := connection.Channel()
	FailOnError(err)
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		RabbitMQQueueName,
		RabbitMQDurable,
		RabbitMQDeleteWhenUnused,
		RabbitMQExecutive,
		RabbitMQNoWait,
		nil, //argument
	)
	FailOnError(err)

	err = channel.Publish(
		RabbitMQExchange,
		queue.Name,
		RabbitMQMandatory,
		RabbitMQImmediate,
		amqp.Publishing{
			ContentType: contentType,
			Body:        []byte(message),
		})
	FailOnError(err)
}

//Handle the response is one of the most repetitive tasks
//So define a function to DRY
func ResponseHandler(result []byte, contentType string, httpStatusCode string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", contentType)
	switch httpStatusCode {
	case "OK":
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusBadGateway)
	}
	w.Write(result)
}

//Show all alive reports to user
func ShowAllAliveReports(response http.ResponseWriter, request *http.Request) {
	//Get all alive reports from database
	reports, _ := models.GetAllAliveReport()
	result, _ := json.Marshal(reports)

	ResponseHandler(result, contentType, "OK", response)
}

//Get Report From User
func GetReportFromUser(response http.ResponseWriter, request *http.Request) {
	//Parse the user's report
	log, err := ioutil.ReadAll(request.Body)
	FailOnError(err)

	var reportLog models.UserReport
	json.Unmarshal([]byte(log), &reportLog)

	reportLogJson, _ := json.Marshal(reportLog.SetNowTime())

	RabbitMQMessaging(string(reportLogJson))
}
