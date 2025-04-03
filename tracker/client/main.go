package main

import (
	"context"
	"log"
	"time"

	"github.com/killiankopp/arago/tracker/config"
	pb "github.com/killiankopp/arago/tracker/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultUUID = "fe65a9a9-0fe3-4df8-b2e0-28564c340b9f"
)

func main() {
	conn, client := setupClientConnection()
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Failed to close connection: %v", err)
		}
	}(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	updateImpression(ctx, client, defaultUUID)
	getImpressionCount(ctx, client, defaultUUID)
}

func setupClientConnection() (*grpc.ClientConn, pb.TrackerServiceClient) {
	creds := insecure.NewCredentials()
	conn, err := grpc.NewClient(config.ServerURI, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}

	client := pb.NewTrackerServiceClient(conn)
	return conn, client
}

func updateImpression(ctx context.Context, client pb.TrackerServiceClient, adUuid string) {
	updateReq := &pb.UpdateImpressionRequest{AdUuid: adUuid}
	updateRes, err := client.UpdateImpression(ctx, updateReq)
	if err != nil {
		log.Fatalf("Failed to update impression: %v", err)
	}
	log.Printf("UpdateImpression success: %v", updateRes.Success)
}

func getImpressionCount(ctx context.Context, client pb.TrackerServiceClient, adUuid string) {
	countReq := &pb.GetImpressionCountRequest{AdUuid: adUuid}
	countRes, err := client.GetImpressionCount(ctx, countReq)
	if err != nil {
		log.Fatalf("Failed to get impression count: %v", err)
	}
	log.Printf("Impression count for %s: %d", countRes.AdUuid, countRes.Count)
}
