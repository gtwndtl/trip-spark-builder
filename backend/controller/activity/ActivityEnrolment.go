package activity

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
	"gorm.io/gorm"
)

type EnrollRequest struct {
	StudentID  uint `json:"studentId" binding:"required"`
	ActivityID uint `json:"activityId" binding:"required"`
}

func EnrollActivity(c *gin.Context) {
	var req EnrollRequest

	db := config.DB()
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if activity exists
	var activity entity.Activity
	if err := db.First(&activity, req.ActivityID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Check if student exists
	var student entity.Student
	if err := db.First(&student, req.StudentID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Enroll the student in the activity
	if err := db.Model(&activity).Association("Student").Append(&student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enroll"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Enrolled successfully"})
}






func UnenrollActivity(c *gin.Context) {
	var req EnrollRequest

	db := config.DB()
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบว่า Activity และ Student มีอยู่จริงหรือไม่
	var activity entity.Activity
	var student entity.Student
	if err := db.First(&activity, req.ActivityID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		return
	}
	if err := db.First(&student, req.StudentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	// ลบข้อมูลจากความสัมพันธ์ many-to-many
	if err := db.Model(&activity).Association("Student").Delete(&student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unenroll"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Unenrolled successfully"})
}
