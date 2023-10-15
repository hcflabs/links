package storage

import (
	"github.com/hcflabs/links/lib/models"
)

type LinksBackend interface {
	CreateOrUpdateLink(link models.Link)
	GetTargetLink(url string)(target *string, permanent bool)
	GetOwnersLinks(owner string)(links []models.Link)
	GetOwnersLinksPaginated(owner string, offset int, pagesize ...int)(links []models.Link)
	GetAllLinksPaginated(offset int, pagesize ...int)(links []models.Link)
	GetLinkMetadata(url string)(link *models.Link)
	DeleteLink(url string)
}
