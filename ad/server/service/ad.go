package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/killiankopp/arago/ad/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"time"

	pb "github.com/killiankopp/arago/ad/proto"
	trackerpb "github.com/killiankopp/arago/tracker/proto"
)

type AdService struct {
	pb.UnimplementedAdServiceServer
	collection  *mongo.Collection
	redisClient *redis.Client
}

func NewAdService(collection *mongo.Collection, redisClient *redis.Client) *AdService {
	return &AdService{collection: collection, redisClient: redisClient}
}

func (s *AdService) CreateAd(ctx context.Context, req *pb.CreateAdRequest) (*pb.CreateAdResponse, error) {
	ad := req.Ad

	if ad.Title == "" || ad.Url == "" {
		return nil, status.Errorf(codes.InvalidArgument, "title and url are required")
	}

	ad.Uuid = uuid.New().String()
	err := s.insertAd(ctx, ad)
	if err != nil {
		return nil, err
	}
	return &pb.CreateAdResponse{Uuid: ad.Uuid}, nil
}

func (s *AdService) insertAd(ctx context.Context, ad *pb.Ad) error {
	_, err := s.collection.InsertOne(ctx, bson.M{
		"_id":         ad.Uuid,
		"title":       ad.Title,
		"description": ad.Description,
		"url":         ad.Url,
	})
	return err
}

func (s *AdService) ReadAd(ctx context.Context, req *pb.AdRequest) (*pb.AdResponse, error) {
	log.Printf("ReadAd called with UUID: %s", req.Uuid)

	ad, err := s.findAdByUUID(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}
	return &pb.AdResponse{Ad: ad}, nil
}

type AdWithoutLock struct {
	Uuid        string `json:"uuid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

func (s *AdService) findAdByUUID(ctx context.Context, uuid string) (*pb.Ad, error) {
	var ad pb.Ad

	val, err := s.redisClient.Get(ctx, uuid).Result()
	if errors.Is(err, redis.Nil) {
		err = s.collection.FindOne(ctx, bson.M{"_id": uuid}).Decode(&ad)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return nil, status.Errorf(codes.NotFound, "ad not found")
			}
			return nil, status.Errorf(codes.Internal, "failed to read ad: %v", err)
		}

		adWithoutLock := AdWithoutLock{
			Uuid:        ad.Uuid,
			Title:       ad.Title,
			Description: ad.Description,
			Url:         ad.Url,
		}

		adBytes, err := json.Marshal(adWithoutLock)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to marshal ad: %v", err)
		}
		err = s.redisClient.Set(ctx, uuid, adBytes, 10*time.Minute).Err()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to cache ad: %v", err)
		}
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get ad from cache: %v", err)
	} else {
		err = json.Unmarshal([]byte(val), &ad)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to unmarshal ad: %v", err)
		}
	}

	return &ad, nil
}

func (s *AdService) ServeAd(ctx context.Context, req *pb.AdRequest) (*pb.AdResponse, error) {
	log.Printf("ServeAd called with UUID: %s", req.Uuid)

	ad, err := s.findAdByUUID(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}

	err = s.trackImpression(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}

	return &pb.AdResponse{Ad: ad}, nil
}

func (s *AdService) trackImpression(ctx context.Context, adUUID string) error {
	conn, err := s.connectToTrackerService()
	if err != nil {
		return err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Failed to close connection: %v", err)
		}
	}(conn)

	trackerClient := trackerpb.NewTrackerServiceClient(conn)
	_, err = trackerClient.UpdateImpression(ctx, &trackerpb.UpdateImpressionRequest{AdUuid: adUUID})
	if err != nil {
		log.Printf("Failed to update impression for ad_uuid %s: %v", adUUID, err)
		return status.Errorf(codes.Internal, "failed to update impression: %v", err)
	}
	return nil
}

func (s *AdService) connectToTrackerService() (*grpc.ClientConn, error) {
	creds := insecure.NewCredentials()
	conn, err := grpc.NewClient(config.ServerPrintURI, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Printf("Failed to connect to tracker service: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to connect to tracker service: %v", err)
	}
	return conn, nil
}
