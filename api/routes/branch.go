package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"merinio/api/models"
	"merinio/api/services"
	"net/http"
)

func GetListBranches(c *gin.Context) {
	branchList, err := services.GetBranchService(nil).GetListBranch()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"branches": branchList})
}

func SaveBranch(c *gin.Context) {
	var param *models.BranchRequest

	err := c.BindJSON(&param)

	if err != nil {
		log.Printf("SaveBranch BinJson: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	branch, err := services.GetBranchService(nil).SaveBranch(param)
	if err != nil {
		log.Printf("SaveBranch: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, branch)
}
