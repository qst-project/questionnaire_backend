package gateway

import (
	"github.com/qst-project/backend.git/pkg"
	"github.com/qst-project/backend.git/pkg/core"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type PostgresClient struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewPostgresClient(config pkg.Config, _logger *log.Logger) (postgresClient PostgresClient, err error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(config.PostgresDsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		return
	}
	postgresClient = PostgresClient{
		db:     db,
		logger: _logger,
	}
	err = postgresClient.Migrate()
	return postgresClient, nil

}

func (c *PostgresClient) Migrate() (err error) {
	err = c.db.AutoMigrate(
		// &core.User{},
		&core.QuestionDB{},
		&core.QuestionnaireDB{},
		&core.RadioPossibleAnswerDB{},
		&core.TextPossibleAnswerDB{},
		&core.CheckboxPossibleAnswerDB{},
		&core.CheckboxAnswerDB{},
		&core.RadioAnswerDB{},
		&core.TextAnswerDB{},
	)
	return
}
