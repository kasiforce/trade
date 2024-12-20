package dao

import (
	"context"
	"github.com/kasiforce/trade/repository/db/model"
	"github.com/kasiforce/trade/types"
	"gorm.io/gorm"
)

type RefundComplaint struct {
	*gorm.DB
}

func NewRefundComplaintByDB(db *gorm.DB) *RefundComplaint {
	return &RefundComplaint{db}
}

func NewRefundComplaint(ctx context.Context) *RefundComplaint {
	return &RefundComplaint{NewDBClient(ctx)}
}

func (rc *RefundComplaint) CountAll() (total int64, err error) {
	err = rc.DB.Model(&model.RefundComplaint{}).Count(&total).Error
	return
}

func (rc *RefundComplaint) FindAll(req types.ShowRefundReq) (refundComplaints []*model.RefundComplaint, err error) {
	query := rc.DB.Model(&model.RefundComplaint{}).
		Joins("JOIN trade_records t ON t.tradeID = refund_complaint.tradeID").
		Joins("JOIN users seller ON seller.userID = t.sellerID").
		Joins("JOIN users buyer ON buyer.userID = t.buyerID").
		Joins("JOIN goods ON goods.goodsID = t.goodsID").
		//Joins("JOIN refund_complaint rc ON rc.tradeID = refund_records.tradeID").
		Select("refund_complaint.tradeID, refund_complaint.cTime, t.orderTime, t.payTime, t.shippingTime, t.turnoverTime, seller.userName AS sellerName, buyer.userName AS buyerName, goods.goodsName AS goodsName, goods.price AS price, t.shippingCost AS shippingCost, refund_complaint.buyerReason AS buyerReason, refund_complaint.sellerReason AS sellerReason, refund_complaint.cStatus AS cStatus, seller.userID AS sellerID, buyer.userID AS buyerID").
		Where("t.status=?", "处理中")
	if req.SearchQuery != "" {
		query = query.Where("refund_complaint.tradeID = ?", req.SearchQuery)
	}
	query = query.Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize)
	err = query.Debug().Find(&refundComplaints).Error
	return
}

func (rc *RefundComplaint) FindByID(id int) (r *model.RefundComplaint, err error) {
	err = rc.DB.Model(&model.RefundComplaint{}).Where("complaintID = ?", id).First(&r).Error
	return
}

func (rc *RefundComplaint) FindByTradeID(tradeID int) (refundComplaints []*model.RefundComplaint, err error) {
	err = rc.DB.Model(&model.RefundComplaint{}).Where("tradeID = ?", tradeID).Find(&refundComplaints).Error
	return
}

func (rc *RefundComplaint) CreateRefundComplaint(r *model.RefundComplaint) (err error) {
	err = rc.DB.Model(&model.RefundComplaint{}).Create(&r).Error
	return
}

func (rc *RefundComplaint) UpdateRefundComplaint(id int, r *model.RefundComplaint) (err error) {
	err = rc.DB.Model(&model.RefundComplaint{}).Where("complaintID = ?", id).Updates(&r).Error
	return
}

func (rc *RefundComplaint) DeleteRefundComplaint(id int) (err error) {
	err = rc.DB.Model(&model.RefundComplaint{}).Where("complaintID = ?", id).Delete(&model.RefundComplaint{}).Error
	return
}
