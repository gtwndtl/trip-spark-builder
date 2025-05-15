package pdf

import (
	"bytes"

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

type UpdateContentRequest3 struct {
	Contents [][]string `json:"contents"`
}

func PatchPDF3(c *gin.Context) {
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

func GenerateUpdatedPDF3(contents [][]string) (bytes.Buffer, error) {

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

func GeneratePDF3() (bytes.Buffer, error){
	
	
	// darkGrayColor := getDarkGrayColor()

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	

	m.AddUTF8Font("THSarabun", consts.Normal, "./font/THSarabun.ttf")
	m.AddUTF8Font("THSarabun", consts.Italic, "./font/THSarabun Italic.ttf")
	m.AddUTF8Font("THSarabun", consts.Bold, "./font/THSarabun Bold.ttf")
	m.AddUTF8Font("THSarabun", consts.BoldItalic, "./font/THSarabun Bold Italic.ttf")
	m.SetDefaultFontFamily("THSarabun")
	m.SetPageMargins(10, 15, 10)



	// Form Header
	m.Row(15, func() {

		m.Col(1, func() { // Add space for the logo
			_ = m.FileImage("sut.jpg", props.Rect{
				Center:  false,
				Percent: 150, // Adjust size as necessary
			})
		})

		m.Col(12, func() {
			m.Text("คำร้องขอชำระเงินค่าลงทะเบียนเรียนล่าช้า                                                                              ท.1", props.Text{
				Top:   0,
				Size:  16,
				Style: consts.Bold,
				Align: consts.Left,
			})
			m.Text("Request for late payment of registration fee", props.Text{
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

			m.Text("inputName", props.Text{
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
			m.Text("inputStudentID", props.Text{
				Top:   7.5,
				Size:  11,
				Align: consts.Left,
				Left:  157,
				Style: consts.Bold, // Bold inputName text
				Color: getBlueColor(),
			})
			m.Text("สังกัดสำนักวิชา the Institute of  ", props.Text{
				Top:   13.5,
				Size:  11,
				Align: consts.Left,
				Left:  10, 
			})
			m.Text(fmt.Sprintf("%v","Faculty"), props.Text{
				Top:   13.5,
				Size:  11,
				Align: consts.Left,
				Left:  50,
				Style: consts.Bold, // Bold inputName text
				Color: getBlueColor(),
			})
			m.Text("สาขาวิชา / หลักสูตร School of  ", props.Text{
				Top:   13.5,
				Size:  11,
				Align: consts.Left,
				Left:  100, 
			})
			m.Text("Major", props.Text{
				Top:   13.5,
				Size:  11,
				Align: consts.Left,
				Left:  140,
				Style: consts.Bold, // Bold inputName text
				Color: getBlueColor(),
			})

			m.Text("มีความประสงค์ขอชำระเงินค่าธรรมเนียมการลงทะเบียนเรียน ภาคการศึกษาที่  ", props.Text{
				Top:   19.5,
				Size:  11,
				Align: consts.Left,
				Left:  10, 
			})

			m.Text("2 / ", props.Text{
				Top:   19.5,
				Size:  11,
				Align: consts.Left,
				Left:  90,
				Style: consts.Bold, // Bold inputName text
				Color: getBlueColor(),
			})

			m.Text("2568", props.Text{
				Top:   19.5,
				Size:  11,
				Align: consts.Left,
				Left:  95,
				Style: consts.Bold, // Bold inputName text
				Color: getBlueColor(),
			})

			m.Text("ล่าช้า", props.Text{
				Top:   19.5,
				Size:  11,
				Align: consts.Left,
				Left:  105, 
			})

			m.Text("เนื่องจาก Because", props.Text{
				Top:   25.5,
				Size:  11,
				Align: consts.Left,
				Left:  10, 
			})

			m.Text("เหตุผลลลลลลล", props.Text{
				Top:   25.5,
				Size:  11,
				Align: consts.Left,
				Left:  35,
				Style: consts.Bold, // Bold inputName text
				Color: getBlueColor(),
			})


			m.Text("ทั้งนี้ ข้าพเจ้าสามารถชำระเงินได้ในวันที่ ", props.Text{
				Top:   31.5,
				Size:  11,
				Align: consts.Left,
				Left:  10, 
			})

			m.Text("2025-12-02", props.Text{
				Top:   31.5,
				Size:  11,
				Align: consts.Left,
				Left:  55,
				Style: consts.Bold, // Bold inputName text
				Color: getBlueColor(),
			})

			m.Text("หรือสามารถชำระเงินได้ทันทีที่ได้รับการอนุมัติ ", props.Text{
				Top:   31.5,
				Size:  11,
				Align: consts.Left,
				Left:  75, 
			})

			m.Text("จึงเรียนมาเพื่อโปรดพิจารณา For your consideration ", props.Text{
				Top:   37.5,
				Size:  11,
				Align: consts.Left,
				Left:  25, 
			})

			m.Text("นักศึกษาลงชื่อ Signature", props.Text{
				Top:   43.5,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Left:  100, 
			})
			m.Text("Sirion srimueang", props.Text{
				Top:   43.5,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Color: getBlueColor(),
				Left:  130, 
			})
			m.Text("โทรศัพท์ Tel. No.", props.Text{
				Top:   49.5,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Left:  100, 
			})
			m.Text("0651018312", props.Text{
				Top:   49.5,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Color: getBlueColor(),
				Left:  128, 
			})

			m.Text("วันที่ Date :", props.Text{
				Top:   49.5,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Left:  148, 
			})

			m.Text(fmt.Sprintf(" %s", "2025-12-30"), props.Text{
				Top:   49.5,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Color: getBlueColor(),
				Left:  170, 
			})

			
		})

		

		
		
	




	})

	return m.Output()

}


func getDarkGrayColor3() color.Color {
	return color.Color{
		Red:   55,
		Green: 55,
		Blue:  55,
	}
}

func getBlueColor3() color.Color {
	return color.Color{
		Red:   10,
		Green: 10,
		Blue:  150,
	}
}

func getRedColor3() color.Color {
	return color.Color{
		Red:   150,
		Green: 10,
		Blue:  10,
	}
}

func CreatePrintStory3(c *gin.Context) {
	// Parse JSON input
	var requestData struct {
		InputName      string 
		InputStudentID string 
		Degree         string 
		Faculty        string 
		Major          string 
		Details        string 
		CourseCode       string 
		CourseTitle      string 
		Group            string 
		OldGroup         string 
		NewGroup         string 
		SpecifyReason    string 
		InputPhoneNumber string 
		Date             time.Time 
		RequestID uint 
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
	pdfData, err := GeneratePDF3(
		// requestData.InputName,
		// requestData.InputStudentID,
		// requestData.Degree,
		// Faculty.FacultyName,
		// Major.MajorName,
		// requestData.Details,
		// requestData.CourseCode,
		// requestData.CourseTitle,
		// requestData.Group,
		// requestData.OldGroup,
		// requestData.NewGroup,
		// requestData.SpecifyReason,
		// requestData.InputPhoneNumber,
		// requestData.Date,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate PDF"})
		return
	}

	// บันทึกไฟล์ตามชื่อ
	filename := fmt.Sprintf("%s_%s.pdf", requestData.InputStudentID, requestData.InputName)
	savePath := filepath.Join("uploads", filename)

	if _, err := os.Stat("UploadsPdf"); os.IsNotExist(err) {
		if err := os.Mkdir("UploadsPdf", os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create directory"})
			return
		}
	}

	if err := os.WriteFile(savePath, pdfData.Bytes(), 0644); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save PDF"})
		return
	}

	// สุ่ม Code pdf 
	newPrintStoryCode := fmt.Sprintf("R%09d", rand.Intn(1000000000))


	var Request entity.Request
	db.Where("id", requestData.RequestID).First(&Request)

	// บันทึกลง databases
	printStory := entity.PrintStory{
		PrintStoryCode: newPrintStoryCode,
		DocumentPath:   savePath,
		CreateAt:       time.Now(),
		RequestID:      requestData.RequestID,
	}

	if err := config.DB().Create(&printStory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save print story"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "PDF generated and saved successfully", "data": printStory})
}