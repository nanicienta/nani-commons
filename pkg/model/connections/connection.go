package connections

type IConnection interface {
	GetConnectionType() string
}

type BaseConnection struct {
	Id       string `json:"id"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
	Ssl      bool   `json:"ssl"`
}
