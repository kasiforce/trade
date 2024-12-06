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
func (c *Comment) GetAllComments(req types.ShowCommentsReq) (r []*types.CommentInfo, total int64, err error) {
	err = c.DB.Model(&model.Comment{}).Preload("User").Preload("Goods").Count(&total).Error
	if err != nil {
		return
	}
	err = c.DB.Model(&model.Comment{}).
		Joins("As co left join users as u on u.userID = co.commentatorID ").
		Joins("left join goods as g on g.goodsID = co.goodsID").
		Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).
		Select("co.commentID as commentID," +
			"g.goodsName as goodsName," +
			"u.userName as commentatorName," +
			"co.commentContent as commentContent," +
			"co.commentTime as commentTime").
		Find(&r).Error

	return
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
func (c *Comment) GetCommentsByUser(id int) (r []types.CommentInfoByID, err error) {
	err = c.DB.Model(&model.Comment{}).Preload("User").Preload("Goods").Where("commentatorID = ?", id).Error
	if err != nil {
		return
	}
	err = c.DB.Model(&model.Comment{}).
		Joins("As co left join users as u on u.userID = co.commentatorID ").
		Joins("left join goods as g on g.goodsID = co.goodsID").
		Where("co.commentatorID = ?", id).
		Select("co.commentID as commentID," +
			"g.goodsID as goodsID," +
			"co.commentatorID as commentatorID," +
			"u.userName as commentatorName," +
			"co.commentContent as commentContent," +
			"co.commentTime as commentTime").
		Find(&r).Error

	return
}
