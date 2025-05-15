package room

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
)

// CreateRoom creates a new room with automatically derived values
func CreateRoom(c *gin.Context) {
	var Room entity.Room

	// Get the room number from the request
	RoomNumber := c.PostForm("RoomNumber")
	if RoomNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณากรอกเลขห้อง"})
		return
	}
	Room.RoomNumber = RoomNumber

	// Check if the room number already exists in the database
	db := config.DB()
	var existingRoom entity.Room
	if err := db.Where("room_number = ?", RoomNumber).First(&existingRoom).Error; err == nil {
		// Room number already exists
		c.JSON(http.StatusBadRequest, gin.H{"error": "ไม่พบหอพักนี้ในระบบ"})
		return
	}

	// Parse room number (assuming room format "2115")
	// Index 1 = Dormitory, Index 2 = Floor, Rest = RoomUnit
	dormitoryDigit := RoomNumber[1] // '2' (for Dormitory)
	floorDigit := RoomNumber[2]    // '1' (for Floor)

	// Convert Dormitory and Floor digits to integers
	DormitoryID, err := strconv.Atoi(string(dormitoryDigit))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ไม่พบหอพักนี้ในระบบ"})
		return
	}

	FloorID, err := strconv.Atoi(string(floorDigit))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ไม่พบชั้นนี้ในระบบ"})
		return
	}

	// Set default values for other fields
	Room.Score = 0
	Room.StatusID = 1 // Assuming "1" means Active status

	// Find the related floor and dormitory records
	var Floor entity.Floor
	if err := db.First(&Floor, FloorID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่พบชั้นนี้ในระบบ"})
		return
	}
	Room.Floor = Floor
	Room.FloorID = Floor.ID

	var Dormitory entity.Dormitory
	if err := db.First(&Dormitory, DormitoryID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่พบหอพักนี้ในระบบ"})
		return
	}
	Room.Dormitory = Dormitory
	Room.DormitoryID = Dormitory.ID

	// Create the room record
	if err := db.Create(&Room).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return success with the created room
	c.JSON(http.StatusCreated, gin.H{"message": "สร้างห้องสำเร็จ", "room": Room})
}

func ListRoom(c *gin.Context) {
    var dormitory entity.Dormitory // ใช้สำหรับตรวจสอบว่ามี dormitory นี้จริงหรือไม่
    var rooms []entity.Room

    dormitoryID := c.Param("id") // รับ DormitoryID จาก URL params

    if dormitoryID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ไม่พบหอพักนี้ในระบบ"})
        return
    }

    // เชื่อมต่อกับฐานข้อมูล
    db := config.DB()

    // ตรวจสอบว่ามี dormitory ตาม id นี้ในตาราง dormitory หรือไม่
    if err := db.Where("id = ? AND deleted_at IS NULL", dormitoryID).First(&dormitory).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบหอพักนี้"})
        return
    }

    // ค้นหา rooms ที่สัมพันธ์กับ dormitory นี้
    results := db.Where("dormitory_id = ? AND deleted_at IS NULL", dormitoryID).Find(&rooms)
    if results.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": results.Error.Error()})
        return
    }

    // หากไม่พบห้องพักที่สัมพันธ์กับ dormitory นี้
    if len(rooms) == 0 {
        c.JSON(http.StatusNoContent, nil) // 204 No Content
        return
    }

    // ส่งข้อมูล rooms กลับในรูปแบบ JSON
    c.JSON(http.StatusOK, rooms) // 200 OK พร้อมข้อมูลห้องพัก
}

  func DeleteRoom(c *gin.Context) {
    // รับค่าจาก URL path parameter
    RoomNumber := c.Param("RoomNumber") // รับค่าจาก path parameter

    db := config.DB()

    // ตรวจสอบว่า RoomNumber มีค่าหรือไม่
    if RoomNumber == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณากรอกเลขห้อง"})
        return
    }

    // ลบข้อมูลห้องที่ตรงกับ RoomNumber
    tx := db.Exec("DELETE FROM Rooms WHERE room_number = ?", RoomNumber)

    // ตรวจสอบข้อผิดพลาดจาก tx.Error
    if tx.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete room", "details": tx.Error.Error()})
        return
    }

    // ตรวจสอบผลลัพธ์ของการลบ
    rowsAffected := tx.RowsAffected
    if rowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบห้องนี้ในระบบ"})
        return
    }

    // ส่งผลลัพธ์กลับไปยัง frontend
    c.JSON(http.StatusOK, gin.H{"message": "ลบห้องพัก", "RoomNumber": RoomNumber})
}

func DeleteAllRoomsInDormitory(c *gin.Context) {
    dormitoryID := c.Param("id")  // รับค่า DormitoryID จาก URL params

    // ตรวจสอบว่า dormitoryID มีค่าหรือไม่
    if dormitoryID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "DormitoryID is required"})
        return
    }

    db := config.DB()

    // ลบห้องพักทั้งหมดที่มี dormitory_id ตรงกับค่าที่รับมา
    tx := db.Where("dormitory_id = ?", dormitoryID).Delete(&entity.Room{})

    // ตรวจสอบข้อผิดพลาดจาก tx.Error
    if tx.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete rooms", "details": tx.Error.Error()})
        return
    }

    // ตรวจสอบผลลัพธ์ของการลบ
    rowsAffected := tx.RowsAffected
    if rowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "No rooms found for this dormitory"})
        return
    }

    // ส่งผลลัพธ์กลับไปยัง frontend
    c.JSON(http.StatusOK, gin.H{"message": "All rooms deleted successfully", "dormitoryID": dormitoryID})
}

func GetAll(c *gin.Context) {


    db := config.DB()
 
 
    var room []entity.Room
    db.Preload("Dormitory").Preload("Status").Find(&room)
 
 
    c.JSON(http.StatusOK, &room)
 
 
 }
