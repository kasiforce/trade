package model

import (
	"time"
)

type Comment struct {
	CommentID      int       `gorm:"primaryKey;autoIncrement;column:commentID"`
	GoodsID        int       `gorm:"not null;column:goodsID"`
	CommentatorID  int       `gorm:"not null;column:commentatorID"`
	CommentContent string    `gorm:"type:text;not null;column:commentContent"`
	CommentTime    time.Time `gorm:"type:datetime;not null;column:commentTime"`
}

func (Comment) TableName() string {
	return "comment"
}
