package dao

import (
	"context"
	//"errors"
	//"fmt"
	"github.com/kasiforce/trade/repository/db/model"
	"github.com/kasiforce/trade/types"
	"gorm.io/gorm"
	"time"
)

type TradeRecords struct {
	*gorm.DB
}

func NewTradeRecordsByDB(db *gorm.DB) *TradeRecords {
	return &TradeRecords{db}
}

func NewTradeRecords(ctx context.Context) *TradeRecords {
	return &TradeRecords{NewDBClient(ctx)}
}

//	func (tr *TradeRecords) FindAll() (tradeRecords []*model.TradeRecords, err error) {
//		err = tr.DB.Model(&model.TradeRecords{}).Find(&tradeRecords).Error
//		return
//	}
//
//	func (tr *TradeRecords) FindByID(id int) (t *model.TradeRecords, err error) {
//		err = tr.DB.Model(&model.TradeRecords{}).Where("tradeID = ?", id).First(&t).Error
//		return
//	}
//
//	func (tr *TradeRecords) FindBySellerID(sellerID int) (tradeRecords []*model.TradeRecords, err error) {
//		err = tr.DB.Model(&model.TradeRecords{}).Where("sellerID = ?", sellerID).Find(&tradeRecords).Error
//		return
//	}
//
//	func (tr *TradeRecords) FindByBuyerID(buyerID int) (tradeRecords []*model.TradeRecords, err error) {
//		err = tr.DB.Model(&model.TradeRecords{}).Where("buyerID = ?", buyerID).Find(&tradeRecords).Error
//		return
//	}
//
//	func (tr *TradeRecords) FindByGoodsID(goodsID int) (tradeRecords []*model.TradeRecords, err error) {
//		err = tr.DB.Model(&model.TradeRecords{}).Where("goodsID = ?", goodsID).Find(&tradeRecords).Error
//		return
//	}
//
//	func (tr *TradeRecords) CreateTradeRecord(t *model.TradeRecords) (err error) {
//		err = tr.DB.Model(&model.TradeRecords{}).Create(&t).Error
//		return
//	}
//
//	func (tr *TradeRecords) UpdateTradeRecord(id int, t *model.TradeRecords) (err error) {
//		err = tr.DB.Model(&model.TradeRecords{}).Where("tradeID = ?", id).Updates(&t).Error
//		return
//	}
//
//	func (tr *TradeRecords) DeleteTradeRecord(id int) (err error) {
//		err = tr.DB.Model(&model.TradeRecords{}).Where("tradeID = ?", id).Delete(&model.TradeRecords{}).Error
//		return
//	}
//
// GetAllOrders 获取所有订单
func (c *TradeRecords) GetAllOrders(req types.ShowOrdersReq) (r []types.OrderInfo, total int64, err error) {
	query := c.DB.Model(&model.TradeRecords{})

	if req.SearchQuery != "" {
		query = query.Where("tradeID = ?", req.SearchQuery)
	}

	err = query.Count(&total).Error
	if err != nil {
		return
	}

	var orders []struct {
		TradeID            int
		SellerID           int
		BuyerID            int
		SellerName         string
		BuyerName          string
		GoodsID            int
		GoodsName          string
		Price              float64
		DeliveryMethod     string
		ShippingCost       float64
		ShippingProvince   string
		ShippingCity       string
		ShippingArea       string
		ShippingDetailArea string
		DeliveryProvince   string
		DeliveryCity       string
		DeliveryArea       string
		DeliveryDetailArea string
		OrderTime          time.Time
		PayTime            time.Time
		ShippingTime       time.Time
		TurnoverTime       time.Time
		Status             string
	}

	err = query.
		Joins("left join users as seller on seller.userID = trade_records.sellerID").
		Joins("left join users as buyer on buyer.userID = trade_records.buyerID").
		Joins("left join goods on goods.goodsID = trade_records.goodsID").
		Joins("left join address as shippingAddr on shippingAddr.addrID = trade_records.shippingAddrID").
		Joins("left join address as deliveryAddr on deliveryAddr.addrID = trade_records.deliveryAddrID").
		Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).
		Select("trade_records.tradeID as TradeID," +
			"trade_records.sellerID as SellerID," +
			"trade_records.buyerID as BuyerID," +
			"seller.userName as SellerName," +
			"buyer.userName as BuyerName," +
			"trade_records.goodsID as GoodsID," +
			"goods.goodsName as GoodsName," +
			"trade_records.turnoverAmount as Price," +
			"trade_records.payMethod as DeliveryMethod," +
			"trade_records.shippingCost as ShippingCost," +
			"shippingAddr.province as ShippingProvince," +
			"shippingAddr.city as ShippingCity," +
			"shippingAddr.districts as ShippingArea," +
			"shippingAddr.address as ShippingDetailArea," +
			"deliveryAddr.province as DeliveryProvince," +
			"deliveryAddr.city as DeliveryCity," +
			"deliveryAddr.districts as DeliveryArea," +
			"deliveryAddr.address as DeliveryDetailArea," +
			"trade_records.orderTime as OrderTime," +
			"trade_records.payTime as PayTime," +
			"trade_records.shippingTime as ShippingTime," +
			"trade_records.turnoverTime as TurnoverTime," +
			"trade_records.status as Status").
		Scan(&orders).Error

	if err != nil {
		return
	}

	for _, order := range orders {
		r = append(r, types.OrderInfo{
			TradeID:        order.TradeID,
			SellerID:       order.SellerID,
			BuyerID:        order.BuyerID,
			SellerName:     order.SellerName,
			BuyerName:      order.BuyerName,
			GoodsID:        order.GoodsID,
			GoodsName:      order.GoodsName,
			Price:          order.Price,
			DeliveryMethod: order.DeliveryMethod,
			ShippingCost:   order.ShippingCost,
			SenderAddress: types.AddressDetail{
				Province:   order.DeliveryProvince,
				City:       order.DeliveryCity,
				Area:       order.DeliveryArea,
				DetailArea: order.DeliveryDetailArea,
			},
			ShippingAddress: types.AddressDetail{
				Province:   order.ShippingProvince,
				City:       order.ShippingCity,
				Area:       order.ShippingArea,
				DetailArea: order.ShippingDetailArea,
			},
			OrderTime:    order.OrderTime,
			PayTime:      order.PayTime,
			ShippingTime: order.ShippingTime,
			TurnoverTime: order.TurnoverTime,
			Status:       order.Status,
		})
	}

	return
}

// UpdateOrderStatus 修改订单状态
func (c *TradeRecords) UpdateOrderStatus(req types.UpdateOrderStatusReq) (resp interface{}, err error) {
	// 加载北京时区
	location := time.FixedZone("Asia/Shanghai", 8*60*60)

	// 更新订单状态
	if req.Status == "未发货" {
		updateData := map[string]interface{}{
			"status":  req.Status,
			"payTime": time.Now().In(location),
		}
		err = c.DB.Model(&model.TradeRecords{}).Where("tradeID = ? AND payTime IS NULL ", req.ID).Updates(updateData).
			Error
	} else if req.Status == "已发货" {
		updateData := map[string]interface{}{
			"status":       req.Status,
			"shippingTime": time.Now().In(location),
		}
		err = c.DB.Model(&model.TradeRecords{}).Where("tradeID = ? AND ShippingTime IS NULL ", req.ID).Updates(updateData).
			Error
	} else if req.Status == "交易完成" || req.Status == "已取消" || req.Status == "已退款" {
		updateData := map[string]interface{}{
			"status":       req.Status,
			"turnoverTime": time.Now().In(location),
		}
		err = c.DB.Model(&model.TradeRecords{}).Where("tradeID = ?", req.ID).Updates(updateData).
			Error
	} else {
		err = c.DB.Model(&model.TradeRecords{}).Where("tradeID = ?", req.ID).Update("status", req.Status).
			Error
	}

	if err != nil {
		return
	}

	// 如果存在退款理由，插入退款申诉
	if req.RefundReason != "" {
		refundComplaint := model.RefundComplaint{
			TradeID:     req.ID,
			BuyerReason: req.RefundReason,
			CTime:       time.Now(),
			CStatus:     0,
		}
		err = c.DB.Create(&refundComplaint).Error
		if err != nil {
			return
		}
	}
	//rejectReason如果存在拒绝退款理由，插入退款申诉
	if req.RejectReason != "" {
		updateData := map[string]interface{}{
			"sellerReason": req.RejectReason,
		}
		err = c.DB.Model(&model.RefundComplaint{}).Where("tradeID = ?", req.ID).Updates(updateData).Error
		if err != nil {
			return
		}
	}

	// 如果存在评价内容，创建评论
	if req.Comment != "" {
		var tradeRecord model.TradeRecords
		err = c.DB.Where("tradeID = ?", req.ID).First(&tradeRecord).Error
		if err != nil {
			return
		}

		comment := model.Comment{
			GoodsID:        tradeRecord.GoodsID,
			CommentatorID:  tradeRecord.BuyerID,
			CommentContent: req.Comment,
			CommentTime:    time.Now().In(location),
		}
		err = c.DB.Create(&comment).Error
		if err != nil {
			return
		}
	}

	resp = types.UpdateOrderStatusResp{
		Status: req.Status,
	}
	return
}

// UpdateOrderAddress 修改订单地址
func (c *TradeRecords) UpdateOrderAddress(req types.UpdateOrderAddressReq) (err error) {
	// 更新订单地址
	err = c.DB.Model(&model.TradeRecords{}).Where("tradeID = ?", req.ID).Updates(map[string]interface{}{
		//"deliveryAddrID": req.Province + req.City + req.Area + req.DetailArea,
		"deliveryAddrID": req.AddrID,
	}).Error
	if err != nil {
		return
	}

	return
}

// CreateOrder 生成订单
func (c *TradeRecords) CreateOrder(req types.CreateOrderReq, id int) (resp interface{}, err error) {

	// 创建订单
	order := model.TradeRecords{
		SellerID:       req.SellerID,
		GoodsID:        req.GoodsID,
		TurnoverAmount: req.Price,
		ShippingCost:   req.ShippingCost,
		OrderTime:      time.Now(),
		Status:         "未付款",
		BuyerID:        id,
	}
	// 处理 ShippingAddrID
	if req.SenderAddrID != 0 {
		order.ShippingAddrID = &req.SenderAddrID // 设置为实际值
	} else {
		order.ShippingAddrID = nil // 设置为 NULL
	}

	// 处理 DeliveryAddrID
	if req.ShippingAddrID != 0 {
		order.DeliveryAddrID = &req.ShippingAddrID // 设置为实际值
	} else {
		order.DeliveryAddrID = nil // 设置为 NULL
	}

	switch req.DeliveryMethod {
	case "无需快递":
		order.PayMethod = 0
	case "自提":
		order.PayMethod = 1
	case "邮寄":
		order.PayMethod = 2
	}
	err = c.DB.Create(&order).Error
	if err != nil {
		return
	}

	resp = types.CreateOrderResp{
		TradeID: order.TradeID,
	}

	return resp, nil
}

// GetMyOrdersPurchased 获取我买到的订单
func (c *TradeRecords) GetMyOrdersPurchased(req types.GetMyOrdersReq, id int) (r []types.GetMyOrderInfo, total int64, err error) {
	query := c.DB.Model(&model.TradeRecords{}).
		Joins("left join users as seller on seller.userID = trade_records.sellerID").
		Joins("left join goods on goods.goodsID = trade_records.goodsID").
		Joins("left join address as shippingAddr on shippingAddr.addrID = trade_records.shippingAddrID").
		Joins("left join address as deliveryAddr on deliveryAddr.addrID = trade_records.deliveryAddrID").
		Where("trade_records.buyerID = ?", id)

	err = query.Count(&total).Error
	if err != nil {
		return
	}

	var orders []struct {
		TradeID            int
		SellerID           int
		SellerName         string
		GoodsID            int
		GoodsName          string
		Price              float64
		DeliveryMethod     string
		ShippingCost       float64
		ShippingProvince   string
		ShippingCity       string
		ShippingArea       string
		ShippingDetailArea string
		DeliveryProvince   string
		DeliveryCity       string
		DeliveryArea       string
		DeliveryDetailArea string
		OrderTime          time.Time
		PayTime            time.Time
		ShippingTime       time.Time
		TurnoverTime       time.Time
		Status             string
		ShippingTel        string
		ShippingName       string
		DeliveryTel        string
		DeliveryName       string
	}

	err = query.Offset((req.Page - 1) * req.PageSize).Limit(req.PageSize).
		Select("trade_records.tradeID as TradeID," +
			"trade_records.sellerID as SellerID," +
			"seller.userName as SellerName," +
			"trade_records.goodsID as GoodsID," +
			"goods.goodsName as GoodsName," +
			"trade_records.turnoverAmount as Price," +
			"trade_records.payMethod as DeliveryMethod," +
			"trade_records.shippingCost as ShippingCost," +
			"shippingAddr.province as ShippingProvince," +
			"shippingAddr.city as ShippingCity," +
			"shippingAddr.districts as ShippingArea," +
			"shippingAddr.address as ShippingDetailArea," +
			"shippingAddr.tel as ShippingTel," +
			"shippingAddr.receiver as ShippingName," +
			"deliveryAddr.province as DeliveryProvince," +
			"deliveryAddr.city as DeliveryCity," +
			"deliveryAddr.districts as DeliveryArea," +
			"deliveryAddr.address as DeliveryDetailArea," +
			"deliveryAddr.tel as DeliveryTel," +
			"deliveryAddr.receiver as DeliveryName," +
			"trade_records.orderTime as OrderTime," +
			"trade_records.payTime as PayTime," +
			"trade_records.shippingTime as ShippingTime," +
			"trade_records.turnoverTime as TurnoverTime," +
			"trade_records.status as Status").
		Scan(&orders).Error

	if err != nil {
		return
	}

	for _, order := range orders {
		r = append(r, types.GetMyOrderInfo{
			TradeID:        order.TradeID,
			SellerID:       order.SellerID,
			SellerName:     order.SellerName,
			GoodsID:        order.GoodsID,
			GoodsName:      order.GoodsName,
			Price:          order.Price,
			DeliveryMethod: order.DeliveryMethod,
			ShippingCost:   order.ShippingCost,
			SenderAddress: types.AddressInfo{
				Province:   order.DeliveryProvince,
				City:       order.DeliveryCity,
				Area:       order.DeliveryArea,
				DetailArea: order.DeliveryDetailArea,
				Tel:        order.DeliveryTel,
				Name:       order.DeliveryName,
			},
			ShippingAddress: types.AddressInfo{
				Province:   order.ShippingProvince,
				City:       order.ShippingCity,
				Area:       order.ShippingArea,
				DetailArea: order.ShippingDetailArea,
				Tel:        order.ShippingTel,
				Name:       order.ShippingName,
			},
			OrderTime:    order.OrderTime.Format("2006-01-02 15:04:05"),
			PayTime:      order.PayTime.Format("2006-01-02 15:04:05"),
			ShippingTime: order.ShippingTime.Format("2006-01-02 15:04:05"),
			TurnoverTime: order.TurnoverTime.Format("2006-01-02 15:04:05"),
			Status:       order.Status,
		})
	}

	return
}

// GetMySoldOrders 获取我卖出的订单
func (c *TradeRecords) GetMySoldOrders(req types.GetMyOrdersReq, id int) (r []types.GetMyOrderInfo, total int64, err error) {
	query := c.DB.Model(&model.TradeRecords{}).
		Joins("left join users as buyer on buyer.userID = trade_records.buyerID").
		Joins("left join goods on goods.goodsID = trade_records.goodsID").
		Joins("left join address as shippingAddr on shippingAddr.addrID = trade_records.shippingAddrID").
		Joins("left join address as deliveryAddr on deliveryAddr.addrID = trade_records.deliveryAddrID").
		Where("trade_records.sellerID = ?", id)

	err = query.Count(&total).Error
	if err != nil {
		return
	}

	var orders []struct {
		TradeID            int
		BuyerID            int
		BuyerName          string
		GoodsID            int
		GoodsName          string
		Price              float64
		DeliveryMethod     string
		ShippingCost       float64
		ShippingProvince   string
		ShippingCity       string
		ShippingArea       string
		ShippingDetailArea string
		DeliveryProvince   string
		DeliveryCity       string
		DeliveryArea       string
		DeliveryDetailArea string
		OrderTime          time.Time
		PayTime            time.Time
		ShippingTime       time.Time
		TurnoverTime       time.Time
		Status             string
		ShippingTel        string
		ShippingName       string
		DeliveryTel        string
		DeliveryName       string
	}

	err = query.Offset((req.Page - 1) * req.PageSize).Limit(req.PageSize).
		Select("trade_records.tradeID as TradeID," +
			"trade_records.buyerID as BuyerID," +
			"buyer.userName as BuyerName," +
			"trade_records.goodsID as GoodsID," +
			"goods.goodsName as GoodsName," +
			"trade_records.turnoverAmount as Price," +
			"trade_records.payMethod as DeliveryMethod," +
			"trade_records.shippingCost as ShippingCost," +
			"shippingAddr.province as ShippingProvince," +
			"shippingAddr.city as ShippingCity," +
			"shippingAddr.districts as ShippingArea," +
			"shippingAddr.address as ShippingDetailArea," +
			"shippingAddr.tel as ShippingTel," +
			"shippingAddr.receiver as ShippingName," +
			"deliveryAddr.province as DeliveryProvince," +
			"deliveryAddr.city as DeliveryCity," +
			"deliveryAddr.districts as DeliveryArea," +
			"deliveryAddr.address as DeliveryDetailArea," +
			"deliveryAddr.tel as DeliveryTel," +
			"deliveryAddr.receiver as DeliveryName," +
			"trade_records.orderTime as OrderTime," +
			"trade_records.payTime as PayTime," +
			"trade_records.shippingTime as ShippingTime," +
			"trade_records.turnoverTime as TurnoverTime," +
			"trade_records.status as Status").
		Scan(&orders).Error

	if err != nil {
		return
	}

	for _, order := range orders {
		r = append(r, types.GetMyOrderInfo{
			TradeID:        order.TradeID,
			SellerID:       order.BuyerID,
			SellerName:     order.BuyerName,
			GoodsID:        order.GoodsID,
			GoodsName:      order.GoodsName,
			Price:          order.Price,
			DeliveryMethod: order.DeliveryMethod,
			ShippingCost:   order.ShippingCost,
			SenderAddress: types.AddressInfo{
				Province:   order.DeliveryProvince,
				City:       order.DeliveryCity,
				Area:       order.DeliveryArea,
				DetailArea: order.DeliveryDetailArea,
				Tel:        order.DeliveryTel,
				Name:       order.DeliveryName,
			},
			ShippingAddress: types.AddressInfo{
				Province:   order.ShippingProvince,
				City:       order.ShippingCity,
				Area:       order.ShippingArea,
				DetailArea: order.ShippingDetailArea,
				Tel:        order.ShippingTel,
				Name:       order.ShippingName,
			},
			OrderTime:    order.OrderTime.Format("2006-01-02 15:04:05"),
			PayTime:      order.PayTime.Format("2006-01-02 15:04:05"),
			ShippingTime: order.ShippingTime.Format("2006-01-02 15:04:05"),
			TurnoverTime: order.TurnoverTime.Format("2006-01-02 15:04:05"),
			Status:       order.Status,
		})
	}

	return
}

// GetOrderDetail 获取订单详情
func (c *TradeRecords) GetOrderDetail(orderID int) (model.TradeRecords, error) {
	var order model.TradeRecords
	err := c.DB.Where("tradeID = ?", orderID).First(&order).Error
	return order, err
}
