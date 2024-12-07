package service

import (
	"context"
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
