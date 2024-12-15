package types

import (
	"time"
)

type RefundInfo struct {
	//RefundID     int       `json:"refundID"`
	TradeID      int       `json:"tradeID"`
	GoodsName    string    `json:"goodsName"`
	Price        float64   `json:"price"`
	ShippingCost float64   `json:"shippingCost"`
	SellerName   string    `json:"sellerName"`
	BuyerName    string    `json:"buyerName"`
	SellerID     int       `json:"sellerID"`
	BuyerID      int       `json:"buyerID"`
	OrderTime    time.Time `json:"orderTime"`
	PayTime      time.Time `json:"payTime"`
	RefundTime   time.Time `json:"refundTime"`
	ShippingTime time.Time `json:"shippingTime"`
	TurnoverTime time.Time `json:"turnoverTime"`
	BuyerReason  string    `json:"buyerReason"`
	SellerReason string    `json:"sellerReason"`
	CStatus      string    `json:"status"`
}

type RefundListResp struct {
	RefundList []RefundInfo `json:"refundList"` // 管理员列表
	Total      int          `json:"total"`      // 总记录数
	PageNum    int          `json:"pageNum"`    // 当前页码
}

type ShowRefundReq struct {
	SearchQuery string `form:"searchQuery" json:"searchQuery"` // 模糊搜索条件
	PageNum     int    `form:"pageNum" json:"pageNum"`         // 当前页码
	PageSize    int    `form:"pageSize" json:"pageSize"`       // 每页记录数
}

type UpdateRefundReq struct {
	Action  string `json:"action"`  // 模糊搜索条件
	TradeID int    `json:"tradeID"` // 当前页码
}
