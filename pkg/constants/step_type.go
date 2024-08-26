package constants

type FeatureName string

const (
	Postgres FeatureName = "postgres"
	MongoDB  FeatureName = "mongodb"
	Redis    FeatureName = "redis"
	HTTP     FeatureName = "http"
	JS       FeatureName = "js"
	Email    FeatureName = "email"
	Queue    FeatureName = "queue"
	API      FeatureName = "api"
	Flow     FeatureName = "flow"
)
