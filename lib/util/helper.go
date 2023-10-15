package util

import (
	"github.com/hcflabs/links/lib/models"
	"github.com/hcflabs/links/lib/storage"
)

func BuildLink(shortURL string, target string) (link models.Link) {
	link = models.Link{
		ShortUrl:    shortURL,
		TargetUrl:   target,
		Owner:       "fake_user@domain.com",
		Description: "this is a fake thing",
	}
	return
}

func InsertLink(backend storage.LinksBackend, shortURL string, target string) {
	backend.CreateOrUpdateLink(BuildLink(shortURL, target))
}
