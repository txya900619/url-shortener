package main

import (
	server_v1 "github.com/txya900619/url-shortener/kgs/internal/api/v1"
	pb_v1 "github.com/txya900619/url-shortener/kgs/pkg/api/v1"

	"fmt"
	"log"
	"net"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func main() {
	viper.AutomaticEnv()

	server := grpc.NewServer(grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor), grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor))

	pb_v1.RegisterKeyServiceServer(server, &server_v1.KeyServiceServer{})

	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(server)

	reflection.Register(server)

	healthServer := health.NewServer()
	healthpb.RegisterHealthServer(server, healthServer)

	addr := fmt.Sprintf(":%d", viper.GetInt("PORT"))
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}
	defer listener.Close()

	if err := server.Serve(listener); err != nil {
		log.Fatal("failed to serve: ", err)
	}

	server.GracefulStop()
}
