package model

import (
	"time"
)

type RefundComplaint struct {
	ComplaintID int       `gorm:"primaryKey;autoIncrement;column:complaintID"`
	TradeID     int       `gorm:"not null;column:tradeID"`
	CReason     string    `gorm:"type:text;not null;column:cReason"`
	CTime       time.Time `gorm:"type:datetime;not null;column:cTime"`
	CStatus     int       `gorm:"type:tinyint;not null;default:0;column:cStatus"`
}

func (RefundComplaint) TableName() string {
	return "refund_complaint"
}
