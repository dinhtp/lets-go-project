package model

import "gorm.io/gorm"

type Task struct {
    gorm.Model
    ProjectID   uint
    Name        string
    Status      string
    Description string
}
