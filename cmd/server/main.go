package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/joho/godotenv"
	"github.com/napakornsk/cv-backend/database"
	pb "github.com/napakornsk/cv-backend/pb"
	"github.com/napakornsk/cv-backend/service/cv"
)

var fullAddr string = fmt.Sprintf("%s:%s",
	os.Getenv("SERVER_ADDR"),
	os.Getenv("SERVER_PORT"),
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Fatalf("SERVER_PORT not set in .env")
	}
	fullAddr := ":" + port

	lis, err := net.Listen("tcp", fullAddr)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", os.Getenv("SERVER_PORT"), err)
	}

	grpcServer := grpc.NewServer()

	database.IintPostgresDB()

	pb.RegisterCVServiceServer(grpcServer, &cv.CVServer{PostgresClient: database.Db})
	log.Printf("Starting gRPC server on PORT: %s\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
