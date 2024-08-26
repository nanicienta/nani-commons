package connections

type DBConnection struct {
	Host     string
	Port     string
	User     string
	Password string
	DB       string
	SSLMode  string
}
