module github.com/dinhtp/lets-go-project

go 1.14

replace github.com/dinhtp/lets-go-pbtype => ../lets-go-pbtype

require (
	github.com/bxcodec/faker/v3 v3.6.0
	github.com/dinhtp/lets-go-pbtype v0.0.0-20211003031624-3f0ff640e3ac
	github.com/go-gormigrate/gormigrate/v2 v2.0.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gogo/protobuf v1.3.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.9.0
	golang.org/x/net v0.0.0-20211205041911-012df41ee64c // indirect
	golang.org/x/sys v0.0.0-20211205182925-97ca703d548d // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211203200212-54befc351ae9 // indirect
	google.golang.org/grpc v1.42.0
	gorm.io/driver/mysql v1.2.1
	gorm.io/gorm v1.22.4
)
