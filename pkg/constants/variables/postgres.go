package variables

type PostgresConstants string

const (
	PostgresHost     PostgresConstants = "host"
	PostgresPort     PostgresConstants = "port"
	PostgresUser     PostgresConstants = "user"
	PostgresPassword PostgresConstants = "password"
	PostgresDB       PostgresConstants = "db"
	PostgresSSLMode  PostgresConstants = "sslmode"
)

var postgresConstantsArray = []PostgresConstants{
	PostgresHost,
	PostgresPort,
	PostgresUser,
	PostgresPassword,
	PostgresDB,
	PostgresSSLMode,
}

func isAPostgresConstant(constant string) bool {
	for _, c := range postgresConstantsArray {
		if string(c) == constant {
			return true
		}
	}
	return false
}
