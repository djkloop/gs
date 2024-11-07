package model

import "time"

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
