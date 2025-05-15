package report

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
)

func GetAll(c *gin.Context) {

	db := config.DB()

	var report []entity.Report
	db.Preload("Users").Preload("Admin").Preload("Dormitory").Preload("Books").Preload("Room").Find(&report)

	c.JSON(http.StatusOK, &report)//ส่งข้อมูลทั้งหมดในตาราง report กลับไป โดยใช้ status 200 JSON ในการตอบกลับ

}

func Delete(c *gin.Context) {

	id := c.Param("id")//เป็นการรับค่า id ที่เราส่งมา โดยใช้ c.Param("id") ซึ่ง id ต้องตรงกับที่เราตั้งไว้ใน path

	db := config.DB()

	if tx := db.Exec("DELETE FROM reports WHERE id = ?", id); tx.RowsAffected == 0 { // ลบข้อมูลในตาราง reports โดยเลือ id ที่ตรงกับที่ส่งมา ถ้าไม่เจอจะส่ง status 400 กลับไป
		//Exec คือการรันคำสั่ง SQL โดยไม่ต้องรับค่ามา และส่งค่ากลับไป
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})//ส่ง status 400 กลับไป JSON ที่บอกว่า id not found

		return

	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})//ส่ง status 200 กลับไป JSON ที่บอกว่า Deleted successful
	//JSON คือการส่งข้อมูลกลับไปในรูปแบบของ JSON ซึ่งจะส่งกลับไปในรูปแบบของ key และ value
}

func Get(c *gin.Context) {

	ID := c.Param("id")//เป็นการรับค่า id ที่เราส่งมา โดยใช้ c.Param("id") ซึ่ง id ต้องตรงกับที่เราตั้งไว้ใน path

	var report entity.Report

	db := config.DB()

	results := db.Preload("Users").Preload("Admin").Preload("Dormitory").Preload("Books").Preload("Room").First(&report, ID)

	if results.Error != nil {

		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})

		return

	}

	if report.ID == 0 {

		c.JSON(http.StatusNoContent, gin.H{})

		return

	}

	c.JSON(http.StatusOK, report)

}

func Update(c *gin.Context) {

	var report entity.Report

	ReportID := c.Param("id")

	db := config.DB()

	result := db.First(&report, ReportID)

	if result.Error != nil {

		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})

		return

	}

	if err := c.ShouldBindJSON(&report); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})

		return

	}

	result = db.Save(&report)

	if result.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})

}

func Create(c *gin.Context) {
	var newReport entity.Report

	// Bind the incoming JSON payload to the newReport struct
	if err := c.ShouldBindJSON(&newReport); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	db := config.DB()

	// Save the new report to the database
	if err := db.Create(&newReport).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create report"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Report created successfully", "report": newReport}) // Changed to http.StatusCreated
}
