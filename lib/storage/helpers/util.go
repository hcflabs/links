package helpers

import (
	"github.com/hcflabs/links/lib/storage"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
)

// "github.com/hcflabs/lib/storage"

func BuildPostgresBackend(config storage.PostgresConfig) (backend storage.PostgresLinksBackend) {
	// fmt.Printf("%+v\n", config)
	dsn := "host=localhost user=postgres password=postgres dbname=hcflinks port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.Host, config.User, config.Password, config.Database, config.Port)
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
