package model

import "time"

type Collection struct {
	GoodsID     int       `gorm:"primary_key;column:goodsID"`
	UserID      int       `gorm:"primary_key;column:userID"`
	CreatedTime time.Time `gorm:"not null;column:createdTime"`
	User        User      `gorm:"foreignKey:UserID;references:UserID"`
	Goods       Goods     `gorm:"foreignKey:GoodsID;references:GoodsID"`
}

func (Collection) TableName() string {
	return "collection"
}
