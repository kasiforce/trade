package model

import "time"

type Goods struct {
	GoodsID        int       `gorm:"primaryKey;autoIncrement;column:goodsID"`
	GoodsName      string    `gorm:"type:varchar(30);not null;column:goodsName"`
	UserID         int       `gorm:"not null;column:userID"`
	Price          float64   `gorm:"type:decimal(10,2);not null;check:price >= 0;column:price"`
	CategoryID     int       `gorm:"not null;column:categoryID"`
	Details        string    `gorm:"type:text;column:details"`
	IsSold         int       `gorm:"type:tinyint;not null;default:0;column:isSold"`
	GoodsImages    string    `gorm:"type:varchar(256);column:goodsImages"`
	CreatedTime    time.Time `gorm:"type:datetime;not null;column:createdTime"`
	DeliveryMethod int       `gorm:"type:int;not null;column:deliveryMethod"`
	//User         User      `gorm:"foreignKey:UserID;references:UserID"`
	//Category     Category  `gorm:"foreignKey:CategoryID;references:CategoryID"`
	//Address      Address   `gorm:"foreignKey:AddressID;references:AddressID"`
	Province     string  `gorm:"size:30;not null;column:province"`
	City         string  `gorm:"size:30;not null;column:city"`
	District     string  `gorm:"size:30;not null;column:districts"`
	Address      string  `gorm:"size:30;not null;column:address"`
	CategoryName string  `gorm:"not null;column:categoryName"`
	UserName     string  `gorm:"type:varchar(30);unique;not null;column:userName"`
	Star         int     `gorm:"foreignKey:GoodsID;references:GoodsID"`
	View         int     `gorm:"-"`
	ShippingCost float64 `gorm:"type:decimal(10,2);not null;default:0;column:shippingCost"`
	IsStarred    bool    `gorm:"type:bool;not null;default:0;column:isStarred"`
	AddrID       int     `gorm:"type:int;not null;default:0;column:addrID"`
	Tel          string  `gorm:"type:string;not null;default:0;column:tel"`
}

func (Goods) TableName() string {
	return "goods"
}
