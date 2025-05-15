package studentevaluation

import (
	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
	// "strconv"
	"net/http"
)

func CreateStudentEvaluation(c *gin.Context) {
	var studentEvaluation entity.StudentEvaluation

	// อ่านข้อมูล JSON จาก Body
	if err := c.ShouldBindJSON(&studentEvaluation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format", "status": 400})
		return
	}

	// ตรวจสอบ TotalScore
	if studentEvaluation.TotalScore < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "TotalScore must be greater than or equal to 0", "status": 400})
		return
	}

	// ตรวจสอบ Grad
	if studentEvaluation.Grad == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Grad is required", "status": 400})
		return
	}

	// ตรวจสอบ StudentID และ CourseID
	if studentEvaluation.StudentID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "StudentID is required", "status": 400})
		return
	}
	if studentEvaluation.CourseID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CourseID is required", "status": 400})
		return
	}

	// เชื่อมต่อกับฐานข้อมูล
	db := config.DB()
	if err := db.Create(&studentEvaluation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create student evaluation", "status": 500})
		return
	}

	// ส่งข้อมูลกลับไปที่ client
	c.JSON(http.StatusCreated, gin.H{
		"message": "Student evaluation created successfully",
		"data":    studentEvaluation,
		"status":  201,
	})
}

// GetStudentEvaluationByID - ดึงข้อมูล StudentEvaluation ตาม ID
func GetStudentEvaluationByID(c *gin.Context) {
	id := c.Param("id")
	db := config.DB()
	var evaluation entity.StudentEvaluation

	if err := db.Preload("Student").Preload("Course").First(&evaluation, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  "ไม่พบข้อมูล",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   evaluation,
	})
}

func ListStudentEvaluation(c *gin.Context) {
	var studentEvaluations []entity.StudentEvaluation

	// เชื่อมต่อกับฐานข้อมูล
	db := config.DB()

	// ดึงข้อมูล StudentEvaluation ทั้งหมด พร้อมกับ Preload ข้อมูลที่เกี่ยวข้อง
	db.Preload("Student").Preload("Course").Find(&studentEvaluations)

	// ส่งข้อมูลกลับในรูปแบบ JSON
	c.JSON(http.StatusOK, &studentEvaluations)
}

// DeleteStudentEvaluation - ลบข้อมูล StudentEvaluation
func DeleteStudentEvaluation(c *gin.Context) {
	id := c.Param("id")
	db := config.DB()

	if err := db.Delete(&entity.StudentEvaluation{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  "ไม่สามารถลบข้อมูลได้",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "ลบข้อมูลสำเร็จ",
	})
}
func AutomaticUpdateGrading(c *gin.Context) {
	// รับค่า StudentEvaluationID และ StudentID จาก JSON
	var input struct {
		StudentEvaluationID uint `json:"StudentEvaluationID" binding:"required"`
		StudentID           uint `json:"StudentID" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format", "status": 400})
		return
	}

	// เชื่อมต่อกับฐานข้อมูล
	db := config.DB()

	// ตรวจสอบว่า StudentEvaluation มีอยู่ในฐานข้อมูลหรือไม่
	var studentEvaluation entity.StudentEvaluation
	if err := db.First(&studentEvaluation, input.StudentEvaluationID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "StudentEvaluation not found", "status": 404})
		return
	}

	// ดึงข้อมูล AssessmentGet ที่มี StudentEvaluationID ตรงกัน
	var assessmentGets []entity.AssessmentGet
	if err := db.Where("student_evaluation_id = ?", input.StudentEvaluationID).Find(&assessmentGets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch AssessmentGet data", "status": 500})
		return
	}

	// คำนวณคะแนนรวม
	totalScore := 0
	for _, assessment := range assessmentGets {
		totalScore += assessment.Score
	}

	// คำนวณเกรดจากคะแนนรวม
	var grade string
	switch {
	case totalScore >= 80:
		grade = "A"
	case totalScore >= 75:
		grade = "B+"
	case totalScore >= 70:
		grade = "B"
	case totalScore >= 65:
		grade = "C+"
	case totalScore >= 60:
		grade = "C"
	case totalScore >= 55:
		grade = "D+"
	case totalScore >= 50:
		grade = "D"
	default:
		grade = "F"
	}

	// อัปเดตคะแนนรวมและเกรดใน StudentEvaluation
	studentEvaluation.TotalScore = totalScore
	studentEvaluation.Grad = grade

	if err := db.Save(&studentEvaluation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update StudentEvaluation", "status": 500})
		return
	}

	// ส่งข้อมูลกลับไปยัง client
	c.JSON(http.StatusOK, gin.H{
		"message": "StudentEvaluation updated successfully",
		"data":    studentEvaluation,
		"status":  200,
	})
}