package models


import (
	"time"
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
	CreatedAt   time.Time
	UpdatedAt   time.Time
}	