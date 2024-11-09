package model

type Category struct {
	CategoryID   int    `gorm:"primary_key;AUTO_INCREMENT;column:categoryID"`
	CategoryName string `gorm:"not null;column:categoryName"`
	Description  string `gorm:"column:descriptions"`
}
