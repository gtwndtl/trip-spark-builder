package building

import (
	"net/http"
    "path/filepath"
	"fmt"
    "strconv"
	"gorm.io/gorm"
	"errors"
    "time"
    // "bytes"
    

	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
    "os"
)
func UploadFile(c *gin.Context) {
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    buildingName := c.PostForm("BuildingName") // รับ BuildingName จาก request
    if buildingName == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "BuildingName is required"})
        return
    }

    // สร้างชื่อไฟล์ใหม่ตาม BuildingName และนามสกุลไฟล์เดิม
    extension := filepath.Ext(file.Filename)
    newFileName := fmt.Sprintf("%s%s", buildingName, extension)

    path := fmt.Sprintf("./uploads/%s", newFileName)
    if err := c.SaveUploadedFile(file, path); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Upload successful", "path": path})
}

func CreateBuilding(c *gin.Context) {
	var Building entity.Building

	// รับไฟล์จาก request
	file, err := c.FormFile("BuildingPicture")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "BuildingPicture is required", "status": 400})
		return
	}

	uploadPath := "./uploads"

	// สร้างชื่อไฟล์ใหม่ไม่ให้ซ้ำกัน
	timestamp := time.Now().Unix()
	extension := filepath.Ext(file.Filename)
	newFileName := fmt.Sprintf("building_%d%s", timestamp, extension)

	// กำหนด path สำหรับบันทึกไฟล์
	BuildingPicture := filepath.Join(uploadPath, newFileName)
	if err := c.SaveUploadedFile(file, BuildingPicture); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file", "status": 500})
		return
	}

	// อ่านข้อมูลอื่น ๆ ของฟอร์ม
	BuildingName := c.PostForm("BuildingName")
	if BuildingName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "BuildingName is required", "status": 400})
		return
	}

	BuildingCode := c.PostForm("BuildingCode")
	if BuildingCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "BuildingCode is required", "status": 400})
		return
	}

	Location := c.PostForm("Location")
	if Location == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Location is required", "status": 400})
		return
	}

	TotalFloors := c.PostForm("TotalFloors")
	totalFloorsInt, err := strconv.Atoi(TotalFloors)
	if err != nil || totalFloorsInt <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid TotalFloors format or value", "status": 400})
		return
	}

	ConditionID := c.PostForm("ConditionID")
	conditionIDInt, err := strconv.Atoi(ConditionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ConditionID format", "status": 400})
		return
	}

	//กำหนดชื่อใหม่("/uploads/%s", newFileName)
	BuildingPicture = fmt.Sprintf("/uploads/%s", newFileName)
	// BuildingPicture = fmt.Sprintf("%s/%s", uploadPath, newFileName)
	// สร้าง Building entity และบันทึกลงฐานข้อมูล
	Building = entity.Building{
		BuildingName:    BuildingName,
		BuildingCode:    BuildingCode,
		Location:        Location,
		TotalFloors:     totalFloorsInt,
		BuildingPicture: BuildingPicture, // บันทึก path ของรูปภาพ
		ConditionID:     uint(conditionIDInt),
	}

	// เชื่อมต่อกับฐานข้อมูล
	db := config.DB()
	if err := db.Create(&Building).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create building", "status": 400})
		return
	}

	// ส่งข้อมูลกลับไปที่ client
	c.JSON(http.StatusCreated, gin.H{"message": "Building created successfully", "data": Building, "status": 201})
}

func GetBuildingsByID(c *gin.Context) {
	// ดึง BuildingID จาก URL params
	buildingID := c.Param("id")

	// ตรวจสอบว่า BuildingID เป็นตัวเลขหรือไม่
	id, err := strconv.ParseUint(buildingID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Building ID"})
		return
	}

	// ดึงข้อมูล Building และห้องเรียนที่เกี่ยวข้อง
	var building entity.Building
	db := config.DB()

	// ดึงข้อมูล Building พร้อมข้อมูลห้องเรียนที่เกี่ยวข้อง
	if err := db.Preload("Classroom").First(&building, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Building not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ส่ง response กลับไปยัง frontend
	c.JSON(http.StatusOK, gin.H{
		"data": building,
	})
}
func GetBuildings(c *gin.Context) {
    var buildings []entity.Building // ใช้ slice เพื่อเก็บหลายรายการ
    db := config.DB()
    results := db.Preload("Condition").Find(&buildings) // ดึงข้อมูลทั้งหมด
    if results.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": results.Error.Error()})
        return
    }
    if len(buildings) == 0 {
        c.JSON(http.StatusNoContent, gin.H{})
        return
    }
    c.JSON(http.StatusOK, buildings) // ส่งข้อมูลกลับในรูปแบบ array
}
func GetBuilding(c *gin.Context) {
    ID := c.Param("id")
    var building entity.Building
    db := config.DB()
    results := db.Preload("Condition").First(&building, ID)
    if results.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
        return
    }
    if building.ID == 0 {
        c.JSON(http.StatusNoContent, gin.H{})
        return
    }
    c.JSON(http.StatusOK, building)
}

func UpdateBuilding(c *gin.Context) {
	// รับ ID ของอาคารจาก URL parameter
	id := c.Param("id")

	var Building entity.Building

	// ตรวจสอบว่า ID ที่ส่งมามีอาคารอยู่ในฐานข้อมูลหรือไม่
	db := config.DB()
	if err := db.Where("id = ?", id).First(&Building).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Building not found", "status": 404})
		return
	}

	// ตรวจสอบว่าไฟล์ใหม่ถูกอัปโหลดมาหรือไม่
	file, err := c.FormFile("BuildingPicture")
	if err == nil {
		// หากมีไฟล์ใหม่ อัปโหลดไฟล์และอัปเดตรูปภาพ
		uploadPath := "./uploads"
		timestamp := time.Now().Unix()
		extension := filepath.Ext(file.Filename)
		newFileName := fmt.Sprintf("building_%d%s", timestamp, extension)
		BuildingPicture := filepath.Join(uploadPath, newFileName)

		// บันทึกไฟล์ใหม่
		if err := c.SaveUploadedFile(file, BuildingPicture); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file", "status": 500})
			return
		}

		// ลบไฟล์เก่าที่อยู่ในระบบ
		if Building.BuildingPicture != "" {
			RemovePic :=  fmt.Sprintf(".%s", Building.BuildingPicture)
			os.Remove(RemovePic)
			// os.Remove(Building.BuildingPicture)
		}
		BuildingPicture = fmt.Sprintf("/uploads/%s", newFileName)
		// BuildingPicture = fmt.Sprintf("%s/%s", uploadPath, newFileName)
		// อัปเดต path ของรูปภาพในฐานข้อมูล
		Building.BuildingPicture = BuildingPicture
	}

	// อ่านข้อมูลที่ส่งมาจากฟอร์มและอัปเดต
	BuildingName := c.PostForm("BuildingName")
	if BuildingName != "" {
		Building.BuildingName = BuildingName
	}

	BuildingCode := c.PostForm("BuildingCode")
	if BuildingCode != "" {
		Building.BuildingCode = BuildingCode
	}

	Location := c.PostForm("Location")
	if Location != "" {
		Building.Location = Location
	}

	TotalFloors := c.PostForm("TotalFloors")
	if TotalFloors != "" {
		totalFloorsInt, err := strconv.Atoi(TotalFloors)
		if err != nil || totalFloorsInt <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid TotalFloors format or value", "status": 400})
			return
		}
		Building.TotalFloors = totalFloorsInt
	}

	ConditionID := c.PostForm("ConditionID")
	if ConditionID != "" {
		conditionIDInt, err := strconv.Atoi(ConditionID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ConditionID format", "status": 400})
			return
		}
		Building.ConditionID = uint(conditionIDInt)
	}

	// บันทึกการเปลี่ยนแปลงลงฐานข้อมูล
	if err := db.Save(&Building).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update building", "status": 500})
		return
	}

	// ส่งข้อมูลกลับไปยัง client
	c.JSON(http.StatusOK, gin.H{"message": "Building updated successfully", "data": Building, "status": 200})
}

func DeleteBuilding(c *gin.Context) {
	// ดึง ID อาคารจาก URL params
	buildingID := c.Param("id")

	// ตรวจสอบว่า buildingID เป็นตัวเลขหรือไม่
	id, err := strconv.ParseUint(buildingID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Building ID"})
		return
	}

	db := config.DB()

	// ตรวจสอบว่าอาคารมีอยู่จริงหรือไม่
	var building entity.Building
	if err := db.First(&building, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Building not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ลบไฟล์รูปภาพของอาคาร (BuildingPicture)
	if building.BuildingPicture != "" {
		deleteBP := fmt.Sprintf(".%s", building.BuildingPicture)
		err := os.Remove(deleteBP) // ลบไฟล์รูปภาพ
		if err != nil && !os.IsNotExist(err) {      // ถ้าไฟล์ไม่พบให้ผ่านไป
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to delete building picture: " + err.Error(),
			})
			return
		}
	}

	// ลบห้องเรียน (Classroom) ที่เกี่ยวข้องกับ BuildingID
	if err := db.Where("building_id = ?", id).Delete(&entity.Classroom{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete related classrooms"})
		return
	}

	// ลบข้อมูลอาคารในฐานข้อมูล
	if err := db.Delete(&building).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete building"})
		return
	}

	// ส่งข้อความยืนยันกลับไปยัง frontend
	c.JSON(http.StatusOK, gin.H{"message": "Building and related data deleted successfully"})
}