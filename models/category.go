package models

import (
    "gorm.io/gorm"
)

type Category struct {
    gorm.Model
    Id int `json:"Id" gorm:"primary_key" `
    Name        string `json:"Name"  binding:"required"`
    Description string `json:"Description"  binding:"required"`
  	Products    []Product // One-to-Many relationship with Product
}
