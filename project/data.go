package project

import (
    "fmt"

    "github.com/dinhtp/lets-go-project/model"
    pb "github.com/dinhtp/lets-go-pbtype/project"
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
