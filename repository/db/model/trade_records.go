package model

import (
	"time"
)

type TradeRecords struct {
	TradeID        int       `gorm:"primaryKey;autoIncrement;column:tradeID"`
	SellerID       int       `gorm:"not null;column:sellerID"`
	BuyerID        int       `gorm:"not null;column:buyerID"`
	GoodsID        int       `gorm:"not null;column:goodsID"`
	TurnoverAmount float64   `gorm:"type:decimal(10,2);not null;check:turnoverAmount >= 0;column:turnoverAmount"`
	ShippingAddrID int       `gorm:"not null;column:shippingAddrID"`
	DeliveryAddrID int       `gorm:"not null;column:deliveryAddrID"`
	OrderTime      time.Time `gorm:"type:datetime;not null;column:orderTime"`
	PayTime        time.Time `gorm:"type:datetime;column:payTime"`
	ShippingTime   time.Time `gorm:"type:datetime;column:shippingTime"`
	TurnoverTime   time.Time `gorm:"type:datetime;column:turnoverTime"`
	PayMethod      int       `gorm:"type:tinyint;not null;default:0;column:payMethod"`
	TrackingNumber string    `gorm:"type:varchar(100);column:trackingNumber"`
	IsReturn       int       `gorm:"type:tinyint;not null;default:0;column:isReturn"`
	Status         string    `gorm:"type:varchar(20);not null;column:status"`
	ShippingCost   float64   `gorm:"type:decimal(10,2);not null;default:0;check:shippingCost >= 0;column:shippingCost"`
}

func (TradeRecords) TableName() string {
	return "trade_records"
}
