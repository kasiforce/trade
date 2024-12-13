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
		// 从请求参数中获取 id
		idStr := c.Query("id")
		var id int
		var err error
		// 如果请求参数为空，则从上下文中获取 id
		if idStr == "" {
			id = c.GetInt("id")
			if err != nil {
				util.LogrusObj.Infoln("Error occurred: id not found in context")
				c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
				return
			}
		} else {
			// 如果请求参数不为空，则将其转换为整数
			id, err = strconv.Atoi(idStr)
			if err != nil {
				util.LogrusObj.Infoln("Error occurred:", err)
				c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
				return
			}
		}
		// 调用服务层方法
		s := service.GetGoodsService()
		resp, err := s.IsSoldGoods(c, id)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		// 返回成功响应
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

func PublishedGoodsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求参数中获取 id
		idStr := c.Query("id")
		var id int
		var err error
		// 如果请求参数为空，则从上下文中获取 id
		if idStr == "" {
			id = c.GetInt("id")
			if err != nil {
				util.LogrusObj.Infoln("Error occurred: id not found in context")
				c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
				return
			}
		} else {
			// 如果请求参数不为空，则将其转换为整数
			id, err = strconv.Atoi(idStr)
			if err != nil {
				util.LogrusObj.Infoln("Error occurred:", err)
				c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
				return
			}
		}
		s := service.GetGoodsService()
		resp, err := s.ShowPublishedGoods(c, id)
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

func ShowGoodsDetailHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ShowDetailReq
		if err := c.ShouldBindQuery(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		s := service.GetGoodsService()
		resp, err := s.ShowGoodsDetail(c, req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

// 发布闲置
func CreateGoodsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.CreateGoodsReq
		if err := c.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		s := service.GetGoodsService()
		resp, err := s.AddGoods(c, req)
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
		// 获取请求中的商品ID
		var req types.ShowDetailReq
		if err := c.ShouldBindQuery(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		// 检查商品ID是否为空
		if req.GoodsID == 0 {
			c.Status(http.StatusBadRequest)
			return
		}
		// 调用服务层方法更新商品的view字段
		s := service.GetGoodsService()
		err := s.IncreaseGoodsView(c.Request.Context(), req)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.Status(http.StatusOK)
	}
}

// 更新商品收藏情况
func UpdateGoodsIsStarredHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取路径参数中的商品ID
		goodsIDStr := c.Param("id")
		// 将 goodsID 从字符串转换为整数
		goodsID, err := strconv.Atoi(goodsIDStr)
		if err != nil {
			util.LogrusObj.Infoln("Invalid GoodsID format:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		// 获取请求体中的 isStarred，绑定 JSON 数据
		var r types.IsStarred
		if err := c.ShouldBindJSON(&r); err != nil {
			util.LogrusObj.Infoln("Invalid JSON body:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}

		// 调用服务层方法更新商品的收藏情况
		s := service.GetGoodsService()
		resp, err := s.UpdateGoodsIsStarred(c, goodsID, r)
		if err != nil {
			util.LogrusObj.Infoln("Failed to update goods starred status:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}

		// 返回成功响应
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

// 修改发布中商品
func UpdateGoodsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求体中的 JSON 数据
		var req types.UpdateGoodsReq
		if err := c.ShouldBindJSON(&req); err != nil {
			util.LogrusObj.Infoln("Invalid JSON body:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}

		// 调用服务层方法更新商品信息
		s := service.GetGoodsService()
		resp, err := s.UpdateGoods(c, req)
		if err != nil {
			util.LogrusObj.Infoln("Failed to update goods:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}

		// 返回成功响应
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}
