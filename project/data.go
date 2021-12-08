package project

import (
    "fmt"
    "time"
    "strconv"

    "gorm.io/gorm"

    pb "github.com/dinhtp/lets-go-pbtype/project"
    "github.com/dinhtp/lets-go-project/model"
)

func prepareDataToResponse(p *model.Project) *pb.Project {
    data := &pb.Project{
        Id:          fmt.Sprintf("%d", p.ID),
        CompanyId:   fmt.Sprintf("%d", p.CompanyID),
        Name:        p.Name,
        Code:        p.Code,
        Status:      p.Status,
        Description: p.Description,
        TotalTask:   "0",
    }
    return data
}

func prepareDataToRequest(p *pb.Project) *model.Project {
    companyid, _ := strconv.Atoi(p.GetCompanyId())
    return &model.Project{
        Model:       gorm.Model{},
        CompanyID:   uint(companyid),
        Name:        p.Name,
        Code:        p.Code,
        Status:      p.Status,
        Description: p.Description,
    }
}

func getModel(id uint, c *model.Project) *model.Project {
    c.ID = id
    return &model.Project{
        Model: gorm.Model{
            ID:        id,
            CreatedAt: time.Time{},
            UpdatedAt: time.Time{},
            DeletedAt: gorm.DeletedAt{},
        },
        CompanyID:   c.CompanyID,
        Name:        c.Name,
        Code:        c.Code,
        Status:      c.Status,
        Description: c.Description,
    }
}
