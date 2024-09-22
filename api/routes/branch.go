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

func GetBranch(c *gin.Context) {
	strBranchId := c.Params.ByName("branch_id")

	gateway, err := services.GetBranchService(nil).GetBranch(strBranchId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gateway)

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
