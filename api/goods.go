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

// AdminShowAllGoodsHandler 获取所有商品（管理员端）
func AdminShowAllGoodsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ShowAllGoodsReq
		if err := c.ShouldBindQuery(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		s := service.GetGoodsService()
		resp, err := s.ShowAllGoods(c, req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

func IsSoldGoodsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.IsSoldGoodsResp
		userIDStr := c.DefaultQuery("id", "")
		if userIDStr == "" {
			req.UserID = 0
		} else {
			// 如果 'id' 存在，尝试将其转换为 int 类型并赋值给 req.UserID
			userID, err := strconv.Atoi(userIDStr)
			if err != nil {
				util.LogrusObj.Infoln("Error occurred:", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userID"})
				return
			}
			req.UserID = userID
		}
		if err := c.ShouldBindQuery(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		s := service.GetGoodsService()
		resp, err := s.IsSoldGoods(c, req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

func PublishedGoodsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.IsSoldGoodsResp
		userIDStr := c.DefaultQuery("id", "")
		if userIDStr == "" {
			req.UserID = 0
		} else {
			// 如果 'id' 存在，尝试将其转换为 int 类型并赋值给 req.UserID
			userID, err := strconv.Atoi(userIDStr)
			if err != nil {
				util.LogrusObj.Infoln("Error occurred:", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userID"})
				return
			}
			req.UserID = userID
		}
		if err := c.ShouldBindQuery(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		s := service.GetGoodsService()
		resp, err := s.ShowPublishedGoods(c, req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

/*
	func ShowGoodsDetailHandler() gin.HandlerFunc {
		return func(c *gin.Context) {
			idStr := c.Param("id")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				util.LogrusObj.Infoln("Invalid product ID:", err)
				c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
				return
			}
			s := service.GetGoodsService()
			resp, err := s.ShowGoodsDetail(c, id)
			if err != nil {
				util.LogrusObj.Infoln("Error occurred:", err)
				c.JSON(http.StatusOK, ErrorResponse(c, err))
				return
			}
			c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
		}
	}

/*

	func FilterGoodsHandler() gin.HandlerFunc {
		return func(c *gin.Context) {
			var req types.ShowGoodsReq
			if err := c.ShouldBindQuery(&req); err != nil {
				util.LogrusObj.Infoln("Error occurred:", err)
				c.JSON(http.StatusOK, ErrorResponse(c, err))
				return
			}
			filter := make(map[string]interface{})
			if req.CategoryID > 0 {
				filter["categoryID"] = req.CategoryID
			}
			if req.IsSold == 0 || req.IsSold == 1 {
				filter["isSold"] = req.IsSold
			}
			if req.PriceMin > 0 {
				filter["price >= ?"] = req.PriceMin
			}
			if req.PriceMax > 0 {
				filter["price <= ?"] = req.PriceMax
			}

			s := service.GetGoodsService()
			resp, err := s.FilterGoods(c, filter, req)
			if err != nil {
				util.LogrusObj.Infoln("Error occurred:", err)
				c.JSON(http.StatusOK, ErrorResponse(c, err))
				return
			}
			c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
		}
	}

	func CreateGoodsHandler() gin.HandlerFunc {
		return func(c *gin.Context) {
			var req types.GoodsInfo
			if err := c.ShouldBind(&req); err != nil {
				util.LogrusObj.Infoln("Error occurred:", err)
				c.JSON(http.StatusOK, ErrorResponse(c, err))
				return
			}
			s := service.GetGoodsService()
			resp, err := s.CreateGoods(c.Request.Context(), req)
			if err != nil {
				util.LogrusObj.Infoln("Error occurred:", err)
				c.JSON(http.StatusOK, ErrorResponse(c, err))
				return
			}
			c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
		}
	}
*/
func DeleteGoodsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		ctx := ctl.NewGoodsContext(c.Request.Context(), &ctl.GoodsInfo{GoodsID: id})
		s := service.GetGoodsService()
		resp, err := s.DeleteGoods(ctx)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}
