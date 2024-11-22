package dao

import (
	"context"
	"github.com/kasiforce/trade/repository/db/model"
	"github.com/kasiforce/trade/types"
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

// FindAll 查询所有商品
func (g *Goods) FindAll(req types.ShowGoodsReq) (goods []model.Goods, err error) {
	db := g.DB
	query := db.Table("goods").Select("goods.*")
	if req.SearchQuery != "" {
		query = query.Where("goods.goodsName LIKE ?", "%"+req.SearchQuery+"%")
	}
	query = query.Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize)
	err = query.Find(&goods).Error
	return
}

// FindByID 获取商品详情
func (g *Goods) FindByID(id int) (good *model.Goods, err error) {
	err = g.DB.Model(&model.Goods{}).Where("goodsID = ?", id).First(&good).Error
	return
}

// FilterGoods 筛选商品
func (g *Goods) FilterGoods(filter map[string]interface{}, req types.ShowGoodsReq) (goods []model.Goods, err error) {
	db := g.DB
	query := db.Table("goods").Select("goods.*")
	for key, value := range filter {
		query = query.Where(key+" = ?", value)
	}
	query = query.Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize)
	err = query.Find(&goods).Error
	return
}

// CreateGoods 创建商品
func (g *Goods) CreateGoods(good map[string]interface{}) (err error) {
	err = g.DB.Model(&model.Goods{}).Create(good).Error
	return
}

// DeleteGoods 删除商品
func (g *Goods) DeleteGoods(id int) (err error) {
	err = g.DB.Model(&model.Goods{}).Delete(&model.Goods{}, id).Error
	return
}
