package models

import "time"




type Order struct {
	// gorm.Model
	OrderID       uint       `json:"orderId" gorm:"primary_key"`
	Client        Client      `json:"client" gorm:"foreignKey:OrderId"`
	CookId         int64        `json:"cook" gorm:"foreignKey:OrderId"`
	CourierId       int64     `json:"courier" gorm:"foreignKey:OrderId"`
	CreatedAt    time.Time    `json:"createdAt"`
	Items        []Item        `json:"items" gorm:"foreignKey:OrderID"`
}



type Item struct {
	// gorm.Model
	ItemID      uint     `json:"ItemId" gorm:"primary_key"`
	MealId      int64    `json:"meal_id" gorm:"index"`
	Quantity    uint     `json:"quantity"` 
	OrderID     uint     `json:"-"`
}

