package main

import (
	"log"
	"net"

	"github.com/killerekt/grpc-prac/config"
	"github.com/killerekt/grpc-prac/internal/db"
	"github.com/killerekt/grpc-prac/pb"
	auth "github.com/killerekt/grpc-prac/pkg"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to get values from the env")
	}

	err = db.DBInit(cfg)
	if err != nil {
		log.Fatalf("Failed to Initialise the DB")
	}

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	db.DB.AutoMigrate(&auth.User{})

	grpcServer := grpc.NewServer()
	authservice := auth.NewAuthService(db.DB)
	pb.RegisterAuthServiceServer(grpcServer, &authservice)

	log.Printf("Services successfully register")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to listen the server on the port")
	}

	log.Println("Server up at running")
}
