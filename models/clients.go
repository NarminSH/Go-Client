package models


import (
	"time"
)



type Client struct {
	ID          uint   `json:"id" gorm:"primary_key"` 
	FirstName   string `json:"firstname" gorm:"type:varchar(50)"`
	LastName    string `json:"lastname" gorm:"type:varchar(50)"`
	Patronymic  string `json:"patronymic" gorm:"type:varchar(70)"`
	Username    string `json:"username" gorm:"type:varchar(70)"`
	PhoneNumber string `json:"phone_number" gorm:"type:varchar(70)"`
	Email       string `json:"email" gorm:"type:varchar(100)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}