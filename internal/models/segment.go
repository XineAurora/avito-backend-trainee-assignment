package models

import "gorm.io/gorm"

type Segment struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}

func CreateSegment(db *gorm.DB, name string) (*Segment, error) {
	segment := Segment{Name: name}
	if tx := db.Create(&segment); tx.Error != nil {
		return nil, tx.Error
	}
	return &segment, nil
}

func DeleteSegment(db *gorm.DB, name string) error {
	tx := db.Where("name=?", name).Delete(&Segment{})
	return tx.Error
}
