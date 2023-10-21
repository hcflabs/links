package storage

import (
	"github.com/hcflabs/links/lib/models"
)

type InMemoryLinksBackend struct {
	LinkMap map[string]models.Link
}


func (s InMemoryLinksBackend) Start() {
	// No Op
}

func (s InMemoryLinksBackend) CreateOrUpdateLink(entry models.Link) {
	s.LinkMap[entry.ShortUrl] = entry
}

func (s InMemoryLinksBackend) GetTargetLink(url string) (target *string, permanent bool) {
	if val, ok := s.LinkMap[url]; ok {
		return &val.TargetUrl, val.LinkOptions.Permanent
	}

	return nil, false
}

func (s InMemoryLinksBackend) GetOwnersLinks(owner string) (links []models.Link) {
	panic("unimplemented")
}

func (s InMemoryLinksBackend) GetLinkMetadata(shortUrl string) (link *models.Link) {
	if val, ok := s.LinkMap[shortUrl]; ok {
		return &val
	}
	return nil
}

func (s InMemoryLinksBackend) DeleteLink(shortUrl string) {
	delete(s.LinkMap, shortUrl)
}

// getAllLinksPaginated implements LinksBackend.
func (s InMemoryLinksBackend) GetAllLinksPaginated(offset int, pagesize int) (links []models.Link) {
	panic("unimplemented")
}

// getOwnersLinksPaginated implements LinksBackend.
func (InMemoryLinksBackend) GetOwnersLinksPaginated(owner string, offset int, pagesize int) (links []models.Link) {
	panic("unimplemented")
}
