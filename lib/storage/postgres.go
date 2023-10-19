package storage

import (
	"fmt"

	"github.com/hcflabs/links/lib/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostgresConfig struct {
	Host     string
	User     string
	Password string
	Database string
	Port     string
}

type PostgresLinksBackend struct {
	DB *gorm.DB
	// Config PostgresConfig
}

func (s PostgresLinksBackend) migrateDb() {
	s.DB.AutoMigrate(&models.Link{})
}

// CreateOrUpdateLink implements LinksBackend.
func (s PostgresLinksBackend) CreateOrUpdateLink(entry models.Link) {
	s.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "short_url"}},
		DoUpdates: clause.AssignmentColumns([]string{"target_url", "owner", "description", "updated_at"}),
	}).Create(&entry)
}

// DeleteLink implements LinksBackend.
func (s PostgresLinksBackend) DeleteLink(url string) {
	// s.DB.Clauses(clause.OnConflict{
	// 	Columns:   []clause.Column{{Name: "id"}},
	// 	DoUpdates: clause.AssignmentColumns([]string{"name", "age"}),
	// }).Create(&entry)
	s.DB.Where(fmt.Sprintf("short_url == %s", url)).Delete(&models.Link{})
}

// GetAllLinksPaginated implements LinksBackend.
func (s PostgresLinksBackend) GetAllLinksPaginated(offset int, pagesize int) (links []models.Link) {
	panic("unimplemented")
}

// GetLinkMetadata implements LinksBackend.
func (s PostgresLinksBackend) GetLinkMetadata(url string) (link *models.Link) {
	panic("unimplemented")
}

// GetOwnersLinks implements LinksBackend.
func (s PostgresLinksBackend) GetOwnersLinks(owner string) (links []models.Link) {


}

// GetOwnersLinksPaginated implements LinksBackend.
func (s PostgresLinksBackend) GetOwnersLinksPaginated(owner string, offset int, pagesize int) (links []models.Link) {
	panic("unimplemented")
}

// GetTargetLink implements LinksBackend.
func (s PostgresLinksBackend) GetTargetLink(url string) (target *string, permanent bool) {
	var link models.Link
	s.DB.First(&link, "short_url", url)
	if &link == nil {
		return nil, false
	}
	
	return &link.TargetUrl, link.Permanent
}
