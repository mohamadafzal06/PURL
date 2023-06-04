package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	Password string `koanf:"password"`
	DB       int    `koanf:"db"`
}

type Redis struct {
	client *redis.Client
}

func New(config Config) Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
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

// func (r *Redis) GetLongURL(ctx context.Context, key string) (string, error) {
// 	panic("not implemented") // TODO: Implement
// }

// func (r *Redis) IsURLInDB(ctx context.Context, url string) (bool, string, error) {
// 	panic("not implemented") // TODO: Implement
// }

// func (r *Redis) IsKeyInDB(ctx context.Context, key string) (bool, string, error) {
// 	panic("not implemented") // TODO: Implement
// }

// func (r *Redis) SetShortURL(ctx context.Context, lurl entity.URL) (uint64, error) {
// 	panic("not implemented") // TODO: Implement
// }
