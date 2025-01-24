#AgentApiGo
AgentApiGo is a consumer application built in Go that interacts with a main API to process job messages from RabbitMQ. The main API will register jobs and send them to a RabbitMQ queue, and AgentApiGo will listen to these messages, validating whether the execution should occur on the specified server.

##Features
Job Consumption: Consumes messages from a RabbitMQ queue.
Job Validation: Validates if the execution should happen on the designated server.
Scalable: Can scale by adding more agents or queues to RabbitMQ.
Reliable Messaging: Utilizes RabbitMQ's persistent messaging to ensure message delivery.
##Architecture
Main API: Registers jobs and sends them to RabbitMQ.
Agent API (AgentApiGo): Consumes the messages from the RabbitMQ queue, validates if the job is intended for the current server, and performs the necessary actions.
#Installation
###Prerequisites
Go 1.18 or later
RabbitMQ instance running (you can use Docker to run RabbitMQ locally)