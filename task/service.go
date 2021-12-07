package task

import (
    "context"

    "gorm.io/gorm"
    "github.com/gogo/protobuf/types"

    pb "github.com/dinhtp/lets-go-pbtype/task"
)

type Service struct {
    db *gorm.DB
}

func NewService(db *gorm.DB) *Service {

    return &Service{db: db}
}

func (s Service) Create(ctx context.Context, r *pb.Task) (*pb.Task, error) {
    return &pb.Task{}, nil
}

func (s Service) Update(ctx context.Context, r *pb.Task) (*pb.Task, error) {
    return &pb.Task{}, nil
}

func (s Service) Get(ctx context.Context, r *pb.OneTaskRequest) (*pb.Task,error) {

    return &pb.Task{}, nil
}

func (s Service) List(ctx context.Context, r *pb.ListTaskRequest) (*pb.ListTaskResponse, error) {
    return &pb.ListTaskResponse{}, nil
}

func (s Service) Delete(ctx context.Context, r *pb.OneTaskRequest) (*types.Empty, error) {
    return &types.Empty{}, nil
}