package controllers

import (
	"fmt"
	// "log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hcflabs/links/lib"
	"github.com/hcflabs/links/lib/generated"
	"github.com/hcflabs/links/lib/storage"
	log "github.com/sirupsen/logrus"
)

type ApiV1Controller struct {
	Backend storage.LinksBackend
}

func init() {
	log.SetLevel(log.InfoLevel) // NEW
}

func (controller ApiV1Controller) GoToLink(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	targetUrl, permanent, err := controller.Backend.GetTargetLink(shortUrl)

	if err != nil {
		log.Error("couldn't get links", err)
		c.Status(http.StatusInternalServerError)
	}

	if targetUrl != nil {
		log.Info(fmt.Sprintf("%s --> %s\n", shortUrl, *targetUrl))
		if permanent {
			c.Redirect(http.StatusPermanentRedirect, *targetUrl)
		} else {
			c.Redirect(http.StatusFound, *targetUrl)
		}
	} else {
		// c.Status(http.StatusNotFound)
		c.Redirect(http.StatusFound, fmt.Sprintf("/admin/new/%s", shortUrl))
	}
}

func (controller ApiV1Controller) CreateOrUpdateLink(c *gin.Context) {
	// TODO Handle force updates, verify etc
	newlink := fromContext(c)

	if lib.FORBIDDEN_LINKS.Has(newlink.ShortUrl) {
		c.String(http.StatusConflict, fmt.Sprintf("%s is a protected route", newlink.ShortUrl))
	} else {
		controller.Backend.CreateOrUpdateLink(newlink)
	}
}

func (controller ApiV1Controller) GetLinkMetadata(c *gin.Context) {
	req_url := c.Param("shortUrl")
	linkMetadata, err := controller.Backend.GetLinkMetadata(req_url)
	if err != nil {
		log.Error(fmt.Sprintf("Error when getting %s", req_url), err)
		c.Status(http.StatusInternalServerError)
	}

	if linkMetadata != nil {
		c.JSON(http.StatusCreated, *linkMetadata)
	} else {
		c.Status(http.StatusNotFound)
	}
}

func (controller ApiV1Controller) DeleteLink(c *gin.Context) {
	controller.Backend.DeleteLink(c.Param("shortUrl"))
	c.Status(http.StatusOK)
}

func (controller ApiV1Controller) GetLinksPaginated(c *gin.Context) {
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
	pages, err := controller.Backend.GetAllLinksPaginated(offset, pagesize)

	if err != nil {
		log.Error(fmt.Sprintf("Trouble getting paginated results for %d", offset), err)
	}

	c.JSON(http.StatusOK, pages)
}

func (controller ApiV1Controller) GetOwnerLinksPaginated(c *gin.Context) {

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
	links, err := controller.Backend.GetOwnersLinksPaginated(owner, offset, pagesize)
	if err != nil {
		log.Error("trouble getting", err)
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, links)
}

func verify() bool {
	return true
}

func fromContext(c *gin.Context) (link *generated.Link) {
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
