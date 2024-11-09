package dao

import (
	"context"
	"github.com/kasiforce/trade/repository/db/model"
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

func (c *Category) FindAll() (ca []*model.Category, err error) {
	err = c.DB.Model(&model.Category{}).Find(&ca).Error
	return
}
