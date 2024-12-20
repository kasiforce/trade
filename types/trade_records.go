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

// UpdateOrderStatusReq 表示修改订单状态的请求
type UpdateOrderStatusReq struct {
	ID           int    `json:"id" binding:"required"`     // 订单编号
	Status       string `json:"status" binding:"required"` // 目标状态
	RejectReason string `json:"rejectReason"`              //拒绝退款理由
	RefundReason string `json:"refundReason"`              // 退款理由
	Comment      string `json:"comment"`                   // 评价内容
}

// UpdateOrderStatusResp 表示修改订单状态的返回信息
type UpdateOrderStatusResp struct {
	Status string `json:"status"` // 更改后的订单状态
}

// UpdateOrderAddressReq 表示修改订单地址的请求
type UpdateOrderAddressReq struct {
	ID int `json:"id" binding:"required"` // 订单号
	//Province   string `json:"province" binding:"required"`   // 省
	//City       string `json:"city" binding:"required"`       // 市
	//Area       string `json:"area" binding:"required"`       // 区
	//DetailArea string `json:"detailArea" binding:"required"` // 详细地址
	AddrID int `json:"addrID"`
}

// CreateOrderReq 表示生成订单的请求
type CreateOrderReq struct {
	SellerID       int     `json:"sellerID" binding:"required"` // 卖家ID
	GoodsID        int     `json:"goodsID" binding:"required"`  // 商品ID
	Price          float64 `json:"price" binding:"required"`    // 价格
	DeliveryMethod string  `json:"deliveryMethod"`              // 交易方式
	ShippingCost   float64 `json:"shippingCost"`                // 运费
	SenderAddrID   int     `json:"SenderAddrID" `               // 发货地址
	ShippingAddrID int     `json:"shippingAddrID"`              // 收货地址
}

// CreateOrderResp 表示生成订单的返回信息
type CreateOrderResp struct {
	TradeID int `json:"tradeID"` // 订单ID
}

// GetMyOrdersReq 表示获取我买到的订单的请求
type GetMyOrdersReq struct {
	Page     int `form:"page" json:"page"`         // 页码，默认为1
	PageSize int `form:"pageSize" json:"pageSize"` // 每页条数，默认5
}

// GetMyOrdersResp 表示获取我买到的订单的返回信息
type GetMyOrdersResp struct {
	Total     int64            `json:"total"`     // 订单总数
	OrderList []GetMyOrderInfo `json:"orderList"` // 订单列表
}

// GetMyOrderInfo GetMyOrderInfo 表示订单信息
type GetMyOrderInfo struct {
	TradeID         int         `json:"tradeID"`         // 订单ID
	SellerID        int         `json:"sellerID"`        // 卖家ID
	SellerName      string      `json:"sellerName"`      // 卖家昵称
	GoodsID         int         `json:"goodsID"`         // 商品ID
	GoodsName       string      `json:"goodsName"`       // 商品名称
	Price           float64     `json:"price"`           // 价格
	DeliveryMethod  string      `json:"deliveryMethod"`  // 交易方式
	ShippingCost    float64     `json:"shippingCost"`    // 运费
	SenderAddress   AddressInfo `json:"senderAddress"`   // 发货地址
	ShippingAddress AddressInfo `json:"shippingAddress"` // 收货地址
	OrderTime       time.Time   `json:"orderTime"`       // 下单时间
	PayTime         time.Time   `json:"payTime"`         // 支付时间
	ShippingTime    time.Time   `json:"shippingTime"`    // 发货时间
	TurnoverTime    time.Time   `json:"turnoverTime"`    // 成交时间
	Status          string      `json:"status"`          // 订单状态
}

// AddressInfo 表示地址信息
type AddressInfo struct {
	AddrID     int    `json:"addrID"`     // 地址ID
	Province   string `json:"province"`   // 省
	City       string `json:"city"`       // 市
	Area       string `json:"area"`       // 区
	DetailArea string `json:"detailArea"` // 详细地址
	Tel        string `json:"tel"`        // 联系电话
	Name       string `json:"name"`       // 联系人
}

// GetOrderDetailReq 表示获取订单详情的请求
type GetOrderDetailReq struct {
	ID int `uri:"id" binding:"required"` // 订单id
}

// GetOrderDetailResp 表示获取订单详情的响应
type GetOrderDetailResp struct {
	ID             int     `json:"id"`             // 订单id
	Countdown      int     `json:"countdown"`      // 支付倒计时
	SellerID       int     `json:"sellerID"`       // 卖家ID
	GoodsID        int     `json:"goodsID"`        // 商品ID
	Price          float64 `json:"price"`          // 价格
	DeliveryMethod string  `json:"deliveryMethod"` // 交易方式
	ShippingCost   float64 `json:"shippingCost"`   // 运费
	SenderAddrID   int     `json:"SenderAddrID"`   // 发货地址ID
	ShippingAddrID int     `json:"shippingAddrID"` // 收货地址ID
}

// PaySuccessReq 支付成功请求
type PaySuccessReq struct {
	TradeID int `json:"tradeId" binding:"required"` // 订单id，必需
}

// PaySuccessResp 支付成功响应
type PaySuccessResp struct {
	Code int    `json:"code"` // 业务状态码，1为成功
	Msg  string `json:"msg"`  // 响应消息
	Data struct {
	} `json:"data"` // 响应结果
}
