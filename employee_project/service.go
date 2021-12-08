package employee_project

import (
    "context"
    "strconv"

    "github.com/gogo/protobuf/types"
    "gorm.io/gorm"

    pb "github.com/dinhtp/lets-go-pbtype/employee_project"
)

type Service struct {
    db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
    return &Service{db: db}
}

func (s Service) Create(ctx context.Context, r *pb.Employee_Project) (*pb.Employee_Project, error) {
    if err := validateOne(r); nil != err {
        return nil, err
    }

    employeeid, _ := strconv.Atoi(r.GetEmployeeId())
    projectid, _ := strconv.Atoi(r.GetProjectId())
    empl_pro, duplicate, err := NewRepository(s.db).checkDuplicateOne(uint(employeeid), uint(projectid))

    if nil != err {
        return nil, err
    }
    if duplicate == 1 {
        return prepareDataToResponse(empl_pro), nil
    }

    employee_project, err := NewRepository(s.db).CreatOne(prepareDataToRequest(r))
    if nil != err {
        return nil, err
    }

    return prepareDataToResponse(employee_project), nil
}

func (s Service) Delete(ctx context.Context, r *pb.Employee_Project) (*types.Empty, error) {
    if err := validateOne(r); nil != err {
        return nil, err
    }

    err := NewRepository(s.db).DeleteOne(r)
    if nil != err {
        return nil, err
    }

    return &types.Empty{}, nil
}
