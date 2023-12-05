package utils

import "math/rand"

type LinkUtils interface {
	GenerateShortUrl() string
}

type linkUtils struct{}

func NewLinksUtils() *linkUtils {
	return &linkUtils{}
}

func (util *linkUtils) GenerateShortUrl() string {
	correctSymbols := "ABCDEFGHIKLMNOPQRSTVXYZabcdefghiklmnopqrstvxyz_0123456789"
	var result string
	for i := 0; i < 10; i++ {
		result = result + string(correctSymbols[rand.Intn(56)])
	}
	return result
}
