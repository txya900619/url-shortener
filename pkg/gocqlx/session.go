package gocqlx

import (
	"fmt"
	"strings"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/spf13/viper"
)

func NewGoCqlxClient() (*gocqlx.Session, error) {
	hosts := strings.Split(viper.GetString("CASSANDRA_HOSTS"), ",")
	fmt.Println(viper.GetString("CASSANDRA_HOSTS"))
	cluster := gocql.NewCluster(hosts...)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: viper.GetString("CASSANDRA_USERNAME"),
		Password: viper.GetString("CASSANDRA_PASSWORD"),
	}
	cluster.ProtoVersion = 4
	cluster.Keyspace = viper.GetString("CASSANDRA_KEYSPACE")
	session, err := gocqlx.WrapSession(cluster.CreateSession())

	return &session, err
}
