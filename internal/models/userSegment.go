package models

import (
	"gorm.io/gorm"
)

type UserSegment struct {
	UserId    int     `gorm:"primaryKey;not null;autoIncrement:false;"`
	SegmentId int     `gorm:"primaryKey;not null;autoIncrement:false;"`
	Segment   Segment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func GetUserActiveSegments(db *gorm.DB, userId int) ([]string, error) {
	userSegments := []string{}
	// SELECT segments.name FROM user_segments JOIN segments ON (segment_id=segments.id) WHERE user_id=$userid;
	tx := db.Model(&UserSegment{}).Select("segments.name").Joins("JOIN segments ON (segment_id=segments.id)").Where("user_id=?", userId).Find(&userSegments)
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

	//TODO: if there is no user segment to delete should return error
	// DELETE FROM user_segments WHERE user_id=$userId AND segment_id IN (SELECT id FROM segments WHERE name IN ($segmentNames));
	tx.Where("user_id=? AND segment_id IN (?)", userId, tx.Model(&Segment{}).Select("id").Where("name IN (?)", exclude)).Delete(&UserSegment{})
	if err := tx.Error; err != nil {
		tx.Rollback()
		return err
	}

	//TODO: create new rows in database
	// tx.Create()

	return tx.Commit().Error
}
