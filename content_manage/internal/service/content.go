package service

import (
	"content_manage/internal/biz"
	"context"
	"fmt"
	"time"

	pb "content_manage/api/v1/cms"
)

type ContentService struct {
	pb.UnimplementedContentServer

	uc *biz.ContentUseCase
}

func NewContentService(uc *biz.ContentUseCase) *ContentService {
	return &ContentService{
		uc: uc,
	}
}

func (s *ContentService) CreateContent(ctx context.Context, req *pb.CreateContentRequest) (*pb.CreateContentReply, error) {

	content := req.GetContent()
	fmt.Println("create content = ", content)
	// 调用biz层方法
	c, err := s.uc.CreateContent(ctx, &biz.Content{
		Title:          content.GetTitle(),
		Description:    content.GetDescription(),
		Author:         content.GetAuthor(),
		VideoURL:       content.GetAuthor(),
		Thumbnail:      content.GetThumbnail(),
		Category:       content.GetCategory(),
		Duration:       time.Duration(content.GetDuration()),
		Resolution:     content.GetResolution(),
		FileSize:       content.GetFileSize(),
		Format:         content.GetFormat(),
		Quality:        content.GetQuality(),
		ApprovalStatus: content.GetApprovalStatus(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateContentReply{
		Message: c.Title,
	}, nil
}
func (s *ContentService) UpdateContent(ctx context.Context, req *pb.UpdateContentRequest) (*pb.UpdateContentReply, error) {
	content := req.GetContent()
	fmt.Println("update content = ", content)

	// 调用biz层方法
	c, err := s.uc.UpdateContent(ctx, &biz.Content{
		ID:             content.GetId(),
		Title:          content.GetTitle(),
		Description:    content.GetDescription(),
		Author:         content.GetAuthor(),
		VideoURL:       content.GetAuthor(),
		Thumbnail:      content.GetThumbnail(),
		Category:       content.GetCategory(),
		Duration:       time.Duration(content.GetDuration()),
		Resolution:     content.GetResolution(),
		FileSize:       content.GetFileSize(),
		Format:         content.GetFormat(),
		Quality:        content.GetQuality(),
		ApprovalStatus: content.GetApprovalStatus(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateContentReply{
		Message: c.Title,
		Id:      c.ID,
	}, nil
}
func (s *ContentService) DeleteContent(ctx context.Context, req *pb.DeleteContentRequest) (*pb.DeleteContentReply, error) {
	return &pb.DeleteContentReply{}, nil
}
func (s *ContentService) GetContent(ctx context.Context, req *pb.GetContentRequest) (*pb.GetContentReply, error) {
	return &pb.GetContentReply{}, nil
}
func (s *ContentService) ListContent(ctx context.Context, req *pb.ListContentRequest) (*pb.ListContentReply, error) {
	return &pb.ListContentReply{}, nil
}
