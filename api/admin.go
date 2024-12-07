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

func ShowAllAdminHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ShowAdminReq
		if err := c.ShouldBindQuery(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		s := service.GetAdminService()
		resp, err := s.ShowAllAdmin(c, req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

func AddAdminHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.AdminInfo
		if err := c.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		s := service.GetAdminService()
		resp, err := s.AddAdmin(c.Request.Context(), req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

func UpdateAdminHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.AdminInfo
		if err := c.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)

		ctx := ctl.NewAdminContext(c.Request.Context(), &ctl.AdminInfo{AdminID: id})
		s := service.GetAdminService()
		resp, err := s.UpdateAdmin(ctx, req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

func DeleteAdminHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		ctx := ctl.NewAdminContext(c.Request.Context(), &ctl.AdminInfo{AdminID: id})
		s := service.GetAdminService()
		resp, err := s.DeleteAdmin(ctx)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

func AdminLoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.AdminLoginReq
		if err := c.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		s := service.GetAdminService()
		resp, err := s.AdminLogin(c, req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}
