package models

import (
	"time"

	"github.com/shopspring/decimal"
)



type Client struct {
	ID          uint   `json:"id" gorm:"primary_key"` 
	FirstName   string `json:"firstname" gorm:"type:varchar(50); not null"`
	LastName    string `json:"lastname" gorm:"type:varchar(50); not null"`
	Patronymic  string `json:"patronymic" gorm:"type:varchar(50)"`
	Username    string `json:"username" gorm:"unique;not null;type:varchar(50)"`
	UserType    string  `json:"usertype" gorm:"type:varchar(4)"`
	PhoneNumber string `json:"phone_number" gorm:"type:varchar(20)"`
	Email       string `json:"email" gorm:"type:varchar(50); not null"`
	Location    string  `json:"location" gorm:"type:varchar(250)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}	



type Base struct{
	ID             uint   `json:"id" gorm:"primary_key"` 
	Patronymic	  string   `json:"patronymic" gorm:"type:varchar(60)"`	
	Username      string   `json:"username" gorm:"unique;type:varchar(200)"`
	FirstName	  string   `json:"firstname" gorm:"type:varchar(150)"`	
	LastName	  string    `json:"lastname" gorm:"type:varchar(150)"`
	Email	      string `json:"email" gorm:"type:varchar(254)"`
	UserType       string  `json:"usertype" gorm:"type:varchar(5)"`
	WorkExperience  int         `json:"work_experience"`
	IsAvailable	 bool    `json:"is_available"`
	Rating	         decimal.Decimal  `json:"rating" gorm:"type:decimal(2,1);"`
	CreatedAt	time.Time	  `json:"created_at" gorm:"not null"`
	UpdatedAt	time.Time	  `json:"updated_at" gorm:"not null"`
}



type Cook struct {
	Base
	Phone     	    string `json:"phone" gorm:"type:varchar(20)"`	
	BirthPlace	    string  `json:"birth_place" gorm:"type:varchar(50)"`
	City	        string  `json:"city" gorm:"type:varchar(50)"`
	ServicePlace	string  `json:"service_place" gorm:"type:varchar(100)"`
	PaymentAddress	string   `json:"payment_address" gorm:"type:varchar(255)"`
}
func (Cook) TableName() string {
	return "cooks_cook"
}

type Courier struct {
	Base
	Phone     	    string `json:"phone" gorm:"type:varchar(12); not null"`	
	Transport	    string  `json:"transport" gorm:"type:varchar(150)"`
	Location    string  `json:"location" gorm:"type:varchar(255)"`
}

func (Courier) TableName() string {
	return "delivery_courier"
}