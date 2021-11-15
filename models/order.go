package models

import (
	"time"
)



type Order struct {
	// gorm.Model
	ID             int64  `json:"Id" gorm:"primary_key"`
	ClientID       int64   `json:"client_id"`
	// Client        ClientID
	CookId                int64     `json:"cook_id"`
	CourierId             int64     `json:"courier_id"`
	DeliveryInformationId int64     `json:"delivery_information_id"`
	Complete              bool      `json:"complete" gorm:"not null"`
	IsRejected            bool      `json:"is_rejected"`
	RejectReason          string    `json:"reject_reason" gorm:"type:varchar(250)"`
	Items                 []Item    `json:"items" gorm:"foreignkey:ID"`
	CreatedAt             time.Time `gorm:"not null"`
	UpdatedAt             time.Time `gorm:"not null"`
}


func (Order) TableName() string {
	return "orders_order"
}


type Item struct {
	// gorm.Model
	ID       int64 `json:"Id" gorm:"primary_key"`
	MealId   int64 `json:"meal_id"`
	Quantity int16 `json:"quantity"`
	OrderID  int64  `json:"-"`
}


func (Item) TableName() string {
	return "orders_orderitem"
}
