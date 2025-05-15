package lecturer

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
)

func GetMajorByFacultyID(c *gin.Context) {
	facultyID := c.Param("id")
	var major []entity.Major

	db := config.DB()
	results := db.Model(&entity.Major{}).Select("id, major_name").Where("faculty_id = ?", facultyID).Find(&major)

	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}

	if len(major) == 0 { 
        c.JSON(http.StatusNoContent, gin.H{"message": "ไม่มีข้อมูล Major สำหรับ facultyID ที่ระบุ"})
        return
    }

	var count int64
	countResults := db.Model(&entity.Major{}).Where("faculty_id = ?", facultyID).Count(&count)

	// ตรวจสอบผลลัพธ์การนับ
	if countResults.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถนับจำนวนคณะได้"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       major,
		"totalCount": count,
	})
}