package config

import (
	"os"
	"strconv"
	"time"
)

func getEnv(key string, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}

var DatabaseHost = getEnv("PURL_DATABASE_HOST", "127.0.0.1")
var DatabasePort = getEnv("PURL_DATABASE_PORT", "postgres")
var DatabasePass = getEnv("PURL_DATABASE_PASS", "admin")

var databaseMaxConnString = getEnv("PURL_DATABASE_MAX_CONN", "mokhtasar")
var DatabaseMaxConn, _ = strconv.Atoi(databaseMaxConnTimeString)

var databaseMaxConnTimeString = getEnv("PURL_DATABASE_MAX_CONN_TIME", "disable")
var databaseMaxConnTimeInt, _ = strconv.Atoi(databaseMaxConnTimeString)
var DatabaseMaxConnTime = time.Duration(databaseMaxConnTimeInt)

type Config struct {
	Host            string        `koanf:"host"`
	Port            int           `koanf:"port"`
	Password        string        `koanf:"password"`
	DB              int           `koanf:"db"`
	MaxIdleConns    int           `koanf:"max_idle_conns"`
	ConnMaxIdleTime time.Duration `koanf:"conn_max_idle_time"`
}
