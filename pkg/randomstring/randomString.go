package randomstring

import (
	"math/rand"
	"time"
)

type RandomGenerator struct {
	Length int
}

func (rg RandomGenerator) GenerateRandom() string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_")
	b := make([]rune, rg.Length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
