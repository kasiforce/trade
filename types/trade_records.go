package types

import "time"

// OrderInfo 表示订单信息
type OrderInfo struct {
	TradeID         int           `json:"tradeID"`         // 交易ID
	SellerID        int           `json:"sellerID"`        // 卖家ID
	BuyerID         int           `json:"buyerID"`         // 买家ID
	SellerName      string        `json:"sellerName"`      // 卖家名
	BuyerName       string        `json:"buyerName"`       // 买家名
	GoodsID         int           `json:"goodsID"`         // 商品ID
	GoodsName       string        `json:"goodsName"`       // 商品名称
	Price           float64       `json:"price"`           // 成交金额
	DeliveryMethod  string        `json:"deliveryMethod"`  // 交易方式
	ShippingCost    float64       `json:"shippingCost"`    // 运费
	SenderAddress   AddressDetail `json:"senderAddress"`   // 发货地址
	ShippingAddress AddressDetail `json:"shippingAddress"` // 收货地址
	OrderTime       time.Time     `json:"orderTime"`       // 下单时间
	PayTime         time.Time     `json:"payTime"`         // 付款时间
	ShippingTime    time.Time     `json:"shippingTime"`    // 发货时间
	TurnoverTime    time.Time     `json:"turnoverTime"`    // 成交时间
	Status          string        `json:"status"`          // 订单状态
}

// AddressDetail  表示地址信息
type AddressDetail struct {
	Province   string `json:"province"`   // 省份
	City       string `json:"city"`       // 城市
	Area       string `json:"area"`       // 区域
	DetailArea string `json:"detailArea"` // 详细地址
}

// OrderListResp 表示订单列表的返回信息
type OrderListResp struct {
	OrderList []OrderInfo `json:"orderList"` // 订单列表
	Total     int64       `json:"total"`     // 总条数
	PageNum   int         `json:"pageNum"`   // 当前页数
}

type ShowOrdersReq struct {
	SearchQuery string `form:"searchQuery" json:"searchQuery"`
	PageNum     int    `form:"pageNum" json:"pageNum"`   // 当前页码
	PageSize    int    `form:"pageSize" json:"pageSize"` // 每页记录数
}
