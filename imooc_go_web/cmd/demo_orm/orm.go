package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Account struct {
	ID        int64     `gorm:"column:id; primary_key"`
	UserId    string    `gorm:"column:user_id"`
	Password  string    `gorm:"column:password"`
	Nickname  string    `gorm:"column:nickname"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (a Account) TableName() string {
	return "cms_account"
}

func main() {
	db := ConnectDB()
	var accounts []Account
	var account Account

	if err := db.Find(&accounts).Error; err != nil {
		fmt.Println("", err)
		return
	}

	fmt.Println(accounts)

	if err := db.Where("id =?", 1).First(&account).Error; err != nil {
		fmt.Println("", err)
		return
	}
	fmt.Println(account)
}

func ConnectDB() *gorm.DB {
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
