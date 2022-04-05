package shorturl

import (
	"context"
	"fmt"
)

type NotFoundError struct {
	ShortUrlId string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("short url '%s' not found", e.ShortUrlId)
}

type Repository interface {
	CreateShortUrl(ctx context.Context, sUrl *ShortUrl) error

	GetShortUrl(ctx context.Context, shortUrlId string) (*ShortUrl, error)

	DeleteShortUrl(ctx context.Context, shortUrlId string) error
}

type CacheRepository interface {
	AddShortUrlCache(ctx context.Context, sUrl *ShortUrl) error

	//return cached origin url (short url id is key)
	GetShortUrlCache(ctx context.Context, shortUrlId string) (string, error)
}
