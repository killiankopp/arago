package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"

	pb "github.com/killiankopp/arago/ad/proto"
	trackerpb "github.com/killiankopp/arago/tracker/proto"
)

type AdService struct {
	pb.UnimplementedAdServiceServer
	collection *mongo.Collection
}

func NewAdService(collection *mongo.Collection) *AdService {
	return &AdService{collection: collection}
}

func (s *AdService) CreateAd(ctx context.Context, req *pb.CreateAdRequest) (*pb.CreateAdResponse, error) {
	ad := req.Ad
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

func (s *AdService) findAdByUUID(ctx context.Context, uuid string) (*pb.Ad, error) {
	var ad pb.Ad
	err := s.collection.FindOne(ctx, bson.M{"_id": uuid}).Decode(&ad)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, status.Errorf(codes.NotFound, "ad not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to read ad: %v", err)
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
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to connect to tracker service: %v", err)
		return status.Errorf(codes.Internal, "failed to connect to tracker service: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

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
