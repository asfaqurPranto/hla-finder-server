package models

import (
	"time"
)

type User struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Name         string `json:"name"`
	Email        string `json:"email" gorm:"unique"`
	Password     string `json:"-"`
	Admin        bool   `default:"false"`
	LastDontated *time.Time 
	City string `default:"Dhaka"`

	//dob uint `json:"age"`
	
}

type LoginRequset struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type RegisterRequset struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	City string `json:"city" binding:"required"`
}

