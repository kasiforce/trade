package dao

import (
	"context"
	"github.com/kasiforce/trade/repository/db/model"
	"gorm.io/gorm"
)

type Address struct {
	*gorm.DB
}

func NewAddressByDB(db *gorm.DB) *Address {
	return &Address{db}
}

func NewAddress(c context.Context) *Address {
	return &Address{NewDBClient(c)}
}

func (a *Address) AddAddr(addr *model.Address) (id int, err error) {
	err = a.DB.Model(&model.Address{}).Create(&addr).Error
	return addr.ID, err
}

func (a *Address) GetAddrByID(ID int) (addr []*model.Address, err error) {
	err = a.DB.Model(&model.Address{}).Where("userID = ?", ID).Find(&addr).Error
	return
}

func (a *Address) UpdateAddr(i int, addr *model.Address) (err error) {
	err = a.DB.Model(&model.Address{}).Where("addrID = ?", i).Updates(&addr).Error
	return
}

func (a *Address) UpdateDefaultAddr(oldID, newID int) (err error) {
	err = a.DB.Model(&model.Address{}).Where("addrID = ?", oldID).Updates(map[string]interface{}{"isDefault": 0}).Error
	if err != nil {
		return err
	}
	err = a.DB.Model(&model.Address{}).Where("addrID = ?", newID).Updates(map[string]interface{}{"isDefault": 1}).Error
	return
}

func (a *Address) DeleteAddr(ID int) (err error) {
	err = a.DB.Model(&model.Address{}).Where("addrID = ?", ID).Delete(&model.Address{}).Error
	return
}
