package dao

import (
	"context"
	"github.com/kasiforce/trade/repository/db/model"
	"github.com/kasiforce/trade/types"
	"gorm.io/gorm"
)

type Category struct {
	*gorm.DB
}

func NewCategoryByDB(db *gorm.DB) *Category {
	return &Category{db}
}

func NewCategory(c context.Context) *Category {
	return &Category{NewDBClient(c)}
}

func (c *Category) CountAll() (total int64, err error) {
	err = c.DB.Model(&model.Category{}).Count(&total).Error
	return
}

func (c *Category) FindAll(req types.ShowCategoryReq) (ca []*model.Category, err error) {
	db := c.DB
	//db = db.Table("category").Select("category.*")
	if req.SearchQuery != "" {
		db = db.Where("categoryName LIKE ?", "%"+req.SearchQuery+"%")
	}
	db = db.Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize)
	err = db.Find(&ca).Error
	return
}

func (c *Category) ShowCategory() (ca []*model.Category, err error) {
	err = c.DB.Model(&model.Category{}).Find(&ca).Error
	return
}

func (c *Category) AddCategory(category *model.Category) (id int, err error) {
	result := c.DB.Model(&model.Category{}).Create(&category)
	if err = result.Error; err != nil {
		return 0, err
	}
	id = category.CategoryID
	return
}

func (c *Category) UpdateCategory(id int, category *model.Category) (err error) {
	err = c.DB.Model(&model.Category{}).Where("categoryID = ?", id).Updates(&category).Error
	return
}

func (c *Category) DeleteCategory(id int) (err error) {
	err = c.DB.Model(&model.Category{}).Delete(&model.Category{}, id).Error
	return
}
