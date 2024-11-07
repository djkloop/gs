package dao

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"imooc_go_web/internal/model"
)

type AccountDao struct {
	db *gorm.DB
}

func NewAccountDao(db *gorm.DB) *AccountDao {
	return &AccountDao{db: db}
}

func (a *AccountDao) IsExist(userId string) (bool, error) {
	var account model.Account
	err := a.db.Where(&model.Account{UserId: userId}).First(&account).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if err != nil {
		fmt.Printf("AccountDao isExist = [%v]:\n", err)
		return false, err
	}

	return true, nil
}

func (a *AccountDao) Create(user *model.Account) error {
	if err := a.db.Create(&user).Error; err != nil {
		fmt.Printf("AccountDao Create = [%v]\n", err)
		return err
	}
	return nil
}

func (a *AccountDao) FirstByUserID(userID string) (*model.Account, error) {
	var account *model.Account
	if err := a.db.Where(&model.Account{UserId: userID}).First(&account).Error; err != nil {
		fmt.Printf("AccountDao FirstByUserID = [%v]\n", err)
		return nil, err
	}
	return account, nil
}
