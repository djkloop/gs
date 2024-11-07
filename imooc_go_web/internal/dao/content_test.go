package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"imooc_go_web/internal/model"
	"testing"
)

func connectDB() *gorm.DB {
	mysqlDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "test_x:7rD4LkY4FJfW7DAe@tcp(43.143.243.166:3306)/test_x?charset=utf8mb4&parseTime=True&loc=Local",
	}))
	if err != nil {
		panic(err)
	}

	db, err := mysqlDB.DB()
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(2)

	mysqlDB = mysqlDB.Debug()

	return mysqlDB
}

func TestContentDao_Create(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		content *model.ContentDetail
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "内容插入测试",
			fields: fields{
				db: connectDB(),
			},
			args: args{
				content: &model.ContentDetail{
					Title: "test-123",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ContentDao{
				db: tt.fields.db,
			}
			if err := c.Create(tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
