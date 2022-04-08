# url-shorterner
URL Shortener using Key Generation Service (KGS)

## TODO
- [ ] add cronjob to delete expired short url
- [ ] better logging
- [ ] repository tests
- [ ] e2e tests
- [ ] cache when not found
- [x] Dockerfile
- [ ] docker-compose.yml
- [x] CI/CD
- [ ] k8s config (maybe in other repo)

## Getting started 
1. Install gowatch
```bash
go install github.com/silenceper/gowatch@latest
```
2. Install generate dependencies
```bash
go install github.com/google/wire/cmd/wire@latest \
&& go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest \
&& go install go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
&& go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
3. Generate needed files
```bash
go generate ./...
```
4. Start PostgreSQL
```bash
docker run --name postgres-docker -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=123456 -d postgres
```
5. Start Redis
```bash
docker run --name redis-docker -p 6379:6379 -d redis redis-server --requirepass "abc123"
```
6. Start Cassandra
```bash
docker run --name cassandra-docker -p 9042:9042 bitnami/cassandra:latest
```
7. Start kgs (if you want to use in production, change KEY_LENGTH to 6 or more)
```bash
cd cmd/kgs \
&& gowatch
```
8. Start shorturl server
```bash
cd cmd/shorturl \
&& gowatch
```

## How to use
See [openapi doc](https://github.com/txya900619/url-shortener/blob/main/api/openapi/shorturl.yml)

## Why use these DBMS
- PostgreSQL (TODO)
- Redis (TODO)
- Cassandra (TODO)