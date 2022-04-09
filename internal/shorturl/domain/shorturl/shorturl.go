package shorturl

import (
	"errors"
	"net/url"
	"time"
)

type ShortUrl struct {
	id        string
	expireAt  time.Time
	originUrl string
}

func NewShortUrl(id string, expireAt time.Time, originUrl string) (*ShortUrl, error) {
	if id == "" {
		return nil, errors.New("empty short url id")
	}
	if _, err := url.ParseRequestURI(originUrl); err != nil {
		return nil, errors.New("invalid origin url")
	}
	if expireAt.IsZero() {
		return nil, errors.New("empty expire at")
	}

	return &ShortUrl{
		id:        id,
		expireAt:  expireAt,
		originUrl: originUrl,
	}, nil
}

func (s *ShortUrl) Id() string {
	return s.id
}

func (s *ShortUrl) ExpireAt() time.Time {
	return s.expireAt
}

func (s *ShortUrl) OriginUrl() string {
	return s.originUrl
}

//check if the short url is expired
func (s *ShortUrl) IsExpired() bool {
	return s.expireAt.Before(time.Now())
}
