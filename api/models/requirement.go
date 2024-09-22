package models

import (
	"gorm.io/gorm"
	"log"
	postgresgorm "merinio/pkg"
	"time"
)

type RequirementManager struct {
	db *gorm.DB
}

type Requirement struct {
	ID       int       `gorm:"primaryKey" json:"id"`
	Name     string    `gorm:"size:255;not null" json:"name"`
	Created  time.Time `gorm:"autoCreateTime" json:"created"`
	Modified time.Time `gorm:"autoUpdateTime" json:"modified"`
	State    int       `gorm:"default:1" json:"state"`
}

func (Requirement) TableName() string {
	return "requirements"
}

func GetRequirementManager(db *gorm.DB) RequirementManager {
	if db != nil {
		return RequirementManager{db: db}
	}
	return RequirementManager{postgresgorm.GetConnection()}
}

func (rs RequirementManager) FindAllWhere(requirementIds []int) ([]Requirement, error) {
	var list []Requirement

	result := rs.db.Find(&list, requirementIds)
	if result.Error != nil {
		log.Printf("Fail to retrieve requirements: %v", result.Error)
		return nil, result.Error
	}
	return list, nil
}

func (rs RequirementManager) MergeRequirements(current, parent []Requirement) []Requirement {
	reqMap := make(map[int]Requirement)

	// Add current requirements to the map
	for _, req := range current {
		reqMap[req.ID] = req
	}

	for _, req := range parent {
		if _, exists := reqMap[req.ID]; !exists {
			reqMap[req.ID] = req
		}
	}

	merged := make([]Requirement, 0, len(reqMap))
	for _, req := range reqMap {
		merged = append(merged, req)
	}

	return merged
}
