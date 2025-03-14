package models

import "gorm.io/gorm"

type Beacon struct {
	gorm.Model
	Mac string
}
