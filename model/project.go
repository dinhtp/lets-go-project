package model

import "gorm.io/gorm"

type Project struct {
    gorm.Model
    CompanyID   int
    Name        string
    Code        string
    Status      string
    Description string
    Task        []Task
}
