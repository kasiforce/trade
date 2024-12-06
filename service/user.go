package service

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kasiforce/trade/pkg/ctl"
	"github.com/kasiforce/trade/pkg/util"
	"github.com/kasiforce/trade/repository/db/dao"
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

func (s *UserService) ShowAllUser(ctx context.Context, req types.ShowUserReq) (resp interface{}, err error) {
	user := dao.NewUser(ctx)
	userList, err := user.FindAll(req)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	var respList []types.UserInfo
	for _, userInfo := range userList {
		respList = append(respList, types.UserInfo{
			UserID:     userInfo.UserID,
			UserName:   userInfo.UserName,
			Password:   userInfo.Passwords,
			SchoolName: userInfo.School.SchoolName,
			Picture:    userInfo.Picture,
			Tel:        userInfo.Tel,
			Mail:       userInfo.Mail,
			Gender:     userInfo.Gender,
			Status:     userInfo.UserStatus,
		})
	}
	var response types.UserListResp
	response.UsersList = respList
	response.PageNum = req.PageNum
	response.Total = len(respList)
	return response, nil
}

//func (s *UserService) ShowUserInfoByID(ctx context.Context) (resp interface{}, err error) {
//	u, err := ctl.GetUserID(ctx)
//	if err != nil {
//		util.LogrusObj.Error(err)
//		return
//	}
//	user := dao.NewUser(ctx)
//	userInfo, err := user.FindByID(u.UserID)
//	if err != nil {
//		util.LogrusObj.Error(err)
//		return
//	}
//	resp = &types.UserInfoResp{
//		UserID:     userInfo.UserID,
//		UserName:   userInfo.UserName,
//		Password:   userInfo.Passwords,
//		SchoolName: userInfo.SchoolName,
//		Picture:    userInfo.Picture,
//		Tel:        userInfo.Tel,
//		Mail:       userInfo.Mail,
//		Gender:     userInfo.Gender,
//		Status:     userInfo.UserStatus,
//	}
//	return
//}

func (s *UserService) AddUser(c context.Context, req types.UserInfo) (resp interface{}, err error) {
	if req.UserName == "" || req.Password == "" || req.SchoolName == "" || req.Mail == "" {
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
	school := dao.NewSchool(c)
	id, err := school.FindSchoolID(req.SchoolName)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	modelUser := map[string]interface{}{
		"userName":   req.UserName,
		"passwords":  req.Password,
		"schoolID":   id,
		"picture":    req.Picture,
		"mail":       req.Mail,
		"gender":     req.Gender,
		"tel":        req.Tel,
		"userStatus": req.Status,
	}
	err = u.CreateUser(modelUser)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return
}

// UpdateUser 管理员更新用户信息
func (s *UserService) UpdateUser(c context.Context, req types.UserInfo) (resp interface{}, err error) {
	//if req.UserName == "" || req.SchoolID == 0 || req.Mail == "" {
	//	err = errors.New("参数不能为空")
	//	return
	//}
	user, err := ctl.GetUserID(c)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	school := dao.NewSchool(c)
	id, err := school.FindSchoolID(req.SchoolName)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	u := dao.NewUser(c)
	modelUser := map[string]interface{}{
		"userName":   req.UserName,
		"passwords":  req.Password,
		"schoolID":   id,
		"picture":    req.Picture,
		"mail":       req.Mail,
		"gender":     req.Gender,
		"tel":        req.Tel,
		"userStatus": req.Status,
	}
	for key, value := range modelUser {
		if value == "" {
			delete(modelUser, key)
		}
	}
	err = u.UpdateUser(user.UserID, modelUser)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return
}

func (s *UserService) DeleteUser(c context.Context) (resp interface{}, err error) {
	user, err := ctl.GetUserID(c)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	u := dao.NewUser(c)
	err = u.DeleteUser(user.UserID)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return
}

func (s *UserService) ShowIntroduction(c *gin.Context) (resp interface{}, err error) {
	id := c.GetInt("id")
	u := dao.NewUser(c)
	user, err := u.FindByID(id)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	resp = types.IntroductionResp{
		UserID:   user.UserID,
		UserName: user.UserName,
		Picture:  user.Picture,
		Gender:   user.Gender,
		School:   user.School.SchoolName,
	}
	return
}

func (s *UserService) ShowUserByID(c *gin.Context) (resp interface{}, err error) {
	id := c.GetInt("id")
	u := dao.NewUser(c)
	user, err := u.FindByID(id)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	resp = types.UserInfo{
		UserID:     user.UserID,
		UserName:   user.UserName,
		Password:   user.Passwords,
		SchoolName: user.School.SchoolName,
		Picture:    user.Picture,
		Tel:        user.Tel,
		Mail:       user.Mail,
		Gender:     user.Gender,
		Status:     user.UserStatus,
	}
	return
}

// Update 用户自己修改信息
func (s *UserService) Update(c context.Context, req types.UpdateUserReq) (resp interface{}, err error) {
	u := dao.NewUser(c)
	modelUser := map[string]interface{}{
		"userName":  req.UserName,
		"passwords": req.Password,
		"picture":   req.Picture,
		"gender":    req.Gender,
		"tel":       req.Tel,
	}
	err = u.UpdateUser(req.UserID, modelUser)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return
}

func (s *UserService) UserLogin(c *gin.Context, req types.UserLoginReq) (resp interface{}, err error) {
	if req.Mail == "" || req.Password == "" {
		err = errors.New("参数不能为空")
		return
	}
	u := dao.NewUser(c)
	user, err := u.CheckMail(req.Mail)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	if user.Passwords != req.Password {
		err = errors.New("密码错误")
		return
	}
	token, err := util.GenerateToken(user.UserID, user.UserName)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	resp = types.UserWithToken{
		UserID:     user.UserID,
		UserName:   user.UserName,
		Password:   user.Passwords,
		SchoolName: user.School.SchoolName,
		Picture:    user.Picture,
		Tel:        user.Tel,
		Mail:       user.Mail,
		Gender:     user.Gender,
		Status:     user.UserStatus,
		Token:      token,
	}
	//middleware.SetToken(c, token)
	return
}
