package storage

import (
	"context"
	"fmt"
	"os"

	// "github.com/georgysavva/scany/v2/pgxscan"
	"github.com/hcflabs/links/lib/models"
	"github.com/sirupsen/logrus"

	"github.com/redis/go-redis/v9"
	// "github.com/jackc/pgx/v5"
	// "strconv"
	// "github.com/jackc/pgx/v5/pgxpool"
)

type PostgresConfig struct {
	Host     string
	User     string
	Password string
	Database string
	Port     string
}

type RedisLinksBackend struct {
	client *redis.Client
}

var log = logrus.New()

func BuildRedisBackend(PostgresConfig) LinksBackend {
	rds := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &RedisLinksBackend{client: rds}
}

// Start implements LinksBackend.
func (*RedisLinksBackend) Start() {
	panic("unimplemented")
}

// GetTargetLink implements LinksBackend.
func (*RedisLinksBackend) GetTargetLink(url string) (target *string, permanent bool) {
	panic("unimplemented")
}

// CreateOrUpdateLink implements LinksBackend.
func (*RedisLinksBackend) CreateOrUpdateLink(link models.InternalLink) {
	panic("unimplemented")
}

// DeleteLink implements LinksBackend.
func (*RedisLinksBackend) DeleteLink(url string) {
	panic("unimplemented")
}

// GetAllLinksPaginated implements LinksBackend.
func (*RedisLinksBackend) GetAllLinksPaginated(offset int, pagesize int) (links []models.InternalLink) {
	panic("unimplemented")
}

// GetLinkMetadata implements LinksBackend.
func (*RedisLinksBackend) GetLinkMetadata(url string) (link *models.InternalLink) {
	panic("unimplemented")
}

// GetOwnersLinks implements LinksBackend.
func (*RedisLinksBackend) GetOwnersLinks(owner string) (links []models.InternalLink) {
	panic("unimplemented")
}

// GetOwnersLinksPaginated implements LinksBackend.
func (*RedisLinksBackend) GetOwnersLinksPaginated(owner string, offset int, pagesize int) (links []models.InternalLink) {
	panic("unimplemented")
}


