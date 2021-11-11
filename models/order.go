package models

import (
	"time"

)




type Order struct {
	// gorm.Model
	ID                    int64          `json:"Id" gorm:"primary_key"`  
	ClientUsername          string            `json:"clientUsername" gorm:"index; type:varchar(50)"`   
	// Client                Client         
	CookId                int64          `json:"cook_id" gorm:"index"`
	CourierId             int64         `json:"courier" gorm:"foreignKey:OrderId"`
	DeliveryInformationId  int64        `json:"delivery_information_id" gorm:"foreignKey:"`
	Complete             bool           `json:"complete"`
	IsRejected           bool           `json:"is_rejected"`
	RejectReason         string         `json:"reject_reason" gorm:"type:varchar(250)"`
	Items        []Item    `json:"items" gorm:"foreignkey:ID"`
	CreatedAt          time.Time        `gorm:"not null"`    
	UpdatedAt          time.Time         `gorm:"not null"` 
}


func (Order) TableName() string {
    return "orders_order"
}


type Item struct {
	// gorm.Model
	ID         int64       `json:"Id" gorm:"primary_key"`
	MealId      int64    `json:"meal_id" gorm:"index"`
	Quantity    int16     `json:"quantity"` 
	OrderID     uint     `json:"-" gorm:"index"`
}

func (Item) TableName() string {
    return "orders_orderitem"
}

