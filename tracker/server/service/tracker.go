package service

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"

	pb "github.com/killiankopp/arago/tracker/proto"
)

type TrackerService struct {
	pb.UnimplementedTrackerServiceServer
	collection *mongo.Collection
}

func NewTrackerService(collection *mongo.Collection) *TrackerService {
	return &TrackerService{collection: collection}
}

func (s *TrackerService) UpdateImpression(ctx context.Context, req *pb.UpdateImpressionRequest) (*pb.UpdateImpressionResponse, error) {
	log.Printf("Received UpdateImpression request for ad_uuid: %s", req.AdUuid)

	filter := bson.M{"ad_uuid": req.AdUuid}
	update := bson.M{
		"$inc": bson.M{"count": 1},
		"$set": bson.M{"timestamp": time.Now().Unix()},
	}
	opts := options.Update().SetUpsert(true)

	_, err := s.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Printf("Failed to update impression for ad_uuid %s: %v", req.AdUuid, err)
		return nil, status.Errorf(codes.Internal, "failed to update impression: %v", err)
	}

	log.Printf("Successfully updated impression for ad_uuid: %s", req.AdUuid)
	return &pb.UpdateImpressionResponse{Success: true}, nil
}

func (s *TrackerService) GetImpressionCount(ctx context.Context, req *pb.GetImpressionCountRequest) (*pb.GetImpressionCountResponse, error) {
	filter := bson.M{"ad_uuid": req.AdUuid}
	var impression pb.AdImpression

	err := s.collection.FindOne(ctx, filter).Decode(&impression)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, status.Errorf(codes.NotFound, "ad_uuid not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get impression count: %v", err)
	}

	return &pb.GetImpressionCountResponse{AdUuid: impression.AdUuid, Count: impression.Count}, nil
}
