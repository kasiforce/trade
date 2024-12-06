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

// NewGoodsByDB 使用指定的数据库连接创建 Goods 实例
func NewGoodsByDB(db *gorm.DB) *Goods {
	return &Goods{db}
}

// NewGoods 使用默认数据库上下文创建 Goods 实例
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

// FilterGoods 按条件筛选商品
func (g *Goods) FilterGoods(filter map[string]interface{}, req types.ShowGoodsReq) (goods []model.Goods, err error) {
	db := g.DB
	query := db.Table("goods").Select("goods.*")
	for key, value := range filter {
		query = query.Where(key+" = ?", value)
	}
	if req.SearchQuery != "" {
		query = query.Where("goods.goodsName LIKE ?", "%"+req.SearchQuery+"%")
	}
	if req.PriceMin > 0 {
		query = query.Where("goods.price >= ?", req.PriceMin)
	}
	if req.PriceMax > 0 {
		query = query.Where("goods.price <= ?", req.PriceMax)
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

// FindByCategoryID 根据分类ID查询商品
func (g *Goods) FindByCategoryID(categoryID, pageNum, pageSize int) (goods []model.Goods, total int64, err error) {
	db := g.DB
	query := db.Table("goods").Where("categoryID = ?", categoryID)
	err = query.Count(&total).Error // 计算总数
	if err != nil {
		return
	}
	err = query.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&goods).Error
	return
}

// AdvancedFilterGoods 综合筛选商品
func (g *Goods) AdvancedFilterGoods(req types.ShowGoodsReq) (goods []model.Goods, total int64, err error) {
	db := g.DB
	query := db.Table("goods")

	// 商品名称模糊查询
	if req.SearchQuery != "" {
		query = query.Where("goodsName LIKE ?", "%"+req.SearchQuery+"%")
	}

	// 分类筛选
	if req.CategoryID > 0 {
		query = query.Where("categoryID = ?", req.CategoryID)
	}

	// 价格区间筛选
	if req.PriceMin > 0 {
		query = query.Where("price >= ?", req.PriceMin)
	}
	if req.PriceMax > 0 {
		query = query.Where("price <= ?", req.PriceMax)
	}

	// 是否已售筛选
	if req.IsSold == 0 || req.IsSold == 1 {
		query = query.Where("isSold = ?", req.IsSold)
	}

	// 统计总记录数
	err = query.Count(&total).Error
	if err != nil {
		return
	}

	// 分页查询
	err = query.Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Find(&goods).Error
	return
}
