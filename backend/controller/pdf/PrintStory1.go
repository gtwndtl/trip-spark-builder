// 2024-12-28 22.08 น.
// ดึง Request ID
package pdf

import (
	"bytes"
	"errors"

	"strings"

	"fmt"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
)

type UpdateContentRequest struct {
	Contents [][]string `json:"contents"`
}

func PatchPDF(c *gin.Context) {
	var request UpdateContentRequest

	// ตรวจสอบข้อมูลที่เข้ามาใน Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// สร้าง PDF ใหม่พร้อมเนื้อหาที่อัพเดต
	newContents := request.Contents
	pdfBuffer, err := GenerateUpdatedPDF(newContents)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating PDF"})
		return
	}

	// ส่ง PDF ที่อัปเดตแล้วกลับไป
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=updated_invoice.pdf")
	c.Data(http.StatusOK, "application/pdf", pdfBuffer.Bytes())
}

func GenerateUpdatedPDF(contents [][]string) (bytes.Buffer, error) {

	// ใช้ฟังก์ชัน GeneratePDF เดิม แต่แทนที่เนื้อหา
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 15, 10)

	// ส่วนของ Header และ Footer
	m.RegisterHeader(func() {
		m.Col(3, func() {
			_ = m.FileImage("sut.png", props.Rect{
				Center:  true,
				Percent: 80,
			})
		})
		m.ColSpace(6)

		m.Row(20, func() {
			m.Col(12, func() {
				m.Text("Updated Invoice", props.Text{
					Size:  16,
					Style: consts.Bold,
					Align: consts.Center,
				})
			})
		})
	})
	m.RegisterFooter(func() {
		m.Row(10, func() {
			m.Col(12, func() {
				m.Text("Updated by API", props.Text{
					Top:   3,
					Style: consts.Italic,
					Size:  10,
					Align: consts.Center,
				})
			})
		})
	})

	return m.Output()
}

func GeneratePDF(inputName, inputStudentID, Degree, Faculty, Major, Details, CourseCode, CourseTitle, Group, OldGroup, NewGroup, SpecifyReason, inputPhoneNumber string, Date time.Time) (bytes.Buffer, error) {
	
	var output bytes.Buffer
	//ตรวจสอบความยาวของ input ที่ส่งมา

	if len(inputName) == 0 { 
		return output, errors.New("InputName is required")
	} 

	if len(inputStudentID) == 0 { 
		return output, errors.New("InputStudentID is required")
	} 

	if len(Degree) == 0 { 
		return output, errors.New("Degree is required")
	} 

	if len(Faculty) == 0 { 
		return output, errors.New("Faculty is required")
	} 
	
	if len(Major) == 0 { 
		return output, errors.New("Major is required")
	} 

	if len(CourseCode) == 0 { 
		return output, errors.New("CourseCode is required")
	} 

	if len(CourseTitle) == 0 { 
		return output, errors.New("CourseTitle is required")
	} 

	if len(Group) == 0 { 
		return output, errors.New("Group is required")
	} 
	
	if (Details) == "เปลี่ยนกลุ่มวิชา Reduce courses" {
		if len(OldGroup) == 0 { 
			return output, errors.New("OldGroup is required")
		} 

		if len(NewGroup) == 0 { 
			return output, errors.New("NewGroup is required")
		} 

		if len(SpecifyReason) == 0 { 
			return output, errors.New("SpecifyReason is required")
		} 
	}


	if len(inputPhoneNumber) == 0 { 
		return output, errors.New("InputPhoneNumber is required")
	} 
	
	
	
	darkGrayColor := getDarkGrayColor()

	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	m.AddUTF8Font("THSarabun", consts.Normal, "./font/THSarabun.ttf")
	m.AddUTF8Font("THSarabun", consts.Italic, "./font/THSarabun Italic.ttf")
	m.AddUTF8Font("THSarabun", consts.Bold, "./font/THSarabun Bold.ttf")
	m.AddUTF8Font("THSarabun", consts.BoldItalic, "./font/THSarabun Bold Italic.ttf")
	m.SetDefaultFontFamily("THSarabun")
	m.SetPageMargins(10, 15, 10)

	if Date.IsZero() {
        Date = time.Now()
    }

	formattedDate := Date.Format("2006-01-02")

	// Form Header
	m.Row(15, func() {

		m.Col(1, func() { // Add space for the logo
			_ = m.FileImage("sut.jpg", props.Rect{
				Center:  false,
				Percent: 150, // Adjust size as necessary
			})
		})

		m.Col(12, func() {
			m.Text("คำร้องขอลงทะเบียนเพิ่ม / เปลี่ยนกลุ่ม / ลดรายวิชา                                                                    ท.1", props.Text{
				Top:   0,
				Size:  16,
				Style: consts.Bold,
				Align: consts.Left,
			})
			m.Text("Request to Register for Additional Credits / to Change Study Group / to Reduce Courses", props.Text{
				Top:   6,
				Size:  16,
				Style: consts.Bold,
				Align: consts.Left,
			})
			m.Row(8, func() {})
		})
	})

	// เส้นคั้นระหว่างส่วนหัวกับรายละเอียดคำร้อง
	m.Line(1, props.Line{

		Width: 0.5, // ความหนาของเส้น 2
	})
	m.Line(1, props.Line{

		Width: 0.5, // ความหนาของเส้น 2
	})

	m.Row(135, func() {
		m.Col(12, func() {
			m.Text("เรียน   อาจารย์ผู้สอน / อาจารย์ผู้รับผิดชอบวิชา    Dear Instructor / Course Coordinator", props.Text{
				Top:   1.5,
				Size:  11,
				Align: consts.Left,
				Left:  3,
				Style: consts.Bold,
			})
			m.Text("ข้าพเจ้า ( นาย / นาง / นางสาว ) I am (Mr. / Mrs. / Miss)", props.Text{
				Top:   7.5,
				Size:  11,
				Align: consts.Left,
				Left:  10, 
			})

			m.Text(inputName, props.Text{
				Top:   7.5,
				Size:  11,
				Align: consts.Left,
				Left:  80,
				Style: consts.Bold, // Bold inputName text
				Color: getBlueColor(),
			})
			m.Text("เลขประจำตัว Student code ", props.Text{
				Top:   7.5,
				Size:  11,
				Align: consts.Left,
				Left:  123, 
			})
			m.Text(inputStudentID, props.Text{
				Top:   7.5,
				Size:  11,
				Align: consts.Left,
				Left:  157,
				Style: consts.Bold, // Bold inputName text
				Color: getBlueColor(),
			})
			m.Text("เป็นนักศึกษาระดับ a student at ", props.Text{
				Top:   13.5,
				Size:  11,
				Align: consts.Left,
				Left:  3, 
			})
			m.Text(Degree, props.Text{
				Top:   13.5,
				Size:  11,
				Align: consts.Left,
				Left:  42,
				Style: consts.Bold, // Bold inputName text
				Color: getBlueColor(),
			})
			m.Text("สังกัดสำนักวิชา the Institute of  ", props.Text{
				Top:   13.5,
				Size:  11,
				Align: consts.Left,
				Left:  61, 
			})
			m.Text(fmt.Sprintf("%v",Faculty), props.Text{
				Top:   13.5,
				Size:  11,
				Align: consts.Left,
				Left:  99,
				Style: consts.Bold, // Bold inputName text
				Color: getBlueColor(),
			})
			m.Text("สาขาวิชา / หลักสูตร School of  ", props.Text{
				Top:   19.5,
				Size:  11,
				Align: consts.Left,
				Left:  3, 
			})
			m.Text(Major, props.Text{
				Top:   19.5,
				Size:  11,
				Align: consts.Left,
				Left:  43,
				Style: consts.Bold, // Bold inputName text
				Color: getBlueColor(),
			})
			m.Text("มีความประสงค์จะลงทะเบียน wish to register : ", props.Text{
				Top:   25.5,
				Size:  11,
				Align: consts.Left,
				Left:  3, 
			})

			// Add some space before the "Transactions" section
			m.Row(33, func() {})

			m.SetBackgroundColor(darkGrayColor)

			//เส้นคั้นตารางตรงกลาง
			m.Text("|", props.Text{
				Top:   5.4,
				Style: consts.Bold,
				Left:  110.7,
				Color: darkGrayColor,
			})
			m.Text("|", props.Text{
				Top:   7.8,
				Style: consts.Bold,
				Left:  110.7,
				Color: darkGrayColor,
			})
			m.Text("|", props.Text{
				Top:   10.2,
				Style: consts.Bold,
				Left:  110.7,
				Color: darkGrayColor,
			})
			m.Text("|", props.Text{
				Top:   12.6,
				Style: consts.Bold,
				Left:  110.7,
				Color: darkGrayColor,
			})
			m.Text("|", props.Text{
				Top:   15,
				Style: consts.Bold,
				Left:  110.7,
				Color: darkGrayColor,
			})
			m.Text("|", props.Text{
				Top:   17.4,
				Style: consts.Bold,
				Left:  110.7,
				Color: darkGrayColor,
			})
			m.Text("|", props.Text{
				Top:   19.8,
				Style: consts.Bold,
				Left:  110.7,
				Color: darkGrayColor,
			})
			m.Text("|", props.Text{
				Top:   22.2,
				Style: consts.Bold,
				Left:  110.7,
				Color: darkGrayColor,
			})
			m.Text("|", props.Text{
				Top:   24.6,
				Style: consts.Bold,
				Left:  110.7,
				Color: darkGrayColor,
			})
			m.Text("|", props.Text{
				Top:   27,
				Style: consts.Bold,
				Left:  110.7,
				Color: darkGrayColor,
			})
			m.Text("|", props.Text{
				Top:   29.4,
				Style: consts.Bold,
				Left:  110.7,
				Color: darkGrayColor,
			})
			m.Text("|", props.Text{
				Top:   31.8,
				Style: consts.Bold,
				Left:  110.7,
				Color: darkGrayColor,
			})


			m.Row(7, func() {
				m.Col(7, func() {
					m.Text("1. รายการ Details", props.Text{
						Top:   0.5,
						Size:  11,
						Style: consts.Bold,
						Align: consts.Center,
						Color: color.NewWhite(),
					})
				})
				m.Col(5, func() {
					m.Text("2. ผลการพิจารณา Decision Made", props.Text{
						Top:   0.5,
						Size:  11,
						Style: consts.Bold,
						Align: consts.Center,
						Color: color.NewWhite(),
					})

				})
				// m.ColSpace(3)

			})
			// Remove background color before this section
			m.SetBackgroundColor(color.NewWhite()) // Set background back to white
			m.Row(1, func() {})

			m.Col(13, func() {
				m.Text("⬛ Approved", props.Text{
					Top:   1,
					Size:  11,
					Align: consts.Left,
					Style: consts.Bold,
					Left:  119, // Adjust positioning as needed
				})
			})
			m.Row(0.5, func() {})
			m.Col(13, func() {
				m.Text("⬛ NotApproved", props.Text{
					Top:   1,
					Size:  11,
					Align: consts.Left,
					Style: consts.Bold,
					Left:  149, // Adjust positioning as needed
				})
			})

			m.Row(0.5, func() {})

			m.Text(Details, props.Text{
				Top:   1.5,
				Size:  11,
				Align: consts.Left,
				Left:  3,
				Style: consts.Bold, // Bold inputName text
				Color: getBlueColor(),
			})

			m.Text("รหัสวิชา Course Code  ", props.Text{
				Top:   1.5,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Left:  44, 
			})
			m.Text(CourseCode, props.Text{
				Top:   1.5,
				Size:  11,
				Align: consts.Left,
				Left:  74,
				Style: consts.Bold, // Bold inputName text
				Color: getBlueColor(),
			})

			m.Row(3, func() {})
			m.Text("ชื่อวิชา(ภาษาอังกฤษ) Course Title", props.Text{
				Top:   20.5,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Left:  3, 
			})
			m.Text(CourseTitle, props.Text{
				Top:   20.5,
				Size:  11,
				Align: consts.Left,
				Left:  46,
				Style: consts.Bold, // Bold inputName text
				Color: getBlueColor(),
			})

			m.Text("เหตุผล (กรณีไม่อนุญาต) :", props.Text{
				Top:   20.5,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Left:  119, 
			})

			//ส่วนของอาจารย์
			m.Text("  ", props.Text{ //เหตุผล ของอาจารย์ กรอกในส่วนของอาจารย์
				Top:   20.5,
				Size:  11,
				Align: consts.Left,
				Left:  149,
				Style: consts.Bold, // Bold inputName text
				Color: getBlueColor(),
			})

			m.Row(4.5, func() {})
			m.Text("กลุ่ม Group No.", props.Text{
				Top:   27.5,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Left:  3, 
			})
			m.Text(Group, props.Text{
				Top:   27.5,
				Size:  11,
				Align: consts.Left,
				Left:  24,
				Style: consts.Bold, // Bold inputName text
				Color: getBlueColor(),
			})

			// ส่วนของอาจารย์
			m.Text("          ", props.Text{ //ชื่ออาจารย์ ของอาจารย์ กรอกในส่วนของอาจารย์
				Top:   27.5,
				Size:  11,
				Align: consts.Left,
				Left:  135,
				Style: consts.Bold, // Bold inputName text
				Color: getBlueColor(),
			})
			m.Row(5, func() {})

			m.Text("กรณีเปลี่ยนกลุ่ม In the case of changing study group ", props.Text{
				Top:   27.5,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Color: getRedColor(),
				Left:  3, 
			})

			m.Text("อาจารย์ผู้สอน / อาจารย์ผู้รับผิดชอบวิชา", props.Text{
				Top:   27.5,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Left:  125, 
			})

			m.Row(5.5, func() {})
			m.Text("กลุ่มเดิมคือกลุ่ม the old group no. is", props.Text{
				Top:   27.5,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Left:  3, 
			})
			m.Text(OldGroup, props.Text{
				Top:   27.5,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Color: getBlueColor(),
				Left:  50, 
			})

			// m.Row(5, func() {})
			m.Text("กลุ่มใหม่คือกลุ่ม the new group no. is", props.Text{
				Top:   27.5,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Left:  55, 
			})
			m.Text(NewGroup, props.Text{
				Top:   27.5,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Color: getBlueColor(),
				Left:  103, 
			})

			m.Text("Instructor / Course coordinator", props.Text{
				Top:   27.5,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Left:  127, 
			})
			m.Row(5.5, func() {})
			m.Text("ระบุเหตุผล Specify reason ", props.Text{
				Top:   27.5,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Left:  3, 
			})
			m.Text(SpecifyReason, props.Text{
				Top:   27.5,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Color: getBlueColor(),
				Left:  38, 
			})
			m.Row(0.6, func() {})
			m.Text("|", props.Text{
				Top:   27.5,
				Style: consts.Bold,
				Left:  110.7,
				Color: darkGrayColor,
			})
			m.Row(1.59, func() {})
			m.Text("|", props.Text{
				Top:   25.5,
				Style: consts.Bold,
				Left:  110.7,
				Color: darkGrayColor,
			})
			m.Row(2, func() {})
			m.Text("|", props.Text{
				Top:   25.5,
				Style: consts.Bold,
				Left:  110.7,
				Color: darkGrayColor,
			})
			m.Row(2.2, func() {})
			m.Text("|", props.Text{
				Top:   25.5,
				Style: consts.Bold,
				Left:  110.7,
				Color: darkGrayColor,
			})
			m.Row(2.4, func() {})
			m.Text("|", props.Text{
				Top:   25.5,
				Style: consts.Bold,
				Left:  110.7,
				Color: darkGrayColor,
			})
			m.Row(2.5, func() {})
			m.Text("|", props.Text{
				Top:   25.5,
				Style: consts.Bold,
				Left:  110.7,
				Color: darkGrayColor,
			})
			m.Row(2.5, func() {})
			m.Text("|", props.Text{
				Top:   25.5,
				Style: consts.Bold,
				Left:  110.7,
				Color: darkGrayColor,
			})

			m.Row(6.5, func() {})
			m.Line(1, props.Line{

				Width: 0.1, // ความหนาของเส้น 2
			})

			m.Row(1, func() {})
			m.Text("จึงเรียนมาเพื่อโปรดพิจารณา For your consideration.", props.Text{
				Top:   27.5,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Left:  23, 
			})
			m.Row(3, func() {})
			m.Text("นักศึกษาลงชื่อ Signature", props.Text{
				Top:   27.5,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Left:  49, 
			})
			m.Text(inputName, props.Text{
				Top:   27.5,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Color: getBlueColor(),
				Left:  80, 
			})
			m.Text("โทรศัพท์ Tel. No.", props.Text{
				Top:   27.5,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Left:  115, 
			})
			m.Text(inputPhoneNumber, props.Text{
				Top:   27.5,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Color: getBlueColor(),
				Left:  138, 
			})

			m.Text("วันที่ Date :", props.Text{
				Top:   27.5,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Left:  158, 
			})

			m.Text(fmt.Sprintf(" %s", formattedDate), props.Text{
				Top:   27.5,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Color: getBlueColor(),
				Left:  174, 
			})

		})

	})

	return m.Output()
}


func getDarkGrayColor() color.Color {
	return color.Color{
		Red:   55,
		Green: 55,
		Blue:  55,
	}
}

func getBlueColor() color.Color {
	return color.Color{
		Red:   10,
		Green: 10,
		Blue:  150,
	}
}

func getRedColor() color.Color {
	return color.Color{
		Red:   150,
		Green: 10,
		Blue:  10,
	}
}

func CreatePrintStory(c *gin.Context) {
	// Parse JSON input
	var requestData struct {
		InputName      string `json:"inputName" valid:"required~InputName is required"`
		InputStudentID string `json:"inputStudentID" valid:"required~InputStudentID is required, matches(^[B]\\d{7}$)~InputStudentID is invalid"`
		Degree         string `json:"degree" valid:"required~Degree is required"`
		Faculty        string `json:"faculty" valid:"required~Faculty is required"`
		Major          string `json:"major" valid:"required~Major is required"`
		Details        string `json:"details" valid:"required~Details is required"`
		CourseCode       string `json:"courseCode" valid:"required~CourseCode is required"`
		CourseTitle      string `json:"courseTitle" valid:"required~CourseTitle is required"`
		Group            string `json:"group" valid:"required~Group is required"`
		OldGroup         string `json:"oldGroup" valid:"required~OldGroup is required"`
		NewGroup         string `json:"newGroup" valid:"required~NewGroup is required"`
		SpecifyReason    string `json:"specifyReason" valid:"required~SpecifyReason is required"`
		InputPhoneNumber string `json:"inputPhoneNumber" valid:"required~InputPhoneNumber is required"`
		Date             time.Time `json:"date"`
		RequestID uint `json:"requestID" valid:"-"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	var Faculty entity.Faculty
	db := config.DB()
	db.Where("id", requestData.Faculty).First(&Faculty)

	var Major entity.Major
	db.Where("id", requestData.Major).First(&Major)


	// Generate PDF
	pdfData, err := GeneratePDF(
		requestData.InputName,
		requestData.InputStudentID,
		requestData.Degree,
		Faculty.FacultyName,
		Major.MajorName,
		requestData.Details,
		requestData.CourseCode,
		requestData.CourseTitle,
		requestData.Group,
		requestData.OldGroup,
		requestData.NewGroup,
		requestData.SpecifyReason,
		requestData.InputPhoneNumber,
		requestData.Date,
	)

	if err != nil {
		// log.Printf("GeneratePDF error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate PDF"})
		return
	}

	// บันทึกไฟล์ตามชื่อ
	filename := fmt.Sprintf("%s_%s.pdf", requestData.InputStudentID, requestData.InputName)
	savePath := filepath.Join("uploads_pdf", filename)

	if _, err := os.Stat("uploads_pdf"); os.IsNotExist(err) {
		if err := os.Mkdir("uploads_pdf", os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create directory"})
			return
		}
	}

	// if err := os.WriteFile(savePath, pdfData.Bytes(), 0644); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save PDF"})
	// 	return
	// }

	// สุ่ม Code pdf 
	newPrintStoryCode := fmt.Sprintf("R%09d", rand.Intn(1000000000))


	var request entity.Request
	db.Where("id", requestData.RequestID).First(&request)

	var course entity.Course
	db.Where("LOWER(course_code) = ?", strings.ToLower(requestData.CourseCode)).First(&course)


	// บันทึกลง databases
	printStory := entity.PrintStory{
		PrintStoryCode: newPrintStoryCode,
		DocumentPath:   savePath,
		DocumentFile: pdfData.Bytes(),
		CreateAt:       time.Now(),
		RequestID:      requestData.RequestID,
	}

	request.CourseID = course.ID
	request.Course = course

	if err := config.DB().Updates(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update course"})
		return
	}

	if err := config.DB().Create(&printStory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save print story"})
		return
	}


	c.JSON(http.StatusOK, gin.H{"message": "PDF generated and saved successfully", "data": printStory})
}



// GET /printstory
func GetPrintStory(c *gin.Context) {
    var printStories []entity.PrintStory 

    db := config.DB()

	
    // ใช้ Preload เพื่อดึงข้อมูลที่เกี่ยวข้องกับ Request
    result := db.Preload("Request").First(&printStories)

    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, printStories)
}




// GET /printstory/:request_id
func GetPrintStoryByRequestID(c *gin.Context) {
    requestID := c.Param("request_id") // รับค่า request_id จาก URL
    var printStories []entity.PrintStory

    db := config.DB()

    // Query to join tables and filter by request_id
    result := db.
        Joins("JOIN requests ON requests.id = print_stories.request_id").
        Where("requests.id = ?", requestID).
        Preload("Request").  // Preload the associated Request data
        Find(&printStories)

    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
        return
    }

    // Return the print stories in the response
    c.JSON(http.StatusOK, gin.H{"printStories": printStories})
}




func GetPDFFile(c *gin.Context) {
    // รับชื่อไฟล์จาก query parameter
    filename := c.Param("filename")
	db := config.DB()

	// src := strings.NewReader("GfG\n") 

    // ระบุ path ของโฟลเดอร์ที่เก็บไฟล์ PDF
    filePath := filepath.Join("uploads_pdf", filename)

	
	var printstory entity.PrintStory
	db.Where("document_path = ?", filePath).First(&printstory)


    // // ตรวจสอบว่าไฟล์มีอยู่จริงหรือไม่
    // if _, err := os.Stat(filePath); os.IsNotExist(err) {
    //     c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
    //     return
    // }

	

	// FileContentType := http.DetectContentType(printstory.DocumentFile) //Get file header

	FileStat := printstory.DocumentFile
	FileSize := len(FileStat)

	Filename := printstory.DocumentPath

	//Set the headers
	// w.Header().Set("Content-Type", FileContentType+";"+Filename)
	// w.Header().Set("Content-Length", string(FileSize))
	

	// io.CopyBuffer(w, src, FileStat)

	c.Header("Content-Disposition", "attachment; filename="+Filename)
	c.Header("Content-Type", "application/text/plain")
	c.Header("Accept-Length", fmt.Sprintf("%d", FileSize))
	c.Writer.Write([]byte(FileStat))
	c.JSON(http.StatusOK, gin.H{
		"msg": "Download file successfully",
	})


    // // ให้บริการไฟล์ PDF
    // c.File(filePath)
}


