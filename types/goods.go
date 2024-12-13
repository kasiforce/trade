package types

import "time"

// 查询所有商品（管理员） 17
type GoodsInfo struct {
	GoodsID        int       `json:"id"`    // 商品ID
	GoodsName      string    `json:"title"` // 商品名称
	Price          float64   `json:"price"` // 价格
	CategoryName   string    `json:"category"`
	Details        string    `json:"description"` // 商品详情
	GoodsImages    string    `json:"imageUrl"`    // 商品图片
	ShippingCost   float64   `json:"shippingCost"`
	UserName       string    `json:"userName"`
	Province       string    `json:"province"`
	City           string    `json:"city"`
	District       string    `json:"area"`
	Address        string    `json:"detailArea"`
	CreatedTime    time.Time `json:"postTime"` // 创建时间
	DeliveryMethod string    `json:"deliveryMethod"`
	Star           int       `json:"stars"`
	View           int       `json:"views"`
	IsSold         int       `json:"isSold"` // 是否已售：0 未售，1 已售
}

// 已售出商品（用户） 6
type GoodsInfo2 struct {
	GoodsID     int       `json:"id"`          // 商品ID
	GoodsName   string    `json:"title"`       // 商品名称
	Price       float64   `json:"price"`       // 价格
	Details     string    `json:"description"` // 商品详情
	GoodsImages string    `json:"imageUrl"`    // 商品图片
	CreatedTime time.Time `json:"postTime"`    // 创建时间
}

// 获取发布中商品（用户） 19
type GoodsInfo3 struct {
	GoodsID        int       `json:"id"`     // 商品ID
	GoodsName      string    `json:"title"`  // 商品名称
	UserID         int       `json:"userID"` // 用户ID
	Price          float64   `json:"price"`  // 价格
	CategoryName   string    `json:"category"`
	Details        string    `json:"description"` // 商品详情
	GoodsImages    string    `json:"imageUrl"`    // 商品图片
	ShippingCost   float64   `json:"shippingCost"`
	UserName       string    `json:"userName"`
	Province       string    `json:"province"`
	City           string    `json:"city"`
	District       string    `json:"area"`
	Address        string    `json:"detailArea"`
	CreatedTime    time.Time `json:"postTime"` // 创建时间
	DeliveryMethod string    `json:"deliveryMethod"`
	Star           int       `json:"stars"`
	View           int       `json:"views"`
	IsSold         int       `json:"isSold"` // 是否已售：0 未售，1 已售
	AddrID         int       `json:"addrID"`
}

// 商品列表 5
type GoodsInfo4 struct {
	GoodsID     int     `json:"id"`      // 商品ID
	GoodsName   string  `json:"name"`    // 商品名称
	Price       float64 `json:"price"`   // 价格
	GoodsImages string  `json:"picture"` // 商品图片
	CategoryID  int     `json:"categoryID"`
}

type GoodsListResp struct {
	ProductList []GoodsInfo `json:"productList"` // 商品列表
	Total       int         `json:"total"`       // 总记录数
	PageNum     int         `json:"pageNum"`     // 当前页码
}

type GoodsListResp2 struct {
	Data []GoodsInfo2 `json:"data"` // 商品列表
}

type ShowGoodsReq struct {
	SearchQuery    string  `form:"searchQuery" json:"searchQuery"` // 商品名称模糊查询
	Page           int     `form:"page" json:"page"`               // 当前页码
	Limit          int     `form:"limit" json:"limit"`             // 每页记录数
	Province       string  `form:"province" json:"province"`
	City           string  `form:"city" json:"city"`
	District       string  `form:"area" json:"area"`
	DeliveryMethod string  `form:"deliveryMethod" json:"deliveryMethod"`
	CategoryID     int     `form:"categoryID" json:"categoryID"`   // 商品分类
	PriceMin       float64 `form:"priceMin" json:"priceMin"`       // 最低价格
	PriceMax       float64 `form:"priceMax" json:"priceMax"`       // 最高价格
	PublishDate    string  `form:"publishDate" json:"publishDate"` // 创建时间
	ShippingCost   int     `form:"shippingCost" json:"shippingCost"`
}

type ShowAllGoodsReq struct {
	SearchQuery string `form:"searchQuery" json:"searchQuery"` // 模糊搜索条件
	PageNum     int    `form:"pageNum" json:"pageNum"`         // 当前页码
	PageSize    int    `form:"pageSize" json:"pageSize"`       // 每页记录数
}

type ShowGoodsListReq struct {
	SearchQuery string `form:"searchQuery" json:"searchQuery"` // 模糊搜索条件
	Page        int    `form:"page" json:"page"`               // 当前页码
	Limit       int    `form:"limit" json:"limit"`             // 每页记录数
	Category    string `form:"category" json:"category"`
}

type ShowGoodsDetail struct {
	GoodsID        int       `json:"id"`    // 商品ID
	GoodsName      string    `json:"title"` // 商品名称
	Price          float64   `json:"price"` // 价格
	CategoryName   string    `json:"category"`
	Details        string    `json:"description"` // 商品详情
	GoodsImages    string    `json:"imageUrl"`    // 商品图片
	ShippingCost   float64   `json:"shippingCost"`
	UserName       string    `json:"userName"`
	UserID         int       `json:"userID"` // 用户ID
	AddrID         int       `json:"addrID"`
	Tel            string    `json:"tel"`
	Province       string    `json:"province"`
	City           string    `json:"city"`
	District       string    `json:"area"`
	Address        string    `json:"detailArea"`
	CreatedTime    time.Time `json:"postTime"` // 创建时间
	DeliveryMethod string    `json:"deliveryMethod"`
	Star           int       `json:"stars"`
	View           int       `json:"views"`
	IsSold         int       `json:"isSold"` // 是否已售：0 未售，1 已售
	IsStarred      bool      `json:"isStarred"`
}

type ShowDetailReq struct {
	GoodsID int `form:"id" json:"id"` // 商品ID
}

type CreateGoodsReq struct {
	GoodsName      string  `json:"title"`
	Details        string  `json:"description"`
	Price          float64 `json:"price"`
	CategoryName   string  `json:"category"`
	GoodsImages    string  `json:"imageUrl"`
	AddrID         int     `json:"addrID"`
	DeliveryMethod string  `json:"deliveryMethod"`
	ShippingCost   float64 `json:"shippingCost"`
}

type IsStarred struct {
	IsStarred bool `json:"isStarred"`
}

type UpdateGoodsReq struct {
	GoodsID        int     `form:"id" json:"id"`       // 商品名称模糊查询
	GoodsName      string  `form:"title" json:"title"` // 当前页码
	Price          float64 `form:"price" json:"price"`
	Category       string  `form:"category" json:"category"`
	Details        string  `form:"description" json:"description"`
	ImageUrl       string  `form:"imageUrl" json:"imageUrl"`
	ShippingCost   float64 `form:"shippingCost" json:"shippingCost"`
	UserName       string  `form:"userName" json:"userName"`
	AddrID         int     `form:"addrID" json:"addrID"`
	DeliveryMethod string  `form:"deliveryMethod" json:"deliveryMethod"`
}
