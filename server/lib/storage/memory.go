package storage

import (
	"errors"

	"github.com/bobg/go-generics/v3/maps"
	"github.com/hcflabs/links/lib/generated"
	//  "golang.org/x/exp/map"
	// log "github.com/sirupsen/logrus"
)

type InMemoryLinksBackend struct {
	LinkMap map[string]generated.Link
}

func (s InMemoryLinksBackend) Start() {
	// No Op
}

func (s InMemoryLinksBackend) CreateOrUpdateLink(entry *generated.Link) error {
	s.LinkMap[entry.ShortUrl] = *entry
	return nil
}

func (s InMemoryLinksBackend) GetTargetLink(url string) (target *string, permanent bool, err error) {
	if val, ok := s.LinkMap[url]; ok {
		// return &val.TargetUrl, val.LinkOptions.Permanent
		return &val.TargetUrl, false, nil
	}

	return nil, false, errors.New("Did not find target")
}

func (s InMemoryLinksBackend) GetOwnersLinks(owner string) (links *[]generated.Link, err error) {
	panic("unimplemented")
}

func (s InMemoryLinksBackend) GetLinkMetadata(shortUrl string) (link *generated.Link, err error) {
	if val, ok := s.LinkMap[shortUrl]; ok {
		return &val, nil
	}
	return nil, errors.New("Did not find target")
}

func (s InMemoryLinksBackend) DeleteLink(shortUrl string) error {
	delete(s.LinkMap, shortUrl)
	return nil
}

// getAllLinksPaginated implements LinksBackend.
func (s InMemoryLinksBackend) GetAllLinksPaginated(offset int, pagesize int) (*[]generated.Link, error) {
	l := maps.Values(s.LinkMap)
	return &l, nil
}

// getOwnersLinksPaginated implements LinksBackend.
func (InMemoryLinksBackend) GetOwnersLinksPaginated(owner string, offset int, pagesize int) (links *[]generated.Link, err error) {
	panic("unimplemented")
}
