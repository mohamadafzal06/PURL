package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/mohamadafzal06/purl/entity"
)

type Postgres struct {
	db *sql.DB
}

func NewDB(username, password, host, port, dbname string) (*Postgres, error) {
	connectionStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		username, password, host, port, dbname,
	)
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return nil, fmt.Errorf("cannot initialize a connection to database: %w\n", err)
	}
	pg := new(Postgres)
	pg.db = db
	return pg, nil
}

func (p *Postgres) Conn() *sql.DB {
	return p.db
}

func (p *Postgres) Ping() error {
	err := p.db.Ping()
	if err != nil {
		log.Fatalf("cannot ping the database: %v\n", err)
		return fmt.Errorf("cannot ping the database: %v\n", err)
	}

	log.Printf("\nPinging DB successfully...\n")

	return nil
}

func (p *Postgres) GetLongURL(ctx context.Context, key string) (string, error) {
	row := p.db.QueryRowContext(ctx, "SELECT url FROM urls WHERE key=$1", key)
	if err := row.Err(); err != nil {
		// TODO: change the return value witht another entity.URL
		return "", fmt.Errorf("cannot retrieve url by this key: %w\n", err)
	}

	var url string
	err := row.Scan(&url)
	if err != nil {
		// TODO: change the return value witht another entity.URL
		return "", fmt.Errorf("error while scanning: %w\n", err)
	}
	return url, nil
}

func (p *Postgres) SetShortURL(ctx context.Context, longUrl entity.URL) (uint64, error) {
	_, err := p.db.ExecContext(ctx, "INSERT INTO urls (key, url) VALUES($1, $2);", longUrl.Key, longUrl.LongURL)
	if err != nil {
		// TODO: change the return value witht another entity.URL
		return 0, fmt.Errorf("cannot set this url into db: %w\n", err)
	}

	row := p.db.QueryRowContext(ctx, "SELECT * FROM urls WHERE key=$1;", longUrl.Key)
	if row.Err() != nil {
		return 0, fmt.Errorf("cannot retrieve the inserted url: %w\n", err)
	}
	var insertedUrl struct {
		id  int64
		key string
		url string
	}
	if err = row.Scan(&insertedUrl.id, &insertedUrl.key, &insertedUrl.url); err != nil {
		return 0, fmt.Errorf("scanning the inserted url failed: %w\n", err)
	}

	insertedID := uint64(insertedUrl.id)

	// TODO: change the return value witht another entity.URL
	return insertedID, nil
}

func (p *Postgres) IsURLInDB(ctx context.Context, url string) (bool, string, error) {
	row := p.db.QueryRowContext(ctx, "SELECT key FROM urls WHERE url=$1", url)

	var key string

	err := row.Scan(&key)

	if err != nil {
		return false, "", fmt.Errorf("error while scanning the result: %w\n", err)
	}

	return true, key, nil
}

func (p *Postgres) IsKeyInDB(ctx context.Context, key string) (bool, string, error) {
	row := p.db.QueryRowContext(ctx, "SELECT url FROM urls WHERE key=$1", key)

	var url string

	err := row.Scan(&url)
	if err != nil {
		return false, "", fmt.Errorf("error while scanning the result: %w\n", err)
	}

	return true, url, nil
}
