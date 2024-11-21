package model

import "time"

type Report struct {
	ReportID    int       `gorm:"primaryKey;autoIncrement;column:reportID"`
	UserID      int       `gorm:"not null;column:userID"`
	GoodsID     int       `gorm:"not null;column:goodsID"`
	Reason      string    `gorm:"type:text;not null;column:reason"`
	CreatedTime time.Time `gorm:"type:datetime;not null;column:createdTime"`
	User        User      `gorm:"foreignKey:UserID;references:UserID"`
	Goods       Goods     `gorm:"foreignKey:GoodsID;references:GoodsID"`
}

func (Report) TableName() string {
	return "report"
}
