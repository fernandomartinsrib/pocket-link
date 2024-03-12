package shortener

import "github.com/jaevor/go-nanoid"

const (
	customUrlLength = 12
	base62Alphabet  = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

func Encode(longUrl string) (string, error) {
	shortUrl, err := nanoid.CustomASCII(base62Alphabet, customUrlLength)
	if err != nil {
		return "", err
	}

	return shortUrl(), nil
}
