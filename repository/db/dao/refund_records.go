package dao

import (
	"context"
	"github.com/kasiforce/trade/repository/db/model"
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

func (r *RefundRecord) FindAll() (refunds []*model.RefundRecord, err error) {
	err = r.DB.Model(&model.RefundRecord{}).Find(&refunds).Error
	return
}

func (r *RefundRecord) FindByID(id int) (refund *model.RefundRecord, err error) {
	err = r.DB.Model(&model.RefundRecord{}).Where("refundID = ?", id).First(&refund).Error
	return
}

func (r *RefundRecord) CreateRefundRecord(refund *model.RefundRecord) (err error) {
	err = r.DB.Model(&model.RefundRecord{}).Create(&refund).Error
	return
}

func (r *RefundRecord) UpdateRefundRecord(id int, refund *model.RefundRecord) (err error) {
	err = r.DB.Model(&model.RefundRecord{}).Where("refundID = ?", id).Updates(refund).Error
	return
}

func (r *RefundRecord) DeleteRefundRecord(id int) (err error) {
	err = r.DB.Model(&model.RefundRecord{}).Where("refundID = ?", id).Delete(&model.RefundRecord{}).Error
	return
}
