package service

import (
	"context"
	"github.com/kasiforce/trade/pkg/util"
	"github.com/kasiforce/trade/repository/db/dao"
	"github.com/kasiforce/trade/types"
	"strconv"
	"sync"
)

var refundServ *RefundService
var refundServOnce sync.Once

type RefundService struct {
}

func GetRefundService() *RefundService {
	refundServOnce.Do(func() {
		refundServ = &RefundService{}
	})
	return refundServ
}

func (s *RefundService) ShowAllRefund(ctx context.Context, req types.ShowRefundReq) (resp interface{}, err error) {
	refund := dao.NewRefundRecord(ctx)
	refundList, err := refund.FindAll(req)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	var respList []types.RefundInfo
	for _, refundInfo := range refundList {
		respList = append(respList, types.RefundInfo{
			RefundID:     refundInfo.RefundID,
			TradeID:      refundInfo.TradeID,
			GoodsName:    refundInfo.GoodsName,
			Price:        refundInfo.Price,
			ShippingCost: refundInfo.ShippingCost,
			SellerName:   refundInfo.SellerName,
			//SellerReason: refundInfo.CReason,
			BuyerName: refundInfo.BuyerName,
			//BuyerReason:  refundInfo.CReason,
			SellerID:     refundInfo.SellerID,
			BuyerID:      refundInfo.BuyerID,
			OrderTime:    refundInfo.OrderTime,
			PayTime:      refundInfo.PayTime,
			RefundTime:   refundInfo.RefundAgreedTime,
			ShippingTime: refundInfo.ShippingTime,
			TurnoverTime: refundInfo.TurnoverTime,
			CStatus:      strconv.Itoa(refundInfo.CStatus),
			BuyerReason:  refundInfo.BuyerReason,
		})
	}
	if respList == nil { // 确保返回空数组而不是 null
		respList = []types.RefundInfo{}
	}
	var response types.RefundListResp
	response.RefundList = respList
	response.PageNum = req.PageNum
	response.Total = len(respList)
	return response, nil
}
