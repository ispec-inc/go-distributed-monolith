package view

import (
	pb "github.com/ispec-inc/monorepo/go/proto/article/view"
	"github.com/ispec-inc/monorepo/go/svc/article/pkg/domain/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewUser(m *model.User) *pb.User {
	return &pb.User{
		Id:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		Email:       m.Email,
		CreatedAt:   timestamppb.New(m.CreatedAt),
		UpdatedAt:   timestamppb.New(m.UpdatedAt),
	}
}

func NewUsers(ms []*model.User) []*pb.User {
	vs := make([]*pb.User, len(ms))
	for i, m := range ms {
		vs[i] = NewUser(m)
	}
	return vs
}
