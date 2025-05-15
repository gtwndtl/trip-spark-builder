package student

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
)

type StudentEducation struct {
	ID          uint   `json:"ID"`
	StudentID  uint   `json:"StudentID"`
	EducationID uint   `json:"EducationID"`
	Instution   string `json:"Instution"`
	DegreeID    uint   `json:"DegreeID"`
	Degree      string `json:"Degree"`
	Certificate string `json:"Certificate"`
}

func CreateStudentEducation(c *gin.Context) {
	var studenteducation entity.StudentEducation

	if err := c.ShouldBindJSON(&studenteducation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "รูปแบบข้อมูลไม่ถูกต้อง",
		})
		return
	}

	db := config.DB()

	var studentID, educationID, degreeID uint

	var student entity.Student
	if studenteducation.StudentID != nil {
		studentID = uint(*studenteducation.StudentID)
		if err := db.First(&student, studentID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "ไม่พบข้อมูลอาจารย์ในระบบ",
			})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ไม่พบข้อมูลอาจารย์",
		})
		return
	}

	var education entity.Education
	if studenteducation.EducationID != nil {
		educationID = uint(*studenteducation.EducationID)
		if err := db.First(&education, educationID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "ไม่พบข้อมูลสถานการศึกษาในระบบ",
			})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ไม่พบข้อมูลสถานการศึกษา",
		})
		return
	}

	var degree entity.Degree
	if studenteducation.DegreeID != nil {
		degreeID = uint(*studenteducation.DegreeID)
		if err := db.First(&degree, degreeID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "ไม่พบข้อมูลระดับการศึกษาในระบบ",
			})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ไม่พบข้อมูลระดับการศึกษา",
		})
		return
	}

	studenteducate := entity.StudentEducation{
		StudentID:  &studentID,
		EducationID: &educationID,
		DegreeID:    &degreeID,
		Certificate: studenteducation.Certificate,
	}

	if err := db.Create(&studenteducate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "ไม่สามารถบันทึกข้อมูลได้",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "ข้อมูลประวัติการศึกษาถูกเพิ่มเรียบร้อยแล้ว",
	})
}

func UpdateStudentEducation(c *gin.Context) {
	var studentEducation entity.StudentEducation
	studentEducationID := c.Param("id")

	db := config.DB()
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถเชื่อมต่อฐานข้อมูลได้"})
		return
	}

	if err := db.First(&studentEducation, studentEducationID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบข้อมูลการศึกษาของนักศึกษาที่ต้องการอัปเดต"})
		return
	}

	var input StudentEducation
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "รูปแบบข้อมูลไม่ถูกต้อง"})
		return
	}

	var existingEducation entity.Education
	if err := db.Where("instution = ?", input.Instution).First(&existingEducation).Error; err == nil {
		studentEducation.EducationID = &existingEducation.ID
	} else {
		newEducation := entity.Education{
			Instution: input.Instution,
		}
		if err := db.Create(&newEducation).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถสร้างข้อมูลสถานศึกษาใหม่ได้"})
			return
		}
		studentEducation.EducationID = &newEducation.ID
	}

	var degree entity.Degree
	if err := db.First(&degree, input.DegreeID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบข้อมูลระดับการศึกษาที่ระบุ"})
		return
	}

	studentEducation.DegreeID = &input.DegreeID
	studentEducation.Certificate = input.Certificate

	if err := db.Save(&studentEducation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถอัปเดตข้อมูลการศึกษาของนักศึกษาได้"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "อัปเดตข้อมูลการศึกษาของนักศึกษาเรียบร้อยแล้ว"})
}

func GetEducationByStudentID(c *gin.Context) {
	ID := c.Param("id")

	db := config.DB()
	query := `
        SELECT 
			se.id AS id,
    		s.id AS student_id,
			e.id AS education_id,
    		e.instution AS instution,
			d.id AS degree_id,
    		d.degree AS degree,
    		se.certificate
		FROM student_educations se
		LEFT JOIN students s ON se.student_id = s.id
		LEFT JOIN educations e ON se.education_id = e.id
		LEFT JOIN degrees d ON se.degree_id = d.id
		WHERE se.student_id = ?`
	var educationRes []StudentEducation

	// Query database
	if err := db.Raw(query, ID).Scan(&educationRes).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	count := len(educationRes)

	c.JSON(http.StatusOK, gin.H{
		"educations": educationRes,
		"count":      count,
	})
}
