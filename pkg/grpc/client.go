package grpc

import (
	"github.com/spf13/viper"
	kgs_grpc "github.com/txya900619/url-shortener/pkg/genproto/kgs"
)

func NewKGSClient() (kgs_grpc.KeyServiceClient, error) {
	conn, err := NewGrpcConn(viper.GetString("KGS_ADDR"))
	if err != nil {
		return nil, err
	}

	return kgs_grpc.NewKeyServiceClient(conn), nil
}
