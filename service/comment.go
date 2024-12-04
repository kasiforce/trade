package service

import (
	"context"
	"github.com/kasiforce/trade/dao"
	"github.com/kasiforce/trade/model"
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

func (s *CommentService) ShowAllComments(ctx context.Context, req types.ShowCommentsReq) (resp interface{}, err error) {
	comments, total, err := dao.GetAllComments(req)
	if err != nil {
		return nil, err
	}

	var commentInfos []types.CommentInfo
	for _, comment := range comments {
		commentInfos = append(commentInfos, types.CommentInfo{
			CommentatorID:   int64(comment.CommentatorID),
			CommentatorName: "TODO: 从用户表获取用户名", // 这里需要从用户表获取用户名
			CommentContent:  comment.CommentContent,
			CommentID:       int64(comment.CommentID),
			CommentTime:     &comment.CommentTime.String(),
			GoodsID:         int64(comment.GoodsID),
		})
	}

	return types.CommentListResp{
		Comments: commentInfos,
		Total:    total,
		PageNum:  req.PageNum,
	}, nil
}

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

func (s *CommentService) DeleteCommentByID(ctx context.Context, commentID int) (resp interface{}, err error) {
	err = dao.DeleteComment(commentID)
	if err != nil {
		return nil, err
	}

	return "Comment deleted successfully", nil
}

func (s *CommentService) ShowCommentsByUser(ctx context.Context, req types.ShowCommentsReq) (resp interface{}, err error) {
	comments, total, err := dao.GetCommentsByUser(req)
	if err != nil {
		return nil, err
	}

	var commentInfos []types.CommentInfo
	for _, comment := range comments {
		commentInfos = append(commentInfos, types.CommentInfo{
			CommentatorID:   int64(comment.CommentatorID),
			CommentatorName: "TODO: 从用户表获取用户名", // 这里需要从用户表获取用户名
			CommentContent:  comment.CommentContent,
			CommentID:       int64(comment.CommentID),
			CommentTime:     comment.CommentTime.String(),
			GoodsID:         int64(comment.GoodsID),
		})
	}

	return types.CommentListResp{
		Comments: commentInfos,
		Total:    total,
		PageNum:  req.PageNum,
	}, nil
}
