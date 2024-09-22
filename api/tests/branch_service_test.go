package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"merinio/api/models"
	"merinio/api/services"
	"testing"
)

// Mocking the BranchManager
type MockBranchManager struct {
	mock.Mock
}

func (m *MockBranchManager) FindAll() ([]models.Branch, error) {
	args := m.Called()
	return args.Get(0).([]models.Branch), args.Error(1)
}

func (m *MockBranchManager) Save(branch *models.Branch) (*models.Branch, error) {
	args := m.Called(branch)
	return args.Get(0).(*models.Branch), args.Error(1)
}

func (m *MockBranchManager) FindByParentId(parentID *int) (models.Branch, error) {
	args := m.Called(parentID)
	return args.Get(0).(models.Branch), args.Error(1)
}

func TestGetListBranch(t *testing.T) {
	mockBranchManager := new(MockBranchManager)
	branchService := services.BranchService{
		BranchManager: mockBranchManager,
	}

	branches := []models.Branch{
		{ID: 1, Name: "Branch 1"},
		{ID: 2, Name: "Branch 2"},
	}

	mockBranchManager.On("FindAll").Return(branches, nil)

	result, err := branchService.GetListBranch()

	assert.NoError(t, err)
	assert.Equal(t, branches, result)
	mockBranchManager.AssertExpectations(t)
}

func TestSaveBranch(t *testing.T) {
	mockBranchManager := new(MockBranchManager)
	branchService := services.BranchService{
		BranchManager: mockBranchManager,
	}

	branchRequest := &models.BranchRequest{
		Name:     "New Branch",
		ParentID: nil,
		IsRoot:   true,
	}

	mockBranchManager.On("Save", mock.Anything).Return(&models.Branch{ID: 1, Name: "New Branch"}, nil)

	result, err := branchService.SaveBranch(branchRequest)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "New Branch", result.Name)
	mockBranchManager.AssertExpectations(t)
}

func TestLoadParentData(t *testing.T) {
	mockBranchManager := new(MockBranchManager)
	branchService := services.BranchService{
		BranchManager: mockBranchManager,
	}

	branch := &models.Branch{ID: 2, ParentID: new(int)}
	*branch.ParentID = 1

	parentBranch := models.Branch{ID: 1, Requirements: []models.Requirement{}, Restrictions: []models.Restriction{}}

	mockBranchManager.On("FindByParentId", branch.ParentID).Return(parentBranch, nil)

	err := branchService.LoadParentData(branch)

	assert.NoError(t, err)
	mockBranchManager.AssertExpectations(t)
}
