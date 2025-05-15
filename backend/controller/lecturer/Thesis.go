package lecturer

import (
	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
	"log"
	"net/http"
)

func CreateThesis(c *gin.Context) {
	var thesis entity.Thesis

	if err := c.ShouldBindJSON(&thesis); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "ข้อมูลไม่ถูกต้อง",
		})
		return
	}

	db := config.DB()

	var existingThesis entity.Thesis
	if err := db.Where("Title = ?", thesis.Title).First(&existingThesis).Error; err == nil {
		log.Printf("สถานการศึกษานี้มีอยู่ในระบบแล้ว: %+v", existingThesis)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"data":    existingThesis, 
		})
		return
	}

	if err := db.Create(&thesis).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  "ไม่สามารถเพิ่มข้อมูลได้",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    thesis, 
	})
}
