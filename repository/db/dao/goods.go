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

// 管理员端查询所有商品
func (g *Goods) AdminFindAll(req types.ShowAllGoodsReq) (goods []model.Goods, err error) {
	db := g.DB
	// 关联查询 goods, users, address 表
	query := db.Table("goods").
		Select(`goods.goodsID, goods.goodsName, goods.userID, goods.price, 
            category.categoryName, goods.details, goods.isSold, goods.goodsImages, 
            goods.createdTime, users.userName, address.province, address.city, address.districts,
            COALESCE(COUNT(collection.goodsID), 0) AS collection,
            GROUP_CONCAT(DISTINCT trade_records.payMethod) AS payMethod`).
		Joins("LEFT JOIN users ON goods.userID = users.userID").
		Joins("LEFT JOIN category ON goods.categoryID = category.categoryID").
		Joins("LEFT JOIN address ON goods.userID = address.userID AND address.isDefault = 1").
		Joins("LEFT JOIN collection ON goods.goodsID = collection.goodsID").
		Joins("LEFT JOIN trade_records ON trade_records.goodsID = goods.goodsID").
		Group("goods.goodsID, goods.goodsName, goods.userID, goods.price, category.categoryName, goods.details, goods.isSold, goods.goodsImages, goods.createdTime, users.userName, address.province, address.city, address.districts")
	if req.SearchQuery != "" {
		query = query.Where("trade_records.tradeID LIKE ?", "%"+req.SearchQuery+"%")
	}

	query = query.Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize)
	err = query.Scan(&goods).Error
	return
}

// FindAll 查询所有商品
func (g *Goods) FindAll(req types.ShowAllGoodsReq) (goods []model.Goods, err error) {
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

// 当前用户筛选已售出商品
func (g *Goods) IsSoldGoods(req types.IsSoldGoodsResp) (goods []model.Goods, err error) {
	db := g.DB
	// 只筛选 isSold = 1 的商品
	query := db.Table("goods").
		Select("goods.goodsID, goods.goodsName, goods.userID, goods.price, goods.details, goods.isSold, goods.goodsImages, goods.createdTime").
		Where("goods.isSold = ?", 1)
	// 如果 req.UserID 不为 0，筛选 goods.userID 或 trade_records.sellerID
	if req.UserID != 0 {
		// 使用 JOIN 查询 trade_records 表，确保 sellerID 与 req.UserID 匹配
		query = query.Joins("JOIN trade_records t ON t.goodsID = goods.goodsID").
			Where("t.sellerID = ?", req.UserID)
	}

	// 执行查询并返回结果
	err = query.Find(&goods).Error
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
