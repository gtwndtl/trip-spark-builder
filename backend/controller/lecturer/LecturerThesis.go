package lecturer

import (
	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
	"net/http"
	"time"
)

type LecturerThesisRes struct {
	ID               uint      `json:"ID"`
	LecturerID       uint      `json:"LecturerID"`
	ThesisID         uint      `json:"ThesisID"`
	Title            string    `json:"Title"`
	URL              string    `json:"URL"`
	RoleOfThesisID   uint      `json:"RoleOfThesisID"`
	RoleOfThesisName string    `json:"RoleOfThesisName"`
	PublicationDate  time.Time `json:"PublicationDate"`
}

func CreateLecturerThesis(c *gin.Context) {
	var lecturerthesis entity.LecturerThesis

	if err := c.ShouldBindJSON(&lecturerthesis); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "รูปแบบข้อมูลไม่ถูกต้อง",
		})
		return
	}

	db := config.DB()

	var lecturerID, thesisID, rolethesisID uint

	var lecturer entity.Lecturer
	if lecturerthesis.LecturerID != nil {
		lecturerID = uint(*lecturerthesis.LecturerID)
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

	var thesis entity.Thesis
	if lecturerthesis.ThesisID != nil {
		thesisID = uint(*lecturerthesis.ThesisID)
		if err := db.First(&thesis, thesisID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "ไม่พบข้อมูลวิทยานิพนธ์ในระบบ",
			})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ไม่พบข้อมูลวิทยานิพนธ์",
		})
		return
	}

	var roleOfThesis entity.RoleOfThesis
	if lecturerthesis.RoleOfThesisID != nil {
		rolethesisID = uint(*lecturerthesis.RoleOfThesisID)
		if err := db.First(&roleOfThesis, rolethesisID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "ไม่พบข้อมูลตำแหน่งในวิทยานิพนธ์ในระบบ",
			})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ไม่พบข้อมูลตำแหน่งในวิทยานิพนธ์",
		})
		return
	}

	lecturerthesisfinish := entity.LecturerThesis{
		LecturerID:     &lecturerID,
		ThesisID:       &thesisID,
		RoleOfThesisID: &rolethesisID,
	}

	if err := db.Create(&lecturerthesisfinish).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "ไม่สามารถบันทึกข้อมูลได้",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "ข้อมูลวิทยานิพนธ์ถูกเพิ่มเรียบร้อยแล้ว",
	})
}

func UpdateLecturerThesis(c *gin.Context) {
	var input LecturerThesisRes
	lecturerThesisID := c.Param("id")

	db := config.DB()
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถเชื่อมต่อกับเซิร์ฟเวอร์"})
		return
	}

	var lecturerThesis entity.LecturerThesis
	if err := db.First(&lecturerThesis, lecturerThesisID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบข้อมูลวิทยานิพนธ์ของอาจารย์"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	if input.ThesisID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ข้อมูล ThesisID ไม่ถูกต้อง"})
		return
	}

	var thesis entity.Thesis
	if err := db.First(&thesis, input.ThesisID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบข้อมูลวิทยานิพนธ์"})
		return
	}

	thesis.Title = input.Title
	thesis.URL = input.URL
	thesis.PublicationDate = input.PublicationDate

	if err := db.Save(&thesis).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "เกิดข้อผิดพลาดในการแก้ไขวิทยานิพนธ์"})
		return
	}

	lecturerThesis.ThesisID = &input.ThesisID
	lecturerThesis.RoleOfThesisID = &input.RoleOfThesisID

	if err := db.Save(&lecturerThesis).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "เกิดข้อผิดพลาดในการแก้ไขวิทยานิพนธ์ของอาจารย์"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "แก้ไขข้อมูลสำเร็จ"})
}

func GetThesisByLecturerID(c *gin.Context) {
	ID := c.Param("id")

	db := config.DB()
	query := `
        SELECT 
			lt.id AS id,
    		l.id AS lecturer_id,
			t.id AS thesis_id,
			t.publication_date as publication_date,
			t.title as title,
			t.url as url,
			r.id as role_of_thesis_id,
			r.role_of_thesis_name

		FROM lecturer_theses lt
		LEFT JOIN lecturers l ON lt.lecturer_id = l.id
		LEFT JOIN theses t ON lt.thesis_id = t.id
		LEFT JOIN role_of_theses r ON lt.role_of_thesis_id = r.id
		WHERE lt.lecturer_id = ?`
	var thesisRes []LecturerThesisRes

	if err := db.Raw(query, ID).Scan(&thesisRes).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	count := len(thesisRes)

	c.JSON(http.StatusOK, gin.H{
		"thesis": thesisRes,
		"count":  count,
	})
}
