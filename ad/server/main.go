package main

import (
	"log"
	"net"

	pb "github.com/killiankopp/arago/ad/proto"
	"github.com/killiankopp/arago/ad/server/db"
	"github.com/killiankopp/arago/ad/server/service"
	"google.golang.org/grpc"
)

const (
	port   = ":50051"
	dbName = "addb"
)

func startGRPCServer(adService pb.AdServiceServer) error {
	s := grpc.NewServer()

	pb.RegisterAdServiceServer(s, adService)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	log.Printf("Server listening on %s", port)

	return s.Serve(lis)
}

func main() {
	client, err := db.ConnectToMongoDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer db.DisconnectFromMongoDB(client)

	redisClient, err := db.ConnectToRedis()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer db.DisconnectFromRedis(redisClient)

	adCollection := client.Database(dbName).Collection("ads")
	adService := service.NewAdService(adCollection, redisClient)

	if err := startGRPCServer(adService); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
