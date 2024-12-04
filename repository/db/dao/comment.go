package dao

import (
	"context"
	"errors"
	"github.com/kasiforce/trade/repository/db/model"
	"github.com/kasiforce/trade/types"
	"gorm.io/gorm"
)

type Comment struct {
	*gorm.DB
}

// NewCommentByDB 通过数据库连接创建 Comment 实例
func NewCommentByDB(db *gorm.DB) *Comment {
	return &Comment{db}
}

// NewComment 通过上下文创建 Comment 实例
func NewComment(ctx context.Context) *Comment {
	return &Comment{NewDBClient(ctx)}
}

// GetAllComments 获取所有评论
func (c *Comment) GetAllComments(req types.ShowCommentsReq) ([]model.Comment, int, error) {
	var comments []model.Comment
	var total int64

	query := c.DB.Model(&model.Comment{})

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
func (c *Comment) CreateComment(comment model.Comment) error {
	return c.DB.Create(&comment).Error
}

// DeleteComment 删除评论
func (c *Comment) DeleteComment(commentID int) error {
	result := c.DB.Delete(&model.Comment{}, commentID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("评论不存在")
	}
	return nil
}

// GetCommentsByUser 根据用户ID获取评论
func (c *Comment) GetCommentsByUser(req types.ShowCommentsReq) ([]model.Comment, int, error) {
	var comments []model.Comment
	var total int64

	query := c.DB.Model(&model.Comment{}).Where("commentatorID = ?", req.CommentatorID)

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
