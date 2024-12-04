package dao

import (
	"errors"
	"github.com/kasiforce/trade/repository/db/model"
	"github.com/kasiforce/trade/types"
	"gorm.io/gorm"
)

var db *gorm.DB

// GetAllComments 获取所有评论
func GetAllComments(req types.ShowCommentsReq) ([]model.Comment, int, error) {
	var comments []model.Comment
	var total int64

	query := db.Model(&model.Comment{})

	if req.GoodsID != 0 {
		query = query.Where("goodsID = ?", req.GoodsID)
	}

	if req.Commentator != "" {
		query = query.Where("commentatorID IN (SELECT userID FROM users WHERE username LIKE ?)", "%"+req.Commentator+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).Find(&comments).Error; err != nil {
		return nil, 0, err
	}

	return comments, int(total), nil
}

// CreateComment 创建评论
func CreateComment(comment model.Comment) error {
	return db.Create(&comment).Error
}

// DeleteComment 删除评论
func DeleteComment(commentID int) error {
	result := db.Delete(&model.Comment{}, commentID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("评论不存在")
	}
	return nil
}

// GetCommentsByUser 根据用户ID获取评论
func GetCommentsByUser(req types.ShowCommentsReq) ([]model.Comment, int, error) {
	var comments []model.Comment
	var total int64

	query := db.Model(&model.Comment{}).Where("commentatorID = ?", req.CommentatorID)

	if req.GoodsID != 0 {
		query = query.Where("goodsID = ?", req.GoodsID)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).Find(&comments).Error; err != nil {
		return nil, 0, err
	}

	return comments, int(total), nil
}
