package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kasiforce/trade/pkg/ctl"
	"github.com/kasiforce/trade/pkg/e"
	"github.com/kasiforce/trade/pkg/util"
	"github.com/kasiforce/trade/service"
	"github.com/kasiforce/trade/types"
	"net/http"
	"strconv"
)

func ShowAllUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ShowUserReq
		if err := c.ShouldBindQuery(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		s := service.GetUserService()
		resp, err := s.ShowAllUser(c, req)
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
		var req types.UserInfo
		if err := c.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		s := service.GetUserService()
		resp, err := s.AddUser(c.Request.Context(), req)
		str, ok := resp.(string) // 使用类型断言来获取字符串
		if ok {
			r := &ctl.Response{Code: e.Error, Data: err, Msg: str}
			c.JSON(http.StatusOK, r)
			return
		}
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

// UpdateUserHandler 管理员更新用户信息
func UpdateUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.UserInfo
		if err := c.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		ctx := ctl.NewContext(c.Request.Context(), &ctl.UserInfo{UserID: id})
		s := service.GetUserService()
		resp, err := s.UpdateUser(ctx, req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

func DeleteUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		ctx := ctl.NewContext(c.Request.Context(), &ctl.UserInfo{UserID: id})
		s := service.GetUserService()
		resp, err := s.DeleteUser(ctx)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

func ShowIntroductionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ID
		if err := c.ShouldBindQuery(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		s := service.GetUserService()
		resp, err := s.ShowIntroduction(c, req.ID)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

func ShowUserByIDHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := service.GetUserService()
		resp, err := s.ShowUserByID(c)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

func UpdateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.UpdateUserReq
		if err := c.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		s := service.GetUserService()
		resp, err := s.Update(c, req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

func UserLoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.UserLoginReq
		if err := c.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		s := service.GetUserService()
		resp, err := s.UserLogin(c, req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

func UserRegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.UserRegisterReq
		if err := c.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		s := service.GetUserService()
		resp, err := s.UserRegister(c, &req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

func SendEmailCodeHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.MailCodeReq
		if err := ctx.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}

		l := service.GetUserService()
		resp, err := l.SendEmailCode(ctx.Request.Context(), &req)
		if err != nil {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

func PwdUpdateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.PwdUpdateReq
		if err := c.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln(err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		l := service.GetUserService()
		resp, err := l.PwdUpdate(c, &req)
		if err != nil {
			util.LogrusObj.Infoln(err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}
