package gocqlx

import (
	"fmt"
	"strings"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/spf13/viper"
)

func newGoCqlxClientWithKeyspace(keyspace string) (*gocqlx.Session, error) {
	hosts := strings.Split(viper.GetString("CASSANDRA_HOSTS"), ",")
	cluster := gocql.NewCluster(hosts...)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: viper.GetString("CASSANDRA_USERNAME"),
		Password: viper.GetString("CASSANDRA_PASSWORD"),
	}
	cluster.ProtoVersion = 4
	cluster.Keyspace = keyspace
	session, err := gocqlx.WrapSession(cluster.CreateSession())

	return &session, err
}

func NewGoCqlxClient() (*gocqlx.Session, error) {
	systemSession, err := newGoCqlxClientWithKeyspace("system")
	if err != nil {
		return nil, err
	}

	err = systemSession.ExecStmt(fmt.Sprintf("CREATE KEYSPACE IF NOT EXISTS %s WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': 1};", viper.GetString("CASSANDRA_KEYSPACE")))
	if err != nil {
		return nil, err
	}

	return newGoCqlxClientWithKeyspace(viper.GetString("CASSANDRA_KEYSPACE"))
}
