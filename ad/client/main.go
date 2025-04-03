package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"log"
	"time"

	"github.com/killiankopp/arago/ad/config"
	pb "github.com/killiankopp/arago/ad/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

const (
	defaultTitle       = "Super Title"
	defaultDescription = "Super Description"
	defaultURL         = "http://superexample.com"
	defaultUUID        = "fe65a9a9-0fe3-4df8-b2e0-28564c340b9f"
)

func main() {
	log.Println("Starting main function")
	conn, err := setupConnection(config.ServerURI)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("failed to close connection: %v", err)
		}
	}(conn)
	c := pb.NewAdServiceClient(conn)

	ctx, cancel := createContext(time.Second)
	defer cancel()

	createAd(c, ctx, defaultTitle, defaultDescription, defaultURL)
	readAd(c, ctx, defaultUUID)
	serveAd(c, ctx, defaultUUID)
	log.Println("Ending main function")
}

func setupConnection(address string) (*grpc.ClientConn, error) {
	log.Println("Setting up connection")
	creds := insecure.NewCredentials()
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}
	log.Println("Connection setup complete")
	return conn, nil
}

func createContext(timeout time.Duration) (context.Context, context.CancelFunc) {
	log.Println("Creating context with timeout")
	return context.WithTimeout(context.Background(), timeout)
}

func createAd(client pb.AdServiceClient, ctx context.Context, title, description, url string) {
	log.Println("Starting createAd function")
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
	log.Println("Ending createAd function")
}

func readAd(client pb.AdServiceClient, ctx context.Context, uuid string) {
	log.Println("Starting readAd function")
	r, err := client.ReadAd(ctx, &pb.AdRequest{
		Uuid: uuid,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			log.Printf("Ad not found: %v", err)
		} else {
			log.Fatalf("could not read ad: %v", err)
		}
		return
	}
	log.Printf("Ad read: %v", r.GetAd())
	log.Println("Ending readAd function")
}

func serveAd(client pb.AdServiceClient, ctx context.Context, uuid string) {
	log.Println("Starting serveAd function")
	r, err := client.ServeAd(ctx, &pb.AdRequest{
		Uuid: uuid,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			log.Printf("Ad not found: %v", err)
		} else {
			log.Fatalf("could not serve ad: %v", err)
		}
		return
	}
	log.Printf("Ad served: %v", r.GetAd())
	log.Println("Ending serveAd function")
}
