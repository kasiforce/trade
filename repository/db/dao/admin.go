package dao

import (
	"context"
	"github.com/kasiforce/trade/repository/db/model"
	"gorm.io/gorm"
)

type Admin struct {
	*gorm.DB
}

func NewAdminByDB(db *gorm.DB) *Admin {
	return &Admin{db}
}

func NewAdmin(ctx context.Context) *Admin {
	return &Admin{NewDBClient(ctx)}
}

func (a *Admin) FindAll() (admins []*model.Admin, err error) {
	err = a.DB.Model(&model.Admin{}).Find(&admins).Error
	return
}

func (a *Admin) FindByID(id int) (admin *model.Admin, err error) {
	err = a.DB.Model(&model.Admin{}).Where("adminID = ?", id).First(&admin).Error
	return
}

func (a *Admin) CreateAdmin(admin *model.Admin) (err error) {
	err = a.DB.Model(&model.Admin{}).Create(&admin).Error
	return
}

func (a *Admin) UpdateAdmin(id int, admin *model.Admin) (err error) {
	err = a.DB.Model(&model.Admin{}).Where("adminID = ?", id).Updates(admin).Error
	return
}

func (a *Admin) DeleteAdmin(id int) (err error) {
	err = a.DB.Model(&model.Admin{}).Where("adminID = ?", id).Delete(&model.Admin{}).Error
	return
}
