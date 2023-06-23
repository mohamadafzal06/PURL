package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mohamadafzal06/purl/entity"
	"github.com/mohamadafzal06/purl/pkg/randomstring"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host            string        `koanf:"host"`
	Port            int           `koanf:"port"`
	Password        string        `koanf:"password"`
	DB              int           `koanf:"db"`
	MaxIdleConns    int           `koanf:"max_idle_conns"`
	ConnMaxIdleTime time.Duration `koanf:"conn_max_idle_time"`
}

type Redis struct {
	client *redis.Client
}

func New(cf Config) Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cf.Host, cf.Port),
		Password: cf.Password,
		DB:       cf.DB,

		// config for Pool
		// 10
		MaxIdleConns: cf.MaxIdleConns,
		// time.Duration(240 * time.Second)
		ConnMaxIdleTime: cf.ConnMaxIdleTime,
	})

	return Redis{client: rdb}
}

func (r Redis) Client() *redis.Client {
	return r.client
}

func (r Redis) Ping(ctx context.Context) error {
	status := r.client.Ping(ctx)
	if err := status.Err(); err != nil {
		return fmt.Errorf("Pinging redis failed: %w\n", err)
	}
	log.Printf("Pinging successfully...")
	return nil
}

func (r Redis) Save(ctx context.Context, url string, expires int64) (string, error) {
	var key string
	rg := randomstring.RandomGenerator{
		Length: 6,
	}

	key = rg.RandomString()
	// ok, err := r.isKeyInDB(ctx, key)
	// if err != nil {
	// 	return "", fmt.Errorf("cannot check the key is in redis or not: %w\n", err)
	// }
	// if !ok {
	// 	key = rg.RandomString()
	// }

	// TODO: set expires.Format properly(dont hardcod)
	shortLink := entity.URL{
		Key:         key,
		OriginalURL: url,
		Expires:     expires,
		Visits:      0,
	}

	// TODO: command is not correct
	_, err := r.client.Do(ctx, "HMSET", shortLink.Key, "url", shortLink.OriginalURL, "expires", shortLink.Expires, "visits", shortLink.Visits).Result()
	if err != nil {
		return "", fmt.Errorf("cannot do the HMSET redis raw command: %w\n", err)
	}

	_, err = r.client.Do(ctx, "EXPIREAT", shortLink.Key, shortLink.Expires).Result()
	if err != nil {
		return "", fmt.Errorf("cannot do the EXPIREAT redis raw command: %w\n", err)

	}

	return shortLink.Key, nil
}

func (r Redis) Load(ctx context.Context, key string) (string, error) {

	url, err := r.client.Do(ctx, "HGET", key, "url").Result()
	if err != nil {
		return "", err
	}
	urlString, ok := url.(string)
	if ok {
		if len(urlString) == 0 {

			// TODO: return better error
			return "", fmt.Errorf("the link is not found: %w\n", err)
		}
	}

	_, err = r.client.Do(ctx, "HINCRBY", key, "visits", 1).Result()

	return urlString, nil
}

func (r Redis) LoadInfo(ctx context.Context, key string) (entity.URL, error) {

	var shortLink entity.URL
	// TODO: is binding all field of shortLink, or key is ignored
	//err := r.client.HGetAll(ctx, key).Scan(&shortLink)
	sRes := r.client.HMGet(ctx, key, "url", "expires", "visits")
	var link entity.URL
	err := sRes.Scan(&link)
	if err != nil {
		return entity.URL{}, fmt.Errorf("the key is not exitst: %w\n", err)
	}

	return shortLink, nil
}

func (r Redis) Close(ctx context.Context) error {
	err := r.client.Close()
	if err != nil {
		return fmt.Errorf("cannot close connection: %w\n", err)
	}

	return nil
}

// func (r Redis) isKeyInDB(ctx context.Context, key string) (bool, error) {
// 	exists, err := r.client.Do(ctx, "EXISTS", key).Result()
// 	if err != nil {
// 		// TODO: log the Error
// 		return false, fmt.Errorf("cannot do EXISTS raw redis command: %w\n", err)
// 	}
// 	res, ok := exists.(int)
// 	if ok {
// 		if res > 0 {
// 			return true, nil
// 		}
// 	}

// 	// TODO: check the result
// 	return false, nil
// }

// func (r Redis) isAvailable(ctx context.Context, key string) bool {
// 	exists, err := r.client.Do(ctx, "EXISTS", key).Result()
// 	if err != nil {
// 		return false
// 	}

// 	// TODO: rest of the code should be refactored
// 	res, ok := exists.(int)
// 	if ok {
// 		if res > 0 {
// 			return true
// 		}
// 		return false
// 	}
// 	return false
// }
