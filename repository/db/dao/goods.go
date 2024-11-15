package dao

import (
	"context"
	"github.com/kasiforce/trade/repository/db/model"
	"gorm.io/gorm"
)

type Goods struct {
	*gorm.DB
}

func NewGoodsByDB(db *gorm.DB) *Goods {
	return &Goods{db}
}

func NewGoods(ctx context.Context) *Goods {
	return &Goods{NewDBClient(ctx)}
}

func (g *Goods) FindAll() (goods []*model.Goods, err error) {
	err = g.DB.Model(&model.Goods{}).Find(&goods).Error
	return
}

func (g *Goods) FindByID(id int) (good *model.Goods, err error) {
	err = g.DB.Model(&model.Goods{}).Where("goodsID = ?", id).First(&good).Error
	return
}

func (g *Goods) CreateGoods(good *model.Goods) (err error) {
	err = g.DB.Model(&model.Goods{}).Create(&good).Error
	return
}

func (g *Goods) UpdateGoods(id int, good *model.Goods) (err error) {
	err = g.DB.Model(&model.Goods{}).Where("goodsID = ?", id).Updates(good).Error
	return
}

func (g *Goods) DeleteGoods(id int) (err error) {
	err = g.DB.Model(&model.Goods{}).Where("goodsID = ?", id).Delete(&model.Goods{}).Error
	return
}
