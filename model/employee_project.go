package model

import "gorm.io/gorm"

type employee_project struct {
    gorm.Model
    EmployeeID uint
    ProjectID  uint
}

