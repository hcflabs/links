package main

import (
	"fmt"
	"net/http"
	"os"

	// "net/http"
	// "github.com/gin-gonic/contrib/static"
	// "github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/hcflabs/links/lib/controllers"
	"github.com/hcflabs/links/lib/models"
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
		config = storage.PostgresConfig{
			Host:     os.Getenv("LINKS_DB_HOST"),
			User:     os.Getenv("LINKS_DB_USER"),
			Password: os.Getenv("LINKS_DB_PASSWORD"),
			Database: os.Getenv("LINKS_DB_PASSWORD"),
			Port: os.Getenv("LINKS_DB_PASSWORD"),
		}
		backend = storage{}
	default:
		fmt.Printf("InMemory Backend loading")
		backend = storage.InMemoryLinksBackend{
			LinkMap: make(map[string]models.Link),
		}
	}

	cfg = ServerConfig{
		Port:           os.Getenv("LINKS_PORT"),
		AdminBuildPath: os.Getenv("LINKS_ADMIN_PATH"),
	}
	return
}

func initLinks(backend storage.LinksBackend) {
	// TODO Remove after Testing
	util.InsertLink(backend, "holdon", "https://www.youtube.com/watch?v=HKK4KmDlj8U")
	util.InsertLink(backend, "great", "https://www.youtube.com/watch?v=kSVQtlQtxCs")
	util.InsertLink(backend, "hcf", "https://haltcatchfire.io")
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
	fmt.Printf("Loaded Config \n%+v\n", cfg)

	fmt.Printf("Loading Test Links")
	initLinks(backend)

	// Set up Server
	router := gin.Default()

	// Init
	api := controllers.ApiController{Backend: backend}
	// Serve frontend static files
	// router.Use(static.Serve("/admin", (fmt.Sprintf("%s", cfg.AdminBuildPath),  http.Dir()))
	router.Static("/admin", fmt.Sprintf("%s", cfg.AdminBuildPath))
	router.GET("/admin", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/admin/index.html")
	})
	// Primary Route
	router.GET("/:shortUrl", api.GetRedirect)
	// Admin API
	router.GET("/api/links/:shortUrl", api.GetLinkMetadata)
	router.POST("/api/links/:shortUrl", api.CreateOrUpdateLink)
	router.DELETE("/api/links/:shortUrl", api.DeleteLink)
	router.GET("/api/links", api.GetLinksPaginated)
	router.GET("/api/owners/:owner/links", api.GetOwnerLinksPaginated)

	router.Run(fmt.Sprintf(":%s", cfg.Port))
}
