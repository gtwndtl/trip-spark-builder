package books

import (
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
)

func GetAll(c *gin.Context) {


	db := config.DB()
 
 
	var books []entity.Books
	db.Preload("Student").Preload("Room").Find(&books)
 
 
	c.JSON(http.StatusOK, &books)
 
 
 }

 func Delete(c *gin.Context) {
    id := c.Param("id")

    db := config.DB()

    // ใช้ `UPDATE` เพื่ออัปเดตฟิลด์ `deleted_at` เป็นเวลาปัจจุบัน
    if tx := db.Exec("UPDATE books SET deleted_at = ? WHERE id = ? AND deleted_at IS NULL", time.Now(), id); tx.RowsAffected == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "id not found or already deleted"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Soft Deleted successfully"})
}

func GetBookByID(c *gin.Context) {
    db := config.DB()

    RoomID := c.Param("id") // id ที่ส่งมาคือ room_id

    var books []entity.Books // เปลี่ยนเป็น slice เพื่อรองรับหลายรายการ

    // Find all books with the specified room_id
    if err := db.Preload("Student").Preload("Room").Where("room_id = ?", RoomID).Find(&books).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "No books found"})
        return
    }


    // Return all matched books as JSON
    c.JSON(http.StatusOK, books)
}

func UpdateRoomStatus(c *gin.Context) {
    db := config.DB()

    // Get the RoomID from the URL parameters
    RoomID := c.Param("id")

    var count int64
    var room entity.Room

    // Count the number of bookings for the specified RoomID
    if err := db.Model(&entity.Books{}).Where("room_id = ?", RoomID).Count(&count).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count bookings"})
        return
    }

    // Update the room status based on the count
    if count >= 3 {
        // If bookings reach or exceed 3, set status_id to 2 (full)
        if err := db.Model(&room).Where("id = ?", RoomID).Update("status_id", 2).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update room status to full"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "Room status updated to full"})
    } else if (count <= 2) {
        // If bookings are less than 3, set status_id to 1 (available)
        if err := db.Model(&room).Where("id = ?", RoomID).Update("status_id", 1).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update room status to available"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "Room status updated to available", "current_bookings": count})
    }
}


func BookingCreate(c *gin.Context) {
	var req struct {
		StudentID uint      `json:"student_id"`
		RoomID    uint      `json:"room_id"`
		BooksTime time.Time `json:"books_time"`
	}
	db := config.DB()

	// ตรวจสอบความถูกต้องของข้อมูลที่ได้รับ
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบว่านักศึกษาคนนี้มีการจองอยู่แล้วหรือไม่
	var existingBooking entity.Books
	if err := db.Where("student_id = ?", req.StudentID).First(&existingBooking).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Student already has a booking"})
		return
	}

	// ตรวจสอบว่าสถานะของห้องว่างหรือไม่
	var room entity.Room
	if err := db.Where("id = ? AND status_id = 1", req.RoomID).First(&room).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Room is not available for booking"})
		return
	}

	// เริ่มต้นการ Transaction
	tx := db.Begin()

	// สร้างข้อมูลการจองหอพัก
	booking := entity.Books{
		StudentID:  req.StudentID,
		RoomID:     req.RoomID,
		BooksTime:  req.BooksTime,
	}
	if err := tx.Create(&booking).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}

	// บันทึก Transaction
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction failed"})
		return
	}

	// ส่งข้อความตอบกลับเมื่อการจองเสร็จสิ้น
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Dormitory booking confirmed successfully",
		"booking": booking,
	})
}

func GetBookByStudentID(c *gin.Context) {
	db := config.DB()

	StudentID := c.Param("id") // id ที่ส่งมาคือ student_id

	var studentBook entity.Books

	// ค้นหา room_id โดยใช้ student_id จากตาราง Books
	if err := db.Where("student_id = ?", StudentID).First(&studentBook).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student's books not found"})
		return
	}

	var books []entity.Books // รองรับหลายรายการ

	// ค้นหาหนังสือทั้งหมดของนักศึกษาที่อยู่ในห้องเดียวกัน
	if err := db.Preload("Student").Preload("Room").
		Where("room_id = ?", studentBook.RoomID).Find(&books).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No books found for students in this room"})
		return
	}

	// ส่งคืน JSON
	c.JSON(http.StatusOK, books)
}
