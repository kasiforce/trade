package model

type Address struct {
	ID        int    `gorm:"primaryKey;autoIncrement;column:addrID"`
	UserID    int    `gorm:"index;column:userID"`
	Name      string `gorm:"size:100;not null;column:receiver"`
	Tel       string `gorm:"size:20;not null;column:tel"`
	Province  string `gorm:"size:30;not null;column:province"`
	City      string `gorm:"size:30;not null;column:city"`
	District  string `gorm:"size:30;not null;column:districts"`
	Address   string `gorm:"size:256;not null;column:address"`
	IsDefault int    `gorm:"column:isDefault"`
}
