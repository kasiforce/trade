package dao

import (
	"context"
	"github.com/kasiforce/trade/repository/db/model"
	"github.com/kasiforce/trade/types"
	"gorm.io/gorm"
)

type Collection struct {
	*gorm.DB
}

func NewCollectionByDB(db *gorm.DB) *Collection {
	return &Collection{db}
}

func NewCollection(ctx context.Context) *Collection {
	return &Collection{NewDBClient(ctx)}
}

func (c *Collection) CountAll(id int) (total int64, err error) {
	err = c.DB.Model(&model.Collection{}).Where("userID=?", id).Count(&total).Error
	return
}

func (c *Collection) FindByID(req types.ShowCollectionReq, id int) (collection []model.Collection, err error) {
	db := c.DB
	query := db.Table("collection").Preload("Goods").Where("userID = ?", id)
	query = query.Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize)
	err = query.Find(&collection).Error
	return
}
