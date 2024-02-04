package storage

import "github.com/hcflabs/links/lib/generated"

type LinksBackend interface {
	Start()
	CreateOrUpdateLink(*generated.Link) error
	GetTargetLink(url string) (*string, bool, error)
	GetOwnersLinks(owner string) (*[]generated.Link, error)
	GetOwnersLinksPaginated(owner string, offset int, pagesize int) (*[]generated.Link, error)
	GetAllLinksPaginated(offset int, pagesize int) (*[]generated.Link, error)
	GetLinkMetadata(url string) (*generated.Link, error)
	DeleteLink(url string) error
	//TODO: Implement Analytics
	// GetMostRecentLinks(offset int, pagesize int)(links []models.InternalLink)
	// GetMostRecentModifiedLinks(offset int, pagesize int)(links []models.InternalLink)
	// GetMostPopularLinks(offset int, pagesize int)(links []models.InternalLink)
}
