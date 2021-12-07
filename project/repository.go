package project

import (
    "gorm.io/gorm"

    "github.com/dinhtp/lets-go-project/model"
)

type Repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {

    return &Repository{db: db}
}

func (r *Repository) FindOne(id int) (*model.Project, error) {
    var result model.Project

    query := r.db.Model(&model.Project{}).Where("id = ?", id).First(&result)

    if err := query.Error; nil != err {
        return nil, err
    }

    return &result, nil
}

