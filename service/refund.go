package service

import (
	"context"
	"fmt"
	"github.com/kasiforce/trade/pkg/util"
	"github.com/kasiforce/trade/repository/db/dao"
	"github.com/kasiforce/trade/types"
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
	refund := dao.NewRefundComplaint(ctx)
	// 获取数据库中所有待退货的商品
	total, err := refund.CountAll()
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}
	refundList, err := refund.FindAll(req)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	var respList []types.RefundInfo
	for _, refundInfo := range refundList {
		// 将 DeliveryMethod 从 int 转换为 string
		var status string
		switch refundInfo.CStatus {
		case 0:
			status = "未处理"
		case 1:
			status = "同意退货"
		case 2:
			status = "拒绝退货"
		default:
			status = "未知"
		}
		respList = append(respList, types.RefundInfo{
			TradeID:      refundInfo.TradeID,
			GoodsName:    refundInfo.GoodsName,
			Price:        refundInfo.Price,
			ShippingCost: refundInfo.ShippingCost,
			SellerName:   refundInfo.SellerName,
			BuyerName:    refundInfo.BuyerName,
			SellerID:     refundInfo.SellerID,
			BuyerID:      refundInfo.BuyerID,
			OrderTime:    refundInfo.OrderTime,
			PayTime:      refundInfo.PayTime,
			RefundTime:   refundInfo.CTime,
			ShippingTime: refundInfo.ShippingTime,
			TurnoverTime: refundInfo.TurnoverTime,
			CStatus:      status,
			BuyerReason:  refundInfo.BuyerReason,
			SellerReason: refundInfo.SellerReason,
		})
	}
	if respList == nil { // 确保返回空数组而不是 null
		respList = []types.RefundInfo{}
	}
	var response types.RefundListResp
	response.RefundList = respList
	response.PageNum = req.PageNum
	response.Total = int(total)
	return response, nil
}

// UpdateRefund 更新售后情况
func (s *RefundService) UpdateRefund(ctx context.Context, req types.UpdateRefundReq) (resp interface{}, err error) {
	action := req.Action // 获取传递过来的售后操作
	tradeID := req.TradeID
	// 检查 TradeID 是否有效
	if tradeID == 0 {
		return nil, fmt.Errorf("TradeID 不能为空")
	}
	refund := dao.NewRefundRecord(ctx)
	// 根据不同的 Action 执行不同的操作
	switch action {
	case "同意退货": // 同意退货
		err = refund.ApproveRefund(ctx, tradeID)
		if err != nil {
			return nil, fmt.Errorf("审批失败: %v", err)
		}
	case "拒绝退货": // 拒绝退货
		err = refund.RejectRefund(ctx, tradeID)
		if err != nil {
			return nil, fmt.Errorf("拒绝失败: %v", err)
		}
	default:
		return nil, fmt.Errorf("无效的操作类型: %s", action)
	}
	// 返回成功响应
	resp = map[string]interface{}{
		"data": fmt.Sprintf("售后状态已更新为 %s", action),
	}
	return resp, nil
}
