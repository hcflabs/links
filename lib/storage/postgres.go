package storage

import (
	"github.com/hcflabs/links/lib/models"
)

type PostgresLinksBackend struct {
	Host     string
	User     string
	Password string
}

func (s PostgresLinksBackend) CreateOrUpdateLink(entry models.Link) {
	panic("Not implemented")
}

func (s PostgresLinksBackend) GetTargetLink(url string) (target *string, permanent bool) {
	panic("Not implemented")
}

func (s PostgresLinksBackend) GetOwnersLinks(owner string) (links []models.Link) {
	panic("Not implemented")

}

func (s PostgresLinksBackend) GetLinkMetadata(shortUrl string) (link *models.Link) {
	panic("Not implemented")

}

func (s PostgresLinksBackend) DeleteLink(shortUrl string) {
	panic("Not implemented")

}

func (PostgresLinksBackend) GetAllLinksPaginated(offset int, pagesize ...int) (links []models.Link) {
	panic("unimplemented")
}

func (PostgresLinksBackend) GetOwnersLinksPaginated(owner string, offset int, pagesize ...int) (links []models.Link) {
	panic("unimplemented")
}
