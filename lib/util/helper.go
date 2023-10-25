package util

import (
	"github.com/hcflabs/links/lib/models"
	"github.com/hcflabs/links/lib/storage"
	// "time"
)

func BuildLink(shortURL string, target string) (link *models.Link) {
	link = &models.Link{
		ShortUrl:  shortURL,
		TargetUrl: target,
		// LinkMetadata: models.LinkMetadata{
		// 	Owner:       "fake_user@domain.com",
		// 	Description: "this is a fake thing",
		// 	CreatedAt:   time.Now(),
		// 	UpdatedAt:   time.Now(),
		// },
		// LinkOptions: models.LinkOptions{
		Permanent: false,
		// },
	}
	return
}

func InsertLink(backend storage.LinksBackend, shortURL string, target string) {
	backend.CreateOrUpdateLink(BuildLink(shortURL, target))
}
