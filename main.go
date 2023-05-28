package main

import (
	"fmt"
	"math/rand"
	"time"
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
	url := "google.com"
	surl := shorten(url)
	lurl := getLongUrl(surl)
	fmt.Println(url, surl, lurl)
}
