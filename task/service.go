package task

import (
    "math"
    "context"
    "strconv"

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
    if err := validateCreate(r); nil != err{
        return nil, err
    }

    task, err := NewRepository(s.db).CreatOne(prepareDataToRequest(r))

    if nil != err {
        return nil, err
    }

    return prepareDataToResponse(task),nil
}

func (s Service) Update(ctx context.Context, r *pb.Task) (*pb.Task, error) {
    if err := validateUpdate(r); nil != err {
        return nil, err
    }

    id, _ := strconv.Atoi(r.GetId())
    task,err := NewRepository(s.db).UpdateOne(id, prepareDataToRequest(r))

    if nil != err {
        return nil,err
    }

    return prepareDataToResponse(task), nil
}

func (s Service) Get(ctx context.Context, r *pb.OneTaskRequest) (*pb.Task, error) {
    if err := validateOne(r); nil != err{
        return nil, err
    }

    id, _ := strconv.Atoi(r.GetId())
    task,err := NewRepository(s.db).FindOne(id)

    if nil != err {
        return nil,err
    }

    return prepareDataToResponse(task), nil
}

func (s Service) List(ctx context.Context, r *pb.ListTaskRequest) (*pb.ListTaskResponse, error) {
    var list []*pb.Task
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

    return &pb.ListTaskResponse{
        Items:      list,
        TotalCount: uint32(count),
        Page:       r.GetPage(),
        Limit:      r.GetLimit(),
        MaxPage:    uint32(math.Ceil(float64(uint32(count)) / float64(r.GetLimit()))),
    }, nil
}

func (s Service) Delete(ctx context.Context, r *pb.OneTaskRequest) (*types.Empty, error) {
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
