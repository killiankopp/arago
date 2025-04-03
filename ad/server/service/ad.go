package service

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	pb "github.com/killiankopp/arago/ad/proto"
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
	_, err := s.collection.InsertOne(ctx, bson.M{
		"_id":         ad.Uuid,
		"title":       ad.Title,
		"description": ad.Description,
		"url":         ad.Url,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateAdResponse{Uuid: ad.Uuid}, nil
}

func (s *AdService) ReadAd(ctx context.Context, req *pb.AdRequest) (*pb.AdResponse, error) {
	// Implement the logic to read an ad
	return &pb.AdResponse{Ad: &pb.Ad{Uuid: req.Uuid}}, nil
}

func (s *AdService) ServeAd(ctx context.Context, req *pb.AdRequest) (*pb.AdResponse, error) {
	// Implement the logic to serve an ad
	return &pb.AdResponse{Ad: &pb.Ad{Uuid: req.Uuid}}, nil
}
