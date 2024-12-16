package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kasiforce/trade/pkg/ctl"
	"github.com/kasiforce/trade/pkg/util"
	"github.com/kasiforce/trade/service"
	"github.com/kasiforce/trade/types"
	"net/http"
	"strconv"
	"time"
)

// ShowAllAnnouncementsHandler 查询所有公告
func ShowAllAnnouncementsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ShowAnnouncementsReq
		if err := c.ShouldBindQuery(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
			return
		}

		s := service.GetAnnouncementService()
		resp, err := s.ShowAllAnnouncements(c, req)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}

		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

// CreateAnnouncementHandler 添加公告
func CreateAnnouncementHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.CreateAnnouncementReq
		if err := c.ShouldBindJSON(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
			return
		}

		s := service.GetAnnouncementService()
		if err := s.CreateAnnouncement(c, req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}

		c.JSON(http.StatusOK, ctl.RespSuccess(c, "公告创建成功"))
	}
}

// UpdateAnnouncementHandler 修改公告
func UpdateAnnouncementHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.UpdateAnnouncementReq
		if err := c.ShouldBindJSON(&req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
			return
		}

		s := service.GetAnnouncementService()
		if err := s.UpdateAnnouncement(c, req); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}

		c.JSON(http.StatusOK, ctl.RespSuccess(c, "公告更新成功"))
	}
}

// DeleteAnnouncementHandler 删除公告
func DeleteAnnouncementHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("announcementID")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
			return
		}

		s := service.GetAnnouncementService()
		if err := s.DeleteAnnouncement(c, id); err != nil {
			util.LogrusObj.Infoln("Error occurred:", err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}

		c.JSON(http.StatusOK, ctl.RespSuccess(c, "公告删除成功"))
	}
}

// SSEAnnouncementsHandler 使用 SSE 推送公告
func SSEAnnouncementsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "text/event-stream")
		c.Header("Cache-Control", "no-cache")
		c.Header("Connection", "keep-alive")
		c.Header("Access-Control-Allow-Origin", "*")

		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				s := service.GetAnnouncementService()
				resp, err := s.ClientShowAllAnnouncements(c)
				if err != nil {
					util.LogrusObj.Infoln("Error occurred:", err)
					continue
				}

				data, _ := resp.([]types.AnnouncementInfo)
				for _, announcement := range data {
					event := fmt.Sprintf("data: %v\n\n", announcement)
					_, err := c.Writer.WriteString(event)
					if err != nil {
						util.LogrusObj.Infoln("Error occurred:", err)
						return
					}
					c.Writer.Flush()
				}
			case <-c.Request.Context().Done():
				return
			}
		}
	}
}
