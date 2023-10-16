package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hcflabs/links/lib/models"
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
	} else {
		c.Status(http.StatusNotFound)
	}
}

func (controller ApiController) CreateOrUpdateLink(c *gin.Context) {
	// TODO Handle force updates, verify etc
	controller.Backend.CreateOrUpdateLink(fromContext(c))
}

func (controller ApiController) GetLinkMetadata(c *gin.Context) {
	linkMetadata := controller.Backend.GetLinkMetadata(c.Param("shortUrl"))
	if linkMetadata != nil {
		c.JSON(http.StatusCreated, *linkMetadata)
	} else {
		c.Status(http.StatusNotFound)
	}
}

func (controller ApiController) DeleteLink(c *gin.Context) {
	controller.Backend.DeleteLink(c.Param("shortUrl"))
	c.Status(http.StatusOK)
}

func (controller ApiController) GetLinksPaginated(c *gin.Context) {
	pagesize, err := strconv.Atoi(c.Query("pagesize"))
	if err != nil {
		// ... handle error
		pagesize = 10
	}

	offset, _ := strconv.Atoi(c.Query("offset"))
	c.JSON(http.StatusOK, controller.Backend.GetAllLinksPaginated(offset, pagesize))
}

func verify() bool {
	return true
}

func fromContext(c *gin.Context) (link models.Link) {
	if err := c.BindJSON(&link); err != nil {

		// TODO Handle?
	}
	return
}
