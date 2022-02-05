package configs

type Config struct {
	LocalPath   string
	PostgresDsn string
}

func NewConfig() (config Config, err error) {
	config = Config{
		LocalPath: "/home/Desktop/development/Go/questionnaire_backend",
		// postgres
		PostgresDsn: "host=localhost port=5432 user=skinny password=12345 dbname=questionnaire_backend",
	}
	return
}
