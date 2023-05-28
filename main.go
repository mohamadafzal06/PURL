package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var db = make(map[string]string)

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func shorten(url string) string {
	sUrl := RandStringRunes(7)
	db[sUrl] = url
	return sUrl
}

func getLongUrl(key string) string {
	lUrl := db[key]
	return lUrl
}

func main() {
	e := echo.New()
	e.GET("/health-check", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	e.Logger.Fatal(e.Start(":8088"))
}
