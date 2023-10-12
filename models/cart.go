package models

import (
    "gorm.io/gorm"
)

type Cart struct {
    gorm.Model
    Id int `json:"Id" gorm:"primary_key" `
    UserId  int // Foreign key to User
    User    User `gorm:"foreignKey:UserId;references:Id;onUpdate:CASCADE;onDelete:SET NULL"`
		ProductId  int // Foreign key to Product
    Product    Product `gorm:"foreignKey:ProductId;references:Id;onUpdate:CASCADE;onDelete:SET NULL"`
}
