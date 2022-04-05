package adapters

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/txya900619/url-shortener/internal/shorturl/domain/shorturl"
)

type RedisCacheShortUrlRepository struct {
	client redis.UniversalClient
}

func NewRedisCacheShortUrlRepository(client redis.UniversalClient) *RedisCacheShortUrlRepository {
	if client == nil {
		panic("redis client is nil")
	}

	return &RedisCacheShortUrlRepository{
		client: client,
	}
}

func (r *RedisCacheShortUrlRepository) AddShortUrlCache(ctx context.Context, sUrl *shorturl.ShortUrl) error {
	return r.client.Set(ctx, sUrl.Id(), sUrl.OriginUrl(), time.Until(sUrl.ExpireAt())).Err()
}

func (r *RedisCacheShortUrlRepository) GetShortUrlCache(ctx context.Context, shortUrlId string) (string, error) {
	return r.client.Get(ctx, shortUrlId).Result()
}
