package storage

import (
	"hcflabs/links/models"
)

type LinksBackend interface {
	createOrUpdateLink(link models.Link)
	getTargetLink(url string)(target string, permanent bool)
	getOwnersLinks(owner string)(links []models.Link)
	getLinkMetadata()(link models.Link)
	deleteLink()
}
