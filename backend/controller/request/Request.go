package request

import (
	"net/http"
	"time"

	"github.com/sut67/team09/config"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/entity"
)

// GET /request/:id
func GetRequestByID(c *gin.Context) {
	ID := c.Param("id")
	var request entity.Request

	db := config.DB()

	// Join table sellers กับ members โดยใช้ member_id
	result := db.First(&request, ID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, request)
}


// POST /Requests
func CreateRequest(c *gin.Context) {
	var request entity.Request

	// Bind the request data to the Request struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB()

	request.RequestDate = time.Now()
	// Create the new Request
	newRequest := entity.Request{
		RequestDate: time.Now(),
		Note : request.Note ,    
		StudentID:request.StudentID, 
		RequestTypeID: request.RequestTypeID,
		StatusRequestID:request.StatusRequestID,
		CourseID: request.CourseID,
	}


	// Save the new Request to the database
	if err := db.Create(&newRequest).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	c.JSON(http.StatusCreated, gin.H{"message": "Created successfully", "data": newRequest})
}


// GET /request
func GetRequest(c *gin.Context) { 
	var request []entity.Request
	

	db := config.DB()
	result := db.Preload("Course").Find(&request)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, request)
}

// PATCH /request
func UpdateRequest(c *gin.Context) {
	var Request entity.Request
	RequestID := c.Param("id")

	db := config.DB()
	result := db.First(&Request, RequestID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}

	if err := c.ShouldBindJSON(&Request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	result = db.Save(&Request)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})
}



// ลบแบบหายไปจากตาราง
func DeleteRequest(c *gin.Context) { //ลบข้อมูลตาม id
	id := c.Param("id")
	db := config.DB()
	if tx := db.Exec("DELETE FROM requests WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})
}


// GET /request/:student_id
func GetRequestByStudentID(c *gin.Context) {

    studentID := c.Param("student_id")

    if studentID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Student ID is required"})
        return
    }

    var requests []entity.Request

    db := config.DB()
    result := db.Preload("Student").Preload("RequestType").Preload("StatusRequest").Preload("Course").Where("student_id = ?", studentID).Find(&requests)


    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
        return
    }


    c.JSON(http.StatusOK, requests)
}




// GET /course/:course_code
func GetCourseByCourseCode(c *gin.Context) {
    courseCode := c.Param("course_code")

    if courseCode == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Course code is required"})
        return
    }

    var course entity.Course

    db := config.DB()

    // Case-insensitive query using LOWER()
    result := db.Where("LOWER(course_code) = LOWER(?)", courseCode).First(&course)

    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, course)
}


// GET /course/:course_code/:group
func GetCourseByCourseCodeAndGroup(c *gin.Context) {
    courseCode := c.Param("course_code")
    group := c.Param("group")

    if courseCode == "" || group == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Course code and group are required"})
        return
    }

    var course entity.Course

    db := config.DB()

    // Escape the reserved keyword `group` using double quotes
    result := db.Where("LOWER(course_code) = LOWER(?) AND \"group\" = ?", courseCode, group).First(&course)

    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Course or group not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, course)
}








func GetCoursesByRequestID(c *gin.Context) {
    requestID := c.Param("id")
    var courses []entity.Course

    db := config.DB()

    // Query to join tables and filter by request_id
    result := db.
        Joins("JOIN requests ON requests.course_id = courses.id").
        Where("requests.id = ?", requestID).
        Preload("Lecturer").
        Find(&courses)

    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"courses": courses})
}




// // GET /request/course/printstory
// func GetRequestAndPrintStoryAndCourse(c *gin.Context) {
//     // รับค่า ID ของคำร้อง
//     requestID := c.Param("id") 

//     // if requestID == "" {
//     //     c.JSON(http.StatusBadRequest, gin.H{"error": "Request ID is required"})
//     //     return
//     // }

//     var requests []entity.Request

//     db := config.DB()

//     // ใช้ Preload เพื่อโหลดข้อมูลที่เกี่ยวข้องกับ Request
//     result := db.
//         Preload("Student").          // โหลดข้อมูล Student
//         Preload("RequestType").      // โหลดข้อมูล RequestType
//         Preload("StatusRequest").    // โหลดข้อมูล StatusRequest
//         Preload("Course").
//         Preload("PrintStory").
// 		Where("id = ?", requestID).
        
//         Find(&requests)

//     if result.Error != nil {
//         c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
//         return
//     }

//     // ส่งข้อมูลทั้งหมดกลับ
//     c.JSON(http.StatusOK, requests)
// }

