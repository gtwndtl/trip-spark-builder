package student

import (
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"

	"github.com/sut67/team09/services"

	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type StudentResponse struct {
	ID              uint   `json:"ID"`
	StudentCode     string `json:"StudentCode"`
	FirstName       string `json:"FirstName"`
	LastName        string `json:"LastName"`
	Email           string `json:"Email"`
	NationalID      string `json:"NationalID"`
	Phone           string `json:"Phone"`
	Profile         string `json:"Profile"`
	BirthDay        string `json:"BirthDay"`
	YearOfStudy     int    `json:"YearOfStudy"`
	GenderID        uint   `json:"GenderID"`
	GenderName      string `json:"GenderName"`
	MajorID         uint   `json:"MajorID"`
	MajorName       string `json:"MajorName"`
	FacultyID       uint   `json:"FacultyID"`
	FacultyName     string `json:"FacultyName"`
	StatusStaffID   uint   `json:"StatusStaffID"`
	StatusStaffName string `json:"StatusStaffName"`
	SemesterID      uint   `json:"SemesterID"`
	Term            int    `json:"Term"`
}

func CreateStudents(c *gin.Context) {
	var studentData entity.Student

	if err := c.ShouldBindJSON(&studentData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ข้อมูลไม่ถูกต้อง",
		})
		return
	}

	db := config.DB()

	var genderID, majorID, statusStaffID, semesterID uint
	statusStaffID = 2

	if studentData.GenderID != nil {
		genderID = uint(*studentData.GenderID)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ไม่พบข้อมูลเพศ",
		})
		return
	}

	if studentData.MajorID != nil {
		majorID = uint(*studentData.MajorID)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ไม่พบข้อมูลสาขา",
		})
		return
	}

	if studentData.SemesterID != 0 {
		semesterID = studentData.SemesterID
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ไม่พบข้อมูลปีการศึกษา",
		})
		return
	}

	tx := db.Begin()

	hashedPassword, err := config.HashPassword(studentData.Password)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "ไม่สามารถเข้ารหัสรหัสผ่านได้",
		})
		return
	}

	student := entity.Student{
		StudentCode:   studentData.StudentCode,
		FirstName:     studentData.FirstName,
		LastName:      studentData.LastName,
		BirthDay:      studentData.BirthDay,
		Email:         studentData.Email,
		Password:      hashedPassword,
		NationalID:    studentData.NationalID,
		Phone:         studentData.Phone,
		Profile:       studentData.Profile,
		YearOfStudy:   studentData.YearOfStudy,
		GenderID:      &genderID,
		MajorID:       &majorID,
		StatusStaffID: &statusStaffID,
		SemesterID:    semesterID,
	}

	if err := tx.Create(&student).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ไม่สามารถเพิ่มข้อมูลนักศึกษาได้",
		})
		return
	}

	tx.Commit()

	c.JSON(http.StatusCreated, gin.H{
		"message": "ข้อมูลนักศึกษาถูกเพิ่มเรียบร้อยแล้ว",
		"data":    student,
	})
}

func UpdateStudent(c *gin.Context) {
	var student entity.Student
	StudentID := c.Param("id")

	db := config.DB()
	result := db.First(&student, StudentID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบข้อมูลนักศึกษา"})
		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	result = db.Save(&student)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "แก้ไขข้อมูลไม่สำเร็จ"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "แก้ไขข้อมูลนักศึกษาสำเร็จ"})
}

func CheckEmailStudent(c *gin.Context) {
	Email := c.Param("email")
	db := config.DB()

	var student entity.Student
	studentResult := db.Where("email = ?", Email).First(&student)

	if studentResult.Error != nil && studentResult.Error != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error": studentResult.Error.Error()})
		return
	}

	if studentResult.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"isValid": false,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"isValid": true,
		})
	}
}

func CheckEmailStudentByID(c *gin.Context) {
	Email := c.Param("email")
	id := c.Param("id")

	db := config.DB()

	var student entity.Student
	studentResult := db.Where("email = ? AND id != ?", Email, id).First(&student)

	if studentResult.Error != nil && studentResult.Error != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error": studentResult.Error.Error()})
		return
	}

	if studentResult.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"isValid": false,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"isValid": true,
		})
	}
}

func CheckStudentCode(c *gin.Context) {
	studentcode := c.Param("studentcode")
	db := config.DB()

	var student entity.Student

	studentResult := db.Where("student_code = ?", studentcode).First(&student)

	if (studentResult.Error != nil && studentResult.Error != gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "เกิดข้อผิดพลาดในการตรวจสอบรหัสนักศึกษา"})
		return
	}

	if studentResult.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"isValid": false,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"isValid": true,
		})
	}
}

func CheckStudentCodeByID(c *gin.Context) {
	studentcode := c.Param("studentcode")
	id := c.Param("id")
	db := config.DB()

	var student entity.Student

	studentResult := db.Where("student_code = ? AND id != ?", studentcode, id).First(&student)

	if (studentResult.Error != nil && studentResult.Error != gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "เกิดข้อผิดพลาดในการตรวจสอบรหัสนักศึกษา"})
		return
	}

	if studentResult.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"isValid": false,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"isValid": true,
		})
	}
}


func GetStudentByID(c *gin.Context) {
	ID := c.Param("id")
	var student StudentResponse

	db := config.DB()

	if err := db.Raw(`
		SELECT 
			s.*,
			g.gender AS gender_name, 
			m.major_name AS major_name, 
			m.faculty_id, 
			f.faculty_name AS faculty_name, 
			st.status_staff AS status_staff_name,
			sem.term as term
		FROM students s
		LEFT JOIN genders g ON s.gender_id = g.id
		LEFT JOIN majors m ON s.major_id = m.id
		LEFT JOIN faculties f ON m.faculty_id = f.id
		LEFT JOIN status_staffs st ON s.status_staff_id = st.id
		LEFT JOIN semesters sem ON s.semester_id = sem.id
		WHERE s.id = ?`, ID).Scan(&student).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if student.ID == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}

	c.JSON(http.StatusOK, student)
}

func GetStudentByMajorID(c *gin.Context) {
	majorID := c.Param("id")
	var students []StudentResponse

	db := config.DB()

	if err := db.Raw(`
		SELECT 
			s.*,
			m.major_name AS major_name, 
			m.faculty_id, 
			f.faculty_name AS faculty_name,
			st.status_staff AS status_staff_name
		FROM students s
		LEFT JOIN majors m ON s.major_id = m.id
		LEFT JOIN faculties f ON m.faculty_id = f.id
		LEFT JOIN status_staffs st ON s.status_staff_id = st.id
		WHERE s.major_id = ?`, majorID).Scan(&students).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if len(students) == 0 {
		c.JSON(http.StatusNoContent, gin.H{"message": "ไม่มีข้อมูลนักศึกษาสำหรับ MajorID ที่ระบุ"})
		return
	}

	var count int64
	countQuery := `SELECT COUNT(*) FROM students WHERE major_id = ?`
	if err := db.Raw(countQuery, majorID).Scan(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถนับจำนวนนักศึกษาได้"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       students,
		"totalCount": count,
	})
}

func DeleteStudent(c *gin.Context) {
	StudentID := c.Param("id")

	db := config.DB()
	tx := db.Begin()

	if err := tx.Unscoped().Where("student_id = ?", StudentID).Delete(&entity.StudentEducation{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถลบข้อมูลในตาราง StudentEducation ได้"})
		return
	}

	if err := tx.Unscoped().Where("id = ?", StudentID).Delete(&entity.Student{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถลบข้อมูลนักศึกษาได้"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "ลบข้อมูลนักศึกษาสำเร็จ"})
}



func GetAll(c *gin.Context) {

	var users []entity.Student

	db := config.DB()

	results := db.Preload("Gender").Preload("Semester").Find(&users)

	if results.Error != nil {

		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})

		return

	}

	c.JSON(http.StatusOK, users)

}

type (
	Authen struct {
		Email string `json:"email"`

		Password string `json:"password"`
	}
)

func SignInStudent(c *gin.Context) {

	var payload Authen

	var user entity.Student

	if err := c.ShouldBindJSON(&payload); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	// ค้นหา user ด้วย Username ที่ผู้ใช้กรอกเข้ามา

	if err := config.DB().Raw("SELECT * FROM students WHERE email = ?", payload.Email).Scan(&user).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	// ตรวจสอบรหัสผ่าน

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "password is incerrect"})

		return

	}

	jwtWrapper := services.JwtWrapper{

		SecretKey: "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",

		Issuer: "AuthService",

		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(user.Email)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "error signing token"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"token_type": "Bearer", "token": signedToken, "id": user.ID, "profile": user.Profile})

}
