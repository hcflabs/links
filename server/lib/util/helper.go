package util

import (
	"github.com/hcflabs/links/lib/generated"
	"github.com/hcflabs/links/lib/storage"
	// "time"
)

func BuildLink(shortURL string, target string) (link *generated.Link) {
	link = &generated.Link{
		ShortUrl:  shortURL,
		TargetUrl: target,
		// LinkMetadata: generated.LinkMetadata{
		// 	Owner:       "fake_user@domain.com",
		// 	Description: "this is a fake thing",
		// 	CreatedAt:   time.Now(),
		// 	UpdatedAt:   time.Now(),
		// },
		// LinkOptions: generated.LinkOptions{
		Permanent: false,
		// },
	}
	return
}

func InsertLink(backend storage.LinksBackend, shortURL string, target string) {
	backend.CreateOrUpdateLink(BuildLink(shortURL, target))
}
