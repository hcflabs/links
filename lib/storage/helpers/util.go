package helpers

import (
	"github.com/hcflabs/links/lib/storage"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
)

// "github.com/hcflabs/lib/storage"

func buildPostgresBackend(config storage.PostgresConfig) (backend storage.PostgresLinksBackend) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.Host, config.Host, config.Password, config.Database, config.Port)
	databaseConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	backend = storage.PostgresLinksBackend{
		DB: databaseConn,
		// Config: config,
	}
	return
}
