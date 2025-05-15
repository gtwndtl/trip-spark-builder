package activity

import (
	"fmt"
	"net/http"

	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
	"gorm.io/gorm"
)

// GET /activity
func GetActivity(c *gin.Context) { // เข้าถึงข้อมูลสินค้าทั้งหมด
	var activity []entity.Activity

	db := config.DB()
	result := db.Find(&activity)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, activity)
}

// GET /activity/:id
func GetActivityById(c *gin.Context) {
    // Extract the ID from the URL parameter
    id := c.Param("id")

    var activity entity.Activity

    // Query the database for the specific activity by ID
    db := config.DB()
    result := db.First(&activity, id) // This will find the first matching record by ID

    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, activity)
}




// POST /activity
func CreateActivity(c *gin.Context) {
    var activity entity.Activity


    activity.ActivityName = c.PostForm("ActivityName")
    activity.Description = c.PostForm("Description")
    activity.Organizer = c.PostForm("Organizer")
    activity.Location = c.PostForm("Location")

    activityDateStr := c.PostForm("ActivityDate")
    activityDate, err := time.Parse("2006-01-02", activityDateStr) 
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Activity Date"})
        return
    }
    activity.ActivityDate = activityDate.In(time.UTC)


    activity.StartTime, _ = time.Parse("15:04:05", c.PostForm("StartTime"))
    activity.EndTime, _ = time.Parse("15:04:05", c.PostForm("EndTime"))
    
    adminIDStr := c.PostForm("AdminID")
    adminID, err := strconv.Atoi(adminIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Admin ID"})
        return
    }
    activity.AdminID = uint(adminID)

    activity.MaxParticipants, _ = strconv.Atoi(c.PostForm("MaxParticipants"))
    activity.StatusActivityID = 1

    // Handle file upload
    file, err := c.FormFile("ActivityPic")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Activity picture is required"})
        return
    }

    // filename := fmt.Sprintf("%s_%d%s", strings.ReplaceAll(activity.ActivityName, " ", "_"), time.Now().Unix(), filepath.Ext(file.Filename))

    currentDate := time.Now().Format("20060102")

    filename := fmt.Sprintf("%s_%s%s", strings.ReplaceAll(activity.ActivityName, " ", "_"), currentDate, filepath.Ext(file.Filename))

    savePath := filepath.Join("uploads_pic_activity", filename)

    if err := c.SaveUploadedFile(file, savePath); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
        return
    }

    // Save file path to database
    activity.ActivityPic = savePath

    db := config.DB()
    if err := db.Create(&activity).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Created success", "data": activity})
}






//อัปเดตได้แล้ว

func UpdateActivity(c *gin.Context) { // Update activity based on ID
    var activity entity.Activity

    ActivityID := c.Param("id")

    // Retrieve existing activity by ID
    db := config.DB()
    result := db.First(&activity, ActivityID)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Activity ID not found"})
        return
    }

    // Update fields from the form
    activity.ActivityName = c.PostForm("ActivityName")
    activity.Description = c.PostForm("Description")
    activity.Organizer = c.PostForm("Organizer")
    activity.Location = c.PostForm("Location")

    activityDateStr := c.PostForm("ActivityDate")
    activityDate, err := time.Parse("2006-01-02", activityDateStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Activity Date"})
        return
    }
    activity.ActivityDate = activityDate.In(time.UTC)

    activity.StartTime, _ = time.Parse("15:04", c.PostForm("StartTime"))
    activity.EndTime, _ = time.Parse("15:04", c.PostForm("EndTime"))

    adminIDStr := c.PostForm("AdminID")
    adminID, err := strconv.Atoi(adminIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Admin ID"})
        return
    }
    activity.AdminID = uint(adminID)

    activity.MaxParticipants, err = strconv.Atoi(c.PostForm("MaxParticipants"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Max Participants"})
        return
    }
    activity.StatusActivityID = 1

    // Handle file upload
    file, err := c.FormFile("ActivityPic")
    if err == nil { // File is optional, handle only if uploaded
        currentDate := time.Now().Format("20060102")
        filename := fmt.Sprintf("%s_%s%s", strings.ReplaceAll(activity.ActivityName, " ", "_"), currentDate, filepath.Ext(file.Filename))
        savePath := filepath.Join("uploads_pic_activity", filename)

        if err := c.SaveUploadedFile(file, savePath); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
            return
        }

        activity.ActivityPic = savePath
    }

    // Save updated activity to the database
    result = db.Save(&activity)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update activity"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Updated successfully"})
}





//ลบแบบหายไปจากตาราง
func DeleteActivitys(c *gin.Context) { //ลบข้อมูลตาม id
	id := c.Param("id")
	db := config.DB()
	if tx := db.Exec("DELETE FROM activities WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})
}




