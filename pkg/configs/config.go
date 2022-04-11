package configs

type Config struct {
	PostgresDsn string
	GrpcTcpPort string
	PemPath     string
	KeyPath     string
}

func NewConfig() (config Config, err error) {
	config = Config{
		PostgresDsn: "host=localhost port=5432 user=questionnaire password=12345 dbname=questionnaire_backend",
		GrpcTcpPort: ":9091",
		PemPath:     "./../cert/server.crt",
		KeyPath:     "./../cert/server.key",
	}
	return
}
