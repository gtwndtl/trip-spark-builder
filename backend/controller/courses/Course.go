package courses

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
)

// ดึงรายการคอร์สทั้งหมด
func ListAllCourse(c *gin.Context) {
	var courses []entity.Course

	// ดึงข้อมูลจากฐานข้อมูล
	db := config.DB()
	if err := db.Preload("Lecturer").Find(&courses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
		return
	}

	c.JSON(http.StatusOK, courses)
}

func ListCourse(c *gin.Context) {
	var courses []entity.Course
	// ดึงข้อมูลที่ course_code ไม่ซ้ำ
	db := config.DB()
	if err := db.Raw(`
		SELECT * 
		FROM courses
		WHERE id IN (
			SELECT MIN(id) 
			FROM courses 
			GROUP BY course_code
		)
	`).Scan(&courses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
		return
	}
	// ส่งข้อมูลกลับ
	c.JSON(http.StatusOK, courses)
}


// เพิ่มคอร์สใหม่
func CreateCourse(c *gin.Context) {
	var course entity.Course

	// ตรวจสอบว่า JSON ถูกต้องหรือไม่
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB()

	// ตรวจสอบว่า CourseCode และ Group มีอยู่ในระบบหรือไม่
	var existingCourse entity.Course
	if err := db.Where("course_code = ? AND `group` = ?", course.CourseCode, course.Group).First(&existingCourse).Error; err == nil {
		// หากพบวิชาที่เหมือนกัน
		c.JSON(http.StatusConflict, gin.H{
			"error": "Course with the same CourseCode and Group already exists",
		})
		return
	}

	// สร้างวิชาใหม่
	if err := db.Create(&course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create course"})
		return
	}

	c.JSON(http.StatusCreated, course)
}


// ดึงข้อมูลคอร์สตาม ID
func GetCourseByID(c *gin.Context) {
	id := c.Param("id")
	var course entity.Course

	// ค้นหาข้อมูลจากฐานข้อมูล
	db := config.DB()
	if err := db.Preload("Lecturer").First(&course, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(http.StatusOK, course)
}

// อัปเดตข้อมูลคอร์ส
func UpdateCourse(c *gin.Context) {
	id := c.Param("id")
	var course entity.Course

	db := config.DB()
	if err := db.First(&course, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	// อัปเดตข้อมูล
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update course"})
		return
	}

	c.JSON(http.StatusOK, course)
}

// ลบคอร์ส
func DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	db := config.DB()

	if err := db.Unscoped().Delete(&entity.Course{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}

func DeleteByCourseCode(c *gin.Context) {
	courseCode := c.Param("course_code")
	db := config.DB()

	if err := db.Unscoped().Where("`course_code` = ?", courseCode).Delete(&entity.Course{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}


// ดึงรายการคอร์สตาม LecturerID
func GetCoursesByLecturerID(c *gin.Context) {
	id := c.Param("id") // รับ LecturerID จาก URL parameter
	var courses []entity.Course

	// ดึงข้อมูลที่ course_code ไม่ซ้ำ
	db := config.DB()
	query := `
		SELECT * 
		FROM courses
		WHERE id IN (
			SELECT MAX(id) 
			FROM courses 
			WHERE lecturer_id = ?
			GROUP BY course_code
		)
	`
	if err := db.Raw(query, id).Scan(&courses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
		return
	}

	// ส่งข้อมูลกลับ
	c.JSON(http.StatusOK, courses)
}

func SearchCourseByKeyword(c *gin.Context) {
    keyword := c.Query("keyword")

    var courses []entity.Course
    db := config.DB()

    // Using the correct syntax for multiple OR conditions in GORM
    results := db.Where("course_name LIKE ? OR course_code LIKE ?", "%"+keyword+"%", "%"+keyword+"%").Find(&courses)

    if results.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": results.Error.Error()})
        return
    }

    if len(courses) == 0 {
        c.JSON(http.StatusNoContent, gin.H{"message": "No course found"})
        return
    }

    c.JSON(http.StatusOK, courses)
}



// ดึงคอร์สที่มี CourseCode ตรงกัน
func GetRelatedCourses(c *gin.Context) {
	courseCode := c.DefaultQuery("courseCode", "") // รับ CourseCode จาก query parameter
	if courseCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CourseCode is required"})
		return
	}

	var relatedCourses []entity.Course

	// ค้นหาคอร์สที่มี CourseCode ตรงกัน
	db := config.DB()
	if err := db.Where("course_code = ?", courseCode).Preload("Lecturer").Find(&relatedCourses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch related courses"})
		return
	}

	// หากไม่พบคอร์สที่ตรงกับ CourseCode
	if len(relatedCourses) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No related courses found"})
		return
	}

	c.JSON(http.StatusOK, relatedCourses)
}

func ToggleCourseStatus(c *gin.Context) {

	courseID := c.Param("id")

	var course entity.Course
	db := config.DB()
	if err := db.First(&course, courseID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	course.Status = !course.Status

	if err := db.Save(&course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update course status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Course status updated successfully",
		"course": course,
	})
}

func GetAll(c *gin.Context) {


	db := config.DB()
 
 
	var course []entity.Course
	db.Find(&course)
 
 
	c.JSON(http.StatusOK, &course)
 
 
 }