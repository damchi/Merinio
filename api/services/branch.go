package services

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"merinio/api/models"
	"strconv"
)

type BranchService struct {
	BranchManager      models.BranchManager
	RequirementManager models.RequirementManager
	RestrictionManager models.RestrictionManager
}

func GetBranchService(db *gorm.DB) BranchService {
	branchManager := models.GetBranchManager(db)
	requirementManager := models.GetRequirementManager(db)
	restrictionManager := models.GetRestrictionManager(db)
	return BranchService{
		branchManager,
		requirementManager,
		restrictionManager,
	}
}

func (bs BranchService) GetListBranch() ([]models.Branch, error) {
	var branches []models.Branch

	branches, err := bs.BranchManager.FindAll()
	if err != nil {
		return branches, err
	}

	for i := range branches {
		if err := bs.LoadParentData(&branches[i]); err != nil {
			return nil, err
		}
	}
	return branches, nil
}

func (bs BranchService) GetBranch(branchId string) (*models.Branch, error) {
	branchIdInt, err := strconv.Atoi(branchId)
	if err != nil {
		log.Printf("Convert branch id to int : %v", err)
		return nil, err
	}

	branches, err := bs.BranchManager.Find(branchIdInt)
	if err != nil {
		return branches, err
	}

	if err := bs.LoadParentData(branches); err != nil {
		return nil, err
	}

	return branches, nil
}

func (bs BranchService) SaveBranch(branchRequest *models.BranchRequest) (*models.Branch, error) {

	branch := &models.Branch{
		Name:     branchRequest.Name,
		ParentID: branchRequest.ParentID,
		IsRoot:   branchRequest.IsRoot,
	}
	if branch.ParentID != nil {
		_, err := bs.BranchManager.FindByParentId(branch.ParentID)
		if err != nil {
			return nil, fmt.Errorf("parent branch with ID %d does not exist", *branch.ParentID)
		}
	}

	if len(branchRequest.Requirement) > 0 {
		var requirements []models.Requirement
		requirements, err := bs.RequirementManager.FindAllWhere(branchRequest.Requirement)
		if err != nil {
			return nil, err
		}
		if len(requirements) > 0 {
			branch.Requirements = requirements
		} else {
			branch.Requirements = nil
		}
	}

	if len(branchRequest.Restriction) > 0 {
		var restrictions []models.Restriction
		restrictions, err := bs.RestrictionManager.FindAllWhere(branchRequest.Restriction)
		if err != nil {
			return nil, err
		}
		if len(restrictions) > 0 {
			branch.Restrictions = restrictions
		} else {
			branch.Restrictions = nil
		}
	}

	result, err := bs.BranchManager.Save(branch)
	if err != nil {
		return branch, err
	}

	return result, nil
}

func (bs BranchService) LoadParentData(branch *models.Branch) error {
	if branch.ParentID == nil {
		return nil
	}

	parent, err := bs.BranchManager.FindByParentId(branch.ParentID)
	if err != nil {
		return err
	}

	// Merge data from parent
	branch.Requirements = bs.RequirementManager.MergeRequirements(branch.Requirements, parent.Requirements)
	branch.Restrictions = bs.RestrictionManager.MergeRestrictions(branch.Restrictions, parent.Restrictions)

	if err := bs.LoadParentData(&parent); err != nil {
		return err
	}

	// Merge again after loading parent's parent
	branch.Requirements = bs.RequirementManager.MergeRequirements(branch.Requirements, parent.Requirements)
	branch.Restrictions = bs.RestrictionManager.MergeRestrictions(branch.Restrictions, parent.Restrictions)

	return nil
}
