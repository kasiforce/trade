package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kasiforce/trade/pkg/ctl"
	"github.com/kasiforce/trade/pkg/util"
	"github.com/kasiforce/trade/service"
	"github.com/kasiforce/trade/types"
	"net/http"
	"strconv"
)

// ShowAllCommentsHandler 显示所有评论
func ShowAllCommentsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ShowCommentsReq
		if err := c.ShouldBindQuery(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		s := service.GetCommentService()
		resp, err := s.ShowAllComments(c, req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

// AddCommentHandler 添加评论
//func AddCommentHandler() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		var req types.CreateCommentReq
//		if err := c.ShouldBind(&req); err != nil {
//			util.LogrusObj.Infoln("Error occurred:", err)
//			c.JSON(http.StatusOK, ErrorResponse(c, err))
//			return
//		}
//		s := service.GetCommentService()
//		resp, err := s.AddComment(c.Request.Context(), req)
//		if err != nil {
//			util.LogrusObj.Infoln("Error occurred:", err)
//			c.JSON(http.StatusOK, ErrorResponse(c, err))
//			return
//		}
//		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
//	}
//}

// DeleteCommentHandler 根据评论ID删除评论
func DeleteCommentHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
			return
		}

		ctx := c.Request.Context()
		s := service.GetCommentService()
		resp, err := s.DeleteCommentByID(ctx, id)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}

		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

// ShowCommentsByUserHandler 根据用户ID显示其发布的评论
func ShowCommentsByUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := service.GetCommentService()
		resp, err := s.ShowCommentsByID(c)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}
