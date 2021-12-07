package project

import (
    "context"
    "strconv"

    "github.com/gogo/protobuf/types"
    "gorm.io/gorm"

    pb "github.com/dinhtp/lets-go-pbtype/project"
)

type Service struct {
    db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
    return &Service{db: db}
}

func (s Service) Create(ctx context.Context, r *pb.Project) (*pb.Project, error) {
    return &pb.Project{}, nil
}

func (s Service) Update(ctx context.Context, r *pb.Project) (*pb.Project, error) {
    return &pb.Project{}, nil
}

func (s Service) Get(ctx context.Context, r *pb.OneProjectRequest) (*pb.Project, error) {
    if err := validateOne(r); nil != err {
        return nil, err
    }

    id, _ := strconv.Atoi(r.GetId())
    project, err := NewRepository(s.db).FindOne(id)

    if nil != err {
        return nil, err
    }

    return prepareDataToResponse(project), nil
}

func (s Service) List(ctx context.Context, r *pb.ListProjectRequest) (*pb.ListProjectResponse, error) {
    return &pb.ListProjectResponse{}, nil
}

func (s Service) Delete(ctx context.Context, r *pb.OneProjectRequest) (*types.Empty, error) {
    return &types.Empty{}, nil
}
