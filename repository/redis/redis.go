package redis

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
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

func (r *Redis) Client() *redis.Client {
	return r.client
}

func (r *Redis) Ping(ctx context.Context) error {
	status := r.client.Ping(ctx)
	if err := status.Err(); err != nil {
		return fmt.Errorf("Pinging redis failed: %w\n", err)
	}
	log.Printf("Pinging successfully...")
	return nil
}

func (r *Redis) isIDInDB(ctx context.Context, id uint64) bool {

	exists, err := r.client.Do(ctx, "EXISTS", "Shortner:"+strconv.FormatUint(id, 10)).Result()
	if err != nil {
		// TODO: log the Error
		return false
	}

	// TODO: check the result
	return exists.(bool)
}

func (r *Redis) Save(ctx context.Context, url string, expires time.Time) (string, error) {
	var id uint64

	for used := true; used; used = r.isIDInDB(ctx, id) {
		id = rand.Uint64()
	}

	// TODO: set expires.Format properly(dont hardcod)
	shortLink := entity.URL{id, url, expires.Format("2006-01-02 15:04:05.728046 +0300 EEST"), 0}

	// TODO: command is not correct
	_, err := r.client.Do(ctx, "HMSET", "Shortener:"+strconv.FormatUint(id, 10), "url", shortLink.OriginalURL, "expires", shortLink.Expires, "visites", shortLink.Visits).Result()
	if err != nil {
		return "", err
	}

	_, err = r.client.Do(ctx, "EXPIREAT", "Shortener:"+strconv.FormatUint(id, 10), expires.Unix()).Result()
	if err != nil {
		return "", err
	}

	return randomstring.Encode(id), nil
}

func (r *Redis) Load(ctx context.Context, code string) (string, error) {

	decodedId, err := randomstring.Decode(code)
	if err != nil {
		return "", err
	}

	//urlString, err := redisClient.String(conn.Do("HGET", "Shortener:"+strconv.FormatUint(decodedId, 10), "url"))
	url, err := r.client.Do(ctx, "HGET", "Shortener:"+strconv.FormatUint(decodedId, 10), "url").Result()
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

	_, err = r.client.Do(ctx, "HINCRBY", "Shortener:"+strconv.FormatUint(decodedId, 10), "visits", 1).Result()

	return urlString, nil
}
func (r *Redis) isAvailable(ctx context.Context, id uint64) bool {
	exists, err := r.client.Do(ctx, "EXISTS", "Shortener:"+strconv.FormatUint(id, 10)).Result()
	if err != nil {
		return false
	}

	// TODO: rest of the code should be refactored
	existsBool, ok := exists.(bool)
	if ok {
		return existsBool
	}
	return false
}

func (r *Redis) LoadInfo(ctx context.Context, code string) (*entity.URL, error) {

	decodedId, err := randomstring.Decode(code)
	if err != nil {
		return nil, err
	}

	var shortLink entity.URL
	err = r.client.HGetAll(ctx, "Shortener:"+strconv.FormatUint(decodedId, 10)).Scan(shortLink)

	if err != nil {
		return nil, fmt.Errorf("the key is not exitst: %w\n", err)
	}

	return &shortLink, nil
}
