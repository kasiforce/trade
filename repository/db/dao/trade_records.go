package dao

import (
	"context"
	"github.com/kasiforce/trade/repository/db/model"
	"gorm.io/gorm"
)

type TradeRecords struct {
	*gorm.DB
}

func NewTradeRecordsByDB(db *gorm.DB) *TradeRecords {
	return &TradeRecords{db}
}

func NewTradeRecords(ctx context.Context) *TradeRecords {
	return &TradeRecords{NewDBClient(ctx)}
}

func (tr *TradeRecords) FindAll() (tradeRecords []*model.TradeRecords, err error) {
	err = tr.DB.Model(&model.TradeRecords{}).Find(&tradeRecords).Error
	return
}

func (tr *TradeRecords) FindByID(id int) (t *model.TradeRecords, err error) {
	err = tr.DB.Model(&model.TradeRecords{}).Where("tradeID = ?", id).First(&t).Error
	return
}

func (tr *TradeRecords) FindBySellerID(sellerID int) (tradeRecords []*model.TradeRecords, err error) {
	err = tr.DB.Model(&model.TradeRecords{}).Where("sellerID = ?", sellerID).Find(&tradeRecords).Error
	return
}

func (tr *TradeRecords) FindByBuyerID(buyerID int) (tradeRecords []*model.TradeRecords, err error) {
	err = tr.DB.Model(&model.TradeRecords{}).Where("buyerID = ?", buyerID).Find(&tradeRecords).Error
	return
}

func (tr *TradeRecords) FindByGoodsID(goodsID int) (tradeRecords []*model.TradeRecords, err error) {
	err = tr.DB.Model(&model.TradeRecords{}).Where("goodsID = ?", goodsID).Find(&tradeRecords).Error
	return
}

func (tr *TradeRecords) CreateTradeRecord(t *model.TradeRecords) (err error) {
	err = tr.DB.Model(&model.TradeRecords{}).Create(&t).Error
	return
}

func (tr *TradeRecords) UpdateTradeRecord(id int, t *model.TradeRecords) (err error) {
	err = tr.DB.Model(&model.TradeRecords{}).Where("tradeID = ?", id).Updates(&t).Error
	return
}

func (tr *TradeRecords) DeleteTradeRecord(id int) (err error) {
	err = tr.DB.Model(&model.TradeRecords{}).Where("tradeID = ?", id).Delete(&model.TradeRecords{}).Error
	return
}
