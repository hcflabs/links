package util

import (
	"github.com/hcflabs/links/lib/generated"
	"github.com/hcflabs/links/lib/storage"
	"google.golang.org/protobuf/types/known/timestamppb"
	// "time"
)

func BuildLink(shortURL string, target string) (link *generated.Link) {
	md := &generated.Link_LinksMetadata{
		Owner:     "",
		Protected: false,
		Private:   false,
		Created:   timestamppb.Now(),
		Modified:  timestamppb.Now(),
	}

	link = &generated.Link{
		ShortUrl:  shortURL,
		TargetUrl: target,
		Metadata:  md,
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
