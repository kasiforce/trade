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
		// 从请求参数中获取 id
		idStr := c.Query("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
			return
		}

		s := service.GetCommentService()
		resp, err := s.ShowCommentsByID(c, id)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

// GetReceivedCommentsHandler 获取别人给该用户的评价
func GetReceivedCommentsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Query("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
			return
		}

		ctx := c.Request.Context()
		s := service.GetCommentService()
		resp, err := s.GetReceivedCommentsByUserID(ctx, id)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}

		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

//// PostCommentHandler 发布评论
//func PostCommentHandler() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		var req types.PostCommentReq
//		if err := c.ShouldBindJSON(&req); err != nil {
//			util.LogrusObj.Infoln("Error occurred:", err)
//			c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
//			return
//		}
//
//		ctx := c.Request.Context()
//		s := service.GetCommentService()
//		resp, err := s.PostComment(ctx, req)
//		if err != nil {
//			util.LogrusObj.Infoln("Error occurred:", err)
//			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
//			return
//		}
//
//		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
//	}
//}
