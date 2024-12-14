package model

import (
	"time"
)

type RefundComplaint struct {
	ComplaintID  int       `gorm:"primaryKey;autoIncrement;column:complaintID"`
	TradeID      int       `gorm:"not null;column:tradeID"`
	BuyerReason  string    `gorm:"type:text;not null;column:buyerReason"`
	CTime        time.Time `gorm:"type:datetime;not null;column:cTime"`
	CStatus      int       `gorm:"type:tinyint;not null;default:0;column:cStatus"`
	SellerReason string    `gorm:"type:text;column:sellerReason"`
	OrderTime    time.Time `gorm:"type:datetime;not null;column:orderTime"`
	PayTime      time.Time `gorm:"type:datetime;column:payTime"`
	ShippingTime time.Time `gorm:"type:datetime;column:shippingTime"`
	TurnoverTime time.Time `gorm:"type:datetime;column:turnoverTime"`
	SellerName   string    `gorm:"not null;column:sellerName"`
	BuyerName    string    `gorm:"not null;column:buyerName"`
	GoodsName    string    `gorm:"not null;column:goodsName"`
	SellerID     int       `gorm:"not null;column:sellerID"`
	BuyerID      int       `gorm:"not null;column:buyerID"`
	GoodsID      int       `gorm:"not null;column:goodsID"`
	ShippingCost float64   `gorm:"type:decimal(10,2);not null;check:turnoverAmount >= 0;column:turnoverAmount"`
	Price        float64   `gorm:"type:decimal(10,2);not null;check:price >= 0;column:price"`
}

func (RefundComplaint) TableName() string {
	return "refund_complaint"
}
