package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kasiforce/trade/pkg/ctl"
	"github.com/kasiforce/trade/pkg/util"
	"github.com/kasiforce/trade/service"
	"net/http"
)

func ShowCategoryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := service.GetCategoryService()
		resp, err := s.ShowCategory(c.Request.Context())
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}
