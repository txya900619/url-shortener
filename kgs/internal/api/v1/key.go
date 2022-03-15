package server

import (
	"context"

	pb_v1 "github.com/txya900619/url-shortener/kgs/pkg/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type KeyServiceServer struct {
	pb_v1.UnimplementedKeyServiceServer
}

func (s *KeyServiceServer) GenerateKey(ctx context.Context, in *emptypb.Empty) (*pb_v1.GenerateKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateKey not implemented")
}

func (s *KeyServiceServer) DeleteKey(ctx context.Context, in *pb_v1.DeleteKeyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteKey not implemented")
}
