package model

import "time"

type RefundRecord struct {
	RefundID          int       `gorm:"primaryKey;autoIncrement;column:refundID"`
	TradeID           int       `gorm:"not null;column:tradeID"`
	PayMethod         int       `gorm:"not null;column:payMethod"`
	RefundAgreedTime  time.Time `gorm:"type:datetime;not null;column:refundAgreedTime"`
	RefundShippedTime time.Time `gorm:"type:datetime;column:refundShippedTime"`
	RefundArrivalTime time.Time `gorm:"type:datetime;column:refundArrivalTime"`
	TrackingNumber    string    `gorm:"type:varchar(100);column:trackingNumber"`
	//Trade             TradeRecords `gorm:"foreignKey:TradeID;references:TradeID"`
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
	BuyerReason  string    `gorm:"column:buyerReason;type:text"`
	SellerReason string    `gorm:"column:sellerReason;type:text"`
	CStatus      int       `gorm:"type:tinyint;not null;default:0;column:cStatus"`
}

func (RefundRecord) TableName() string {
	return "refund_records"
}
