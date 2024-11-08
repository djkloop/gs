package data

import (
	"content_manage/internal/conf"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewContentRepo)

// Data .
type Data struct {
	DB *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	var appData = &Data{}

	// 连接数据库
	connectCmsDB(appData, c)

	return appData, cleanup, nil
}

func connectCmsDB(appData *Data, c *conf.Data) {
	mysqlDB, err := gorm.Open(mysql.Open(c.GetDatabase().GetSource()))
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
	appData.DB = mysqlDB
}
