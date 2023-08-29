package models

import (
	"gorm.io/gorm"
)

type Segment struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}

func CreateSegment(db *gorm.DB, name string) (*Segment, error) {
	segment := Segment{Name: name}
	// if segment was deleted before, restore it
	if err := db.Unscoped().Where("name=?", name).First(&segment).Error; err == nil && segment.DeletedAt.Valid {
		err = db.Unscoped().Model(&segment).Update("deleted_at", nil).Error
		if err != nil {
			return nil, err
		}
		return &segment, nil
	}
	//
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
