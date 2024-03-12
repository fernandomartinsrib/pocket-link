package shortener

import (
	"fmt"
	"pocket-link/config"

	"github.com/jaevor/go-nanoid"
)

const (
	customUrlLength = 12
	base62Alphabet  = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

func Encode(longUrl string) (string, error) {
	config, err := config.LoadConfig("../..")
	if err != nil {
		return "", err
	}

	enc, err := nanoid.CustomASCII(base62Alphabet, customUrlLength)
	if err != nil {
		return "", err
	}

	shortUrl := fmt.Sprintf("%s/%s", config.Domain, enc())
	return shortUrl, nil
}
