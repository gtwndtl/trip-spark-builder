package paymentcourse

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
)

func GetAll(c *gin.Context) {


	db := config.DB()
 
 
	var paycourse []entity.PaymentCourse
	db.Preload("Enrollment").Preload("Course").Preload("Payment").Preload("Users").Find(&paycourse)
 
 
	c.JSON(http.StatusOK, &paycourse)
 
 
 }

func Create(c *gin.Context) {
	var newPaymentCourse entity.PaymentCourse

	// Bind the incoming JSON payload to the newPaymentCourse struct
	if err := c.ShouldBindJSON(&newPaymentCourse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	// Validate that required fields are present
	if newPaymentCourse.PaymentID == 0 || newPaymentCourse.CourseID == 0 || newPaymentCourse.EnrollmentID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields (PaymentID, CourseID, EnrollmentID)"})
		return
	}

	// Optionally, check if PaymentID, CourseID, and EnrollmentID exist in the database
	db := config.DB()

	var payment entity.Payment
	if err := db.First(&payment, newPaymentCourse.PaymentID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid PaymentID"})
		return
	}

	var course entity.Course
	if err := db.First(&course, newPaymentCourse.CourseID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid CourseID"})
		return
	}

	var enrollment entity.Enrollment
	if err := db.First(&enrollment, newPaymentCourse.EnrollmentID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid EnrollmentID"})
		return
	}

	// Save the new PaymentCourse to the database
	if err := db.Create(&newPaymentCourse).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create PaymentCourse"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "PaymentCourse created successfully", "PaymentCourse": newPaymentCourse})
}
