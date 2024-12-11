package dao

import (
	"context"

	"github.com/kasiforce/trade/repository/db/model"
	"github.com/kasiforce/trade/types"
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

func (admin *Admin) CountAll() (total int64, err error) {
	err = admin.DB.Model(&model.Admin{}).Count(&total).Error
	return
}

func (admin *Admin) FindAll(req types.ShowAdminReq) (admins []model.Admin, err error) {
	db := admin.DB
	query := db.Table("admins").Select("admins.*")
	if req.SearchQuery != "" {
		query = query.Where("admins.adminName LIKE ?", "%"+req.SearchQuery+"%")
	}
	query = query.Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize)
	err = query.Find(&admins).Error
	return
}

func (admin *Admin) FindByID(id int) (a *model.Admin, err error) {
	err = admin.DB.Model(&model.Admin{}).Where("adminID = ?", id).First(&a).Error
	return
}

func (admin *Admin) FindByName(name string) (exist bool, err error) {
	var cnt int64
	err = admin.DB.Model(&model.Admin{}).Where("adminName = ?", name).Count(&cnt).Error
	if cnt == 0 {
		return false, err
	}
	return true, err
}

func (admin *Admin) FindByMail(mail string) (exist bool, err error) {
	var cnt int64
	err = admin.DB.Model(&model.Admin{}).Where("mail = ?", mail).Count(&cnt).Error
	if cnt == 0 {
		return false, err
	}
	return true, err
}

func (admin *Admin) CreateAdmin(data map[string]interface{}) (adminID int, err error) {
	result := admin.DB.Model(&model.Admin{}).Create(data)
	if result.Error != nil {
		return 0, result.Error
	}
	// 获取插入记录的主键 ID
	createdAdmin := model.Admin{}
	if err := admin.DB.Last(&createdAdmin).Error; err != nil {
		return 0, err
	}

	adminID = createdAdmin.AdminID // 直接获取插入记录的主键ID
	return adminID, nil
}

func (admin *Admin) UpdateAdmin(id int, data map[string]interface{}) (err error) {
	err = admin.DB.Model(&model.Admin{}).Where("adminID = ?", id).Updates(data).Error
	return
}

func (admin *Admin) DeleteAdmin(id int) (err error) {
	err = admin.DB.Model(&model.Admin{}).Delete(&model.Admin{}, id).Error
	return
}

// CheckMail 登录时检查邮箱是否存在，若存在查出用户名、id、密码
func (admin *Admin) CheckMail(mail string) (a *model.Admin, err error) {
	err = admin.DB.Model(&model.Admin{}).Where("mail = ?", mail).First(&a).Error
	return
}
