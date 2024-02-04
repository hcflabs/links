package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	// "net/http"
	// "github.com/gin-gonic/contrib/static"
	// "github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/hcflabs/links/lib/controllers"
	"github.com/hcflabs/links/lib/generated"
	"github.com/hcflabs/links/lib/middleware"
	"github.com/hcflabs/links/lib/storage"
	"github.com/hcflabs/links/lib/util"
	// "gorm.io/driver/postgres"
)

type ServerConfig struct {
	Port           string
	AdminBuildPath string
}

func loadConfig() (cfg ServerConfig, backend storage.LinksBackend) {
	switch backend_option := os.Getenv("LINKS_BACKEND"); backend_option {
	case "postgres":
		fmt.Printf("Postgres Backend loading")
		config := storage.PostgresConfig{
			Host:     os.Getenv("LINKS_DB_HOST"),
			User:     os.Getenv("LINKS_DB_USER"),
			Password: os.Getenv("LINKS_DB_PASSWORD"),
			Database: os.Getenv("LINKS_DB_DATABASE"),
			Port:     os.Getenv("LINKS_DB_PORT"),
		}

		backend = storage.BuildPostgresBackend(config)
	default:
		fmt.Printf("InMemory Backend loading")
		backend = storage.InMemoryLinksBackend{
			LinkMap: make(map[string]generated.Link),
		}
	}

	cfg = ServerConfig{
		Port:           os.Getenv("LINKS_PORT"),
		AdminBuildPath: os.Getenv("LINKS_ADMIN_PORT"),
	}
	return
}

func initLinks(backend storage.LinksBackend) {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
	log.Info("Migrating Schema")
	backend.Start()
	log.Info("Loading Test Links")
	// TODO Remove after Testing
	util.InsertLink(backend, "holdon", "https://www.youtube.com/watch?v=HKK4KmDlj8U")
	util.InsertLink(backend, "great", "https://www.youtube.com/watch?v=kSVQtlQtxCs")
	util.InsertLink(backend, "hcf", "https://haltcatchfire.io")
}

func getBaseRouter() *gin.Engine {
	// Set up Server
	router := gin.Default()
	router.Use(gin.Recovery())                 // NEW
	router.Use(middleware.LoggingMiddleware()) // NEW
	return router
}

func buildService(api controllers.ApiV1Controller) *gin.Engine {
	router := getBaseRouter()

	router.GET("/:shortUrl", api.GoToLink)

	return router
}

func buildAPIServer(api controllers.ApiV1Controller) *gin.Engine {
	router := getBaseRouter()
	// Admin API

	v1 := router.Group("/api/v1")
	{
		v1.GET("/links/:shortUrl", api.GetLinkMetadata)
		v1.POST("/links/:shortUrl", api.CreateOrUpdateLink)
		v1.DELETE("/links/:shortUrl", api.DeleteLink)
		v1.GET("/links", api.GetLinksPaginated)
		v1.GET("/owners/:owner/links", api.GetOwnerLinksPaginated)
	}
	return router
}

func main() {

	// // Migrate Database if Needed
	// db, err := sql.Open("postgres", "postgres://localhost:5432/database?sslmode=enable")
	// driver, err := postgres.WithInstance(db, &postgres.Config{})
	// m, err := migrate.NewWithDatabaseInstance(
	//     "file:///migrations",
	//     "postgres", driver)
	// m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run

	cfg, backend := loadConfig()
	log.WithFields(log.Fields{
		"config": cfg,
	}).Info("Loading Config")

	initLinks(backend)

	// Init
	api := controllers.ApiV1Controller{Backend: backend}

	service := buildService(api)
	apiServer := buildAPIServer(api)

	go service.Run(fmt.Sprintf(":%s", cfg.Port))
	apiServer.Run(fmt.Sprintf(":%s", cfg.Port))

}
