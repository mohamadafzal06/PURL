package randomstring

import (
	"errors"
	"math"
	"os"
	"strings"

	"github.com/google/uuid"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length   = uint64(len(alphabet))
)

type RandomGenerator struct {
	Length int
}

func (rg RandomGenerator) RandomString() string {
	uuidStr := uuid.New().String()[:rg.Length]
	return uuidStr
}

func (rg RandomGenerator) Encode(number uint64) string {
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(11)

	for ; number > 0; number = number / length {
		encodedBuilder.WriteByte(alphabet[(number % length)])
	}

	return encodedBuilder.String()
}

func (rg RandomGenerator) Decode(encoded string) (uint64, error) {
	var number uint64

	for i, symbol := range encoded {
		alphabeticPosition := strings.IndexRune(alphabet, symbol)

		if alphabeticPosition == -1 {
			return uint64(alphabeticPosition), errors.New("invalid character: " + string(symbol))
		}
		number += uint64(alphabeticPosition) * uint64(math.Pow(float64(length), float64(i)))
	}

	return number, nil
}

func (rg RandomGenerator) EnforceHTTP(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}

func (rg RandomGenerator) RemoveDomainError(url string) bool {
	if url == os.Getenv("DOMAIN") {
		return false
	}

	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0]

	if newURL == os.Getenv("DOMAIN") {
		return false
	}

	return true
}
