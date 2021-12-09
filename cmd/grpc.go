package cmd

import (
    "fmt"
    "net"
    "os"
    "time"
    "context"
    "syscall"
    "os/signal"

    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "google.golang.org/grpc"
    "github.com/sirupsen/logrus"
    _ "github.com/go-sql-driver/mysql"

    ppb "github.com/dinhtp/lets-go-pbtype/project"
    tpb "github.com/dinhtp/lets-go-pbtype/task"
    epb "github.com/dinhtp/lets-go-pbtype/employee_project"
    "github.com/dinhtp/lets-go-project/employee_project"
    "github.com/dinhtp/lets-go-project/project"
    "github.com/dinhtp/lets-go-project/task"

)

var grpcCmd = &cobra.Command{
    Use:   "grpc",
    Short: "go project service serve grpc command",
    Run:   runGrpcCommand,
}

func init() {
    serveCmd.AddCommand(grpcCmd)

    grpcCmd.Flags().StringP("backend", "", "core-tax-grpc-address", "gRPC address")
    grpcCmd.Flags().StringP("mysqlDsn", "", "mysql-dsn", "mysql connection string")

    _ = viper.BindPFlag("backend", grpcCmd.Flags().Lookup("backend"))
    _ = viper.BindPFlag("mysqlDsn", grpcCmd.Flags().Lookup("mysqlDsn"))
}

func runGrpcCommand(cmd *cobra.Command, args []string) {
    ctx := context.Background()
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

    // init DB Connection
    mysqlChan := make(chan *gorm.DB, 1)
    go initializeDbConnection("mysqlDsn", c, mysqlChan)
    orm := <-mysqlChan

    // services
    grpcServer := grpc.NewServer()
    grpcServer = initializeServices(orm, grpcServer)

    // init GRPC backend
    grpcAddr := viper.GetString("backend")
    lis, err := net.Listen("tcp", grpcAddr)
    if err != nil {
        panic(err)
    }

    // Serve GRPC
    go func() {
        err = grpcServer.Serve(lis)
        if err != nil {
            panic(err)
        }
    }()

    logrus.WithFields(logrus.Fields{
        "service": "go-project-service",
        "type":    "grpc",
    }).Info("go project service server started")

    <-c
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    logrus.WithFields(logrus.Fields{
        "service": "go-project-service",
        "type":    "grpc",
    }).Info("go project service gracefully shutdowns")
}

func initializeDbConnection(mysqlDsnField string, c chan os.Signal, mysqlChan chan *gorm.DB) {
    mysqlDsn := viper.GetString(mysqlDsnField)
    orm, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})
    if nil != err {
        fmt.Println(err)
        c <- syscall.SIGTERM
    }

    sqlDB, err := orm.DB()
    if nil != err {
        panic(err)
    }

    sqlDB.SetConnMaxLifetime(300 * time.Minute)
    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetMaxOpenConns(15)

    fmt.Println(fmt.Sprintf("MySQL connection established"))

    mysqlChan <- orm
}

func initializeServices(orm *gorm.DB, grpcServer *grpc.Server) *grpc.Server {
    projectService := project.NewService(orm)
    taskService := task.NewService(orm)
    employeeProjectService := employee_project.NewService(orm)

    ppb.RegisterProjectServiceServer(grpcServer, projectService)
    tpb.RegisterTaskServiceServer(grpcServer, taskService)
    epb.RegisterEmployeeServiceServer(grpcServer, employeeProjectService)

    return grpcServer
}
