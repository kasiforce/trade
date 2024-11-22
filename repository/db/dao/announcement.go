package dao

import (
	"context"
	"github.com/kasiforce/trade/repository/db/model"
	"gorm.io/gorm"
)

type Announcement struct {
	*gorm.DB
}

func NewAnnouncementByDB(db *gorm.DB) *Announcement {
	return &Announcement{db}
}

func NewAnnouncement(ctx context.Context) *Announcement {
	return &Announcement{NewDBClient(ctx)}
}

func (announcement *Announcement) FindAll() (announcements []*model.Announcement, err error) {
	err = announcement.DB.Model(&model.Announcement{}).Find(&announcements).Error
	return
}

func (announcement *Announcement) FindByID(id int) (a *model.Announcement, err error) {
	err = announcement.DB.Model(&model.Announcement{}).Where("announcementID = ?", id).First(&a).Error
	return
}

func (announcement *Announcement) FindByTitle(title string) (exist bool, err error) {
	var cnt int64
	err = announcement.DB.Model(&model.Announcement{}).Where("anTitle = ?", title).Count(&cnt).Error
	if cnt == 0 {
		return false, err
	}
	return true, err
}

func (announcement *Announcement) CreateAnnouncement(a *model.Announcement) (err error) {
	err = announcement.DB.Model(&model.Announcement{}).Create(&a).Error
	return
}

func (announcement *Announcement) UpdateAnnouncement(id int, a *model.Announcement) (err error) {
	err = announcement.DB.Model(&model.Announcement{}).Where("announcementID = ?", id).Updates(&a).Error
	return
}

func (announcement *Announcement) DeleteAnnouncement(id int) (err error) {
	err = announcement.DB.Model(&model.Announcement{}).Where("announcementID = ?", id).Delete(&model.Announcement{}).Error
	return
}
