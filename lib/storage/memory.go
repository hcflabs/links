package storage

import (
	"github.com/hcflabs/links/lib/models"
)

type InMemoryLinksBackend struct {
	LinkMap map[string]models.Link
}

func (r InMemoryLinksBackend) CreateOrUpdateLink(entry models.Link) {
	r.LinkMap[entry.ShortUrl] = entry
}

func (r InMemoryLinksBackend) GetTargetLink(url string) (target *string, permanent bool) {
	if val, ok := r.LinkMap[url]; ok {
		return &val.TargetUrl, val.Permanent
	}

	return nil, false
}

func (r InMemoryLinksBackend) GetOwnersLinks(owner string) (links []models.Link) {

	for _, v := range r.LinkMap {
		if v.Owner == owner {
			links = append(links, v)
		}
	}
	return
}

func (r InMemoryLinksBackend) GetLinkMetadata(shortUrl string) (link *models.Link) {
	if val, ok := r.LinkMap[shortUrl]; ok {
		return &val
	}
	return nil
}

func (r InMemoryLinksBackend) DeleteLink(shortUrl string) {
	delete(r.LinkMap, shortUrl)
}

// getAllLinksPaginated implements LinksBackend.
func (InMemoryLinksBackend) GetAllLinksPaginated(offset int, pagesize ...int) (links []models.Link) {
	panic("unimplemented")
}

// getOwnersLinksPaginated implements LinksBackend.
func (InMemoryLinksBackend) GetOwnersLinksPaginated(owner string, offset int, pagesize ...int) (links []models.Link) {
	panic("unimplemented")
}
