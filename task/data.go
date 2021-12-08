package task

import (
    "fmt"
    "time"
    "strconv"

    "gorm.io/gorm"

    pb "github.com/dinhtp/lets-go-pbtype/task"
    "github.com/dinhtp/lets-go-project/model"
)

func prepareDataToResponse(t *model.Task) *pb.Task {
    data := &pb.Task{
        Id:          fmt.Sprintf("%d", t.ID),
        ProjectId:   fmt.Sprintf("%d", t.ProjectID),
        Name:        t.Name,
        Status:      t.Status,
        Description: t.Description,
    }
    return data
}

func prepareDataToRequest(p *pb.Task) *model.Task {
    projectId, _ := strconv.Atoi(p.GetProjectId())
    return &model.Task{
        Model:       gorm.Model{},
        ProjectID:   uint(projectId),
        Name:        p.Name,
        Status:      p.Status,
        Description: p.Description,
    }
}

func getModel(id uint, c *model.Task) *model.Task {
    c.ID = id

    return &model.Task{
        Model: gorm.Model{
            ID:        id,
            CreatedAt: time.Time{},
            UpdatedAt: time.Time{},
            DeletedAt: gorm.DeletedAt{},
        },
        ProjectID:   c.ProjectID,
        Name:        c.Name,
        Status:      c.Status,
        Description: c.Description,
    }
}
