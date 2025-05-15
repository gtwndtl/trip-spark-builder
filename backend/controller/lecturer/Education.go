package lecturer

import (
	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
	"log"
	"net/http"
)

func CreateEducation(c *gin.Context) {
	var education entity.Education

	if err := c.ShouldBindJSON(&education); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "ข้อมูลไม่ถูกต้อง",
		})
		return
	}

	db := config.DB()

	var existingEducation entity.Education
	if err := db.Where("instution = ?", education.Instution).First(&existingEducation).Error; err == nil {
		log.Printf("สถานการศึกษานี้มีอยู่ในระบบแล้ว: %+v", existingEducation)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"data":    existingEducation, 
		})
		return
	}

	if err := db.Create(&education).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  "ไม่สามารถเพิ่มข้อมูลได้",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    education,
	})
}
