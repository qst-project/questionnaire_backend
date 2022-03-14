package configs

type Config struct {
	PostgresDsn string
	GrpcTcpPort string
}

func NewConfig() (config Config, err error) {
	config = Config{
		PostgresDsn: "host=localhost port=5432 user=questionnaire password=test dbname=questionnaire_test",
		GrpcTcpPort: ":9090",
	}
	return
}
