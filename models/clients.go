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
	Location    string  `json:"location" gorm:"type:varchar(250)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}	


// type Cook struct {
// 	ID             uint   `json:"id" gorm:"primary_key"` 
// 	patronymic	  string `json:"firstname" gorm:"type:varchar(50); not null"`	
// 	username      string `json:"username" gorm:"unique;type:varchar(50); not null"`
// 	first_name	  string `json:"firstname" gorm:"type:varchar(50); not null"`	
// 	last_name	  string    `json:"lastname" gorm:"type:varchar(50); not null"`
// 	email	      string `json:"email" gorm:"type:varchar(50); not null"`
// 	phone     	    string `json:"phone" gorm:"type:varchar(20)"`	
// 	user_type       string  `json:"usertype" gorm:"type:varchar(4)"`
// 	birth_place	    string  `json:"birth_place" gorm:"type:varchar(50)"`
// 	city	        string  `json:"city" gorm:"type:varchar(50)"`
// 	service_place	string  `json:"service_place" gorm:"type:varchar(100)"`
// 	payment_address	
// 	work_experience		
// 	rating	
// 	is_available	
// 	created_at	timestamptz	
// 	updated_at	timestamptz	

// }