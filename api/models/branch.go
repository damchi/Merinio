package models

import (
	"gorm.io/gorm"
	"log"
	postgresgorm "merinio/pkg"
	"time"
)

type branchManagerImpl struct {
	db *gorm.DB
}

type BranchManager interface {
	FindAll() ([]Branch, error)
	Find(branchId int) (*Branch, error)
	Save(branch *Branch) (*Branch, error)
	FindByParentId(parentID *int) (Branch, error)
}

type Branch struct {
	ID           int           `gorm:"primaryKey" json:"id"`
	Name         string        `gorm:"size:255;not null" json:"name"`
	ParentID     *int          `json:"parent_id"`
	Parent       *Branch       `gorm:"foreignKey:ParentID" json:"parent"`
	Created      time.Time     `gorm:"autoCreateTime" json:"created"`
	Modified     time.Time     `gorm:"autoUpdateTime" json:"modified"`
	State        int           `gorm:"default:1" json:"state"`
	IsRoot       bool          `gorm:"default:false" json:"is_root"`
	Requirements []Requirement `gorm:"many2many:branch_requirements" json:"requirements"`
	Restrictions []Restriction `gorm:"many2many:branch_restrictions" json:"restrictions"`
}

type BranchRequest struct {
	Requirement []int  `json:"requirements"`
	Restriction []int  `json:"restrictions"`
	Name        string `gorm:"size:255;not null" json:"name"`
	ParentID    *int   `json:"parent_id"`
	IsRoot      bool   `json:"is_root"`
}

func (Branch) TableName() string {
	return "branches"
}

func GetBranchManager(db *gorm.DB) BranchManager {
	if db != nil {
		return &branchManagerImpl{db: db}
	}
	return &branchManagerImpl{db: postgresgorm.GetConnection()}
}

func (m *branchManagerImpl) FindAll() ([]Branch, error) {
	var branches []Branch

	result := m.db.Preload("Requirements").Preload("Restrictions").Find(&branches)
	if result.Error != nil {
		log.Printf("Fail to retrieve branches: %v", result.Error)
		return nil, result.Error
	}
	return branches, nil
}

func (m *branchManagerImpl) Find(branchId int) (*Branch, error) {
	var branch Branch

	result := m.db.Where("id = ?", branchId).Find(&branch)

	if result.Error != nil {
		log.Printf("Fail to retrieve branch: %v", result.Error)
		return nil, result.Error
	}

	return &branch, nil
}

func (m *branchManagerImpl) FindByParentId(parentId *int) (Branch, error) {
	var parent Branch

	result := m.db.Preload("Requirements").Preload("Restrictions").First(&parent, parentId)
	if result.Error != nil {
		log.Printf("Fail to retrieve branches: %v", result.Error)
		return parent, result.Error
	}
	return parent, nil
}

func (m *branchManagerImpl) Save(branch *Branch) (*Branch, error) {
	result := m.db.Create(branch)
	if result.Error != nil {
		log.Printf("Fail to save branch: %v", result.Error)
		return nil, result.Error
	}
	return branch, nil
}
