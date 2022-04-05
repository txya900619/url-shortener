//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/scylladb/gocqlx/v2"
	"github.com/txya900619/url-shortener/internal/shorturl/adapters"
	"github.com/txya900619/url-shortener/internal/shorturl/app"
	"github.com/txya900619/url-shortener/internal/shorturl/domain/shorturl"
	"github.com/txya900619/url-shortener/internal/shorturl/ports"
	"github.com/txya900619/url-shortener/pkg/goredis"
	"github.com/txya900619/url-shortener/pkg/grpc"
)

var (
	repositorySet      = wire.NewSet(adapters.NewCassandraShortUrlRepository, wire.Bind(new(shorturl.Repository), new(*adapters.CassandraShortUrlRepository)))
	keyServiceSet      = wire.NewSet(grpc.NewKGSClient, adapters.NewKeyGrpc, wire.Bind(new(app.KeyService), new(adapters.KeyGrpc)))
	cacheRepositorySet = wire.NewSet(adapters.NewRedisCacheShortUrlRepository, wire.Bind(new(shorturl.CacheRepository), new(*adapters.RedisCacheShortUrlRepository)))
)

func initHttpServer(gocqlxClient *gocqlx.Session) (*ports.HttpServer, error) {
	wire.Build(ports.NewHttpServer, app.NewShortUrlService, repositorySet, keyServiceSet, cacheRepositorySet, goredis.NewGoRedisClient)

	return &ports.HttpServer{}, nil
}
