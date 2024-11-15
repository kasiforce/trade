package model

import "time"

type RefundRecord struct {
	RefundID          int
	TradeID           int
	PayMethod         int
	RefundAgreedTime  time.Time
	RefundShippedTime time.Time
	RefundArrivalTime time.Time
	TrackingNumber    string
}

func (RefundRecord) TableName() string {
	return "refund_records"
}
