package models

import "gorm.io/gorm"

type UserSegment struct {
	UserId    int     `gorm:"primaryKey;not null;autoIncrement:false;"`
	SegmentId int     `gorm:"primaryKey;not null;autoIncrement:false;"`
	Segment   Segment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func GetUserActiveSegments(db *gorm.DB) ([]UserSegment, error) {
	return nil, nil
}

func UpdateUserActiveSegments(db *gorm.DB, include []string, exclude []string) error {
	return nil
}
