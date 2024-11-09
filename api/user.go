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

func ShowAllUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := service.GetUserService()
		resp, err := s.ShowAllUser(c)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

func ShowUserInfoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDstr := c.Param("id")
		userID, err := strconv.Atoi(userIDstr)
		ctx := ctl.NewContext(c.Request.Context(), &ctl.UserInfo{UserID: userID})
		s := service.GetUserService()
		resp, err := s.ShowUserInfoByID(ctx)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

func AddUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.UserAddReq
		if err := c.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		s := service.GetUserService()
		resp, err := s.AddUser(c.Request.Context(), &req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

func UpdateUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.UserInfoUpdateReq
		if err := c.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		ctx := ctl.NewContext(c.Request.Context(), &ctl.UserInfo{UserID: id})
		s := service.GetUserService()
		resp, err := s.UpdateUser(ctx, &req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}
