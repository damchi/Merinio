package models

import (
	"gorm.io/gorm"
	"log"
	postgresgorm "merinio/pkg"
	"time"
)

type RestrictionManager struct {
	db *gorm.DB
}

type Restriction struct {
	ID       int       `gorm:"primaryKey" json:"id"`
	Name     string    `gorm:"size:255;not null" json:"name"`
	Created  time.Time `gorm:"autoCreateTime" json:"created"`
	Modified time.Time `gorm:"autoUpdateTime" json:"modified"`
	State    int       `gorm:"default:1" json:"state"`
}

func (Restriction) TableName() string {
	return "restrictions"
}

func GetRestrictionManager(db *gorm.DB) RestrictionManager {
	if db != nil {
		return RestrictionManager{db: db}
	}
	return RestrictionManager{postgresgorm.GetConnection()}
}

func (rs RestrictionManager) FindAllWhere(restrictionIds []int) ([]Restriction, error) {
	var list []Restriction

	result := rs.db.Find(&list, restrictionIds)
	if result.Error != nil {
		log.Printf("Fail to retrieve restriction: %v", result.Error)
		return nil, result.Error
	}
	return list, nil
}

func (rs RestrictionManager) MergeRestrictions(current, parent []Restriction) []Restriction {
	restrMap := make(map[int]Restriction)

	// Add current restrictions to the map
	for _, restr := range current {
		restrMap[restr.ID] = restr
	}

	for _, restr := range parent {
		if _, exists := restrMap[restr.ID]; !exists {
			restrMap[restr.ID] = restr
		}
	}

	merged := make([]Restriction, 0, len(restrMap))
	for _, restr := range restrMap {
		merged = append(merged, restr)
	}

	return merged
}
