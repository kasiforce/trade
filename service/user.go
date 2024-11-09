package service

import (
	"context"
	"errors"
	"github.com/kasiforce/trade/pkg/ctl"
	"github.com/kasiforce/trade/pkg/util"
	"github.com/kasiforce/trade/repository/db/dao"
	"github.com/kasiforce/trade/repository/db/model"
	"github.com/kasiforce/trade/types"
	"sync"
)

var userServ *UserService
var userServOnce sync.Once

type UserService struct {
}

func GetUserService() *UserService {
	userServOnce.Do(func() {
		userServ = &UserService{}
	})
	return userServ
}

func (s *UserService) ShowAllUser(ctx context.Context) (resp interface{}, err error) {
	user := dao.NewUser(ctx)
	userList, err := user.FindAll()
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	var respList []types.UserInfoResp
	for _, userInfo := range userList {
		respList = append(respList, types.UserInfoResp{
			UserID:   userInfo.UserID,
			UserName: userInfo.UserName,
			Password: userInfo.Passwords,
			SchoolID: userInfo.SchoolID,
			Picture:  userInfo.Picture,
			Tel:      userInfo.Tel,
			Mail:     userInfo.Mail,
			Gender:   userInfo.Gender,
			Status:   userInfo.UserStatus,
		})
	}
	return respList, nil
}

func (s *UserService) ShowUserInfoByID(ctx context.Context) (resp interface{}, err error) {
	u, err := ctl.GetUserID(ctx)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	user := dao.NewUser(ctx)
	userInfo, err := user.FindByID(u.UserID)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	resp = &types.UserInfoResp{
		UserID:   userInfo.UserID,
		UserName: userInfo.UserName,
		Password: userInfo.Passwords,
		SchoolID: userInfo.SchoolID,
		Picture:  userInfo.Picture,
		Tel:      userInfo.Tel,
		Mail:     userInfo.Mail,
		Gender:   userInfo.Gender,
		Status:   userInfo.UserStatus,
	}
	return
}

func (s *UserService) AddUser(c context.Context, req *types.UserAddReq) (resp interface{}, err error) {
	if req.UserName == "" || req.Password == "" || req.SchoolID == 0 || req.Mail == "" {
		err = errors.New("参数不能为空")
		return
	}
	u := dao.NewUser(c)
	exist, err := u.FindByName(req.UserName)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	if exist {
		err = errors.New("用户名已存在")
		return
	}
	modelUser := &model.User{
		UserName:  req.UserName,
		Passwords: req.Password,
		SchoolID:  req.SchoolID,
		Mail:      req.Mail,
	}
	err = u.CreateUser(modelUser)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return
}

func (s *UserService) UpdateUser(c context.Context, req *types.UserInfoUpdateReq) (resp interface{}, err error) {
	//if req.UserName == "" || req.SchoolID == 0 || req.Mail == "" {
	//	err = errors.New("参数不能为空")
	//	return
	//}
	user, err := ctl.GetUserID(c)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	u := dao.NewUser(c)
	modelUser := &model.User{
		UserName:   req.UserName,
		SchoolID:   req.SchoolID,
		Picture:    req.Picture,
		Tel:        req.Tel,
		Mail:       req.Mail,
		Gender:     req.Gender,
		UserStatus: req.Status,
	}
	err = u.UpdateUser(user.UserID, modelUser)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return
}
