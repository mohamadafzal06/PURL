package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mohamadafzal06/purl/delivery/httpserver"
)

func main() {

	var srv httpserver.Service

	app := echo.New()

	httpserver.NewHealth().Register(app.Group(""))
	httpserver.New(srv).Register(app.Group("api"))

	if err := app.Start(":1996"); !errors.Is(err, http.ErrServerClosed) {
		log.Println("echo initialization failed")
	}

}
