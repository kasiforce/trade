package dao

import (
	"context"
	"github.com/kasiforce/trade/repository/db/model"
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

func (rc *RefundComplaint) FindAll() (refundComplaints []*model.RefundComplaint, err error) {
	err = rc.DB.Model(&model.RefundComplaint{}).Find(&refundComplaints).Error
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
