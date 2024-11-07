package services

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	goflow "github.com/s8sg/goflow/v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"imooc_go_web/internal/process"
	"imooc_go_web/internal/services/cms/handle"
)

type CmsApp struct {
	CmsHandle   *handle.CmsHandle
	DB          *gorm.DB
	RDB         *redis.Client
	flowService *goflow.FlowService
}

func NewCmsApp() *CmsApp {
	app := &CmsApp{}
	connectCmsDB(app)
	connectCmsRDB(app)
	connectFlowServices(app)
	app.CmsHandle = handle.NewCmsHandle(app.DB, app.RDB, app.flowService)
	go func() {
		process.ExecContentFlow(app.DB)
	}()
	return app
}

func connectCmsDB(app *CmsApp) {
	mysqlDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "test_x:7rD4LkY4FJfW7DAe@tcp(43.143.243.166:3306)/test_x?charset=utf8mb4&parseTime=True&loc=Local",
	}))
	if err != nil {
		fmt.Printf("connectCmsDB = [%v]\n", err)
		panic(err)
	}

	db, err := mysqlDB.DB()
	if err != nil {
		fmt.Printf("connectCmsDB = [%v]\n", err)
		panic(err)
	}

	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(2)

	app.DB = mysqlDB
}

func connectCmsRDB(app *CmsApp) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "43.143.243.166:6379",
		Password: "bS9@xG2?", // no password set
		DB:       0,          // use default DB
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	app.RDB = rdb
}

func connectFlowServices(app *CmsApp) {
	fs := &goflow.FlowService{
		Port:              8999,
		RedisURL:          "43.143.243.166:6379",
		RedisPassword:     "bS9@xG2?",
		WorkerConcurrency: 5,
	}
	app.flowService = fs
}
