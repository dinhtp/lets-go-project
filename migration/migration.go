package migration

import (
    "github.com/go-gormigrate/gormigrate/v2"
    "gorm.io/gorm"

    "github.com/dinhtp/lets-go-project/migration/versions"
)

func Migrate(db *gorm.DB) error {

    m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
        {
            ID:      "20211207091900",
            Migrate: versions.Version20211207091900,
        },
    })

    return m.Migrate()
}
