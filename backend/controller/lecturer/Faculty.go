package lecturer

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
)

func GetFaculty(c *gin.Context) {
	var faculty []entity.Faculty

	db := config.DB()
	results := db.Model(&entity.Faculty{}).Select("id, faculty_name").Find(&faculty)

	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}

	var count int64
	countResults := db.Model(&entity.Faculty{}).Count(&count)

	// ตรวจสอบผลลัพธ์การนับ
	if countResults.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถนับจำนวนคณะได้"})
		return
	}

	// ส่งข้อมูล faculty และจำนวนคณะกลับไปที่ client
	c.JSON(http.StatusOK, gin.H{
		"data":       faculty,
		"totalCount": count,
	})
}