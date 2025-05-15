package dormitory

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
)

func CreateDorm(c *gin.Context) {
	var Dormitory entity.Dormitory

	// รับไฟล์จาก request
	file, err := c.FormFile("DormPic")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "DormPic is required","status": 400 })
		return
	}

	// สร้างชื่อไฟล์ใหม่ให้ไม่ซ้ำกัน
	timestamp := time.Now().Unix()
	extension := filepath.Ext(file.Filename)
	newFileName := fmt.Sprintf("dorm_%d%s", timestamp, extension)

	// กำหนด path สำหรับบันทึกไฟล์
	DormPic := fmt.Sprintf("./uploads/%s", newFileName) // สร้างไดเรกทอรี uploads ในโปรเจกต์
	if err := c.SaveUploadedFile(file, DormPic); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file","status": 500 })
		return
	}

	// อ่านข้อมูลอื่น ๆ ของฟอร์ม
	DormName := c.PostForm("DormName")
	if DormName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "DormName is required","status": 400 })
		return
	}

	DormDescription := c.PostForm("DormDescription")
	DormType := c.PostForm("DormType")
	Price := c.PostForm("Price")

	// แปลงราคาจาก string เป็น int
	PriceInt, err := strconv.Atoi(Price)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price format","status": 400 })
		return
	}

	// สร้าง Dormitory entity และบันทึกลงฐานข้อมูล
	Dormitory = entity.Dormitory{
		DormName:        DormName,
		DormDescription: DormDescription,
		DormType:        DormType,
		DormPic:         DormPic, // บันทึก path ของรูปภาพ
		Price:           PriceInt,
	}

	// เชื่อมต่อกับฐานข้อมูล
	db := config.DB()
	if err := db.Create(&Dormitory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create dormitory","status": 400 })
		return
	}

	// ส่งข้อมูลกลับไปที่ client
	c.JSON(http.StatusCreated, gin.H{"message": "Dormitory created successfully", "data": Dormitory,"status": 201 })
}

func DeleteDorm(c *gin.Context) {
	id := c.Param("id") // รับค่า id จาก URL params
	db := config.DB()

	// ตรวจสอบว่า id มีค่าเป็นตัวเลขหรือไม่
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	// เริ่ม transaction เพื่อให้การลบเป็น atomic
	tx := db.Begin()

	// Soft Delete ห้องพักที่สัมพันธ์กับหอพัก (ตั้งค่า deleted_at เป็นเวลาปัจจุบัน)
	if err := tx.Exec("UPDATE Rooms SET deleted_at = ? WHERE dormitory_id = ?", time.Now(), id).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to soft delete rooms", "details": err.Error()})
		return
	}

	// Soft Delete ข้อมูลหอพัก (ตั้งค่า deleted_at เป็นเวลาปัจจุบัน)
	if tx := tx.Exec("UPDATE Dormitories SET deleted_at = ? WHERE id = ?", time.Now(), id); tx.RowsAffected == 0 {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}

	// Commit transaction เมื่อสำเร็จ
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "transaction failed", "details": err.Error()})
		return
	}

	// ส่งผลลัพธ์กลับไปยัง frontend
	c.JSON(http.StatusOK, gin.H{"message": "Soft deleted successfully", "status": 200})
}




  func UpdateDorm(c *gin.Context) {
	var Dormitory entity.Dormitory
	id := c.Param("id") // รับ id จากพารามิเตอร์
	db := config.DB()

	// ค้นหา Dormitory ในฐานข้อมูล
	if err := db.First(&Dormitory, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Dorm not found","status": 404 })
		return
	}

	// ตรวจสอบว่ามีการอัปโหลดรูปภาพใหม่หรือไม่
	file, err := c.FormFile("DormPic")
	if err == nil { // มีการอัปโหลดไฟล์ใหม่
		// สร้างชื่อไฟล์ใหม่ไม่ให้ซ้ำกัน
		timestamp := time.Now().Unix()
		extension := filepath.Ext(file.Filename)
		newFileName := fmt.Sprintf("dorm_%d%s", timestamp, extension)

		// กำหนด path สำหรับบันทึกไฟล์
		DormPic := fmt.Sprintf("./uploads/%s", newFileName)
		if err := c.SaveUploadedFile(file, DormPic); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save dormpic","status": 500 })
			return
		}

		// อัปเดต path ของรูปภาพในฐานข้อมูล
		Dormitory.DormPic = DormPic
	}

	// อ่านข้อมูลอื่น ๆ จาก form-data
	DormName := c.PostForm("DormName")
	DormDescription := c.PostForm("DormDescription")
	DormType := c.PostForm("DormType")
	Price := c.PostForm("Price")

	if DormName != "" {
		Dormitory.DormName = DormName
	}
	if DormDescription != "" {
		Dormitory.DormDescription = DormDescription
	}
	if DormType != "" {
		Dormitory.DormType = DormType
	}
	if Price != "" {
		PriceInt, err := strconv.Atoi(Price)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price format","status": 400 })
			return
		}
		Dormitory.Price = PriceInt
	}

	// บันทึกการอัปเดตในฐานข้อมูล
	if err := db.Save(&Dormitory).Error; err != nil {
		fmt.Println("Failed to update database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update dormitory","status": 500 })
		return
	}
	fmt.Println("Updated Dormitory:", Dormitory)

	c.JSON(http.StatusOK, gin.H{"message": "Dormitory updated successfully", "data": Dormitory,"status": 200 })
}

func ListDorm(c *gin.Context) {
	var Dormitory []entity.Dormitory
  
	// เชื่อมต่อกับฐานข้อมูล
	db := config.DB()
  
	// ดึงข้อมูลจากฐานข้อมูล
	results := db.Find(&Dormitory)
	if results.Error != nil {
	  // ถ้ามีข้อผิดพลาดในการดึงข้อมูล
	  c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error(),"status": 404 })
	  return
	}

	baseURL := "http://localhost:8000" // URL ของ Backend
	for i := range Dormitory {
		if Dormitory[i].DormPic != "" {
			Dormitory[i].DormPic = fmt.Sprintf("%s/%s", baseURL, Dormitory[i].DormPic[2:]) 
		}
	}
  
	// ส่งข้อมูลที่ดึงมาเป็น JSON
	c.JSON(http.StatusOK, Dormitory)
  }

  func GetAll(c *gin.Context) {


	db := config.DB()
 
 
	var dormitory []entity.Dormitory
	db.Find(&dormitory)
 
 
	c.JSON(http.StatusOK, &dormitory)
 
 
 }