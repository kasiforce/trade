package types

import "time"

// AnnouncementInfo 表示公告信息
type AnnouncementInfo struct {
	AnnouncementID int       `json:"announcementID"` // 公告ID
	AnTitle        string    `json:"anTitle"`        // 公告标题
	AnContent      string    `json:"anContent"`      // 公告内容
	AnTime         time.Time `json:"anTime"`         // 发布时间
}

// ShowAnnouncementsReq 表示查询公告列表的请求
type ShowAnnouncementsReq struct {
	SearchQuery string `form:"searchQuery" json:"searchQuery"`
	PageNum     int    `form:"pageNum" json:"pageNum"`   // 当前页码
	PageSize    int    `form:"pageSize" json:"pageSize"` // 每页记录数
}

// CreateAnnouncementReq 表示创建公告的请求
type CreateAnnouncementReq struct {
	AnTitle   string `json:"anTitle" binding:"required"`   // 公告标题
	AnContent string `json:"anContent" binding:"required"` // 公告内容
}

// UpdateAnnouncementReq 表示更新公告的请求
type UpdateAnnouncementReq struct {
	AnnouncementID int    `json:"announcementID" binding:"required"` // 公告ID
	AnTitle        string `json:"anTitle"`                           // 公告标题
	AnContent      string `json:"anContent"`                         // 公告内容
}

// AnnouncementListResp 公告列表的响应
type AnnouncementListResp struct {
	AnnouncementList []AnnouncementInfo `json:"announcementList"` // 公告列表
	Total            int64              `json:"total"`            // 总记录数
	PageNum          int                `json:"pageNum"`          // 当前页码
}
