package adapters

import (
	"context"

	kgs_grpc "github.com/txya900619/url-shortener/pkg/genproto/kgs"
	"google.golang.org/protobuf/types/known/emptypb"
)

type KeyGrpc struct {
	client kgs_grpc.KeyServiceClient
}

func NewKeyGrpc(client kgs_grpc.KeyServiceClient) KeyGrpc {
	if client == nil {
		panic("key service client is nil")
	}

	return KeyGrpc{client: client}
}

// return generated key
func (s KeyGrpc) GenerateKey(ctx context.Context) (string, error) {
	resp, err := s.client.GenerateKey(ctx, &emptypb.Empty{})
	if err != nil {
		return "", err
	}

	return resp.GetKey(), nil
}
