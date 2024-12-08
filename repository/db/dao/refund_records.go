package dao

import (
	"context"
	"github.com/kasiforce/trade/repository/db/model"
	"github.com/kasiforce/trade/types"
	"gorm.io/gorm"
)

type RefundRecord struct {
	*gorm.DB
}

func NewRefundRecordByDB(db *gorm.DB) *RefundRecord {
	return &RefundRecord{db}
}

func NewRefundRecord(ctx context.Context) *RefundRecord {
	return &RefundRecord{NewDBClient(ctx)}
}

func (r *RefundRecord) FindAll(req types.ShowRefundReq) (refunds []*model.RefundRecord, err error) {
	query := r.DB.Model(&model.RefundRecord{}).
		Joins("JOIN trade_records t ON t.tradeID = refund_records.tradeID").
		Joins("JOIN users seller ON seller.userID = t.sellerID").
		Joins("JOIN users buyer ON buyer.userID = t.buyerID").
		Joins("JOIN goods ON goods.goodsID = t.goodsID").
		Joins("JOIN refund_complaint rc ON rc.tradeID = refund_records.tradeID").
		Select("refund_records.refundID, refund_records.tradeID, refund_records.payMethod, refund_records.refundAgreedTime, refund_records.refundShippedTime, refund_records.refundArrivalTime, refund_records.trackingNumber, t.orderTime, t.payTime, t.shippingTime, t.turnoverTime, seller.userName AS sellerName, buyer.userName AS buyerName, goods.goodsName AS goodsName, goods.price AS price, (t.turnoverAmount - goods.price) AS shippingCost, rc.buyerReason AS buyerReason, rc.cStatus AS status, seller.userID AS sellerID, buyer.userID AS buyerID")
	//Select("rc.cReason AS reason, rc.cStatus AS status")
	if req.SearchQuery != "" {
		query = query.Where("refund_records.refundID LIKE ?", "%"+req.SearchQuery+"%")
	}
	query = query.Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize)
	err = query.Debug().Find(&refunds).Error
	return
}

func (r *RefundRecord) FindByID(id int) (refund *model.RefundRecord, err error) {
	err = r.DB.Model(&model.RefundRecord{}).Where("refundID = ?", id).First(&refund).Error
	return
}

/*
func (r *RefundRecord) CreateRefundRecord(refund *model.RefundRecord) (err error) {
	err = r.DB.Model(&model.RefundRecord{}).Create(&refund).Error
	return
}

func (r *RefundRecord) UpdateRefundRecord(id int, refund *model.RefundRecord) (err error) {
	err = r.DB.Model(&model.RefundRecord{}).Where("refundID = ?", id).Updates(refund).Error
	return
}
*/
