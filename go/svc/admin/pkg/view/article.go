package view

import (
	pb "github.com/yamad07/monorepo/go/proto/admin/view"
	"github.com/yamad07/monorepo/go/svc/admin/pkg/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewArticle(m *model.Article) *pb.Article {
	return &pb.Article{
		Id:        m.ID,
		UserId:    m.UserID,
		Title:     m.Title,
		Body:      m.Body,
		CreatedAt: timestamppb.New(m.CreatedAt),
		UpdatedAt: timestamppb.New(m.UpdatedAt),
	}
}

func NewArticles(ms *model.Articles) []*pb.Article {
	vs := make([]*pb.Article, len(*ms))
	for i, m := range *ms {
		vs[i] = NewArticle(&m)
	}
	return vs
}
