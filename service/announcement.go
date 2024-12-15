package service

import (
	"context"
	"github.com/kasiforce/trade/pkg/util"
	"github.com/kasiforce/trade/repository/db/dao"
	"github.com/kasiforce/trade/types"
	"sync"
)

var announcementServ *AnnouncementService
var announcementServOnce sync.Once

type AnnouncementService struct{}

func GetAnnouncementService() *AnnouncementService {
	announcementServOnce.Do(func() {
		announcementServ = &AnnouncementService{}
	})
	return announcementServ
}

// ShowAllAnnouncements 获取所有公告
func (s *AnnouncementService) ShowAllAnnouncements(ctx context.Context, req types.ShowAnnouncementsReq) (resp interface{}, err error) {
	announcement := dao.NewAnnouncement(ctx)
	announcements, total, err := announcement.GetAllAnnouncements(req)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}

	resp = &types.AnnouncementListResp{
		AnnouncementList: announcements,
		Total:            total,
		PageNum:          req.PageNum,
	}
	return resp, nil
}

func (s *AnnouncementService) ClientShowAllAnnouncements(ctx context.Context) (resp interface{}, err error) {
	announcement := dao.NewAnnouncement(ctx)
	announcements, err := announcement.ClientGetAllAnnouncements()
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}
	return announcements, nil
}

// CreateAnnouncement 创建公告
func (s *AnnouncementService) CreateAnnouncement(ctx context.Context, req types.CreateAnnouncementReq) error {
	return dao.NewAnnouncement(ctx).CreateAnnouncement(req)
}

// UpdateAnnouncement 更新公告
func (s *AnnouncementService) UpdateAnnouncement(ctx context.Context, req types.UpdateAnnouncementReq) error {
	return dao.NewAnnouncement(ctx).UpdateAnnouncement(req)
}

// DeleteAnnouncement 删除公告
func (s *AnnouncementService) DeleteAnnouncement(ctx context.Context, announcementID int) error {
	return dao.NewAnnouncement(ctx).DeleteAnnouncement(announcementID)
}
