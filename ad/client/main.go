package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"log"
	"time"

	pb "github.com/killiankopp/arago/ad/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

const (
	address            = "localhost:50051"
	defaultTitle       = "Super Title"
	defaultDescription = "Super Description"
	defaultURL         = "http://superexample.com"
	defaultUUID        = "118bc050-ce8f-4210-9b88-adb5958f2c00"
)

func main() {
	conn, err := setupConnection(address)
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
}

func setupConnection(address string) (*grpc.ClientConn, error) {
	creds := insecure.NewCredentials()
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(creds))
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

func readAd(client pb.AdServiceClient, ctx context.Context, uuid string) {
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
}

func serveAd(client pb.AdServiceClient, ctx context.Context, uuid string) {
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
}
