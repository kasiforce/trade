package model

import (
	"time"
)

type Announcement struct {
	AnnouncementID int       `gorm:"primaryKey;autoIncrement;column:announcementID"`
	AnTitle        string    `gorm:"type:text;not null;column:anTitle"`
	AnContent      string    `gorm:"type:text;not null;column:anContent"`
	AnTime         time.Time `gorm:"type:datetime;not null;column:anTime"`
}

func (Announcement) TableName() string {
	return "announcement"
}
