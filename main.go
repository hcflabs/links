package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	// c "./lib/config"
	// "github.com/spf13/viper"
	// "database/sql"
	// _ "github.com/lib/pq"
	// "github.com/golang-migrate/migrate/v4"
	// "github.com/golang-migrate/migrate/v4/database/postgres"
	// _ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gin-gonic/contrib/static"
)

type DatabaseConfig struct {
	host     string
	user     string
	password string
}

type ServerConfig struct {
	port     string
	dbConfig DatabaseConfig
}

func load_config() (cfg ServerConfig) {
	if os.Getenv("LINKS_BACKEND") == "postgres" {
	
	}
	var db_config DatabaseConfig

	switch backend := os.Getenv("LINKS_BACKEND"); backend {
	case "postgres" :
	
		db_config = DatabaseConfig{
			host:     os.Getenv("LINKS_DB_HOST"),
			user:     os.Getenv("LINKS_DB_USER"),
			password: os.Getenv("LINKS_DB_PASSWORD"),
		}
		fmt.Println("postgres config: %s@%s", db_config.user, db_config.host)
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", backend)
	}

	cfg = ServerConfig{
		port:     os.Getenv("LINKS_PORT"),
		dbConfig: db_config,
	}
	return
}

func main() {

	// // Migrate Database if Needed
	// db, err := sql.Open("postgres", "postgres://localhost:5432/database?sslmode=enable")
	// driver, err := postgres.WithInstance(db, &postgres.Config{})
	// m, err := migrate.NewWithDatabaseInstance(
	//     "file:///migrations",
	//     "postgres", driver)
	// m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run

	cfg := load_config()

	fmt.Printf("Loaded Config:")
	fmt.Printf("%+v\n", cfg)

	// Set up Server
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/admin", static.LocalFile("./client/build", true)))

	router.GET("/holdon", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "https://www.youtube.com/watch?v=HKK4KmDlj8U")
	})

	router.Run(fmt.Sprintf(":%s", cfg.port))
}
