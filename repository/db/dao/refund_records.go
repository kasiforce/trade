package dao

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/kasiforce/trade/repository/db/model"
	"github.com/kasiforce/trade/types"
	"gorm.io/gorm"
)

type RefundRecord struct {
	*gorm.DB
}

func (r *RefundRecord) CountAll() (total int64, err error) {
	err = r.DB.Model(&model.RefundRecord{}).Count(&total).Error
	return
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
		Select("refund_records.refundID, refund_records.tradeID, refund_records.payMethod, refund_records.refundAgreedTime, refund_records.refundShippedTime, refund_records.refundArrivalTime, refund_records.trackingNumber, t.orderTime, t.payTime, t.shippingTime, t.turnoverTime, seller.userName AS sellerName, buyer.userName AS buyerName, goods.goodsName AS goodsName, goods.price AS price, t.shippingCost AS shippingCost, rc.buyerReason AS buyerReason, rc.sellerReason AS sellerReason, rc.cStatus AS cStatus, seller.userID AS sellerID, buyer.userID AS buyerID")
	if req.SearchQuery != "" {
		query = query.Where("refund_records.tradeID = ?", req.SearchQuery)
	}
	query = query.Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize)
	err = query.Debug().Find(&refunds).Error
	return
}

func (r *RefundRecord) FindByID(id int) (refund *model.RefundRecord, err error) {
	err = r.DB.Model(&model.RefundRecord{}).Where("refundID = ?", id).First(&refund).Error
	return
}

func (r *RefundRecord) ApproveRefund(ctx context.Context, tradeID int) error {
	tx := r.DB.WithContext(ctx).Begin()
	// 更新 trade_records 表
	if err := tx.Exec("UPDATE trade_records SET status = '已退款' WHERE tradeID = ?", tradeID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("更新订单状态为 '已退款' 失败: %v", err)
	}
	// 更新 refund_complaint 表
	if err := tx.Exec("UPDATE refund_complaint SET cstatus = 1 WHERE tradeID = ?", tradeID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("更新售后状态为 '同意退货' 失败: %v", err)
	}
	return tx.Commit().Error
}

func (r *RefundRecord) RejectRefund(ctx context.Context, tradeID int) error {
	tx := r.DB.WithContext(ctx).Begin()
	// 查询发货时间
	var shipTime sql.NullTime
	if err := tx.Raw("SELECT shippingTime FROM trade_records WHERE tradeID = ?", tradeID).Scan(&shipTime).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("查询发货时间失败: %v", err)
	}
	// 更新 trade_records 表状态
	status := "未发货"
	if shipTime.Valid {
		status = "已发货"
	}
	if err := tx.Exec("UPDATE trade_records SET status = ? WHERE tradeID = ?", status, tradeID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("更新订单状态失败: %v", err)
	}
	// 更新 refund_complaint 表
	if err := tx.Exec("UPDATE refund_complaint SET cstatus = 2 WHERE tradeID = ?", tradeID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("更新售后状态为 '拒绝退货' 失败: %v", err)
	}
	return tx.Commit().Error
}
