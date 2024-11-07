package dao

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"imooc_go_web/internal/model"
)

type ContentDao struct {
	db *gorm.DB
}

func NewContentDao(db *gorm.DB) *ContentDao {
	return &ContentDao{
		db: db,
	}
}

func (c *ContentDao) Create(content *model.ContentDetail) (int, error) {
	if err := c.db.Create(&content).Error; err != nil {
		fmt.Printf("create content failed, err:%v\n", err)
		return 0, err
	}
	return content.ID, nil
}

func (c *ContentDao) IsExist(contentID int) (bool, error) {
	var content model.ContentDetail
	if err := c.db.Where("id =?", contentID).First(&content).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		fmt.Printf("query content failed, err:%v\n", err)
		return false, err
	}
	return true, nil
}

func (c *ContentDao) Update(contentID int, content *model.ContentDetail) error {
	if err := c.db.Where("id = ?", contentID).Updates(&content).Error; err != nil {
		fmt.Printf("update content failed, err:%v\n", err)
		return err
	}
	return nil
}

func (c *ContentDao) Delete(contentID int) error {
	if err := c.db.Where("id = ? ", contentID).Delete(&model.ContentDetail{}).Error; err != nil {
		fmt.Printf("delete content failed, err:%v\n", err)
		return err
	}
	return nil
}

type SearchParams struct {
	ID       int
	Title    string
	Author   string
	Page     int
	PageSize int
}

func (c *ContentDao) Search(params *SearchParams) ([]*model.ContentDetail, int64, error) {
	query := c.db.Model(&model.ContentDetail{})

	// 构建查询条件
	if params.ID != 0 {
		query = query.Where("id = ?", params.ID)
	}
	if params.Title != "" {
		query = query.Where("title LIKE ?", "%"+params.Title+"%")
	}
	if params.Author != "" {
		query = query.Where("author LIKE ?", "%"+params.Author+"%")
	}

	// 数据总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		fmt.Printf("count content failed, err:%v\n", err)
		return nil, 0, err
	}

	// 计算分页参数
	var page, pageSize = 1, 10
	if params.Page > 0 {
		page = params.Page
	}
	if params.PageSize > 0 {
		pageSize = params.PageSize
	}
	offset := (page - 1) * pageSize

	// 分页查询
	var contents []*model.ContentDetail
	if err := query.Offset(offset).Limit(pageSize).Find(&contents).Error; err != nil {
		fmt.Printf("query content failed, err:%v\n", err)
		return nil, 0, err
	}

	return contents, total, nil
}

func (c *ContentDao) First(id int) (*model.ContentDetail, error) {
	var detail model.ContentDetail
	if err := c.db.Where("id = ?", id).First(&detail).Error; err != nil {
		fmt.Printf("query content failed, err:%v\n", err)
		return &detail, err
	}
	return &detail, nil
}

func (c *ContentDao) UpdateByID(contentID int, column string, value interface{}) error {
	var content model.ContentDetail
	if err := c.db.Model(&content).Where("id =?", contentID).Update(column, value).Error; err != nil {
		fmt.Printf("update! content failed, err:%v\n", err)
		return err
	}
	return nil
}
