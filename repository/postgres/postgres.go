package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/mohamadafzal06/purl/entity"
)

var (
	ErrDBInitialization          = errors.New("Cannot initialize a connection to database.")
	ErrDBPing                    = errors.New("Cannot ping the database.")
	ErrKeyDoesNotExist           = errors.New("The Key does not exist.")
	ErrURLDoesNotExist           = errors.New("The URL does not exist.")
	ErrDuplicatedKey             = errors.New("This Key is already exist.")
	ErrKeyIsWrong                = errors.New("Cannot retrieve URL by this key.")
	ErrDBScanning                = errors.New("Cannot scanning the result of the query.")
	ErrInsertURLFailed           = errors.New("Cannot insert this (key, url) into db.")
	ErrRetireveInsertedURLFailed = errors.New("Cannot retrieve (key, url) from DB.")
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
		log.Fatalf("cannot initialize a connection to database: %v\n", err)
		return nil, ErrDBInitialization
	}

	log.Printf("\nConnect to DB successfully...\n")

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
		return ErrDBPing
	}

	log.Printf("\nPinging DB successfully...\n")

	return nil
}

func (p *Postgres) GetLongURL(ctx context.Context, key string) (string, error) {
	row := p.db.QueryRowContext(ctx, "SELECT url FROM urls WHERE key=$1", key)
	if err := row.Err(); err != nil {
		// TODO: change the return value witht another entity.URL
		return "", ErrKeyIsWrong
	}

	var url string
	err := row.Scan(&url)
	if err != nil {
		// TODO: change the return value witht another entity.URL
		return "", ErrDBScanning
	}
	return url, nil
}

func (p *Postgres) SetShortURL(ctx context.Context, longUrl entity.URL) (uint64, error) {
	_, err := p.db.ExecContext(ctx, "INSERT INTO urls (key, url) VALUES($1, $2);", longUrl.Key, longUrl.LongURL)
	if err != nil {
		// TODO: change the return value witht another entity.URL
		return 0, ErrInsertURLFailed
	}

	row := p.db.QueryRowContext(ctx, "SELECT * FROM urls WHERE key=$1;", longUrl.Key)
	if row.Err() != nil {
		return 0, ErrRetireveInsertedURLFailed
	}
	var insertedUrl struct {
		id  int64
		key string
		url string
	}
	if err = row.Scan(&insertedUrl.id, &insertedUrl.key, &insertedUrl.url); err != nil {
		return 0, ErrDBScanning
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
		return false, "", ErrURLDoesNotExist
	}

	return true, key, nil
}

func (p *Postgres) IsKeyInDB(ctx context.Context, key string) (bool, string, error) {
	row := p.db.QueryRowContext(ctx, "SELECT url FROM urls WHERE key=$1", key)

	var url string

	err := row.Scan(&url)
	if err != nil {
		return false, "", ErrKeyDoesNotExist
	}

	return true, url, nil
}
