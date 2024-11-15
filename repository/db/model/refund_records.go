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
}

func (RefundRecord) TableName() string {
	return "refund_records"
}
