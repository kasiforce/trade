package model

type Admin struct {
	AdminID   int    `gorm:"primaryKey;autoIncrement;column:adminID"`
	Password  string `gorm:"type:varchar(30);not null;column:passwords"`
	AdminName string `gorm:"type:varchar(30);not null;column:adminName"`
	Tel       string `gorm:"type:varchar(20);column:tel"`
	Mail      string `gorm:"type:varchar(40);unique;not null;column:mail"`
	Gender    int    `gorm:"type:tinyint;column:gender"`
	Age       int    `gorm:"type:int;check:age >= 0;column:age"`
}

func (Admin) TableName() string {
	return "admins"
}
