package app

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/txya900619/url-shortener/internal/shorturl/domain/shorturl"
	pkg_errors "github.com/txya900619/url-shortener/pkg/errors"
)

type ShortUrlService struct {
	shortUrlRepo      shorturl.Repository
	shortUrlCacheRepo shorturl.CacheRepository
	keyService        KeyService
}

func NewShortUrlService(shortUrlRepo shorturl.Repository, shortUrlCacheRepo shorturl.CacheRepository, keyService KeyService) ShortUrlService {
	if shortUrlRepo == nil {
		panic("shorturl repository is nil")
	}

	if shortUrlCacheRepo == nil {
		panic("shorturl cache repository is nil")
	}

	if keyService == nil {
		panic("key service is nil")
	}

	return ShortUrlService{
		shortUrlRepo:      shortUrlRepo,
		shortUrlCacheRepo: shortUrlCacheRepo,
		keyService:        keyService,
	}
}

// return created short url id
func (s *ShortUrlService) CreateShortUrl(ctx context.Context, expireAt time.Time, originUrl string) (string, error) {
	//get key from kgs
	shortUrlId, err := s.keyService.GenerateKey(ctx)
	if err != nil {
		return "", errors.Wrap(err, "unable to generate key")
	}

	shortUrl, err := shorturl.NewShortUrl(shortUrlId, expireAt, originUrl)
	if err != nil {
		return "", err
	}

	//add short url to db
	err = s.shortUrlRepo.CreateShortUrl(ctx, shortUrl)
	if err != nil {
		return "", err
	}

	return shortUrlId, nil
}

func (s *ShortUrlService) RedirectToOriginUrl(ctx context.Context, shortUrlId string) (string, error) {
	//TODO: maybe not found need cache too
	originUrl, err := s.shortUrlCacheRepo.GetShortUrlCache(ctx, shortUrlId)
	if err != nil {
		if err != redis.Nil {
			return "", err
		}

		//if not found in cache, get from db
		shortUrl, err := s.shortUrlRepo.GetShortUrl(ctx, shortUrlId)
		if err != nil {
			//if not found in db, return NotFoundError
			if _, ok := err.(shorturl.NotFoundError); ok {
				return "", pkg_errors.NewNotFoundError(err.Error(), "not-found")
			}
			return "", err
		}

		//if short url is expired, return NotFoundError
		if shortUrl.IsExpired() {
			err := s.shortUrlRepo.DeleteShortUrl(ctx, shortUrlId)
			if err != nil {
				log.Fatal("unable to delete short url")
			}

			return "", pkg_errors.NewNotFoundError(err.Error(), "not-found")
		}

		//if found in db and not expired, add to cache
		err = s.shortUrlCacheRepo.AddShortUrlCache(ctx, shortUrl)
		if err != nil {
			log.Fatal("unable to add short url to cache")
		}

		originUrl = shortUrl.OriginUrl()
	}

	return originUrl, nil
}
