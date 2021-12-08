package project

import (
    "fmt"
    "strconv"
    "strings"

    "gorm.io/gorm"

    "github.com/dinhtp/lets-go-project/model"
    pb "github.com/dinhtp/lets-go-pbtype/project"
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

func (r *Repository) DeleteOne(id int) error {
    var result model.Project

    query := r.db.Model(&model.Project{}).Where("id=?", id).Delete(&result)
    if err := query.Error; nil != err {
        return err
    }

    return nil
}

func (r *Repository) UpdateOne(id int, p *model.Project) (*model.Project, error) {
    query := r.db.Model(&model.Project{}).Where("id=?", id).UpdateColumns(getModel(uint(id), p))

    if err := query.Error; nil != err {
        return nil, err
    }

    return p, nil
}

func (r *Repository) CreatOne(p *model.Project) (*model.Project, error) {
    query := r.db.Model(&model.Project{}).Create(p)

    if err := query.Error; nil != err {
        return nil, err
    }

    return p, nil
}

func (r *Repository) ListAll(req *pb.ListProjectRequest) ([]*model.Project, int64, error) {
    var count int64
    var list []*model.Project

    sql := ""
    limit := int(req.GetLimit())
    offset := limit * int(req.GetPage()-1)
    companyId, _ := strconv.Atoi(req.GetCompanyId())

    if req.GetSearchField() != "" && req.GetSearchValue() != "" {
        searchFields := strings.Split(req.GetSearchField(), ",")
        searchValue := fmt.Sprintf("'%%%s%%'", req.GetSearchValue())

        for idx, field := range searchFields {
            if idx == 0 {
                sql += fmt.Sprintf("%s LIKE %s", field, searchValue)
                continue
            }
            sql += fmt.Sprintf(" OR %s LIKE %s", field, searchValue)
        }
    }

    listQuery := r.db.Model(&model.Project{}).Select("*").Limit(limit).Offset(offset)
    countQuery := r.db.Model(&model.Project{}).Select("id")

    if req.GetCompanyId() != "" {
        listQuery = listQuery.Where("company_id = ?", uint(companyId))
        countQuery = countQuery.Where("company_id = ?", uint(companyId))
    }

    if sql != "" {
        countQuery = countQuery.Where(sql)
        listQuery = listQuery.Where(sql)
    }

    if err := countQuery.Count(&count).Error; nil != err {
        return nil, 0, err
    }

    if err := listQuery.Find(&list).Limit(limit).Offset(offset).Error; nil != err {
        return nil, 0, err
    }

    return list, count, nil
}