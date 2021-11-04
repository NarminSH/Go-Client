package models

import "time"




type Order struct {
	// gorm.Model
	OrderID      uint      `json:"orderId" gorm:"primary_key"`
	CustomerName string    `json:"customerName"`
	CreatedAt    time.Time `json:"createdAt"`
	Items        []Item    `json:"items" gorm:"foreignkey:OrderID"`
}



type Item struct {
	// gorm.Model
	ItemID       uint   `json:"ItemId" gorm:"primary_key"`
	MealId      uint     `json:"-"`
	Quantity    uint     `json:"quantity"`
	OrderID     uint     `json:"-"`
}