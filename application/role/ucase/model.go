package ucase

import (
	"taufikRahadi/fiber-boilerplate/database"

	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name string `gorm:"column:name;index"`
}

func InitModel() {
	database.DB.AutoMigrate(&Role{})
}
