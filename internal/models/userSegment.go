package models

import (
	"errors"

	"gorm.io/gorm"
)

type UserSegment struct {
	UserId    int     `gorm:"primaryKey;not null;autoIncrement:false;"`
	SegmentId int     `gorm:"primaryKey;not null;autoIncrement:false;"`
	Segment   Segment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func GetUserActiveSegments(db *gorm.DB, userId int) ([]Segment, error) {
	userSegments := []Segment{}
	// SELECT segments.id, segments.name FROM user_segments JOIN segments ON (segment_id=segments.id) WHERE user_id=$userid;
	tx := db.Model(&UserSegment{}).Select("segments.id, segments.name").Joins("JOIN segments ON (segment_id=segments.id)").Where("user_id=?", userId).Find(&userSegments)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return userSegments, nil
}

func UpdateUserActiveSegments(db *gorm.DB, userId int, include []string, exclude []string) error {
	tx := db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	for _, segmentName := range exclude {
		var segment Segment
		if err := tx.Where("name=?", segmentName).First(&segment).Error; err != nil {
			tx.Rollback()
			return errors.New("one of `exclude` segments doesn't exist")
		}
		userSegment := UserSegment{UserId: userId, Segment: segment}
		if err := tx.Where("user_id=? AND segment_id=?", userId, segment.ID).First(&userSegment).Error; err != nil {
			tx.Rollback()
			return errors.New("user don't participate in one of `exclude` segments")
		} else {
			tx.Delete(&userSegment)
		}
	}

	for _, segmentName := range include {
		var segment Segment
		if err := tx.Where("name=?", segmentName).First(&segment).Error; err != nil {
			tx.Rollback()
			return errors.New("one of `include` segments doesn't exist")
		}
		userSegment := UserSegment{UserId: userId, Segment: segment}
		if err := tx.Create(&userSegment).Error; err != nil {
			tx.Rollback()
			return errors.New("user already participate in one of `include` segments")
		}
	}

	return tx.Commit().Error
}
