package shorturl_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/txya900619/url-shortener/internal/shorturl/domain/shorturl"
)

func TestNewShortURL(t *testing.T) {
	t.Parallel()
	id := "id"
	expireAt := time.Now().Add(time.Hour)
	originUrl := "https://google.com"

	sURL, err := shorturl.NewShortUrl(id, expireAt, originUrl)
	require.NoError(t, err)

	assert.Equal(t, id, sURL.Id())
	assert.Equal(t, expireAt, sURL.ExpireAt())
	assert.Equal(t, originUrl, sURL.OriginUrl())
}

func TestNewShortURL_invalid(t *testing.T) {
	t.Parallel()
	id := "id"
	expireAt := time.Now().Add(time.Hour)
	originUrl := "https://google.com"

	_, err := shorturl.NewShortUrl("", expireAt, originUrl)
	assert.Error(t, err)

	_, err = shorturl.NewShortUrl(id, time.Time{}, originUrl)
	assert.Error(t, err)

	_, err = shorturl.NewShortUrl(id, time.Now().Add(-time.Second), originUrl)
	assert.Error(t, err)

	_, err = shorturl.NewShortUrl(id, expireAt, "")
	assert.Error(t, err)

	_, err = shorturl.NewShortUrl(id, expireAt, "a/a")
	assert.Error(t, err)
}
