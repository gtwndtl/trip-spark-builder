package schedule

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
)

func CreateSchedule(c *gin.Context) {
	var schedule entity.Schedule
	// Binding JSON data to the schedule object
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert StartTime and EndTime from string to time.Time with format HH:mm
	parsedStartTime, err := time.Parse("15:04", schedule.StartTime.Format("15:04"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid StartTime format. Expected format is HH:mm"})
		return
	}

	parsedEndTime, err := time.Parse("15:04", schedule.EndTime.Format("15:04"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid EndTime format. Expected format is HH:mm"})
		return
	}

	// เพิ่ม 7 ชั่วโมง (UTC+7)
	parsedStartTime = parsedStartTime.Add(time.Hour * 7)
	parsedEndTime = parsedEndTime.Add(time.Hour * 7)

	// Validate the schedule times (StartTime should be before EndTime)
	if err := schedule.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the parsed times to the schedule
	schedule.StartTime = parsedStartTime
	schedule.EndTime = parsedEndTime

	// Save the schedule to the database
	db := config.DB()
	if err := db.Create(&schedule).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create schedule"})
		return
	}

	// Respond with the created schedule
	c.JSON(http.StatusCreated, schedule)
}



func ListSchedules(c *gin.Context) {
	var schedules []entity.Schedule
	db := config.DB()

	if err := db.Preload("Course").Preload("Day").Preload("Classroom").Find(&schedules).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch schedules"})
		return
	}

	c.JSON(http.StatusOK, schedules)
}

func GetScheduleByID(c *gin.Context) {
	id := c.Param("id")
	var schedule entity.Schedule
	db := config.DB()

	if err := db.Preload("Course").Preload("Day").Preload("Classroom").First(&schedule, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		return
	}

	c.JSON(http.StatusOK, schedule)
}

func UpdateSchedule(c *gin.Context) {
	id := c.Param("id")
	var schedule entity.Schedule
	db := config.DB()

	// Bind JSON data to a temporary schedule object
	var inputSchedule entity.Schedule
	if err := c.ShouldBindJSON(&inputSchedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert StartTime and EndTime from string to time.Time with format HH:mm
	parsedStartTime, err := time.Parse("15:04", inputSchedule.StartTime.Format("15:04"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid StartTime format. Expected format is HH:mm"})
		return
	}

	parsedEndTime, err := time.Parse("15:04", inputSchedule.EndTime.Format("15:04"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid EndTime format. Expected format is HH:mm"})
		return
	}

	// เพิ่ม 7 ชั่วโมง (UTC+7)
	parsedStartTime = parsedStartTime.Add(time.Hour * 7)
	parsedEndTime = parsedEndTime.Add(time.Hour * 7)

	// ค้นหา schedule จากฐานข้อมูล
	if err := db.First(&schedule, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		return
	}

	// ตรวจสอบการซ้ำของวัน เวลา และห้อง
	var conflictingSchedule entity.Schedule
	if err := db.Where("day_id = ? AND classroom_id = ? AND id != ? AND ((start_time < ? AND end_time > ?) OR (start_time < ? AND end_time > ?))",
		inputSchedule.DayID, inputSchedule.ClassroomID, id,
		parsedEndTime, parsedStartTime, parsedStartTime, parsedEndTime).First(&conflictingSchedule).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Conflicting schedule exists for the same day, time, and room"})
		return
	}

	// Update schedule fields
	schedule.StartTime = parsedStartTime
	schedule.EndTime = parsedEndTime
	schedule.CourseID = inputSchedule.CourseID
	schedule.DayID = inputSchedule.DayID
	schedule.ClassroomID = inputSchedule.ClassroomID

	// Validate the updated schedule
	if err := schedule.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// บันทึกข้อมูลในฐานข้อมูล
	if err := db.Save(&schedule).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update schedule"})
		return
	}

	// ส่งข้อมูลที่อัปเดตกลับไป
	c.JSON(http.StatusOK, schedule)
}

func DeleteSchedule(c *gin.Context) {
	id := c.Param("id")
	db := config.DB()

	if err := db.Delete(&entity.Schedule{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete schedule"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Schedule deleted successfully"})
}

func GetScheduleByCourseID(c *gin.Context){
	id := c.Param("id") 
	var schedule []entity.Schedule
	
	db := config.DB()
	if err := db.Where("course_id = ?", id).Preload("Course").Find(&schedule).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch schedule"})
		return
	}

	c.JSON(http.StatusOK, schedule)
}

func GetScheduleByCourseCode(c *gin.Context) {
    courseCode := c.Param("code")
    var schedules []entity.Schedule
    db := config.DB()

    if err := db.Preload("Course", "course_code = ?", courseCode).
        Preload("Classroom").
        Preload("Day").
        Where("course_id IN (SELECT id FROM courses WHERE course_code = ?)", courseCode).
        Find(&schedules).
        Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch schedules", "details": err.Error()})
        return
    }

    // Return the results
    c.JSON(http.StatusOK, schedules)
}


func DeleteScheduleByCourseCode(c *gin.Context) {
    courseCode := c.Param("code")
    db := config.DB()

    // Query to delete schedules based on course code
    query := `
        DELETE FROM schedules 
		WHERE course_id IN (
		SELECT id FROM courses WHERE course_code = ?
		);
    `

    result := db.Exec(query, courseCode)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete schedules"})
        return
    }

    if result.RowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"message": "No schedules found for this course code"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Schedules deleted successfully"})
}


func GetScheduleByStudentID(c *gin.Context) {
	// รับ StudentID จากพารามิเตอร์
	studentID := c.Param("id")

	// ตรวจสอบว่า StudentID เป็นเลขหรือไม่
	if studentID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "StudentID is required"})
		return
	}

	var enrollments []entity.Enrollment
	db := config.DB()

	// ดึงข้อมูล Enrollment ที่มี StudentID ตรงกัน
	if err := db.Preload("Course").Where("student_id = ?", studentID).Find(&enrollments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch enrollments"})
		return
	}

	var schedules []entity.Schedule

	// ดึงข้อมูล Schedule ที่เกี่ยวข้องกับ Courses ใน Enrollment
	for _, enrollment := range enrollments {
		var courseSchedules []entity.Schedule
		if err := db.Preload("Course").
			Preload("Day").
			Preload("Classroom").
			Where("course_id = ?", enrollment.CourseID).
			Find(&courseSchedules).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch schedules"})
			return
		}
		schedules = append(schedules, courseSchedules...)
	}

	// ส่งข้อมูล schedules กลับไป
	c.JSON(http.StatusOK, schedules)
}



