package day

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
)

// ดึงรายการ Day ทั้งหมด
func ListDays(c *gin.Context) {
	var days []entity.Day
	db := config.DB()

	if err := db.Find(&days).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch days"})
		return
	}

	c.JSON(http.StatusOK, days)
}

// ดึง Day ตาม ID
func GetDayByID(c *gin.Context) {
	id := c.Param("id")
	var day entity.Day
	db := config.DB()

	if err := db.First(&day, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Day not found"})
		return
	}

	c.JSON(http.StatusOK, day)
}

