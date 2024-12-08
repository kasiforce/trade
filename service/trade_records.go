package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"sync"
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
