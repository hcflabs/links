package storage

import (
	"hcflabs/links/models"
)

type InMemoryLinksBackend struct {
	linkMap map[string]models.Link
}

func (r InMemoryLinksBackend) createOrUpdateLink(entry models.Link) {
	r.linkMap[link.shortUrl] = entry
}

func (r InMemoryLinksBackend) getTargetLink(url string) (target string, permanent bool) {

	return r.linkMap[url].TargetUrl, false
}

func (r InMemoryLinksBackend) getOwnersLinks(owner string) (links []models.Link) {
	result := []models.Link{}

	for _, v := range r.linkMap {
		if v.Owner == owner {
			result = append(result, v)
		}
	}
}

func (r InMemoryLinksBackend) getLinkMetadata(shortUrl string) (link models.Link) {
	return r.linkMap[shortUrl]
}

func (r InMemoryLinksBackend) deleteLink(shortUrl string) {
	delete(r.linkMap, shortUrl)
}
