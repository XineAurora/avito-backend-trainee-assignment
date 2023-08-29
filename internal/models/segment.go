package models

import (
	"gorm.io/gorm"
)

type Segment struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"unique;not null"`
}

func CreateSegment(db *gorm.DB, name string) (*Segment, error) {
	segment := Segment{Name: name}
	if err := db.Create(&segment).Error; err != nil {
		return nil, err
	}
	return &segment, nil
}

func DeleteSegment(db *gorm.DB, name string) error {
	var segment Segment
	if err := db.Where("name=?", name).First(&segment).Error; err != nil {
		return err
	}
	if err := db.Delete(&segment).Error; err != nil {
		return err
	}
	return nil
}
