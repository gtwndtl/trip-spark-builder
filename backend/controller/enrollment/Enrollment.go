package enrollment

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
)

func ListEnrollments(c *gin.Context) {
	var enrollments []entity.Enrollment
	db := config.DB()

	if err := db.Find(&enrollments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch enrollments"})
		return
	}

	c.JSON(http.StatusOK, enrollments)
}

func GetEnrollmentByStudentID(c *gin.Context) {
	id := c.Param("id")
	var enrollments []entity.Enrollment
	db := config.DB()

	if err := db.Where("student_id = ?", id).Find(&enrollments).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "enrollments not found"})
		return
	}

	c.JSON(http.StatusOK, enrollments)
}

func GetEnrollmentByCourseID(c *gin.Context) {
	id := c.Param("id")
	var enrollments []entity.Enrollment
	db := config.DB()

	if err := db.Where("course_id = ?", id).Find(&enrollments).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "enrollments not found"})
		return
	}

	c.JSON(http.StatusOK, enrollments)
}

func CreateEnrollment(c *gin.Context) {
	var enrollment entity.Enrollment
	db := config.DB()

	if err := c.ShouldBindJSON(&enrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	if err := db.Create(&enrollment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create enrollment"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": enrollment})
}

func GetAll(c *gin.Context) {


	db := config.DB()
 
 
	var enrollment []entity.Enrollment
	db.Preload("Student").Preload("Course").Find(&enrollment)
 
 
	c.JSON(http.StatusOK, &enrollment)
 
 
 }