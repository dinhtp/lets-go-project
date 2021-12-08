package model

import "gorm.io/gorm"

type Employee_Project struct {
    gorm.Model
    EmployeeID uint
    ProjectID  uint
}

