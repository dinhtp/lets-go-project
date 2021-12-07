package model

import "gorm.io/gorm"

type Task struct {
    gorm.Model
    ProjectID   int
    EmployeeID  int
    Name        string
    Status      string
    Description string
}
