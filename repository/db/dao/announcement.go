package dao

import (
	"context"
	"errors"
	"fmt"
	"github.com/kasiforce/trade/repository/db/model"
	"github.com/kasiforce/trade/types"
	"gorm.io/gorm"
	"time"
)

type Announcement struct {
	*gorm.DB
}

// NewAnnouncementByDB 通过数据库连接创建 Announcement 实例
func NewAnnouncementByDB(db *gorm.DB) *Announcement {
	return &Announcement{db}
}

// NewAnnouncement 通过上下文创建 Announcement 实例
func NewAnnouncement(ctx context.Context) *Announcement {
	return &Announcement{NewDBClient(ctx)}
}

func (a *Announcement) GetAllAnnouncements(req types.ShowAnnouncementsReq) (r []types.AnnouncementInfo, total int64, err error) {
	// 构建查询条件
	query := a.DB.Model(&model.Announcement{})

	// 如果提供了搜索查询，则添加模糊查询条件
	if req.SearchQuery != "" {
		query = query.Where("anTitle LIKE ?", "%"+req.SearchQuery+"%").Count(&total).
			Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).
			Select("announcementID as AnnouncementID," +
				"anTitle as AnTitle," +
				"anContent as AnContent," +
				"anTime as AnTime")
	} else {
		// 构建查询
		query = query.Count(&total).
			Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).
			Select("announcementID as AnnouncementID," +
				"anTitle as AnTitle," +
				"anContent as AnContent," +
				"anTime as AnTime")
	}

	// 执行查询
	err = query.Find(&r).Error
	if err != nil {
		return
	}

	// 打印 r 的值以进行调试
	fmt.Printf("Debug: r = %+v\n", r)

	return
}

func (a *Announcement) ClientGetAllAnnouncements() (r []types.AnnouncementInfo, err error) {
	// 构建查询条件
	query := a.DB.Model(&model.Announcement{})

	// 构建查询
	query = query.
		Select("announcementID as AnnouncementID," +
			"anTitle as AnTitle," +
			"anContent as AnContent," +
					"anTime as AnTime").
		Order("anTime DESC"). // 按 anTime 降序排序
		Limit(3)              // 限制返回最新的3条

	// 执行查询
	err = query.Find(&r).Error
	if err != nil {
		return
	}

	// 打印 r 的值以进行调试
	fmt.Printf("Debug: r = %+v\n", r)
	return
}

// CreateAnnouncement 创建公告
func (a *Announcement) CreateAnnouncement(req types.CreateAnnouncementReq) error {
	newAnnouncement := model.Announcement{
		AnTitle:   req.AnTitle,
		AnContent: req.AnContent,
		AnTime:    time.Now(), // 设置当前时间为发布时间
	}

	result := a.DB.Create(&newAnnouncement)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// UpdateAnnouncement 更新公告
func (a *Announcement) UpdateAnnouncement(req types.UpdateAnnouncementReq) error {
	updateData := map[string]interface{}{
		"AnTitle":   req.AnTitle,
		"AnContent": req.AnContent,
	}

	result := a.DB.Model(&model.Announcement{}).
		Where("announcementID = ?", req.AnnouncementID).
		Updates(updateData)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("公告不存在")
	}

	return nil
}

// DeleteAnnouncement 删除公告
func (a *Announcement) DeleteAnnouncement(announcementID int) error {
	result := a.DB.Delete(&model.Announcement{}, announcementID)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("公告不存在")
	}

	return nil
}
