package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kasiforce/trade/middleware"
	"github.com/kasiforce/trade/pkg/util"
	"github.com/kasiforce/trade/repository/db/dao"
	"github.com/kasiforce/trade/types"
	"sync"
)

var adminServ *AdminService
var adminServOnce sync.Once

type AdminService struct {
}

func GetAdminService() *AdminService {
	adminServOnce.Do(func() {
		adminServ = &AdminService{}
	})
	return adminServ
}

func (s *AdminService) ShowAllAdmin(ctx context.Context, req types.ShowAdminReq) (resp interface{}, err error) {
	admin := dao.NewAdmin(ctx)
	adminList, err := admin.FindAll(req)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	var respList []types.AdminInfo
	for _, adminInfo := range adminList {
		respList = append(respList, types.AdminInfo{
			AdminID:   adminInfo.AdminID,
			AdminName: adminInfo.AdminName,
			Password:  adminInfo.Password,
			Email:     adminInfo.Email,
			Role:      adminInfo.Role,
		})
	}
	var response types.AdminListResp
	response.AdminsList = respList
	response.PageNum = req.PageNum
	response.Total = len(respList)
	return response, nil
}

func (s *AdminService) AddAdmin(ctx context.Context, req types.AdminInfo) (resp interface{}, err error) {
	if req.AdminName == "" || req.Password == "" || req.Email == "" {
		err = errors.New("参数不能为空")
		return
	}
	a := dao.NewAdmin(ctx)
	exist, err := a.FindByName(req.AdminName)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	if exist {
		err = errors.New("管理员名已存在")
		return
	}
	modelAdmin := map[string]interface{}{
		"adminName": req.AdminName,
		"password":  req.Password,
		"email":     req.Email,
		"role":      req.Role,
	}
	err = a.CreateAdmin(modelAdmin)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return
}

func (s *AdminService) UpdateAdmin(ctx context.Context, req types.AdminInfo) (resp interface{}, err error) {
	admin, err := ctl.GetAdminID(ctx)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	a := dao.NewAdmin(ctx)
	modelAdmin := map[string]interface{}{
		"adminName": req.AdminName,
		"password":  req.Password,
		"email":     req.Email,
		"role":      req.Role,
	}
	for key, value := range modelAdmin {
		if value == "" {
			delete(modelAdmin, key)
		}
	}
	err = a.UpdateAdmin(admin.AdminID, modelAdmin)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return
}

func (s *AdminService) DeleteAdmin(ctx context.Context) (resp interface{}, err error) {
	admin, err := ctl.GetAdminID(ctx)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	a := dao.NewAdmin(ctx)
	err = a.DeleteAdmin(admin.AdminID)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return
}

func (s *AdminService) AdminLogin(c *gin.Context, req types.AdminLoginReq) (resp interface{}, err error) {
	if req.Email == "" || req.Password == "" {
		err = errors.New("参数不能为空")
		return
	}
	a := dao.NewAdmin(c)
	admin, err := a.CheckEmail(req.Email)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	if admin.Password != req.Password {
		err = errors.New("密码错误")
		return
	}
	token, err := util.GenerateToken(admin.AdminID, admin.AdminName)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	middleware.SetToken(c, token)
	return
}
