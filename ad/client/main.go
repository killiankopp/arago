package main

import (
	"context"
	"log"
	"time"

	pb "github.com/killiankopp/arago/ad/proto"
	"google.golang.org/grpc"
)

const (
	address            = "localhost:50051"
	defaultTitle       = "New Ad Title"
	defaultDescription = "New Ad Description"
	defaultURL         = "http://example.com"
)

func main() {
	conn, err := setupConnection(address)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAdServiceClient(conn)

	ctx, cancel := createContext(time.Second)
	defer cancel()

	createAd(c, ctx, defaultTitle, defaultDescription, defaultURL)
}

func setupConnection(address string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func createContext(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout)
}

func createAd(client pb.AdServiceClient, ctx context.Context, title, description, url string) {
	r, err := client.CreateAd(ctx, &pb.CreateAdRequest{
		Ad: &pb.Ad{
			Title:       title,
			Description: description,
			Url:         url,
		},
	})
	if err != nil {
		log.Fatalf("could not create ad: %v", err)
	}
	log.Printf("Ad created: %s", r.GetUuid())
}
