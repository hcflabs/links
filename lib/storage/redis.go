package storage

import (
	"context"
	b64 "encoding/base64"
	"fmt"
	// "net/url"
	// "os"
	"strconv"

	// "github.com/georgysavva/scany/v2/pgxscan"
	"github.com/hcflabs/links/lib/models"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"

	"github.com/redis/go-redis/v9"
	// "github.com/jackc/pgx/v5"
	// "strconv"
	// "github.com/jackc/pgx/v5/pgxpool"
)

type RedisConfig struct {
	Host     string
	User     string
	Password string
	Database string
	Port     string
}

type RedisLinksBackend struct {
	client *redis.Client
}

var redislog = logrus.New()
var DEFAULT_DATABASE = 0

func BuildRedisBackend(config RedisConfig) LinksBackend {

	database, dberr := strconv.Atoi(config.Database)

	if dberr != nil {
		log.Error("Issues getting database, using default", dberr)
		database = DEFAULT_DATABASE
	}

	rds := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password: config.Password, // no password set
		DB:       database,        // use default DB
	})

	return &RedisLinksBackend{client: rds}
}

// Start implements LinksBackend.
func (r *RedisLinksBackend) Start() {
	// panic("unimplemented")
}

// GetTargetLink implements LinksBackend.
func (r *RedisLinksBackend) GetTargetLink(url string) (target *string, permanent bool) {
	result, err := r.client.Get(context.Background(), url).Result()
	if err != nil {
		redislog.Error("Issues with getting key", err)
		return nil, false
	}
	li := fromB64Str(result)
	return &li.TargetUrl, li.Permanent
}

// CreateOrUpdateLink implements LinksBackend.
func (r *RedisLinksBackend) CreateOrUpdateLink(link *models.Link) {
	err := r.client.Set(context.Background(), link.ShortUrl, toB64Str(link), 0).Err()
	if err != nil {
		panic(err)
	}
}

// DeleteLink implements LinksBackend.
func (r *RedisLinksBackend) DeleteLink(url string) {
	result, err := r.client.Del(context.Background(), url).Result()
	if err != nil {
		redislog.Error("Issues with deleting key", err)
		return
	}
	redislog.Info(fmt.Sprintf("Deleted %s with %d entries"), url, result)
}

// GetAllLinksPaginated implements LinksBackend.
func (r *RedisLinksBackend) GetAllLinksPaginated(offset int, pagesize int) (links *[]models.Link) {
	panic("unimplemented")
}

// GetLinkMetadata implements LinksBackend.
func (r *RedisLinksBackend) GetLinkMetadata(url string) (link *models.Link) {
	result, err := r.client.Get(context.Background(), url).Result()
	if err != nil {
		redislog.Error("Issues with getting key", err)
		return nil
	}
	li := fromB64Str(result)
	return li
}

// GetOwnersLinks implements LinksBackend.
func (r *RedisLinksBackend) GetOwnersLinks(owner string) (links *[]models.Link) {
	panic("unimplemented")
}

// GetOwnersLinksPaginated implements LinksBackend.
func (r *RedisLinksBackend) GetOwnersLinksPaginated(owner string, offset int, pagesize int) (links *[]models.Link) {
	panic("unimplemented")
}

func toB64Str(link *models.Link) string {

	bytes, err := proto.Marshal(link)
	if err != nil {
		redislog.Error("Trouble Unmarsellhing", err)
		panic("bad unmarshell")
	}
	return b64.StdEncoding.EncodeToString(bytes)

}

func fromB64Str(encoded string) (link *models.Link) {
	sDec, err := b64.StdEncoding.DecodeString(encoded)

	if err != nil {
		redislog.Error("Trouble Decoding", err)
		panic("bad decode")
	}
	lclink := &models.Link{}
	err = proto.Unmarshal(sDec, lclink)
	if err != nil {
		redislog.Error("Trouble Unmarshal", err)
		panic("bad Unmarshal")
	}

	return lclink
}
