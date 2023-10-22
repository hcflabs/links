package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hcflabs/links/lib"
	"github.com/hcflabs/links/lib/models"
	"github.com/hcflabs/links/lib/storage"
	"github.com/sirupsen/logrus"
)

type ApiController struct {
	Backend storage.LinksBackend
}

var log = logrus.New()


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
	newlink := fromContext(c)

	if lib.FORBIDDEN_LINKS.Has(newlink.ShortUrl) {
		c.String(http.StatusConflict, fmt.Sprintf("%s is a protected route", newlink.ShortUrl))
	} else {
		controller.Backend.CreateOrUpdateLink(newlink)
	}
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

	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		// ... handle error
		offset = 0
	}
	c.JSON(http.StatusOK, controller.Backend.GetAllLinksPaginated(offset, pagesize))
}

func (controller ApiController) GetOwnerLinksPaginated(c *gin.Context) {
	
	pagesize, err := strconv.Atoi(c.Query("pagesize"))
	if err != nil {
		// ... handle error
		pagesize = 10
	}

	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		// ... handle error
		offset = 0
	}

	owner := c.Param("owner")
	c.JSON(http.StatusOK, controller.Backend.GetOwnersLinksPaginated(owner, offset, pagesize))
}

func verify() bool {
	return true
}

func fromContext(c *gin.Context) (link models.InternalLink) {
	if err := c.BindJSON(&link); err != nil {

		// TODO Handle?
		panic("bad bind!")
	}

	if c.ShouldBindJSON(&link) == nil {
		log.Println(link.ShortUrl)
	} else {
		panic("bad bind!")
	}
	return
}
