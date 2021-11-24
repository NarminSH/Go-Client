package models

import (
	"time"
)



type Order struct {
	// gorm.Model
	ID             int64  `json:"Id" gorm:"primary_key"`
	ClientID       int   `json:"client_id"`
	// Client        Client
	CookId                int64     `json:"cook_id"`
	// Cook      []Cook
	CourierId             int64     `json:"courier_id"`
	// Courier       []Courier
	DeliveryInformationId int64     `json:"delivery_information_id"`
	Complete              bool      `json:"complete" gorm:"not null"`
	IsRejected            bool      `json:"is_rejected"`
	RejectReason          string    `json:"reject_reason" gorm:"type:varchar(250)"`
	Items                 []Item    `json:"items" gorm:"foreignkey:ID"`
	CreatedAt             time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt             time.Time `json:"updated_at" gorm:"not null"`
}


func (Order) TableName() string {
	return "orders_order"
}


type Item struct {
	// gorm.Model
	ID         int64      `json:"Id" gorm:"primary_key"`
	MealId     int64      `json:"meal_id" gorm:"not null"`
	Quantity   int16      `json:"quantity" gorm:"not null"`
	OrderID    int64      `json:"-" gorm:"not null"`
}


func (Item) TableName() string {
	return "orders_orderitem"
}
