package dao

import (
	"context"

	"github.com/kasiforce/trade/repository/db/model"
	"gorm.io/gorm"
)

type Comment struct {
	*gorm.DB
}

func NewCommentByDB(db *gorm.DB) *Comment {
	return &Comment{db}
}

func NewComment(ctx context.Context) *Comment {
	return &Comment{NewDBClient(ctx)}
}

func (comment *Comment) FindAll() (comments []*model.Comment, err error) {
	err = comment.DB.Model(&model.Comment{}).Find(&comments).Error
	return
}

func (comment *Comment) FindByID(id int) (c *model.Comment, err error) {
	err = comment.DB.Model(&model.Comment{}).Where("commentID = ?", id).First(&c).Error
	return
}

func (comment *Comment) FindByGoodsID(goodsID int) (comments []*model.Comment, err error) {
	err = comment.DB.Model(&model.Comment{}).Where("goodsID = ?", goodsID).Find(&comments).Error
	return
}

func (comment *Comment) CreateComment(c *model.Comment) (err error) {
	err = comment.DB.Model(&model.Comment{}).Create(&c).Error
	return
}

func (comment *Comment) UpdateComment(id int, c *model.Comment) (err error) {
	err = comment.DB.Model(&model.Comment{}).Where("commentID = ?", id).Updates(&c).Error
	return
}
