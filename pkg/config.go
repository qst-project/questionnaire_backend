package pkg

type Config struct {
	PostgresDsn string
	GrpcTcpPort string
}

func NewConfig() (config Config, err error) {
	config = Config{
		PostgresDsn: "host=postgres port=5432 user=qst password=qst_pwd dbname=qst_db connect_timeout=50",
		GrpcTcpPort: ":9091",
	}
	return
}
