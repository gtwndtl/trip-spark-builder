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

type UpdateContentRequest2 struct {
	Contents [][]string `json:"contents"`
}

func PatchPDF2(c *gin.Context) {
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

func GenerateUpdatedPDF2(contents [][]string) (bytes.Buffer, error) {

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

func GeneratePDF2(inputName, inputStudentID,  Faculty, Major, EducationSector, datestring, SpecifyReason, inputPhoneNumber string, Date time.Time) (bytes.Buffer, error){
	
	
	darkGrayColor := getDarkGrayColor()

	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	if Date.IsZero() {
        Date = time.Now()
    }

	formattedDate := Date.Format("2006-01-02")

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

			m.Text(EducationSector, props.Text{
				Top:   19.5,
				Size:  11,
				Align: consts.Left,
				Left:  90,
				Style: consts.Bold, // Bold inputName text
				Color: getBlueColor(),
			})

			m.Text("", props.Text{
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

			m.Text(SpecifyReason, props.Text{
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

			m.Text(datestring, props.Text{
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
			m.Text(inputName, props.Text{
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
			m.Text(inputPhoneNumber, props.Text{
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

			m.Text(fmt.Sprintf(" %s", formattedDate), props.Text{
				Top:   49.5,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Color: getBlueColor(),
				Left:  170, 
			})



			m.Row(58, func() {})

			m.SetBackgroundColor(darkGrayColor)
						//เส้นคั้นตารางตรงกลาง
			m.Text("|", props.Text{
				Top:   5.4,
				Style: consts.Bold,
				Left:  78.7,
				Color: darkGrayColor,
			})
			m.Text("|", props.Text{
				Top:   7.8,
				Style: consts.Bold,
				Left:  78.7,
				Color: darkGrayColor,
			})
			m.Text("|", props.Text{
				Top:   10.2,
				Style: consts.Bold,
				Left:  78.7,
				Color: darkGrayColor,
			})
			m.Text("|", props.Text{
				Top:   12.6,
				Style: consts.Bold,
				Left:  78.7,
				Color: darkGrayColor,
			})
			m.Text("|", props.Text{
				Top:   15,
				Style: consts.Bold,
				Left:  78.7,
				Color: darkGrayColor,
			})
			m.Text("|", props.Text{
				Top:   17.4,
				Style: consts.Bold,
				Left:  78.7,
				Color: darkGrayColor,
			})
			m.Text("|", props.Text{
				Top:   19.8,
				Style: consts.Bold,
				Left:  78.7,
				Color: darkGrayColor,
			})
			m.Text("|", props.Text{
				Top:   22.2,
				Style: consts.Bold,
				Left:  78.7,
				Color: darkGrayColor,
			})
			m.Text("|", props.Text{
				Top:   24.6,
				Style: consts.Bold,
				Left:  78.7,
				Color: darkGrayColor,
			})
			m.Text("|", props.Text{
				Top:   27,
				Style: consts.Bold,
				Left:  78.7,
				Color: darkGrayColor,
			})
			m.Text("|", props.Text{
				Top:   29.4,
				Style: consts.Bold,
				Left:  78.7,
				Color: darkGrayColor,
			})
			m.Text("|", props.Text{
				Top:   31.8,
				Style: consts.Bold,
				Left:  78.7,
				Color: darkGrayColor,
			})


			m.Row(7, func() {
				m.Col(5, func() {
					m.Text("1. ความเห็นของเจ้าหน้าที่", props.Text{
						Top:   0.5,
						Size:  11,
						Style: consts.Bold,
						Align: consts.Center,
						Color: color.NewWhite(),
					})
				})
				m.Col(7, func() {
					m.Text("2. ผลการพิจารณาของผู้อำนวยการศูนย์บริการการศึกษา", props.Text{
						Top:   0.5,
						Size:  11,
						Style: consts.Bold,
						Align: consts.Center,
						Color: color.NewWhite(),
					})

				})

			})


			m.SetBackgroundColor(color.NewWhite()) // Set background back to white
			m.Row(1, func() {})

			m.Col(13, func() {
				m.Text("อนุมัติพร้อมค่าปรับ Approved, with payable fineved", props.Text{
					Top:   1,
					Size:  11,
					Align: consts.Left,
					Style: consts.Bold,
					Left:  105, // Adjust positioning as needed
				})


			})

			m.Row(0.5, func() {})

			m.Text("_____________________________________________  ", props.Text{
				Top:   1.5,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Left:  6, 
			})

			m.Row(3.5, func() {})
			m.Text("_____________________________________________  ", props.Text{
				Top:   1.5,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Left:  6, 
			})
			m.Text("ลงชื่อ Signature", props.Text{
				Top:   3.0,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Left:  115, // Adjust positioning as needed
			})
			m.Text("", props.Text{
				Top:   3.0,
				Size:  11,
				Align: consts.Left,
				Left:  135, // Adjust positioning as needed
			})

			m.Row(5.0, func() {})
			m.Text("(ผู้ช่วยศาสตราจารย์ ดร. ประพัฒน์ เป็นตามวา)", props.Text{
				Top:   3.0,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Left:  110, // Adjust positioning as needed
			})

			m.Row(4.5, func() {})
			m.Text("ลงชื่อ Signature", props.Text{
				Top:   3.0,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Left:  6, // Adjust positioning as needed
			})

			m.Text("ผู้อำนวยการศูนย์บริการการศึกษา", props.Text{
				Top:   3.0,
				Size:  11,
				Align: consts.Left,
				Style: consts.Bold,
				Left:  118, // Adjust positioning as needed
			})
			
		})

		
	})

	return m.Output()

}


func getDarkGrayColor2() color.Color {
	return color.Color{
		Red:   55,
		Green: 55,
		Blue:  55,
	}
}

func getBlueColor2() color.Color {
	return color.Color{
		Red:   10,
		Green: 10,
		Blue:  150,
	}
}

func getRedColor2() color.Color {
	return color.Color{
		Red:   150,
		Green: 10,
		Blue:  10,
	}
}

func CreatePrintStory2(c *gin.Context) {
	// Parse JSON input
	var requestData struct {
		InputName      string `json:"inputName"`
		InputStudentID string `json:"InputStudentID"`
		Faculty        string `json:"faculty"`
		Major          string `json:"major"`
		EducationSector string`json:"educationSector"`
		datestring      string `json:"datestring"`
		SpecifyReason    string `json:"specifyReason"`
		InputPhoneNumber string `json:"inputPhoneNumber"`
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
	pdfData, err := GeneratePDF2(
		requestData.InputName,
		requestData.InputStudentID,
		Faculty.FacultyName,
		Major.MajorName,
		requestData.EducationSector,
		requestData.datestring,
		requestData.SpecifyReason,
		requestData.InputPhoneNumber,
		requestData.Date,
	)

	if err != nil {
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


	// สุ่ม Code pdf 
	newPrintStoryCode := fmt.Sprintf("R%09d", rand.Intn(1000000000))


	var Request entity.Request
	db.Where("id", requestData.RequestID).First(&Request)

	// บันทึกลง databases
	printStory := entity.PrintStory{
		PrintStoryCode: newPrintStoryCode,
		DocumentPath:   savePath,
		DocumentFile: pdfData.Bytes(),
		CreateAt:       time.Now(),
		RequestID:      requestData.RequestID,
	}

	if err := config.DB().Create(&printStory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save print story"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "PDF generated and saved successfully", "data": printStory})
}