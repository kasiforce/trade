package dao

import (
	"context"
	"github.com/kasiforce/trade/repository/db/model"
	"github.com/kasiforce/trade/types"
	"gorm.io/gorm"
)

type User struct {
	*gorm.DB
}

func NewUserByDB(db *gorm.DB) *User {
	return &User{db}
}

func NewUser(ctx context.Context) *User {
	return &User{NewDBClient(ctx)}
}

func (user *User) CountAll() (total int64, err error) {
	err = user.DB.Model(&model.User{}).Count(&total).Error
	return
}
func (user *User) FindAll(req types.ShowUserReq) (users []model.User, err error) {
	db := user.DB
	query := db.Table("users").Select("users.*").Preload("School")
	if req.SearchQuery != "" {
		query = query.Where("users.userName LIKE ?", "%"+req.SearchQuery+"%")
	}
	query = query.Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize)
	err = query.Find(&users).Error
	return
}

func (user *User) FindByID(id int) (u *model.User, err error) {
	err = user.DB.Model(&model.User{}).Preload("School").Where("userID = ?", id).First(&u).Error
	return
}

func (user *User) FindByName(name string) (exist bool, err error) {
	var cnt int64
	err = user.DB.Model(&model.User{}).Where("userName = ?", name).Count(&cnt).Error
	if cnt == 0 {
		return false, err
	}
	return true, err
}

func (user *User) FindByMail(mail string) (exist bool, err error) {
	var cnt int64
	err = user.DB.Model(&model.User{}).Where("mail = ?", mail).Count(&cnt).Error
	if cnt == 0 {
		return false, err
	}
	return true, err
}

func (user *User) CreateUser(u *model.User) (id int, err error) {
	result := user.DB.Model(&model.User{}).Create(u)
	if err = result.Error; err != nil {
		return 0, err
	}
	id = u.UserID
	return
}

func (user *User) UpdateUser(id int, u map[string]interface{}) (err error) {
	err = user.DB.Model(&model.User{}).Where("userID = ?", id).Updates(u).Error
	return
}

func (user *User) DeleteUser(id int) (err error) {
	err = user.DB.Model(&model.User{}).Delete(&model.User{}, id).Error
	return
}

// CheckMail 登录时检查邮箱是否存在，若存在查出用户名、id、密码
func (user *User) CheckMail(mail string) (u *model.User, err error) {
	err = user.DB.Model(&model.User{}).Preload("School").Where("mail = ?", mail).First(&u).Error
	return
}

func (user *User) UpdatePwd(mail string, pwd string) (err error) {
	err = user.DB.Model(&model.User{}).Where("mail = ?", mail).Update("passwords", pwd).Error
	return
}
