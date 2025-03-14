package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Rut      string `gorm:"unique:TRUE;"`
	Nombre   string
	Email    string `gorm:"unique:TRUE"`
	Celular  string
	BeaconID int
	Beacon   Beacon `gorm:"foreignKey:Mac;references:Mac;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
