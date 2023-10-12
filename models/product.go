package models

import (
    "gorm.io/gorm"
)

type Product struct {
    gorm.Model
    Id int `json:"Id" gorm:"primary_key" `
    Name   string `json:"Name" binding:"required,min=2"`
    Description string `json:"Description" binding:"required,min=2"`
    Price int `json:"Price" binding:"required,min=1"`
  	Carts    []Cart // One-to-Many relationship with Cart
    CategoryId  int // Foreign key to Category
    Category    Category `gorm:"foreignKey:CategoryId;references:Id;onUpdate:CASCADE;onDelete:SET NULL" json:"category"`
}
