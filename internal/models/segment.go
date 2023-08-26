package models

import "gorm.io/gorm"

type Segment struct {
	gorm.Model
	Name string `gorm:"unique,not null"`
}
