package dao

import (
	"context"
	"errors"
	"github.com/kasiforce/trade/repository/db/model"
	"github.com/kasiforce/trade/types"
	"gorm.io/gorm"
	"strconv"
	"strings"
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
            COALESCE(COUNT(collection.goodsID), 0) AS star, goods.deliveryMethod, 
			trade_records.shippingCost AS shippingCost`).
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
func (g *Goods) FindAll(req types.ShowGoodsListReq) (goods []model.Goods, err error) {
	db := g.DB
	query := db.Table("goods").Select("goods.*")
	if req.SearchQuery != "" {
		query = query.Where("goods.goodsName LIKE ?", "%"+req.SearchQuery+"%")
	}
	if req.Category != "0" {
		// 将 req.Category 转换为 int 类型
		categoryID, err := strconv.Atoi(req.Category)
		if err != nil {
			return nil, err
		}
		// 添加 categoryID 等于 req.Category 的条件
		query = query.Where("goods.categoryID = ?", categoryID)
	}
	// 添加 issold 不等于 1 的条件
	query = query.Where("goods.isSold != ?", 1)
	query = query.Limit(req.Limit).Offset((req.Page - 1) * req.Limit)
	err = query.Find(&goods).Error
	return
}

// 获取商品详情
func (g *Goods) FindByID(id int) (goods []model.Goods, err error) {
	db := g.DB
	// 关联查询 goods, users, address 表
	query := db.Table("goods").
		Select(`goods.goodsID, goods.goodsName, goods.userID, goods.price,
            category.categoryName, goods.details, goods.isSold, goods.goodsImages,
            goods.createdTime, users.userName, address.province, address.city, address.districts,
            COALESCE(COUNT(collection.goodsID), 0) AS star, goods.deliveryMethod,
            trade_records.shippingCost AS shippingCost`).
		Joins("LEFT JOIN users ON goods.userID = users.userID").
		Joins("LEFT JOIN category ON goods.categoryID = category.categoryID").
		Joins("LEFT JOIN address ON goods.userID = address.userID AND address.isDefault = 1").
		Joins("LEFT JOIN collection ON goods.goodsID = collection.goodsID").
		Joins("LEFT JOIN trade_records ON trade_records.goodsID = goods.goodsID").
		Group("goods.goodsID, goods.goodsName, goods.userID, goods.price, category.categoryName, goods.details, goods.isSold, goods.goodsImages, goods.createdTime, users.userName, address.province, address.city, address.districts")
	if id != 0 {
		query = query.Where("goods.userID = ?", id)
	}
	err = query.Find(&goods).Error
	return
}

// 当前用户筛选已售出商品
func (g *Goods) IsSoldGoods(id int) (goods []model.Goods, err error) {
	db := g.DB
	// 只筛选 isSold = 1 的商品
	query := db.Table("goods").
		Select("goods.goodsID, goods.goodsName, goods.userID, goods.price, goods.details, goods.isSold, goods.goodsImages, goods.createdTime").
		Where("goods.isSold = ?", 1)
	// 如果 req.UserID 不为 0，筛选 goods.userID 或 trade_records.sellerID
	if id != 0 {
		// 使用 JOIN 查询 trade_records 表，确保 sellerID 与 req.UserID 匹配
		query = query.Joins("JOIN trade_records t ON t.goodsID = goods.goodsID").
			Where("t.sellerID = ?", id)
	}

	// 执行查询并返回结果
	err = query.Find(&goods).Error
	return
}

// 用户查询自己发布的所有商品
func (g *Goods) UserFindAll(id int) (goods []model.Goods, err error) {
	db := g.DB
	// 关联查询 goods, users, address 表
	query := db.Table("goods").
		Select(`goods.goodsID, goods.goodsName, goods.userID, goods.price, 
            category.categoryName, goods.details, goods.isSold, goods.goodsImages, 
            goods.createdTime, users.userName, address.province, address.city, address.districts,
            COALESCE(COUNT(collection.goodsID), 0) AS star, goods.deliveryMethod,
            trade_records.shippingCost AS shippingCost`).
		Joins("LEFT JOIN users ON goods.userID = users.userID").
		Joins("LEFT JOIN category ON goods.categoryID = category.categoryID").
		Joins("LEFT JOIN address ON goods.userID = address.userID AND address.isDefault = 1").
		Joins("LEFT JOIN collection ON goods.goodsID = collection.goodsID").
		Joins("LEFT JOIN trade_records ON trade_records.goodsID = goods.goodsID").
		Group("goods.goodsID, goods.goodsName, goods.userID, goods.price, category.categoryName, goods.details, goods.isSold, goods.goodsImages, goods.createdTime, users.userName, address.province, address.city, address.districts")

	if id != 0 {
		query = query.Where("goods.userID = ?", id)
	}
	err = query.Scan(&goods).Error
	return
}

// FilterGoods 按条件筛选商品
func (g *Goods) FilterGoods(req types.ShowGoodsReq) (goods []model.Goods, err error) {
	db := g.DB
	query := db.Table("goods").
		Select(`goods.goodsID, goods.goodsName, goods.userID, goods.price, 
            category.categoryName, goods.details, goods.isSold, goods.goodsImages, 
            goods.createdTime, users.userName, address.province, address.city, address.districts,
            COALESCE(COUNT(collection.goodsID), 0) AS star, goods.deliveryMethod, 
			trade_records.shippingCost AS shippingCost`).
		Joins("LEFT JOIN users ON goods.userID = users.userID").
		Joins("LEFT JOIN category ON goods.categoryID = category.categoryID").
		Joins("LEFT JOIN address ON goods.userID = address.userID AND address.isDefault = 1").
		Joins("LEFT JOIN collection ON goods.goodsID = collection.goodsID").
		Joins("LEFT JOIN trade_records ON trade_records.goodsID = goods.goodsID").
		Group("goods.goodsID, goods.goodsName, goods.userID, goods.price, category.categoryName, goods.details, goods.isSold, goods.goodsImages, goods.createdTime, users.userName, address.province, address.city, address.districts")
	if req.SearchQuery != "" {
		query = query.Where("goods.goodsName LIKE ?", "%"+req.SearchQuery+"%")
	}
	if req.PriceMin > 0 {
		query = query.Where("goods.price >= ?", req.PriceMin)
	}
	if req.PriceMax > 0 {
		query = query.Where("goods.price <= ?", req.PriceMax)
	}
	if req.Province != "" {
		query = query.Where("address.province = ?", req.Province)
	}
	if req.City != "" {
		query = query.Where("address.city = ?", req.City)
	}
	if req.District != "" {
		query = query.Where("address.district = ?", req.District)
	}
	if req.DeliveryMethod != "" {
		var deliveryMethod int
		switch req.DeliveryMethod {
		case "0", "无需快递":

			deliveryMethod = 0
		case "1", "自提":

			deliveryMethod = 1
		case "2", "邮寄":
			deliveryMethod = 2
		default:
			return nil, fmt.Errorf("当前配送方式不存在: %s", req.DeliveryMethod)
		}
		query = query.Where("goods.deliveryMethod = ?", deliveryMethod)
	}
	if req.CategoryID > 0 {
		query = query.Where("goods.categoryID = ?", req.CategoryID)
	}
	if req.PublishDate != "" {
		dateRange := strings.Split(req.PublishDate, ",")
		if len(dateRange) == 2 {
			startDate := dateRange[0]
			endDate := dateRange[1]
			query = query.Where("goods.createdTime BETWEEN ? AND ?", startDate, endDate)
		}
	}
	if req.ShippingCost > 0 {
		query = query.Where("trade_records.shippingCost = ?", req.ShippingCost)
	}
	query = query.Limit(req.Limit).Offset((req.Page - 1) * req.Limit)
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
	result := g.DB.Delete(&model.Goods{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("商品不存在")
	}
	return nil
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

// 更新view
func (g *Goods) IncreaseView(goodsID uint) error {
	return g.DB.Model(&model.Goods{}).Where("goodsID = ?", goodsID).UpdateColumn("view", gorm.Expr("view + 1")).Error
}
