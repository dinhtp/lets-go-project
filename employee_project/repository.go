package employee_project

import (
    "strconv"

    "gorm.io/gorm"

    pb "github.com/dinhtp/lets-go-pbtype/employee_project"
    "github.com/dinhtp/lets-go-project/model"
)

type Repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
    return &Repository{db: db}
}

func (r *Repository) CreatOne(ep *model.EmployeeProject) (*model.EmployeeProject, error) {

    query := r.db.Where(ep).FirstOrCreate(&ep)

    if err := query.Error; nil != err {
        return nil, err
    }

    return ep, nil
}

func (r *Repository) DeleteOne(e *pb.Employee_Project) error {
    employeeId, _ := strconv.Atoi(e.GetEmployeeId())
    projectId, _ := strconv.Atoi(e.GetProjectId())
    query := r.db.Delete(&model.EmployeeProject{}, "employee_id = ?  AND project_id = ?", employeeId, projectId)
    if err := query.Error; nil != err {
        return err
    }

    return nil
}
