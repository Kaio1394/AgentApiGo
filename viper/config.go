package viper

type Config struct {
	App App
	Server
	DataBase
	RabbitConfig
}
type App struct {
	Name string
}

type Server struct {
	Host string
	Port string
}

type DataBase struct {
	TypeDatabase     string
	StringConnection string
}

type RabbitConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	QueueDeploy  string
	QueueHistory string
}
