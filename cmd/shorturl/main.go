package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"github.com/txya900619/url-shortener/internal/shorturl/ports"
	"github.com/txya900619/url-shortener/pkg/ginx"
	"github.com/txya900619/url-shortener/pkg/gocqlx"
)

func main() {
	viper.AutomaticEnv()

	router := ginx.NewEngine()

	gocqlxClient, err := gocqlx.NewGoCqlxClient()
	if err != nil {
		log.Fatal(err)
	}

	err = gocqlxClient.ExecStmt("CREATE KEYSPACE IF NOT EXISTS shorturl WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': 1};")
	if err != nil {
		log.Fatal(err)
	}

	err = gocqlxClient.ExecStmt("CREATE TABLE IF NOT EXISTS shorturl.shorturl ( id text PRIMARY KEY, expire_at timestamp, origin_url text);")
	if err != nil {
		log.Fatal(err)
	}

	httpServer, err := initHttpServer(gocqlxClient)
	if err != nil {
		panic(err)
	}

	ports.RegisterHandlers(router, httpServer)

	router.Run(fmt.Sprintf(":%s", viper.GetString("HTTP_PORT")))
}
