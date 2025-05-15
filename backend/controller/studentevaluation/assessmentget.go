package studentevaluation

import (
	"net/http"
	// "strconv"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
)

func CreateAssessmentGet(c *gin.Context) {
	var assessmentGet entity.AssessmentGet

	// อ่านข้อมูล JSON จาก Body
	if err := c.ShouldBindJSON(&assessmentGet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format", "status": 400})
		return
	}

	// ตรวจสอบ Score
	if assessmentGet.Score < 0 || assessmentGet.Score > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Score must be between 0 and 100", "status": 400})
		return
	}

	// ตรวจสอบ StudentEvaluationID
	if assessmentGet.StudentEvaluationID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "StudentEvaluationID is required", "status": 400})
		return
	}

	// ตรวจสอบ AssessmentTypeID
	if assessmentGet.AssessmentTypeID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "AssessmentTypeID is required", "status": 400})
		return
	}

	// เชื่อมต่อกับฐานข้อมูล
	db := config.DB()

	// ตรวจสอบว่า StudentEvaluation มีอยู่ในฐานข้อมูลหรือไม่
	var studentEvaluation entity.StudentEvaluation
	if err := db.First(&studentEvaluation, assessmentGet.StudentEvaluationID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "StudentEvaluation not found", "status": 404})
		return
	}

	// ตรวจสอบว่า AssessmentType มีอยู่ในฐานข้อมูลหรือไม่
	var assessmentType entity.AssessmentType
	if err := db.First(&assessmentType, assessmentGet.AssessmentTypeID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "AssessmentType not found", "status": 404})
		return
	}

	// สร้าง AssessmentGet
	if err := db.Create(&assessmentGet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create assessment get", "status": 500})
		return
	}
	
	// ส่งข้อมูลกลับไปที่ client
	c.JSON(http.StatusCreated, gin.H{
		"message": "Assessment get created successfully",
		"data":    assessmentGet,
		"status":  201,
	})
}
func DeleteAssessmentGet(c *gin.Context) {
	id := c.Param("id") // รับ ID จากพารามิเตอร์
	db := config.DB()

	// ตรวจสอบว่าข้อมูล AssessmentGet มีอยู่ในฐานข้อมูลหรือไม่
	var assessmentGet entity.AssessmentGet
	if err := db.First(&assessmentGet, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  "ไม่พบข้อมูล AssessmentGet",
		})
		return
	}

	// ลบข้อมูล AssessmentGet
	if err := db.Delete(&assessmentGet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  "ไม่สามารถลบข้อมูลได้",
		})
		return
	}

	// ส่งข้อความตอบกลับเมื่อการลบสำเร็จ
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "ลบข้อมูล AssessmentGet สำเร็จ",
	})
}
func UpdateAssessmentGet(c *gin.Context) {
	// รับ ID ของ AssessmentGet จาก URL parameter
	id := c.Param("id")
	fmt.Printf("Debug: Received ID = %s\n", id)

	var assessmentGet entity.AssessmentGet

	// ตรวจสอบว่า ID ที่ส่งมามี AssessmentGet อยู่ในฐานข้อมูลหรือไม่
	db := config.DB()
	if err := db.Where("id = ?", id).First(&assessmentGet).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "AssessmentGet not found", "status": 404})
		return
	}

	// Debug: พิมพ์ข้อมูล AssessmentGet ก่อนการอัปเดต
	fmt.Printf("Debug: Existing AssessmentGet = %+v\n", assessmentGet)

	// อ่านข้อมูล JSON จาก Body
	var input struct {
		Score              *int `json:"Score"`
		StudentEvaluationID *uint `json:"StudentEvaluationID"`
		AssessmentTypeID    *uint `json:"AssessmentTypeID"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format", "status": 400})
		return
	}

	// Debug: พิมพ์ข้อมูลที่รับมาจาก JSON
	fmt.Printf("Debug: Received Input = %+v\n", input)

	// ตรวจสอบและอัปเดต Score
	if input.Score != nil {
		if *input.Score < 0 || *input.Score > 100 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Score must be between 0 and 100", "status": 400})
			return
		}
		assessmentGet.Score = *input.Score
	}

	// ตรวจสอบและอัปเดต StudentEvaluationID
	if input.StudentEvaluationID != nil {
		var studentEvaluation entity.StudentEvaluation
		if err := db.First(&studentEvaluation, *input.StudentEvaluationID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "StudentEvaluation not found", "status": 404})
			return
		}
		assessmentGet.StudentEvaluationID = *input.StudentEvaluationID
	}

	// ตรวจสอบและอัปเดต AssessmentTypeID
	if input.AssessmentTypeID != nil {
		var assessmentType entity.AssessmentType
		if err := db.First(&assessmentType, *input.AssessmentTypeID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "AssessmentType not found", "status": 404})
			return
		}
		assessmentGet.AssessmentTypeID = *input.AssessmentTypeID
	}

	// Debug: พิมพ์ข้อมูล AssessmentGet หลังการอัปเดต
	fmt.Printf("Debug: Updated AssessmentGet = %+v\n", assessmentGet)

	// บันทึกการเปลี่ยนแปลงลงฐานข้อมูล
	if err := db.Save(&assessmentGet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update AssessmentGet", "status": 500})
		return
	}

	// ส่งข้อมูลกลับไปยัง client
	c.JSON(http.StatusOK, gin.H{
		"message": "AssessmentGet updated successfully",
		"data":    assessmentGet,
		"status":  200,
	})
}
func GetAssessmentGetByID(c *gin.Context) {
	id := c.Param("id")
	db := config.DB()
	var assessmentGet entity.AssessmentGet

	// ค้นหา AssessmentGet พร้อมข้อมูลที่เกี่ยวข้อง (Preload)
	if err := db.Preload("StudentEvaluation").Preload("AssessmentType").First(&assessmentGet, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  "ไม่พบข้อมูล AssessmentGet",
		})
		return
	}

	// ส่งข้อมูลกลับในรูปแบบ JSON
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   assessmentGet,
	})
}

// ListAssessmentGet - ดึงข้อมูล AssessmentGet ทั้งหมด
func ListAssessmentGet(c *gin.Context) {
	db := config.DB()
	var assessmentGets []entity.AssessmentGet

	// ดึงข้อมูล AssessmentGet ทั้งหมด พร้อมข้อมูลที่เกี่ยวข้อง (Preload)
	if err := db.Preload("StudentEvaluation").Preload("AssessmentType").Find(&assessmentGets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  "ไม่สามารถดึงข้อมูล AssessmentGet ได้",
		})
		return
	}

	// ส่งข้อมูลกลับในรูปแบบ JSON
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   assessmentGets,
	})
}