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

func ShowCategoryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ShowCategoryReq
		if err := c.ShouldBindQuery(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		println(req.PageSize)
		s := service.GetCategoryService()
		resp, err := s.ShowCategory(c.Request.Context(), req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))

	}
}
func ShowUserCategoryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := service.GetCategoryService()
		resp, err := s.ShowUserCategory(c.Request.Context())
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}
func AddCategoryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.CategoryInfo
		if err := c.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		s := service.GetCategoryService()
		resp, err := s.AddCategory(c.Request.Context(), req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

func UpdateCategoryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.CategoryInfo
		if err := c.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		s := service.GetCategoryService()
		resp, err := s.UpdateCategory(c.Request.Context(), req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

func DeleteCategoryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := service.GetCategoryService()
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		resp, err := s.DeleteCategory(c, id)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}
