package configs

type Config struct {
	PostgresDsn string
	GrpcTcpPort string
}

func NewConfig() (config Config, err error) {
	config = Config{
		PostgresDsn: "host=localhost port=5432 user=skinny password=12345 dbname=questionnaire_backend",
		GrpcTcpPort: ":9091",
	}
	return
}
