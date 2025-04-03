package main

import (
	"log"
	"net"

	"github.com/killiankopp/arago/tracker/config"
	pb "github.com/killiankopp/arago/tracker/proto"
	"github.com/killiankopp/arago/tracker/server/db"
	"github.com/killiankopp/arago/tracker/server/service"
	"google.golang.org/grpc"
)

func startGRPCServer(trackerService pb.TrackerServiceServer) error {
	s := grpc.NewServer()

	pb.RegisterTrackerServiceServer(s, trackerService)

	lis, err := net.Listen("tcp", config.ServerURI)
	if err != nil {
		return err
	}

	log.Printf("Server listening on %s", config.ServerURI)

	return s.Serve(lis)
}

func main() {
	client, err := db.ConnectToMongoDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer db.DisconnectFromMongoDB(client)

	adCollection := client.Database(config.DBName).Collection(config.CollectionName)
	trackerService := service.NewTrackerService(adCollection)

	if err := startGRPCServer(trackerService); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
