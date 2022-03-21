package repository

import (
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/configs"
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/models"
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

func NewPostgresClient(config configs.Config, _logger *log.Logger) (postgresClient PostgresClient, err error) {
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
		// &models.User{},
		&models.Question{},
		&models.Questionnaire{},
		&models.RadioPossibleAnswer{},
		&models.TextPossibleAnswer{},
		&models.CheckboxPossibleAnswer{},
		&models.CheckboxAnswer{},
		&models.RadioAnswer{},
		&models.TextAnswer{},
	)
	return
}
