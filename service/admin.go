package service

import (
	"context"
	"errors"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/kasiforce/trade/pkg/ctl"
	"github.com/kasiforce/trade/pkg/util"
	"github.com/kasiforce/trade/repository/db/dao"
	"github.com/kasiforce/trade/types"
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
	// 获取数据库中所有管理员的总数
	total, err := admin.CountAll()
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}
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
			Passwords: adminInfo.Password,
			Tel:       adminInfo.Tel,
			Mail:      adminInfo.Mail,
			Gender:    adminInfo.Gender,
			Age:       adminInfo.Age,
		})
	}
	if respList == nil { // 确保返回空数组而不是 null
		respList = []types.AdminInfo{}
	}
	var response types.AdminListResp
	response.AdminList = respList
	response.PageNum = req.PageNum
	response.Total = int(total)
	return response, nil
}

func (s *AdminService) AddAdmin(ctx context.Context, req types.AdminInfo) (resp interface{}, err error) {
	if req.AdminName == "" || req.Passwords == "" || req.Mail == "" {
		err = errors.New("参数不能为空")
		return nil, err
	}
	a := dao.NewAdmin(ctx)
	exist, err := a.FindByName(req.AdminName)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}
	if exist {
		err = errors.New("管理员名已存在")
		return nil, err
	}
	exist1, err1 := a.FindByMail(req.Mail)
	if err1 != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}
	if exist1 {
		err = errors.New("管理员邮箱已存在")
		return nil, err
	}
	modelAdmin := map[string]interface{}{
		"adminName": req.AdminName,
		"passwords": req.Passwords,
		"tel":       req.Tel,
		"mail":      req.Mail,
		"gender":    req.Gender,
		"age":       req.Age,
	}
	adminID, err := a.CreateAdmin(modelAdmin)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}
	// 创建返回结构，包含 adminID
	resp = map[string]interface{}{
		"adminID": adminID,
	}
	return resp, nil
}

func (s *AdminService) UpdateAdmin(ctx context.Context, req types.AdminInfo) (resp interface{}, err error) {
	admin, err := ctl.GetAdminID(ctx)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}
	a := dao.NewAdmin(ctx)
	exist, err := a.FindByID(admin.AdminID)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}
	if exist == nil {
		err = errors.New("管理员不存在")
		return nil, err
	}
	modelAdmin := map[string]interface{}{
		"adminName": req.AdminName,
		"passwords": req.Passwords,
		"mail":      req.Mail,
		"tel":       req.Tel,
		"gender":    req.Gender,
		"age":       req.Age,
	}
	for key, value := range modelAdmin {
		if value == "" {
			delete(modelAdmin, key)
		}
	}
	err = a.UpdateAdmin(admin.AdminID, modelAdmin)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}
	// 创建一个空的返回结构
	resp = map[string]interface{}{}
	return resp, nil
}

func (s *AdminService) DeleteAdmin(ctx context.Context) (resp interface{}, err error) {
	admin, err := ctl.GetAdminID(ctx)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}
	a := dao.NewAdmin(ctx)
	// 检查管理员是否存在
	exist, err := a.FindByID(admin.AdminID)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}
	if exist == nil {
		err = errors.New("管理员不存在")
		return nil, err
	}
	err = a.DeleteAdmin(admin.AdminID)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}
	// 创建一个空的返回结构
	resp = map[string]interface{}{}
	return resp, nil
}

func (s *AdminService) AdminLogin(c *gin.Context, req types.AdminLoginReq) (resp interface{}, err error) {
	if req.Mail == "" || req.Password == "" {
		err = errors.New("参数不能为空")
		return
	}
	a := dao.NewAdmin(c)
	admin, err := a.CheckMail(req.Mail)
	if err != nil {
		util.LogrusObj.Error(err)
		err = errors.New("邮箱不存在")
		return
	}

	if admin.Password != req.Password {
		err = errors.New("密码错误")
		return
	}

	token, err := util.GenerateToken(admin.AdminID, admin.AdminName)
	if err != nil {
		util.LogrusObj.Error(err)
		err = errors.New("生成Token失败")
		return
	}

	// 返回响应对象
	resp = map[string]interface{}{
		"adminID":   admin.AdminID,
		"adminName": admin.AdminName,
		"password":  admin.Password,
		"tel":       admin.Tel,
		"mail":      admin.Mail,
		"gender":    admin.Gender,
		"age":       admin.Age,
		"token":     token,
	}
	return
}
