package adapters

import (
	"context"
	"errors"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/table"

	"github.com/txya900619/url-shortener/internal/shorturl/domain/shorturl"
)

var cassandraShortUrlTable = table.New(table.Metadata{
	Name: "shorturl",
	Columns: []string{
		"id",
		"expire_at",
		"origin_url",
	},
	PartKey: []string{"id"},
	SortKey: []string{},
})

type cassandraShortUrl struct {
	Id        string    `db:"id"`
	ExpireAt  time.Time `db:"expire_at"`
	OriginUrl string    `db:"origin_url"`
}

type CassandraShortUrlRepository struct {
	session *gocqlx.Session
}

func NewCassandraShortUrlRepository(session *gocqlx.Session) *CassandraShortUrlRepository {
	if session == nil {
		panic("gocqlx session is nil")
	}

	return &CassandraShortUrlRepository{
		session: session,
	}
}

func (r *CassandraShortUrlRepository) CreateShortUrl(_ context.Context, sURL *shorturl.ShortUrl) error {
	cassandraShortUrl := cassandraShortUrl{
		Id:        sURL.Id(),
		ExpireAt:  sURL.ExpireAt(),
		OriginUrl: sURL.OriginUrl(),
	}

	return r.session.Query(cassandraShortUrlTable.Insert()).BindStruct(cassandraShortUrl).ExecRelease()
}

func (r *CassandraShortUrlRepository) GetShortUrl(_ context.Context, shortUrlId string) (*shorturl.ShortUrl, error) {
	cassandraShortUrl := cassandraShortUrl{
		Id: shortUrlId,
	}

	err := r.session.Query(cassandraShortUrlTable.Get()).BindStruct(cassandraShortUrl).GetRelease(&cassandraShortUrl)
	if err != nil {
		//if not found, return shorturl domain's NotFoundError
		if errors.Is(err, gocql.ErrNotFound) {
			return nil, shorturl.NotFoundError{
				ShortUrlId: shortUrlId,
			}
		}
		return nil, err
	}

	return shorturl.NewShortUrl(cassandraShortUrl.Id, cassandraShortUrl.ExpireAt, cassandraShortUrl.OriginUrl)
}

func (r *CassandraShortUrlRepository) DeleteShortUrl(_ context.Context, shortUrlId string) error {
	cassandraShortUrl := cassandraShortUrl{
		Id: shortUrlId,
	}

	return r.session.Query(cassandraShortUrlTable.Delete()).BindStruct(cassandraShortUrl).ExecRelease()
}
