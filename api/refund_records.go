package api

import (
	"github.com/cghdjvjg/trade/pkg/ctl"
	"github.com/cghdjvjg/trade/pkg/util"
	"github.com/cghdjvjg/trade/service"
	"github.com/cghdjvjg/trade/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowAllrefundHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ShowRefundReq
		if err := c.ShouldBindQuery(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		s := service.GetRefundService()
		resp, err := s.ShowAllRefund(c, req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}
