package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kasiforce/trade/pkg/ctl"
	"github.com/kasiforce/trade/pkg/util"
	"github.com/kasiforce/trade/service"
	"github.com/kasiforce/trade/types"
	"net/http"
	//"strconv"
)

// GetAllOrdersHandler 查询所有订单
func GetAllOrdersHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ShowOrdersReq
		if err := c.ShouldBindQuery(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		ctx := c.Request.Context()
		s := service.GetTrade_recordsService()
		resp, err := s.GetAllOrders(ctx, req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}

		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

// UpdateOrderStatusHandler 修改订单状态
func UpdateOrderStatusHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.UpdateOrderStatusReq
		if err := c.ShouldBindJSON(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
			return
		}

		ctx := c.Request.Context()
		s := service.GetTrade_recordsService()
		resp, err := s.UpdateOrderStatus(ctx, req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}

		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

// UpdateOrderAddressHandler 修改订单地址
func UpdateOrderAddressHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.UpdateOrderAddressReq
		if err := c.ShouldBindJSON(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
			return
		}

		ctx := c.Request.Context()
		s := service.GetTrade_recordsService()
		resp, err := s.UpdateOrderAddress(ctx, req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}

		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

// CreateOrderHandler 生成订单
func CreateOrderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.CreateOrderReq
		if err := c.ShouldBindJSON(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
			return
		}

		//ctx := c.Request.Context()
		s := service.GetTrade_recordsService()
		resp, err := s.CreateOrder(c, req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}

		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

// GetMyOrdersHandler 获取我买到的订单
func GetMyOrdersHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.GetMyOrdersReq
		if err := c.ShouldBindQuery(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
			return
		}

		//ctx := c.Request.Context()
		s := service.GetTrade_recordsService()
		resp, err := s.GetMyOrders(c, req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}

		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

// GetMySoldOrdersHandler 获取我卖出的订单
func GetMySoldOrdersHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.GetMyOrdersReq
		if err := c.ShouldBindQuery(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
			return
		}

		s := service.GetTrade_recordsService()
		resp, err := s.GetMySoldOrders(c, req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}

		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

// GetOrderDetailHandler 获取订单详情
func GetOrderDetailHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.GetOrderDetailReq
		if err := c.ShouldBindUri(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
			return
		}

		s := service.GetTrade_recordsService()
		resp, err := s.GetOrderDetail(c, req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}

		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

// PaySuccessHandler 支付成功后改变订单状态为未发货
func PaySuccessHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.PaySuccessReq
		if err := c.ShouldBindJSON(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
			return
		}
		ctx := c.Request.Context()
		s := service.GetTrade_recordsService()
		resp, err := s.PaySuccess(ctx, req)

		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}
