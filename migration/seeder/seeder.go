package seeder

import (
    "gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
    err := CreateFakeProject(db)
    err1 := CreateFakeTask(db)
    if nil != err {
        return err
    }
    if nil != err1 {
        return err1
    }
    return nil
}

func CreateFakeProject(db *gorm.DB) error {
    for i := 1; i <= 10; i++ {
        co, err := FakeProject()
        if nil != err {
            return err
        }

        co.ID = uint(i)
        err = db.Create(&co).Error
        if nil != err {
            return err
        }
    }
    return nil
}

func CreateFakeTask(db *gorm.DB) error {
    for i := 1; i <= 5; i++ {
        co, err := FakeTask()
        if nil != err {
            return err
        }

        co.ID = uint(i)
        err = db.Create(&co).Error
        if nil != err {
            return err
        }
    }
    return nil
}