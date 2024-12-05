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

//// AddComment 添加评论
//func (s *CommentService) AddComment(ctx context.Context, req types.CreateCommentReq) (resp interface{}, err error) {
//	comment := model.Comment{
//		GoodsID:        int(req.GoodsID),
//		CommentatorID:  int(req.CommentatorID),
//		CommentContent: req.CommentContent,
//		CommentTime:    time.Now(),
//	}
//
//	err = dao.CreateComment(comment)
//	if err != nil {
//		return nil, err
//	}
//
//	return "Comment added successfully", nil
//}

// DeleteCommentByID 删除评论
func (s *CommentService) DeleteCommentByID(ctx context.Context, commentID int) (resp interface{}, err error) {
	err = dao.NewComment(ctx).DeleteComment(commentID)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return "Comment deleted successfully", nil
}

// ShowCommentsByID 根据用户ID获取评论
func (s *CommentService) ShowCommentsByID(c *gin.Context) (resp interface{}, err error) {
	id := c.GetInt("id")
	u := dao.NewComment(c)
	comments, err := u.GetCommentsByUser(id)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	resp = &comments
	return
}
