package models

import "gorm.io/gorm"

type User struct{
	gorm.Model
	Id int `json:"Id" gorm:"primary_key" `
	Name string `json:"Name" uri:"name" binding:"required,min=3"`
	Email string `json:"Email" binding:"required,min=3"`
	Password string `json:"Password" binding:"required,min=3"`
	Carts    []Cart // One-to-Many relationship with Cart
}
