package db

import (
	"context"
	"pocket-link/pkg/shortener"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUrl(t *testing.T) {
	longUrl := "http://huge-url.com"

	shortUrl, err := shortener.Encode(longUrl)
	require.NoError(t, err)
	require.NotEmpty(t, shortUrl)

	arg := CreateUrlParams{
		LongUrl:  longUrl,
		ShortUrl: shortUrl,
	}

	storedUrl, err := testQueries.CreateUrl(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, storedUrl)
	require.Equal(t, shortUrl, storedUrl.ShortUrl)
	require.Equal(t, longUrl, storedUrl.LongUrl)
}
