package model

type User struct {
	UserID     int    `gorm:"primaryKey;autoIncrement;column:userID"`
	UserName   string `gorm:"type:varchar(30);unique;not null;column:userName"`
	Passwords  string `gorm:"type:varchar(30);not null;column:passwords"`
	SchoolID   int    `gorm:"not null;column:schoolID"`
	Picture    string `gorm:"type:varchar(256);column:picture"`
	Tel        string `gorm:"type:varchar(20);column:tel"`
	Mail       string `gorm:"type:varchar(40);unique;not null;column:mail"`
	Gender     int    `gorm:"type:tinyint;column:gender"`
	UserStatus int    `gorm:"type:tinyint;column:userStatus"`
	School     School `gorm:"foreignKey:SchoolID;references:SchoolID"`
}

func (User) TableName() string {
	return "users"
}
