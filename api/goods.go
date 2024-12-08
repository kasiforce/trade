package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kasiforce/trade/pkg/ctl"
	"github.com/kasiforce/trade/pkg/util"
	"github.com/kasiforce/trade/service"
	"github.com/kasiforce/trade/types"
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
		s := service.GetGoodsService()
		resp, err := s.IsSoldGoods(c)
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
		s := service.GetGoodsService()
		resp, err := s.ShowPublishedGoods(c)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

func DeleteGoodsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		ctx := ctl.NewGoodsContext(c.Request.Context(), &ctl.GoodsInfo{GoodsID: id})
		s := service.GetGoodsService()
		resp, err := s.DeleteGoods(ctx, id)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

// ShowAllGoodsHandler 获取商品列表
func ShowAllGoodsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ShowGoodsListReq
		if err := c.ShouldBindQuery(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		s := service.GetGoodsService()
		resp, err := s.ShowGoodsList(c, req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

// 筛选商品
func FilterGoodsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ShowGoodsReq
		if err := c.ShouldBindQuery(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		s := service.GetGoodsService()
		resp, err := s.FilterGoods(c, req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

// 更新view
func IncreaseGoodsViewHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取商品ID
		goodsID := c.Param("id")
		if goodsID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "商品ID不能为空"})
			return
		}

		// 将商品ID转换为uint类型
		id, err := strconv.ParseUint(goodsID, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效的商品ID"})
			return
		}

		// 调用服务层方法更新商品的view字段
		s := service.GetGoodsService()
		err = s.IncreaseGoodsView(c.Request.Context(), uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "更新商品view失败"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "商品view更新成功"})
	}
}
