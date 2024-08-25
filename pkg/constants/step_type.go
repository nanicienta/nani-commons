package constants

type StepType string

const (
	Postgres StepType = "postgres"
	MongoDB  StepType = "mongodb"
	Redis    StepType = "redis"
	HTTP     StepType = "http"
	JS       StepType = "js"
	Email    StepType = "email"
	Queue    StepType = "queue"
	API      StepType = "api"
)
