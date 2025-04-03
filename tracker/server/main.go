package main

import (
	"log"
	"net"

	pb "github.com/killiankopp/arago/tracker/proto"
	"github.com/killiankopp/arago/tracker/server/db"
	"github.com/killiankopp/arago/tracker/server/service"
	"google.golang.org/grpc"
)

const (
	port   = ":50052"
	dbName = "addb"
)

type server struct {
	pb.UnimplementedTrackerServiceServer
}

func startGRPCServer(trackerService pb.TrackerServiceServer) error {
	s := grpc.NewServer()

	pb.RegisterTrackerServiceServer(s, trackerService)

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

	adCollection := client.Database(dbName).Collection("prints")
	trackerService := service.NewTrackerService(adCollection)

	if err := startGRPCServer(trackerService); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
