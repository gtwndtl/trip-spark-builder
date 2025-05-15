package classroom

import (
	"net/http"
	"strconv"
	"gorm.io/gorm"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
)

func CreateClassroom(c *gin.Context) {
	var classroom entity.Classroom
	db := config.DB()

	// Bind JSON payload ไปที่ struct classroom
	if err := c.ShouldBindJSON(&classroom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบ BuildingID
	var building entity.Building
	if err := db.First(&building, classroom.BuildingID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Building not found"})
		return
	}

	// ตรวจสอบ ConditionID
	var condition entity.Condition
	if err := db.First(&condition, classroom.ConditionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Condition not found"})
		return
	}

	// แปลงค่า EndTimeClassroomString เป็น time.Time
	if classroom.EndTimeClassroomString != "" {
		parsedTime, err := time.Parse("15:04", classroom.EndTimeClassroomString)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time format. Use HH:mm"})
			return
		}
		classroom.EndTimeClassroom = parsedTime
	}

	// กรองค่าที่รับจาก frontend และสร้าง Classroom ใหม่
	u := entity.Classroom{
		RoomNumber:       classroom.RoomNumber,
		Capacity:         classroom.Capacity,
		Status:           classroom.Status,
		Floor:            classroom.Floor,
		BuildingID:       classroom.BuildingID,
		Building:         building,
		ConditionID:      classroom.ConditionID,
		Condition:        condition,
		EndTimeClassroom: classroom.EndTimeClassroom,
	}

	// บันทึกข้อมูล u ลงในฐานข้อมูล
	if err := db.Create(&u).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Classroom created successfully",
		"data":    u,
	})
}

func GetLisClassroom(c *gin.Context) {
    var classrooms []entity.Classroom

	db := config.DB()
	results := db.Preload("Condition").Preload("Building").Find(&classrooms)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, classrooms)
}

func UpdateClassroom(c *gin.Context) {
	var classroom entity.Classroom
	classroomID := c.Param("id")
	db := config.DB()

	// ดึงข้อมูล Classroom ที่จะอัปเดต
	result := db.First(&classroom, classroomID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}

	// Bind JSON payload ไปที่ struct classroom
	if err := c.ShouldBindJSON(&classroom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	// แปลงค่า EndTimeClassroomString เป็น time.Time
	if classroom.EndTimeClassroomString != "" {
		parsedTime, err := time.Parse("15:04", classroom.EndTimeClassroomString)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time format. Use HH:mm"})
			return
		}
		classroom.EndTimeClassroom = parsedTime
	}

	// บันทึกข้อมูล
	result = db.Save(&classroom)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated successfully"})
}

func GetClassroomByID(c *gin.Context) {
	// ดึง ClassroomID จาก URL params
	classroomID := c.Param("id")

	// ตรวจสอบว่า ClassroomID เป็นตัวเลขหรือไม่
	id, err := strconv.ParseUint(classroomID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Classroom ID"})
		return
	}

	// ดึงข้อมูล Classroom และข้อมูลที่เกี่ยวข้อง
	var classroom entity.Classroom
	db := config.DB()

	// ดึงข้อมูล Classroom พร้อมข้อมูลที่เกี่ยวข้อง (Building, Condition)
	if err := db.Preload("Building").Preload("Condition").First(&classroom, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Classroom not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ส่ง response กลับไปยัง frontend
	c.JSON(http.StatusOK, gin.H{
		"data": classroom,
	})
}

func DeleteClassroom(c *gin.Context) {
    var classroom entity.Classroom
	classroomID := c.Param("id")
	db := config.DB()
	db.First(&classroom, classroomID)
	if classroom.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "promotion not found"})
		return
	}
	if err := db.Delete(&classroom).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Classroom deleted successfully"})
}

func GetLisClassroomReady(c *gin.Context) {
    var classrooms []entity.Classroom

    db := config.DB()
    results := db.Preload("Condition").Preload("Building").
        Where("status = ? OR status = ?", "Ready to work", "พร้อมใช้งาน").
        Find(&classrooms)
        
    if results.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
        return
    }
    c.JSON(http.StatusOK, classrooms)
}
