package lecturer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
)

type LecturerEducation struct {
	ID          uint   `json:"ID"`
	LecturerID  uint   `json:"LecturerID"`
	EducationID uint   `json:"EducationID"`
	Instution   string `json:"Instution"`
	DegreeID    uint   `json:"DegreeID"`
	Degree      string `json:"Degree"`
	Certificate string `json:"Certificate"`
}

func CreateLecturerEducation(c *gin.Context) {
	var lecturereducation entity.LecturerEducation

	if err := c.ShouldBindJSON(&lecturereducation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "รูปแบบข้อมูลไม่ถูกต้อง",
		})
		return
	}

	db := config.DB()

	var lecturerID, educationID, degreeID uint

	var lecturer entity.Lecturer
	if lecturereducation.LecturerID != nil {
		lecturerID = uint(*lecturereducation.LecturerID)
		if err := db.First(&lecturer, lecturerID).Error; err != nil {
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
	if lecturereducation.EducationID != nil {
		educationID = uint(*lecturereducation.EducationID)
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
	if lecturereducation.DegreeID != nil {
		degreeID = uint(*lecturereducation.DegreeID)
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

	lecturereducate := entity.LecturerEducation{
		LecturerID:  &lecturerID,
		EducationID: &educationID,
		DegreeID:    &degreeID,
		Certificate: lecturereducation.Certificate,
	}

	if err := db.Create(&lecturereducate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "ไม่สามารถบันทึกข้อมูลได้",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "ข้อมูลประวัติการศึกษาถูกเพิ่มเรียบร้อยแล้ว",
	})
}

func UpdateLecturerEducation(c *gin.Context) {
	var lecturerEducation entity.LecturerEducation
	lecturerEducationID := c.Param("id")

	db := config.DB()
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถเชื่อมต่อฐานข้อมูลได้"})
		return
	}

	if err := db.First(&lecturerEducation, lecturerEducationID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบข้อมูลการศึกษาของอาจารย์ที่ต้องการอัปเดต"})
		return
	}

	var input LecturerEducation
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "รูปแบบข้อมูลไม่ถูกต้อง"})
		return
	}

	var existingEducation entity.Education
	if err := db.Where("instution = ?", input.Instution).First(&existingEducation).Error; err == nil {
		lecturerEducation.EducationID = &existingEducation.ID
	} else {
		newEducation := entity.Education{
			Instution: input.Instution,
		}
		if err := db.Create(&newEducation).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถสร้างข้อมูลสถานศึกษาใหม่ได้"})
			return
		}
		lecturerEducation.EducationID = &newEducation.ID
	}

	var degree entity.Degree
	if err := db.First(&degree, input.DegreeID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบข้อมูลระดับการศึกษาที่ระบุ"})
		return
	}

	lecturerEducation.DegreeID = &input.DegreeID
	lecturerEducation.Certificate = input.Certificate

	if err := db.Save(&lecturerEducation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถอัปเดตข้อมูลการศึกษาของอาจารย์ได้"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "อัปเดตข้อมูลการศึกษาของอาจารย์เรียบร้อยแล้ว"})
}

func GetEducationByLecturerID(c *gin.Context) {
	ID := c.Param("id")

	db := config.DB()
	query := `
        SELECT 
			le.id AS id,
    		l.id AS lecturer_id,
			e.id AS education_id,
    		e.instution AS instution,
			d.id AS degree_id,
    		d.degree AS degree,
    		le.certificate
		FROM lecturer_educations le
		LEFT JOIN lecturers l ON le.lecturer_id = l.id
		LEFT JOIN educations e ON le.education_id = e.id
		LEFT JOIN degrees d ON le.degree_id = d.id
		WHERE le.lecturer_id = ?
		ORDER BY degree_id`
	var educationRes []LecturerEducation

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
