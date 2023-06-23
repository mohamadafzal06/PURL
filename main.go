package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mohamadafzal06/purl/delivery/httpserver"
	"github.com/mohamadafzal06/purl/pkg/randomstring"
	"github.com/mohamadafzal06/purl/repository/redis"
	"github.com/mohamadafzal06/purl/service"
)

func main() {

	dbConfig := redis.Config{
		Host:            "localhost",
		Port:            6380,
		Password:        "",
		MaxIdleConns:    10,
		ConnMaxIdleTime: time.Duration(240 * time.Second),
	}
	rp := redis.New(dbConfig)

	rg := randomstring.RandomGenerator{
		Length: 6,
	}

	srv := service.New(rp, rg)

	app := echo.New()

	httpserver.NewHealth(rp).Register(app.Group(""))

	// TODO: check for implementation of interface
	httpserver.NewServer("localhost", "1996", srv).Register(app.Group("/api"))

	if err := app.Start(":1996"); !errors.Is(err, http.ErrServerClosed) {
		log.Println("echo initialization failed")
	}
}
