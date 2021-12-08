package project

import (
    "math"
    "context"
    "strconv"

    "gorm.io/gorm"
    "github.com/gogo/protobuf/types"

    pb "github.com/dinhtp/lets-go-pbtype/project"
)

type Service struct {
    db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
    return &Service{db: db}
}

func (s Service) Create(ctx context.Context, r *pb.Project) (*pb.Project, error) {
    if err := validateCreate(r); nil != err{
        return nil, err
    }

    project, err := NewRepository(s.db).CreatOne(prepareDataToRequest(r))

    if nil != err {
        return nil, err
    }

    return prepareDataToResponse(project),nil
}

func (s Service) Update(ctx context.Context, r *pb.Project) (*pb.Project, error) {
    if err := validateUpdate(r); nil != err {
        return nil, err
    }

    id, _ := strconv.Atoi(r.GetId())
    project,err := NewRepository(s.db).UpdateOne(id, prepareDataToRequest(r))

    if nil != err {
        return nil,err
    }

    return prepareDataToResponse(project), nil
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
    var list []*pb.Project
    if err := validateList(r); nil != err {
        return nil, err
    }

    company, count, err := NewRepository(s.db).ListAll(r)

    if nil != err {
        return nil, err
    }
    for i := 0; i < len(company); i++ {
        list = append(list, prepareDataToResponse(company[i]))
    }

    return &pb.ListProjectResponse{
        Items:      list,
        TotalCount: uint32(count),
        Page:       r.GetPage(),
        Limit:      r.GetLimit(),
        MaxPage:    uint32(math.Ceil(float64(uint32(count)) / float64(r.GetLimit()))),
    }, nil
}

func (s Service) Delete(ctx context.Context, r *pb.OneProjectRequest) (*types.Empty, error) {
    if err := validateOne(r); nil != err {
        return nil, err
    }

    id, _ := strconv.Atoi(r.GetId())
    err := NewRepository(s.db).DeleteOne(id)

    if nil != err {
        return nil, err
    }

    return &types.Empty{}, nil
}
