package storage

import (
	"github.com/hcflabs/links/lib/models"
)

type LinksBackend interface {
	Start()
	CreateOrUpdateLink(link models.InternalLink)
	GetTargetLink(url string) (target *string, permanent bool)
	GetOwnersLinks(owner string) (links []models.InternalLink)
	GetOwnersLinksPaginated(owner string, offset int, pagesize int) (links []models.InternalLink)
	GetAllLinksPaginated(offset int, pagesize int) (links []models.InternalLink)
	GetLinkMetadata(url string) (link *models.InternalLink)
	DeleteLink(url string)
	//TODO: Implement Analytics
	// GetMostRecentLinks(offset int, pagesize int)(links []models.InternalLink)
	// GetMostRecentModifiedLinks(offset int, pagesize int)(links []models.InternalLink)
	// GetMostPopularLinks(offset int, pagesize int)(links []models.InternalLink)
}
