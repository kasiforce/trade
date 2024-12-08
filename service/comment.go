package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/kasiforce/trade/pkg/util"
	"github.com/kasiforce/trade/repository/db/dao"
	"github.com/kasiforce/trade/types"
	"sync"
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
		util.LogrusObj.Error(err)
		return
	}
	resp = &types.CommentListResp{
		CommentList: comments,
		Total:       total,
		PageNum:     req.PageNum,
	}
	return
}

// DeleteCommentByID 删除评论
func (s *CommentService) DeleteCommentByID(ctx context.Context, commentID int) (resp interface{}, err error) {
	err = dao.NewComment(ctx).DeleteComment(commentID)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}

	// 返回一个对象类型的响应
	resp = map[string]string{
		"message": "Comment deleted successfully",
	}
	return resp, nil
}

// ShowCommentsByID 根据用户ID获取评论
func (s *CommentService) ShowCommentsByID(c *gin.Context, id int) (resp interface{}, err error) {
	u := dao.NewComment(c)
	resp, err = u.GetCommentsByUser(id)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return
}

// GetReceivedCommentsByUserID 根据用户ID获取收到的评价
func (s *CommentService) GetReceivedCommentsByUserID(ctx context.Context, userID int) (resp interface{}, err error) {
	u := dao.NewComment(ctx)
	resp, err = u.GetReceivedComments(userID)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return
}

//// PostComment 发布评论
//func (s *CommentService) PostComment(ctx context.Context, req types.PostCommentReq) (resp interface{}, err error) {
//	u := dao.NewComment(ctx)
//	err = u.CreateComment(req)
//	if err != nil {
//		util.LogrusObj.Error(err)
//		return
//	}
//	resp = map[string]string{
//		"message": "Comment posted successfully",
//	}
//	return resp, nil
//}
