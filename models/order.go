package models

import (
	"time"

)




type Order struct {
	// gorm.Model
	ID                    int64          `json:"Id" gorm:"primary_key"`  
	ClientId              int64            `json:"client_id" gorm:"index"`   
	Client                Client         `json:"client" gorm:"foreignKey:ClientId"`
	CookId                int64          `json:"cook_id" gorm:"index"`
	CourierId             int64         `json:"courier" gorm:"foreignKey:OrderId"`
	DeliveryInformationId  int64        `json:"delivery_information_id" gorm:"foreignKey:"`
	CustomerFirstName      string       `json:"customer_first_name" gorm:"type:varchar(150); not null"`
	CustomerLastName      string        `json:"customer_last_name" gorm:"type:varchar(150); not null"`
	CustomerPhone         string        `json:"customer_phone" gorm:"type:varchar(12); not null"`
	CustomerLocation      string        `json:"customer_location" gorm:"type:varchar(150); not null"`
	CustomerEmail        string         `json:"customer_email" gorm:"type:varchar(254); not null"`
	Complete             bool           `json:"complete"`
	IsRejected           bool           `json:"is_rejected"`
	RejectReason         string         `json:"reject_reason" gorm:"type:varchar(250)"`
	CreatedAt          time.Time        `gorm:"not null"`    
	UpdatedAt          time.Time         `gorm:"not null"` 
}


func (Order) TableName() string {
    return "orders_order"
}


// type Item struct {
// 	// gorm.Model
// 	ItemID      uint     `json:"ItemId" gorm:"primary_key"`
// 	MealId      int64    `json:"meal_id" gorm:"index"`
// 	Quantity    uint     `json:"quantity"` 
// 	OrderID     uint     `json:"-"`
// }

