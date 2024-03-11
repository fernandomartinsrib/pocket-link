package shortener

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncode(t *testing.T) {
	shortUrl, err := Encode("http://huge-url.com")
	fmt.Println(shortUrl)
	require.NoError(t, err)
	require.NotEmpty(t, shortUrl)
}
