package controllers
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

func CreateOrUpdateLink(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "https://www.youtube.com/watch?v=HKK4KmDlj8U")
}

func GetLinkMetadata(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "https://www.youtube.com/watch?v=HKK4KmDlj8U")
}

func GetLinkMetadata(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "https://www.youtube.com/watch?v=HKK4KmDlj8U")
}