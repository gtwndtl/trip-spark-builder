package lecturer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
	"github.com/sut67/team09/services"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LecturerResponse struct {
	ID              uint   `json:"ID"`
	LecturerCode    string `json:"LecturerCode"`
	FirstName       string `json:"FirstName"`
	LastName        string `json:"LastName"`
	Email           string `json:"Email"`
	NationalID      string `json:"NationalID"`
	Phone           string `json:"Phone"`
	Profile         string `json:"Profile"`
	BirthDay        string `json:"BirthDay"`
	DateEmployed    string `json:"DateEmployed"`
	PositionID      uint   `json:"PositionID"`
	PositionName    string `json:"PositionName"`
	GenderID        uint   `json:"GenderID"`
	GenderName      string `json:"GenderName"`
	MajorID         uint   `json:"MajorID"`
	MajorName       string `json:"MajorName"`
	FacultyID       uint   `json:"FacultyID"`
	FacultyName     string `json:"FacultyName"`
	StatusStaffID   uint   `json:"StatusStaffID"`
	StatusStaffName string `json:"StatusStaffName"`
}

func CreateLecturer(c *gin.Context) {
	var lecturerData entity.Lecturer

	if err := c.ShouldBindJSON(&lecturerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ข้อมูลไม่ถูกต้อง",
		})
		return
	}

	db := config.DB()

	var genderID, positionID, majorID, statusStaffID uint
	statusStaffID = 1

	if lecturerData.GenderID != nil {
		genderID = uint(*lecturerData.GenderID)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ไม่พบข้อมูลเพศ",
		})
		return
	}

	if lecturerData.PositionID != nil {
		positionID = uint(*lecturerData.PositionID)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ไม่พบข้อมูลตำแหน่ง",
		})
		return
	}

	if lecturerData.MajorID != nil {
		majorID = uint(*lecturerData.MajorID)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ไม่พบข้อมูลสาขา",
		})
		return
	}

	tx := db.Begin()

	hashedPassword, err := config.HashPassword(lecturerData.Password)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "ไม่สามารถเข้ารหัสรหัสผ่านได้",
		})
		return
	}

	lecturer := entity.Lecturer{
		LecturerCode:  lecturerData.LecturerCode,
		FirstName:     lecturerData.FirstName,
		LastName:      lecturerData.LastName,
		BirthDay:      lecturerData.BirthDay,
		Email:         lecturerData.Email,
		Password:      hashedPassword,
		NationalID:    lecturerData.NationalID,
		Phone:         lecturerData.Phone,
		Profile:       lecturerData.Profile,
		DateEmployed:  lecturerData.DateEmployed,
		GenderID:      &genderID,
		PositionID:    &positionID,
		MajorID:       &majorID,
		StatusStaffID: &statusStaffID,
	}

	if err := tx.Create(&lecturer).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ไม่สามารถเพิ่มข้อมูลอาจารย์ได้",
		})
		return
	}

	tx.Commit()

	c.JSON(http.StatusCreated, gin.H{
		"message": "ข้อมูลอาจารย์ถูกเพิ่มเรียบร้อยแล้ว",
		"data":    lecturer,
	})
}

func UpdateLecturer(c *gin.Context) {
	var lecturer entity.Lecturer
	LecturerID := c.Param("id")

	db := config.DB()
	result := db.First(&lecturer, LecturerID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบข้อมูลอาจารย์"})
		return
	}

	if err := c.ShouldBindJSON(&lecturer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	result = db.Save(&lecturer)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "แก้ไขข้อมูลไม่สำเร็จ"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "แก้ไขข้อมูลอาจารย์สำเร็จ"})
}

func CheckEmail(c *gin.Context) {
	Email := c.Param("email")
	db := config.DB()

	var lecturer entity.Lecturer
	lecturerResult := db.Where("email = ?", Email).First(&lecturer)

	if lecturerResult.Error != nil && lecturerResult.Error != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error": lecturerResult.Error.Error()})
		return
	}

	if lecturerResult.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"isValid": false,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"isValid": true,
		})
	}
}

func CheckPhone(c *gin.Context) {
	phone := c.Param("phone")
	db := config.DB()

	var lecturer entity.Lecturer
	var student entity.Student

	lecturerResult := db.Where("phone = ?", phone).First(&lecturer)
	studentResult := db.Where("phone = ?", phone).First(&student)

	if (lecturerResult.Error != nil && lecturerResult.Error != gorm.ErrRecordNotFound) ||
		(studentResult.Error != nil && studentResult.Error != gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "เกิดข้อผิดพลาดในการตรวจสอบเบอร์โทรศัพท์"})
		return
	}

	if lecturerResult.RowsAffected > 0 || studentResult.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"isValid": false,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"isValid": true,
		})
	}
}

func CheckLecturerCode(c *gin.Context) {
	lecturercode := c.Param("lecturercode")
	db := config.DB()

	var lecturer entity.Lecturer

	lecturerResult := db.Where("lecturer_code = ?", lecturercode).First(&lecturer)

	if (lecturerResult.Error != nil && lecturerResult.Error != gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "เกิดข้อผิดพลาดในการตรวจสอบรหัสอาจารย์"})
		return
	}

	if lecturerResult.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"isValid": false,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"isValid": true,
		})
	}
}

func CheckNationalID(c *gin.Context) {
	nationalID := c.Param("nationalID")
	db := config.DB()

	var lecturer entity.Lecturer
	var student entity.Student

	lecturerResult := db.Where("national_id = ?", nationalID).First(&lecturer)
	studentResult := db.Where("national_id = ?", nationalID).First(&student)

	if (lecturerResult.Error != nil && lecturerResult.Error != gorm.ErrRecordNotFound) ||
		(studentResult.Error != nil && studentResult.Error != gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "เกิดข้อผิดพลาดในการตรวจสอบเลขบัตรประชาชน"})
		return
	}

	if lecturerResult.RowsAffected > 0 || studentResult.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"isValid": false,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"isValid": true,
		})
	}
}

func CheckEmailByID(c *gin.Context) {
	Email := c.Param("email")
	id := c.Param("id")

	db := config.DB()

	var lecturer entity.Lecturer
	lecturerResult := db.Where("email = ? AND id != ?", Email, id).First(&lecturer)

	if lecturerResult.Error != nil && lecturerResult.Error != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error": lecturerResult.Error.Error()})
		return
	}

	if lecturerResult.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"isValid": false,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"isValid": true,
		})
	}
}

func CheckPhoneByID(c *gin.Context) {
	phone := c.Param("phone")
	id := c.Param("id")
	db := config.DB()

	var lecturer entity.Lecturer
	var student entity.Student

	lecturerResult := db.Where("phone = ? AND id != ?", phone, id).First(&lecturer)
	studentResult := db.Where("phone = ? AND id != ?", phone, id).First(&student)

	if (lecturerResult.Error != nil && lecturerResult.Error != gorm.ErrRecordNotFound) ||
		(studentResult.Error != nil && studentResult.Error != gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "เกิดข้อผิดพลาดในการตรวจสอบเบอร์โทรศัพท์"})
		return
	}

	if lecturerResult.RowsAffected > 0 || studentResult.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"isValid": false,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"isValid": true,
		})
	}
}

func CheckLecturerCodeByID(c *gin.Context) {
	lecturercode := c.Param("lecturercode")
	id := c.Param("id")
	db := config.DB()

	var lecturer entity.Lecturer

	lecturerResult := db.Where("lecturer_code = ? AND id != ?", lecturercode, id).First(&lecturer)

	if (lecturerResult.Error != nil && lecturerResult.Error != gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "เกิดข้อผิดพลาดในการตรวจสอบรหัสอาจารย์"})
		return
	}

	if lecturerResult.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"isValid": false,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"isValid": true,
		})
	}
}

func CheckNationalIDByID(c *gin.Context) {
	nationalID := c.Param("nationalid")
	id := c.Param("id")
	db := config.DB()

	var lecturer entity.Lecturer
	var student entity.Student

	lecturerResult := db.Where("national_id = ? AND id != ?", nationalID, id).First(&lecturer)
	studentResult := db.Where("national_id = ? AND id != ?", nationalID, id).First(&student)

	if (lecturerResult.Error != nil && lecturerResult.Error != gorm.ErrRecordNotFound) ||
		(studentResult.Error != nil && studentResult.Error != gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "เกิดข้อผิดพลาดในการตรวจสอบเลขบัตรประชาชน"})
		return
	}

	if lecturerResult.RowsAffected > 0 || studentResult.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"isValid": false,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"isValid": true,
		})
	}
}

func GetLecturerByID(c *gin.Context) {
	ID := c.Param("id")
	var lecturer LecturerResponse

	db := config.DB()

	if err := db.Raw(`
		SELECT 
			l.*,
			p.position AS position_name, 
			g.gender AS gender_name, 
			m.major_name AS major_name, 
			m.faculty_id, 
			f.faculty_name AS faculty_name, 
			s.status_staff AS status_staff_name
		FROM lecturers l
		LEFT JOIN positions p ON l.position_id = p.id
		LEFT JOIN genders g ON l.gender_id = g.id
		LEFT JOIN majors m ON l.major_id = m.id
		LEFT JOIN faculties f ON m.faculty_id = f.id
		LEFT JOIN status_staffs s ON l.status_staff_id = s.id
		WHERE l.id = ?`, ID).Scan(&lecturer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if lecturer.ID == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}

	c.JSON(http.StatusOK, lecturer)
}

func GetLecturerByMajorID(c *gin.Context) {
	majorID := c.Param("id")
	var lecturer []LecturerResponse

	db := config.DB()

	if err := db.Raw(`
		SELECT 
			l.*,
			m.major_name AS major_name, 
			m.faculty_id, 
			f.faculty_name AS faculty_name,
			st.status_staff AS status_staff_name
		FROM lecturers l
		LEFT JOIN majors m ON l.major_id = m.id
		LEFT JOIN faculties f ON m.faculty_id = f.id
		LEFT JOIN status_staffs st ON l.status_staff_id = st.id
		WHERE l.major_id = ?`, majorID).Scan(&lecturer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if len(lecturer) == 0 {
		c.JSON(http.StatusNoContent, gin.H{"message": "ไม่มีข้อมูลนักศึกษาสำหรับ MajorID ที่ระบุ"})
		return
	}

	var count int64
	countQuery := `SELECT COUNT(*) FROM lecturers WHERE major_id = ?`
	if err := db.Raw(countQuery, majorID).Scan(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถนับจำนวนอาจารย์ได้"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       lecturer,
		"totalCount": count,
	})
}

func DeleteLecturer(c *gin.Context) {
	LecturerID := c.Param("id")

	db := config.DB()
	tx := db.Begin()

	// ลบข้อมูลใน LecturerThesis แบบ Hard Delete
	if err := tx.Unscoped().Where("lecturer_id = ?", LecturerID).Delete(&entity.LecturerThesis{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถลบข้อมูลในตาราง LecturerThesis ได้"})
		return
	}

	// ลบข้อมูลใน LecturerEducation แบบ Hard Delete
	if err := tx.Unscoped().Where("lecturer_id = ?", LecturerID).Delete(&entity.LecturerEducation{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถลบข้อมูลในตาราง LecturerEducation ได้"})
		return
	}

	// ลบข้อมูลใน Lecturer แบบ Hard Delete
	if err := tx.Unscoped().Where("id = ?", LecturerID).Delete(&entity.Lecturer{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถลบข้อมูลอาจารย์ได้"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "ลบข้อมูลอาจารย์สำเร็จ"})
}


type (
	Authen struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func SignInLecturer(c *gin.Context) {
	var payload Authen
	var user entity.Lecturer

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา user ด้วย email
	if err := config.DB().Raw("SELECT * FROM lecturers WHERE email = ?", payload.Email).Scan(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "lecturer not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "password is incerrect"})

		return

	}

	// สร้าง JWT token
	jwtWrapper := services.JwtWrapper{
		SecretKey:       "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error signing token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token_type": "Bearer", "token": signedToken, "id": user.ID, "profile": user.Profile})
}
