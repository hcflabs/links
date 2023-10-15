package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hcflabs/links/lib/storage"
)

type ApiController struct {
	Backend storage.LinksBackend
}

func (controller ApiController) GetRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	targetUrl, permanent := controller.Backend.GetTargetLink(shortUrl)

	if targetUrl != nil {
		fmt.Printf("%s --> %s\n", shortUrl, *targetUrl)
		if permanent {
			c.Redirect(http.StatusPermanentRedirect, *targetUrl)
		} else {
			c.Redirect(http.StatusTemporaryRedirect, *targetUrl)
		}
	}
	c.Status(http.StatusNotFound)

}

func (controller ApiController) CreateOrUpdateLink(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "https://www.youtube.com/watch?v=HKK4KmDlj8U")
}

func (controller ApiController) GetLinkMetadata(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "https://www.youtube.com/watch?v=HKK4KmDlj8U")
}

func (controller ApiController) DeleteLink(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "https://www.youtube.com/watch?v=HKK4KmDlj8U")
}

func (controller ApiController) GetLinksPaginated(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "https://www.youtube.com/watch?v=HKK4KmDlj8U")
}
