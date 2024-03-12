package shortener

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncode(t *testing.T) {
	shortUrl, err := Encode("http://huge-url.com")
	require.NoError(t, err)
	require.NotEmpty(t, shortUrl)
}
