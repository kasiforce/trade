package dao

import (
	"context"
	"github.com/kasiforce/trade/repository/db/model"
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

func (user *User) FindAll() (users []*model.User, err error) {
	err = user.DB.Model(&model.User{}).Find(&users).Error
	return
}

func (user *User) FindByID(id int) (u *model.User, err error) {
	err = user.DB.Model(&model.User{}).Where("userID = ?", id).First(&u).Error
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
func (user *User) CreateUser(u *model.User) (err error) {
	err = user.DB.Model(&model.User{}).Create(&u).Error
	return
}

func (user *User) UpdateUser(id int, u *model.User) (err error) {
	err = user.DB.Model(&model.User{}).Where("userID = ?", id).Updates(&u).Error
	return
}
