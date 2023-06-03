package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/mohamadafzal06/purl/delivery/httpserver"
	"github.com/mohamadafzal06/purl/pkg/randomstring"
	"github.com/mohamadafzal06/purl/repository/postgres"
	"github.com/mohamadafzal06/purl/service"
)

func main() {

	repo, err := postgres.NewDB("postgres", "postgres", "127.0.0.1", "5432", "purl")
	if err != nil {
		log.Fatal("cannot initialize repository")
	}

	rg := randomstring.RandomGenerator{
		Length: 6,
	}

	srv := service.New(repo, rg)

	app := echo.New()

	httpserver.NewHealth().Register(app.Group(""))

	// TODO: check for implementation of interface
	httpserver.NewServer(srv).Register(app.Group("/api"))

	if err := app.Start(":1996"); !errors.Is(err, http.ErrServerClosed) {
		log.Println("echo initialization failed")
	}
}
