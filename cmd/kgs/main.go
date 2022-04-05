package main

import (
	"github.com/txya900619/url-shortener/internal/kgs"
	ports "github.com/txya900619/url-shortener/internal/kgs/ports"
	"github.com/txya900619/url-shortener/internal/kgs/schema"
	kgs_grpc "github.com/txya900619/url-shortener/pkg/genproto/kgs"
	"github.com/txya900619/url-shortener/pkg/gorm"

	"fmt"
	"log"
	"net"

	"github.com/spf13/viper"
	"github.com/txya900619/url-shortener/pkg/grpc"
)

func main() {
	viper.AutomaticEnv()

	server := grpc.NewGrpcServer()

	db, err := gorm.Open()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&schema.UnusedKey{}, &schema.UsedKey{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	err = kgs.InsertUnusedKeys(db)
	if err != nil {
		log.Fatalf("failed to generate unused keys: %v", err)
	}

	keyServiceServer, err := ports.NewKeyServiceServer(db)
	if err != nil {
		log.Fatalf("failed to create key service server: %v", err)
	}

	kgs_grpc.RegisterKeyServiceServer(server, keyServiceServer)

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
