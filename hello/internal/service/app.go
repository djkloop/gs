package service

import (
	"context"

	pb "hello/api/operate/v1"
)

type AppService struct {
	pb.UnimplementedAppServer
}

func NewAppService() *AppService {
	return &AppService{}
}

func (s *AppService) CreateContent(ctx context.Context, req *pb.CreateContentReq) (*pb.CreateContentRep, error) {
	return &pb.CreateContentRep{}, nil
}
