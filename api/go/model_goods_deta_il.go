/*
 * 二手交易平台
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// GoodsDetaIl - 商品详情
type GoodsDetaIl struct {

	// 商品id
	Id int32 `json:"id"`

	// 商品名
	Name string `json:"name"`

	// 价格
	Price float32 `json:"price"`

	// 分类ID
	CategoryID int32 `json:"categoryID"`

	// 物品描述
	Describe string `json:"describe"`

	// 图片URL
	Image string `json:"image"`

	// 邮费
	ShippingCost float32 `json:"shippingCost,omitempty"`

	// 发布者名
	UserName string `json:"userName"`

	// 发货地址，仅需省市区
	Address string `json:"address"`

	// 发布时间
	PostTime string `json:"postTime"`

	// 配送方式
	DeliveryMethod string `json:"deliveryMethod"`

	// 浏览量
	Views int32 `json:"views"`

	// 收藏量
	Stars int32 `json:"stars"`

	// 是否卖出，1为已卖出，0为未卖出
	IsSold int32 `json:"isSold"`
}
