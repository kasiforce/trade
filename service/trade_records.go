package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"sync"
	"time"

	//"github.com/gin-gonic/gin"
	"github.com/kasiforce/trade/pkg/util"
	"github.com/kasiforce/trade/repository/db/dao"
	"github.com/kasiforce/trade/types"
)

var trade_recordsServ *Trade_recordsService
var trade_recordsServOnce sync.Once

type Trade_recordsService struct {
}

func GetTrade_recordsService() *Trade_recordsService {
	trade_recordsServOnce.Do(func() {
		trade_recordsServ = &Trade_recordsService{}
	})
	return trade_recordsServ
}

// GetAllOrders 获取所有订单
func (s *Trade_recordsService) GetAllOrders(ctx context.Context, req types.ShowOrdersReq) (resp interface{}, err error) {
	u := dao.NewTradeRecords(ctx)
	orders, total, err := u.GetAllOrders(req)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	resp = &types.OrderListResp{
		OrderList: orders,
		Total:     total,
		PageNum:   req.PageNum,
	}
	return
}

// UpdateOrderStatus 修改订单状态
func (s *Trade_recordsService) UpdateOrderStatus(ctx context.Context, req types.UpdateOrderStatusReq) (resp interface{}, err error) {
	u := dao.NewTradeRecords(ctx)
	resp, err = u.UpdateOrderStatus(req)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return
}

// UpdateOrderAddress 修改订单地址
func (s *Trade_recordsService) UpdateOrderAddress(ctx context.Context, req types.UpdateOrderAddressReq) (resp interface{}, err error) {
	u := dao.NewTradeRecords(ctx)
	err = u.UpdateOrderAddress(req)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	// 返回一个对象类型的响应
	resp = map[string]string{
		"message": "OrderAddress updated successfully",
	}
	return resp, nil
}

// CreateOrder 生成订单
func (s *Trade_recordsService) CreateOrder(ctx *gin.Context, req types.CreateOrderReq) (resp interface{}, err error) {
	id := ctx.GetInt("id")
	u := dao.NewTradeRecords(ctx)
	resp, err = u.CreateOrder(req, id)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return
}

// GetMyOrders 获取我买到的订单
func (s *Trade_recordsService) GetMyOrders(ctx *gin.Context, req types.GetMyOrdersReq) (resp interface{}, err error) {
	id := ctx.GetInt("id")
	u := dao.NewTradeRecords(ctx)
	orders, total, err := u.GetMyOrdersPurchased(req, id)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	resp = &types.GetMyOrdersResp{
		Total:     total,
		OrderList: orders,
	}
	return
}

// GetMySoldOrders 获取我卖出的订单
func (s *Trade_recordsService) GetMySoldOrders(ctx *gin.Context, req types.GetMyOrdersReq) (resp interface{}, err error) {
	id := ctx.GetInt("id")
	u := dao.NewTradeRecords(ctx)
	orders, total, err := u.GetMySoldOrders(req, id)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	resp = &types.GetMyOrdersResp{
		Total:     total,
		OrderList: orders,
	}
	return
}

// GetOrderDetail 获取订单详情
func (s *Trade_recordsService) GetOrderDetail(ctx *gin.Context, req types.GetOrderDetailReq) (resp types.GetOrderDetailResp, err error) {
	// 获取订单ID
	orderID := req.ID

	// 查询订单详情
	u := dao.NewTradeRecords(ctx)
	order, err := u.GetOrderDetail(orderID)
	if err != nil {
		util.LogrusObj.Error(err)
		return resp, err
	}

	// 计算倒计时
	countdown := calculateCountdown(order.OrderTime)

	resp = types.GetOrderDetailResp{
		ID:             order.TradeID,
		Countdown:      countdown,
		SellerID:       order.SellerID,
		GoodsID:        order.GoodsID,
		Price:          order.TurnoverAmount,
		DeliveryMethod: getDeliveryMethod(order.PayMethod),
		ShippingCost:   order.ShippingCost,
	}
	// 检查 ShippingAddrID 是否为 nil
	if order.ShippingAddrID != nil {
		resp.SenderAddrID = *order.ShippingAddrID
	} else {
		resp.SenderAddrID = 0
	}
	// 检查 DeliveryAddrID 是否为 nil
	if order.DeliveryAddrID != nil {
		resp.ShippingAddrID = *order.DeliveryAddrID
	} else {
		resp.ShippingAddrID = 0
	}

	return resp, nil
}

// calculateCountdown 计算倒计时
func calculateCountdown(orderTime time.Time) int {
	const countdownDuration = 300 // 倒计时总时长为300秒
	elapsedTime := time.Since(orderTime).Seconds()
	remainingTime := countdownDuration - int(elapsedTime)
	if remainingTime < 0 {
		return -1
	}
	return remainingTime
}

// getDeliveryMethod 根据支付方式获取交易方式
func getDeliveryMethod(payMethod int) string {
	switch payMethod {
	case 0:
		return "无需快递"
	case 1:
		return "自提"
	case 2:
		return "邮寄"
	default:
		return "未知"
	}
}

func (s *Trade_recordsService) PaySuccess(ctx context.Context, req types.PaySuccessReq) (resp interface{}, err error) {
	u := dao.NewTradeRecords(ctx)
	err = u.UpdateOrderStatusToUnshipped(req)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	resp = map[string]string{
		"message": "支付成功，订单状态已更新为未发货",
	}
	return
}
