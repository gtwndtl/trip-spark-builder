package lecturer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
)

func GetStatusStaffLecturer(c *gin.Context) {
	var statusstaff []entity.StatusStaff
	db := config.DB()

	if err := db.Where("id != ? AND id != ?", 2, 3).Find(&statusstaff).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "เกิดข้อผิดพลาดในการโหลดข้อมูล"})
		return
	}

	c.JSON(http.StatusOK, &statusstaff)
}

func GetStatusStaffStudent(c *gin.Context) {
	var statusstaff []entity.StatusStaff
	db := config.DB()

	if err := db.Where("id != ?", 1).Find(&statusstaff).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "เกิดข้อผิดพลาดในการโหลดข้อมูล"})
		return
	}

	c.JSON(http.StatusOK, &statusstaff)
}