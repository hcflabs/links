package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/hcflabs/links/lib/models"
	"github.com/sirupsen/logrus"

	// "github.com/jackc/pgx/v5"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresConfig struct {
	Host     string
	User     string
	Password string
	Database string
	Port     string
}

type PostgresLinksBackend struct {
	DB *pgxpool.Pool
	// Config PostgresConfig
}

// Create a new instance of the logger. You can have any number of instances.
var log = logrus.New()

func BuildPostgresBackend(localConfig PostgresConfig) (backend PostgresLinksBackend) {
	// fmt.Printf("%+v\n", config)
	// dsn := "host=localhost user=postgres password=postgres dbname=hcflinks port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.Host, config.User, config.Password, config.Database, config.Port)

	// "postgres://username:password@localhost:5432/database_name"
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", localConfig.User, localConfig.Password, localConfig.Host, localConfig.Port, localConfig.Database)
    poolconfiL, err := pgxpool.ParseConfig(connStr)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to parse config: %v\n", err)
        os.Exit(1)
    }
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	ctx := context.Background()
	db, _ := pgxpool.New(ctx, builtURL)

    looger := &log.Logger{
        Out:          os.Stderr,
        Formatter:    new(log.JSONFormatter),
        Hooks:        make(log.LevelHooks),
        Level:        log.InfoLevel,
        ExitFunc:     os.Exit,
        ReportCaller: false,
    }
    poolconfiL.ConnConfig.Logger = log(looger)
    conn, err := pgxpool.ConnectConfig(context.Background(), poolconfiL)

	// db, _ := gorm.Open(postgres.New(postgres.Config{
	// 	DSN: dsn,
	// 	PreferSimpleProtocol: true, // disables implicit prepared statement usage
	// 	// DriverName: "github.com/lib/pq",
	//   }), &gorm.Config{})
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.Host, config.User, config.Password, config.Database, config.Port)

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",config.User,config.Password,config.Host, config.Port, config.Database )
	// db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	backend = PostgresLinksBackend{DB: conn}

	return
}

func (s PostgresLinksBackend) Start() {
	// Create Table
	s.DB.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS "links" (
		short_url VARCHAR (250) UNIQUE NOT NULL,
		target_url VARCHAR (2048) NOT NULL,
		permanent BOOLEAN NOT NULL,
		protected BOOLEAN NOT NULL,
		owned_by VARCHAR (255) NOT NULL,
		created TIMESTAMP NOT NULL default current_timestamp,
		modified TIMESTAMP NOT NULL default current_timestamp,
		description VARCHAR (500)
	);`)
	// Create Index
	s.DB.Exec(context.Background(), `CREATE UNIQUE INDEX IF NOT EXISTS short_url_index ON links(short_url);`)
	// Write Update Function
	s.DB.Exec(context.Background(), `
	CREATE OR REPLACE FUNCTION update_modified_column() RETURNS TRIGGER AS $$ BEGIN NEW.modified = now();
	RETURN NEW;
	END;
	$$ language 'plpgsql';
	CREATE TRIGGER update_links_modtime BEFORE
	UPDATE ON links FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
	`)
}

// GetTargetLink implements LinksBackend.
func (s PostgresLinksBackend) GetTargetLink(url string) (target *string, permanent bool) {
	// s.DB.First(&link, "short_url", url)
	// .Scan(&target, permanent)
	target_struct := struct {
		target_url string
		permanent  bool
	}{}

	pgxscan.Select(context.Background(), s.DB, &target_struct, "select target_url, permanent from links where short_url=$1", url)
	return &target_struct.target_url, target_struct.permanent

}

// CreateOrUpdateLink implements LinksBackend.
func (s PostgresLinksBackend) CreateOrUpdateLink(entry models.InternalLink) {
	// s.DB.Clauses(clause.OnConflict{
	// 	Columns:   []clause.Column{{Name: "short_url"}},
	// 	// DoUpdates: clause.AssignmentColumns([]string{"target_url", owned_by, "description", "updated_at"}),
	// 	DoUpdates: clause.AssignmentColumns([]string{"target_url", "updated_at"}),
	// }).Create(&entry)
	if _, err := s.DB.Exec(context.Background(),
		`insert into links(short_url, target_url, permanent, protected, owned_by, description) values ($1, $2, $3, $4, $5, $6)
		on conflict (short_url) do update set target_url=excluded.target_url, owned_by=excluded.owned_by, permanent=excluded.permanent, 
		protected=excluded.protected, description=excluded.description;`,
		entry.ShortUrl, entry.TargetUrl, entry.Permanent, entry.Permanent, strconv.FormatBool(false), entry.Description); err != nil {
		panic(err)
	}
}

// DeleteLink implements LinksBackend.
func (s PostgresLinksBackend) DeleteLink(url string) {
	// s.DB.Clauses(clause.OnConflict{
	// 	Columns:   []clause.Column{{Name: "id"}},
	// 	DoUpdates: clause.AssignmentColumns([]string{"name", "age"}),
	// }).Create(&entry)
	// s.DB.Where(fmt.Sprintf("short_url == %s", url)).Delete(&models.Link{})
	if _, err := s.DB.Exec(context.Background(), "delete from links where short_url=$1", url); err != nil {
		panic("Internal server error")
	}
}

// GetAllLinksPaginated implements LinksBackend.
func (s PostgresLinksBackend) GetAllLinksPaginated(offset int, pagesize int) (links []models.InternalLink) {
	panic("unimplemented")
}

// GetLinkMetadata implements LinksBackend.
func (s PostgresLinksBackend) GetLinkMetadata(url string) (link *models.InternalLink) {
	// s.DB.First(&link, "short_url", url)
	// var link []*models.InternalLink
	pgxscan.Select(context.Background(), s.DB, &link,
		`SELECT short_url, target_url, permanent, created, modified, description, owner 
		FROM links WHERE short_url=$1`, url)

	// if &link == nil {
	// 	return nil
	// }

	return
}

// GetOwnersLinks implements LinksBackend.
func (s PostgresLinksBackend) GetOwnersLinks(owner string) (links []models.InternalLink) {

	panic("unimplemented")

}

// GetOwnersLinksPaginated implements LinksBackend.
func (s PostgresLinksBackend) GetOwnersLinksPaginated(owner string, offset int, pagesize int) (links []models.InternalLink) {
	panic("unimplemented")
}
