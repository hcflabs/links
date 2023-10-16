package storage_test

import (
	"github.com/hcflabs/links/lib/models"
	"github.com/hcflabs/links/lib/storage"
	"github.com/hcflabs/links/lib/util"
	"testing"
)

func GetEmpty() (backend storage.InMemoryLinksBackend) {
	backend = storage.InMemoryLinksBackend{
		LinkMap: make(map[string]models.Link),
	}
	return
}

func TestFoo(t *testing.T) {
	fake := GetEmpty()
	util.InsertLink(fake, "holdon", "holdonredirect")
	util.InsertLink(fake, "great", "greatredirect")
	util.InsertLink(fake, "hcf", "https://haltcatchfire.io")

	v, _ := fake.GetTargetLink("holdon")
	if *v != "holdonredirect" {
		t.Errorf("Redirect Broken %s", *v)
	}

	fake.DeleteLink("holdon")
	v, _ = fake.GetTargetLink("holdon")
	if v != nil {
		t.Errorf("Expected Deleted %s", *v)
	}

}
