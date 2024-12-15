package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kasiforce/trade/pkg/ctl"
	"github.com/kasiforce/trade/pkg/util"
	"github.com/kasiforce/trade/repository/db/dao"
	"github.com/kasiforce/trade/service/pay"
	"github.com/kasiforce/trade/types"
	"github.com/smartwalle/alipay/v3"
	"net/http"
	"strconv"
)

// AlipayHandler 支付宝支付接口
func AlipayHandler(c *gin.Context) {
	// 获取查询参数
	orderIDStr := c.Query("orderId")
	redirectURL := c.Query("redirect")

	// 如果 orderId 为空，返回错误
	if orderIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 0,
			"msg":  "orderId is required",
			"data": nil,
		})
		return
	}

	// 将 orderIDStr 转换为整数
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 0,
			"msg":  "invalid orderId",
			"data": nil,
		})
		return
	}

	// 查询订单信息
	ctx := c.Request.Context()
	tradeRecordsDao := dao.NewTradeRecords(ctx)
	order, err := tradeRecordsDao.GetOrderDetail(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 0,
			"msg":  "failed to get order details",
			"data": nil,
		})
		return
	}

	// 生成支付宝支付请求
	var p = alipay.TradePagePay{}
	p.NotifyURL = pay.GetServerDomain() + "/alipay/notify" // 异步通知地址
	p.ReturnURL = redirectURL                              // 支付后跳转地址
	p.Subject = "订单支付" + strconv.Itoa(order.TradeID)
	p.OutTradeNo = strconv.Itoa(order.TradeID)
	p.TotalAmount = strconv.FormatFloat(order.TurnoverAmount, 'f', 2, 64)
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	// 获取支付链接
	url, err := pay.Client.TradePagePay(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 0,
			"msg":  "failed to generate alipay url",
			"data": nil,
		})
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, url.String())
}

func AlipayNotifyHandler(c *gin.Context) {
	// 解析支付宝通知
	if err := c.Request.ParseForm(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 0,
			"msg":  "failed to parse form",
			"data": nil,
		})
		return
	}

	// 验证签名
	if err := pay.Client.VerifySign(c.Request.Form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 0,
			"msg":  "invalid signature",
			"data": nil,
		})
		return
	}

	// 处理支付结果
	outTradeNo := c.Request.Form.Get("out_trade_no")
	tradeStatus := c.Request.Form.Get("trade_status")

	// 更新订单状态
	orderID, _ := strconv.Atoi(outTradeNo)
	if tradeStatus == "TRADE_SUCCESS" {
		// 更新订单状态为已支付
		// 这里调用你的订单更新逻辑
		req := types.UpdateOrderStatusReq{
			ID:     orderID,
			Status: "未发货",
		}
		u := dao.NewTradeRecords(c)
		resp, err := u.UpdateOrderStatus(req)
		if err != nil {
			util.LogrusObj.Error(err)
			return
		}
		// 返回支付宝成功响应
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}
