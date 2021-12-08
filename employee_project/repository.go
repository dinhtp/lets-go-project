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
    db = db.Debug()
    return &Repository{db: db}
}

func (r *Repository) CreatOne(ep *model.Employee_Project) (*model.Employee_Project, error) {
    query := r.db.Model(&model.Employee_Project{}).Create(ep)

    if err := query.Error; nil != err {
        return nil, err
    }

    return ep, nil
}

func (r *Repository) DeleteOne(e *pb.Employee_Project) error {
    employeeid, _ := strconv.Atoi(e.GetEmployeeId())
    projectid, _ := strconv.Atoi(e.GetProjectId())
    query := r.db.Delete(&model.Employee_Project{}, "employee_id = ?  AND project_id = ?", employeeid, projectid)
    if err := query.Error; nil != err {
        return err
    }

    return nil
}

func (r *Repository) checkDuplicateOne(employeeid uint, projectid uint) (*model.Employee_Project, int, error) {
    var result *model.Employee_Project
    query := r.db.Model(&model.Employee_Project{}).Where("employee_id = ? AND project_id = ?", employeeid, projectid).Find(&result)
    if err := query.Error; nil != err {
        return nil, 0, err
    }
    if result.EmployeeID == employeeid && result.ProjectID == projectid {
        return result, 1, nil
    } else {
        return nil, 0, nil
    }
}
