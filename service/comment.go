package service

import (
	"context"
	"github.com/kasiforce/trade/repository/db/dao"
	"github.com/kasiforce/trade/repository/db/model"
	"github.com/kasiforce/trade/types"
	"sync"
	"time"
)

var commentServ *CommentService
var commentServOnce sync.Once

type CommentService struct {
}

func GetCommentService() *CommentService {
	commentServOnce.Do(func() {
		commentServ = &CommentService{}
	})
	return commentServ
}

// ShowAllComments 获取所有评论
func (s *CommentService) ShowAllComments(ctx context.Context, req types.ShowCommentsReq) (resp interface{}, err error) {
	comment := dao.NewComment(ctx)
	comments, total, err := comment.GetAllComments(req)
	if err != nil {
		return nil, err
	}

	var commentInfos []types.CommentInfo
	for _, commentInfo := range comments {
		commentInfos = append(commentInfos, types.CommentInfo{
			CommentatorID:   commentInfo.CommentatorID,
			CommentatorName: commentInfo.CommentatorName, // 这里需要从用户表获取用户名
			CommentContent:  commentInfo.CommentContent,
			CommentID:       commentInfo.CommentID,
			CommentTime:     commentInfo.CommentTime,
			GoodsID:         commentInfo.GoodsID,
		})
	}

	return types.CommentListResp{
		Comments: commentInfos,
		Total:    total,
		PageNum:  req.PageNum,
	}, nil
}

// AddComment 添加评论
func (s *CommentService) AddComment(ctx context.Context, req types.CreateCommentReq) (resp interface{}, err error) {
	comment := model.Comment{
		GoodsID:        int(req.GoodsID),
		CommentatorID:  int(req.CommentatorID),
		CommentContent: req.CommentContent,
		CommentTime:    time.Now(),
	}

	err = dao.CreateComment(comment)
	if err != nil {
		return nil, err
	}

	return "Comment added successfully", nil
}

// DeleteCommentByID 删除评论
func (s *CommentService) DeleteCommentByID(ctx context.Context, commentID int) (resp interface{}, err error) {
	err = dao.DeleteComment(commentID)
	if err != nil {
		return nil, err
	}

	return "Comment deleted successfully", nil
}

// ShowCommentsByID 根据用户ID获取评论
func (s *CommentService) ShowCommentsByID(ctx context.Context, req types.ShowCommentsReq) (resp interface{}, err error) {
	comment := dao.NewComment(ctx)
	comments, total, err := comment.GetCommentsByUser(req)
	if err != nil {
		return nil, err
	}

	var commentInfos []types.CommentInfo
	for _, commentInfo := range comments {
		commentInfos = append(commentInfos, types.CommentInfo{
			CommentatorID:   commentInfo.CommentatorID,
			CommentatorName: commentInfo.CommentatorName, // 这里需要从用户表获取用户名
			CommentContent:  commentInfo.CommentContent,
			CommentID:       commentInfo.CommentID,
			CommentTime:     commentInfo.CommentTime,
			GoodsID:         commentInfo.GoodsID,
		})
	}

	return types.CommentListResp{
		Comments: commentInfos,
		Total:    total,
		PageNum:  req.PageNum,
	}, nil
}
