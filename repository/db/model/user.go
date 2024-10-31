package model

type User struct {
	userID     int    `gorm:"primaryKey;autoIncrement;column:userID"`
	userName   string `gorm:"type:varchar(30);unique;not null;column:userName"`
	passwords  string `gorm:"type:varchar(30);not null;column:passwords"`
	schoolID   int    `gorm:"not null;column:schoolID"`
	picture    string `gorm:"type:varchar(256);column:picture"`
	tel        string `gorm:"type:varchar(20);column:tel"`
	mail       string `gorm:"type:varchar(40);unique;not null;column:mail"`
	gender     int    `gorm:"type:tinyint;column:gender"`
	userStatus int    `gorm:"type:tinyint;column:userStatus;default:1"`
}

func (User) TableName() string {
	return "users"
}
