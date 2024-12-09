package dao

import (
	"context"
	"errors"
	"fmt"
	"github.com/kasiforce/trade/pkg/util"
	"github.com/kasiforce/trade/repository/db/model"
	"github.com/kasiforce/trade/types"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
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
            goods.createdTime, users.userName, address.province, address.city, address.districts, address.address,
            COALESCE(COUNT(collection.goodsID), 0) AS star, goods.deliveryMethod, goods.shippingCost`).
		Joins("LEFT JOIN users ON goods.userID = users.userID").
		Joins("LEFT JOIN category ON goods.categoryID = category.categoryID").
		Joins("LEFT JOIN address ON goods.addrID = address.addrID AND address.isDefault = 1").
		Joins("LEFT JOIN collection ON goods.goodsID = collection.goodsID").
		Joins("LEFT JOIN trade_records ON trade_records.goodsID = goods.goodsID").
		Group("goods.goodsID, goods.goodsName, goods.userID, goods.price, category.categoryName, goods.details, goods.isSold, goods.goodsImages, goods.createdTime, users.userName, address.province, address.city, address.districts, address.address")
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
            goods.createdTime, users.userName, address.province, address.city, address.districts, address.address,
            COALESCE(COUNT(collection.goodsID), 0) AS star, goods.deliveryMethod, goods.shippingCost`).
		Joins("LEFT JOIN users ON goods.userID = users.userID").
		Joins("LEFT JOIN category ON goods.categoryID = category.categoryID").
		Joins("LEFT JOIN address ON goods.addrID = address.addrID AND address.isDefault = 1").
		Joins("LEFT JOIN collection ON goods.goodsID = collection.goodsID").
		Joins("LEFT JOIN trade_records ON trade_records.goodsID = goods.goodsID").
		Group("goods.goodsID, goods.goodsName, goods.userID, goods.price, category.categoryName, goods.details, goods.isSold, goods.goodsImages, goods.createdTime, users.userName, address.province, address.city, address.districts, address.address")
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
            goods.createdTime, users.userName, address.province, address.city, address.districts, address.address,
            COALESCE(COUNT(collection.goodsID), 0) AS star, goods.deliveryMethod, goods.shippingCost`).
		Joins("LEFT JOIN users ON goods.userID = users.userID").
		Joins("LEFT JOIN category ON goods.categoryID = category.categoryID").
		Joins("LEFT JOIN address ON goods.addrID = address.addrID AND address.isDefault = 1").
		Joins("LEFT JOIN collection ON goods.goodsID = collection.goodsID").
		Joins("LEFT JOIN trade_records ON trade_records.goodsID = goods.goodsID").
		Group("goods.goodsID, goods.goodsName, goods.userID, goods.price, category.categoryName, goods.details, goods.isSold, goods.goodsImages, goods.createdTime, users.userName, address.province, address.city, address.districts, address.address")

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
            goods.createdTime, users.userName, address.province, address.city, address.districts, address.address,
            COALESCE(COUNT(collection.goodsID), 0) AS star, goods.deliveryMethod, goods.shippingCost`).
		Joins("LEFT JOIN users ON goods.userID = users.userID").
		Joins("LEFT JOIN category ON goods.categoryID = category.categoryID").
		Joins("LEFT JOIN address ON goods.addrID = address.addrID AND address.isDefault = 1").
		Joins("LEFT JOIN collection ON goods.goodsID = collection.goodsID").
		Joins("LEFT JOIN trade_records ON trade_records.goodsID = goods.goodsID").
		Group("goods.goodsID, goods.goodsName, goods.userID, goods.price, category.categoryName, goods.details, goods.isSold, goods.goodsImages, goods.createdTime, users.userName, address.province, address.city, address.districts, address.address")
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
		query = query.Where("goods.shippingCost = ?", req.ShippingCost)
	}
	query = query.Limit(req.Limit).Offset((req.Page - 1) * req.Limit)
	err = query.Find(&goods).Error
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

// 获取商品详情
func (g *Goods) ShowGoodsDetail(req types.ShowDetailReq, userid int) (goods model.Goods, err error) {
	db := g.DB
	// 关联查询 goods, users, address 表
	query := db.Table("goods").
		Select("goods.goodsID, goods.goodsName, goods.userID, goods.price, "+
			"category.categoryName, goods.details, goods.isSold, goods.goodsImages, "+
			"goods.createdTime, users.userName, address.province, address.city, address.districts, address.address,"+
			"COALESCE(COUNT(collection.goodsID), 0) AS star, goods.deliveryMethod, goods.shippingCost,"+
			"users.tel AS tel, "+
			"address.addrID AS addrID, "+
			"CASE WHEN EXISTS (SELECT 1 FROM collection WHERE collection.goodsID = goods.goodsID AND collection.userID = ?) THEN TRUE ELSE FALSE END AS isStarred", userid).
		Joins("LEFT JOIN users ON goods.userID = users.userID").
		Joins("LEFT JOIN category ON goods.categoryID = category.categoryID").
		Joins("LEFT JOIN address ON goods.addrID = address.addrID AND address.isDefault = 1").
		Joins("LEFT JOIN collection ON goods.goodsID = collection.goodsID").
		Joins("LEFT JOIN trade_records ON trade_records.goodsID = goods.goodsID").
		Group("goods.goodsID, goods.goodsName, goods.userID, goods.price, category.categoryName, goods.details, goods.isSold, goods.goodsImages, goods.createdTime, users.userName, address.province, address.city, address.districts, address.address, users.tel, address.addrID")
	if req.GoodsID > 0 {
		query = query.Where("goods.goodsID = ?", req.GoodsID)
	}

	err = query.Scan(&goods).Error
	return
}

// 发布闲置
func (g *Goods) CreateGoods(req types.CreateGoodsReq, userid int) (int, error) {
	db := g.DB
	// 查询 address 表中是否已存在相同地址
	var addrID int
	err := db.Table("address").
		Select("addrID").
		Where("userID = ? AND province = ? AND city = ? AND districts = ? AND address = ?",
			userid, req.Province, req.City, req.District, req.Address).
		Scan(&addrID).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		util.LogrusObj.Error(err)
		return 0, err
	}
	// 如果地址不存在，插入新地址
	if addrID == 0 {
		// 查询用户电话和姓名
		var user struct {
			Tel      string `gorm:"column:tel"`
			UserName string `gorm:"column:userName"`
		}
		err = db.Table("users").
			Select("tel, userName").
			Where("userID = ?", userid).
			Scan(&user).Error
		if err != nil {
			util.LogrusObj.Error("查询用户电话和姓名失败:", err)
			return 0, err
		}
		if user.Tel == "" || user.UserName == "" {
			return 0, errors.New("用户的 tel 或 userName 为空")
		}
		// 插入新地址
		err = db.Exec("INSERT INTO address (userID, province, city, districts, address, tel, receiver, isDefault) VALUES (?, ?, ?, ?, ?, ?, ?, 0)",
			userid, req.Province, req.City, req.District, req.Address, user.Tel, user.UserName).Error
		if err != nil {
			util.LogrusObj.Error("插入新地址失败:", err)
			return 0, err
		}
		// 获取新插入的地址 ID
		createdAddress := model.Address{}
		if err := db.Last(&createdAddress).Error; err != nil {
			util.LogrusObj.Error("获取新地址 ID 失败:", err)
			return 0, err
		}
		addrID = createdAddress.ID // 获取插入的地址 ID
		util.LogrusObj.Info("新地址 ID:", addrID)
	}

	// 转换 deliveryMethod 字段
	var deliveryMethod int
	switch req.DeliveryMethod {
	case "无需快递":
		deliveryMethod = 0
	case "自提":
		deliveryMethod = 1
	case "邮寄":
		deliveryMethod = 2
	}

	// 插入商品记录，并返回插入的商品 ID
	err = db.Exec(
		"INSERT INTO goods (userID, goodsName, details, price, categoryID, goodsImages, createdTime, deliveryMethod, shippingCost, addrID) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		userid, req.GoodsName, req.Details, req.Price, req.CategoryID, req.GoodsImages, time.Now(), deliveryMethod, req.ShippingCost, addrID).Error
	if err != nil {
		util.LogrusObj.Error(err)
		return 0, err
	}

	// 获取新插入的商品 ID
	createdGoods := model.Goods{}
	if err := db.Last(&createdGoods).Error; err != nil {
		util.LogrusObj.Error("获取新商品 ID 失败:", err)
		return 0, err
	}
	goodsID := createdGoods.GoodsID // 获取插入的商品 ID

	util.LogrusObj.Info("新商品 ID:", goodsID)
	return goodsID, nil
}
