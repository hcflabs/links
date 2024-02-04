package storage

import (
	"context"
	b64 "encoding/base64"
	// "errors"
	"fmt"

	// "net/url"
	// "os"
	"strconv"

	// "github.com/georgysavva/scany/v2/pgxscan"

	"github.com/hcflabs/links/lib/generated"
	log "github.com/sirupsen/logrus"
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
func (r *RedisLinksBackend) GetTargetLink(url string) (target *string, permanent bool, err error) {
	result, err := r.client.Get(context.Background(), url).Result()
	if err != nil {
		log.Error("Issues with getting key", err)
		return nil, false, err
	}
	li := fromB64Str(result)
	return &li.TargetUrl, li.Permanent, nil
}

// CreateOrUpdateLink implements LinksBackend.
func (r *RedisLinksBackend) CreateOrUpdateLink(link *generated.Link) error {
	return r.client.Set(context.Background(), link.ShortUrl, toB64Str(link), 0).Err()

}

// DeleteLink implements LinksBackend.
func (r *RedisLinksBackend) DeleteLink(url string) error {
	result, err := r.client.Del(context.Background(), url).Result()
	if err != nil {
		log.Error("Issues with deleting key", err)
		return err
	}
	log.Info(fmt.Sprintf("Deleted %s with %d entries"), url, result)
	return nil
}

// GetAllLinksPaginated implements LinksBackend.
func (r *RedisLinksBackend) GetAllLinksPaginated(offset int, pagesize int) (links *[]generated.Link, err error) {
	panic("unimplemented")
}

// GetLinkMetadata implements LinksBackend.
func (r *RedisLinksBackend) GetLinkMetadata(url string) (link *generated.Link, err error) {
	result, err := r.client.Get(context.Background(), url).Result()
	if err != nil {
		log.Error("Issues with getting key", err)
		return nil, err
	}
	li := fromB64Str(result)
	return li, nil
}

// GetOwnersLinks implements LinksBackend.
func (r *RedisLinksBackend) GetOwnersLinks(owner string) (links *[]generated.Link, err error) {
	panic("unimplemented")
}

// GetOwnersLinksPaginated implements LinksBackend.
func (r *RedisLinksBackend) GetOwnersLinksPaginated(owner string, offset int, pagesize int) (links *[]generated.Link, err error) {
	panic("unimplemented")
}

func toB64Str(link *generated.Link) string {

	bytes, err := proto.Marshal(link)
	if err != nil {
		log.Error("Trouble Unmarsellhing", err)
		panic("bad unmarshell")
	}
	return b64.StdEncoding.EncodeToString(bytes)

}

func fromB64Str(encoded string) (link *generated.Link) {
	sDec, err := b64.StdEncoding.DecodeString(encoded)

	if err != nil {
		log.Error("Trouble Decoding", err)
		panic("bad decode")
	}
	lclink := &generated.Link{}
	err = proto.Unmarshal(sDec, lclink)
	if err != nil {
		log.Error("Trouble Unmarshal", err)
		panic("bad Unmarshal")
	}

	return lclink
}
