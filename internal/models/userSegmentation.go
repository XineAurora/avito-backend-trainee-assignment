package models

import "gorm.io/gorm"

type UserSegment struct {
	gorm.Model
	UserId    int     `gorm:"not null"`
	SegmentId int     `gorm:"not null,constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Segment   Segment `json:"-"`
}
