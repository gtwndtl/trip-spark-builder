package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/controller"
	"github.com/sut67/team09/controller/activity"
	"github.com/sut67/team09/controller/admin"
	"github.com/sut67/team09/controller/books"
	"github.com/sut67/team09/controller/enrollment"
	"github.com/sut67/team09/controller/report"
	"github.com/sut67/team09/controller/request"
	"github.com/sut67/team09/controller/student"

	"github.com/sut67/team09/controller/building"
	"github.com/sut67/team09/controller/classroom"
	"github.com/sut67/team09/controller/condition"
	"github.com/sut67/team09/controller/courses"
	"github.com/sut67/team09/controller/day"
	"github.com/sut67/team09/controller/dormitory"
	"github.com/sut67/team09/controller/lecturer"
	"github.com/sut67/team09/controller/payment"
	"github.com/sut67/team09/controller/pdf"
	"github.com/sut67/team09/controller/room"
	"github.com/sut67/team09/controller/schedule"
	"github.com/sut67/team09/controller/studentevaluation"
	"github.com/sut67/team09/middlewares"

	//"github.com/sut67/team09/controller/enrollment"
	//"github.com/sut67/team09/middlewares"
	"github.com/sut67/team09/controller/paymentcourse"
)

const PORT = "8000"

func main() {
	// open connection database
	config.ConnectionDB()

	// Generate databases
	config.SetupDatabase()

	config.ImportCoursesFromExcel("course.xlsx")

	r := gin.Default()
	r.Use(CORSMiddleware())

	r.Static("/uploads", "./uploads")
	r.Static("/uploads_pdf", "./uploads_pdf")
	r.Static("/uploads_pic_activity", "./uploads_pic_activity")

	r.POST("/signinstudent", student.SignInStudent)

	r.POST("/signinadmin", admin.SignInAdmin)

	r.POST("/signinlecturer", lecturer.SignInLecturer)

	router := r.Group("/")
	{
		router.Use(middlewares.Authorizes())

		r.POST("/signin", controller.SignIn)
		//router.Use(middlewares.Authorizes())
		router.GET("/users", student.GetAll)
		router.GET("/admins", admin.GetAll)
		router.GET("/rooms", room.GetAll)
		router.GET("/dormitorys", dormitory.GetAll)
		router.GET("/books", books.GetAll)
		router.DELETE("/book/:id", books.Delete)
		router.POST("/payment", payment.Create)
		router.GET("/payments", payment.GetAll)
		router.GET("/payment/:id", payment.Get)
		router.PUT("/payment/:id", payment.Update) //payment
		router.GET("/enrollments", enrollment.GetAll)
		router.GET("/course", courses.GetAll)
		router.GET("/paymentcourse", paymentcourse.GetAll)
		router.POST("/paycourse", paymentcourse.Create)

		r.POST("/lecturer", lecturer.CreateLecturer)
		r.PUT("/lecturer/:id", lecturer.UpdateLecturer)
		r.DELETE("/deleteLecturer/:id", lecturer.DeleteLecturer)
		r.GET("/checkEmailLecturer/:email", lecturer.CheckEmail) 
		r.GET("/lecturerby/:id", lecturer.GetLecturerByID)
		r.GET("/faculty", lecturer.GetFaculty)
		r.GET("/majorbyfacultyid/:id", lecturer.GetMajorByFacultyID)
		r.GET("/gender", lecturer.GetGenders)
		r.GET("/lecturerbymajorid/:id", lecturer.GetLecturerByMajorID)
		r.GET("/degree", lecturer.GetDegrees)
		r.GET("/position", lecturer.GetPositions)
		r.GET("/positionLecturer", lecturer.GetPositionLecturer)
		r.GET("/statusstaff", lecturer.GetStatusStaffLecturer)
		r.GET("/degreewithoutstudent", lecturer.GetDegreesWithOutStudent)
		r.POST("/createeducation", lecturer.CreateEducation)
		r.POST("/createlecturereducation", lecturer.CreateLecturerEducation)
		r.POST("/createlecturerthesis", lecturer.CreateLecturerThesis)
		r.POST("/createthesis", lecturer.CreateThesis)
		r.GET("/roleofthesis", lecturer.GetRoleOfThesis)
		r.GET("/GetEducation/:id", lecturer.GetEducationByLecturerID)
		r.PUT("/updatelecturereducation/:id", lecturer.UpdateLecturerEducation)
		r.GET("/getThesis/:id", lecturer.GetThesisByLecturerID)
		r.PUT("/updatelecturerThesis/:id", lecturer.UpdateLecturerThesis)
		r.GET("/checkphone/:phone", lecturer.CheckPhone)
		r.GET("/checknationalid/:nationalid", lecturer.CheckNationalID)
		r.GET("/checklecturer_code/:lecturercode", lecturer.CheckLecturerCode)
		r.GET("/checkemail_lecturerByID/:id/:email", lecturer.CheckEmailByID)
		r.GET("/checklecturercodeByID/:id/:lecturercode", lecturer.CheckLecturerCodeByID)
		r.GET("/checknationalidByid/:id/:nationalid", lecturer.CheckNationalIDByID)
		r.GET("/checkphoneByid/:id/:phone", lecturer.CheckPhoneByID)
		r.GET("/statusstaff_student", lecturer.GetStatusStaffStudent)

		r.GET("/getstudent/:id", student.GetStudentByID)
		r.POST("/createstudent", student.CreateStudents)
		r.GET("/getstudentBymajorID/:id", student.GetStudentByMajorID)
		r.GET("/semester", student.GetSemester)
		r.GET("/check-email-student/:email", student.CheckEmailStudent)
		r.GET("/check-email-studentByid/:id/:email", student.CheckEmailStudentByID)
		r.POST("/create-student-education", student.CreateStudentEducation)
		r.GET("/check-student-code/:studentcode", student.CheckStudentCode)
		r.GET("/check_studentcodeByID/:id/:studentcode", student.CheckStudentCodeByID)
		r.PUT("/editStudent/:id", student.UpdateStudent)
		r.PUT("/edit_studentEducation/:id", student.UpdateStudentEducation)
		r.DELETE("/deleteStudent/:id", student.DeleteStudent)
		r.GET("/getstudent_education/:id", student.GetEducationByStudentID)

		r.GET("/allcourses", courses.ListAllCourse)
		r.GET("/courses", courses.ListCourse)
		r.GET("/courses/:id", courses.GetCourseByID)
		r.GET("/courses/related", courses.GetRelatedCourses)
		r.GET("/courses/search", courses.SearchCourseByKeyword)
		r.PUT("/courses/:id", courses.UpdateCourse)
		r.PUT("/courses/:id/toggle", courses.ToggleCourseStatus)
		r.POST("/courses", courses.CreateCourse)
		r.DELETE("/courses/:id", courses.DeleteCourse)
		r.DELETE("/courses/delete/:course_code", courses.DeleteByCourseCode)

		r.GET("/courses/lecturer/:id", courses.GetCoursesByLecturerID)

		r.GET("/schedule", schedule.ListSchedules)
		r.GET("/schedule/:id", schedule.GetScheduleByID)
		r.GET("/schedule/course/id/:id", schedule.GetScheduleByCourseID)
		r.GET("/schedule/course/code/:code", schedule.GetScheduleByCourseCode)
		r.GET("/schedule/enrollment/:id", schedule.GetScheduleByStudentID)
		r.POST("/schedule", schedule.CreateSchedule)
		r.POST("/schedule/:id", schedule.UpdateSchedule)
		r.DELETE("/schedule/:id", schedule.DeleteSchedule)
		r.DELETE("/schedule/delete/:code", schedule.DeleteScheduleByCourseCode)
  
		r.GET("/enrollment", enrollment.ListEnrollments)
		r.GET("/enrollment/student/:id", enrollment.GetEnrollmentByStudentID)
		r.GET("/enrollment/course/:id", enrollment.GetEnrollmentByCourseID)
		r.POST("/enrollment", enrollment.CreateEnrollment)

		r.GET("/day", day.ListDays)
		r.GET("/d/:id", day.GetDayByID)

		r.GET("/condition", condition.ListCondition)
		r.GET("/classroom", classroom.GetLisClassroom)
		r.GET("/classroom/ready", classroom.GetLisClassroomReady)
		r.POST("/classroom", classroom.CreateClassroom)
		r.PUT("/classroom/:id", classroom.UpdateClassroom)
		r.GET("/classroom/:id", classroom.GetClassroomByID)
		r.DELETE("/classroom/:id", classroom.DeleteClassroom)

		r.PUT("/building/:id", building.UpdateBuilding)
		r.POST("/upload", building.UploadFile)
		r.POST("/building", building.CreateBuilding)
		r.GET("/building", building.GetBuildings)
		r.GET("/building/:id", building.GetBuilding)
		r.DELETE("/building/:id", building.DeleteBuilding)

		r.GET("/assessmenttype", studentevaluation.ListAssessmentType)
		r.GET("/studentevaluation", studentevaluation.ListStudentEvaluation)
		r.POST("/studentevaluation", studentevaluation.CreateStudentEvaluation)
		r.GET("/studentevaluation/:id", studentevaluation.GetStudentEvaluationByID)
		r.DELETE("/studentevaluation/:id", studentevaluation.DeleteStudentEvaluation)
		r.PUT("/studentevaluationgrading/:id", studentevaluation.AutomaticUpdateGrading)

		r.GET("/assessmentget", studentevaluation.ListAssessmentGet)
		r.GET("/assessmentget/:id", studentevaluation.GetAssessmentGetByID)
		r.PUT("/assessmentget/:id", studentevaluation.UpdateAssessmentGet)
		r.POST("/assessmentget", studentevaluation.CreateAssessmentGet)
		r.DELETE("/assessmentget/:id", studentevaluation.DeleteAssessmentGet)

		router.GET("/reports", report.GetAll)
		router.GET("/report/:id", report.Get)
		router.DELETE("/report/:id", report.Delete)
		router.POST("/report", report.Create)
		router.PUT("/report/:id", report.Update) //report

		r.POST("/Dormitory", dormitory.CreateDorm)
		r.GET("/Dormitory", dormitory.ListDorm)
		r.PATCH("/Dormitory/:id", dormitory.UpdateDorm)
		r.DELETE("/Dormitory/:id", dormitory.DeleteDorm)
		r.POST("/Room", room.CreateRoom)
		r.GET("/Room/:id", room.ListRoom)
		r.DELETE("/Room/Delete/:RoomNumber", room.DeleteRoom)
		r.DELETE("/Dormitory/DeleteAllRooms/:id", room.DeleteAllRoomsInDormitory)
		r.GET("/books/:id", books.GetBookByID)
		r.GET("/books/student/:id", books.GetBookByStudentID)
		r.PATCH("/Room/status/:id", books.UpdateRoomStatus)
		r.POST("/booking", books.BookingCreate)

		//Jiw Up ระบบยื่นคำร้องสร้าง pdf
		r.GET("/print", func(c *gin.Context) {
			inputName := ""
			inputStudentID := ""
			Degree := ""
			Faculty := ""
			Major := ""
			Details := ""
			CourseCode := ""
			CourseTitle := ""
			Group := ""
			OldGroup := ""
			NewGroup := ""
			SpecifyReason := ""
			inputPhoneNumber := ""
			Date := time.Now()

			pdfData, err := pdf.GeneratePDF(inputName, inputStudentID, Degree, Faculty, Major, Details, CourseCode, CourseTitle, Group, OldGroup, NewGroup, SpecifyReason, inputPhoneNumber, Date)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.Data(http.StatusOK, "application/pdf", pdfData.Bytes())
		})


		r.PATCH("/pdf", pdf.PatchPDF)
		r.POST("/CreatePrintStory", pdf.CreatePrintStory)
		r.GET("/printstory", pdf.GetPrintStory)
		r.GET("/printstory/:request_id", pdf.GetPrintStoryByRequestID)
		r.GET("/pdf/:filename", pdf.GetPDFFile)

		//Request
		r.GET("/request", request.GetRequest)
		r.GET("/request/:id", request.GetRequestByID)
		r.GET("/requesttype", request.GetRequestTypes)
		r.GET("/statusRequest", request.GetStatusRequest)
		r.PATCH("/add/:id", request.UpdateRequest)
		r.POST("/createrequest", request.CreateRequest)
		r.DELETE("/deleterequest/:id", request.DeleteRequest)
		r.GET("/request/student/:student_id", request.GetRequestByStudentID)
		r.GET("/courses/request/:id", request.GetCoursesByRequestID)
		r.GET("/course/:course_code", request.GetCourseByCourseCode)
		r.GET("/course/:course_code/:group", request.GetCourseByCourseCodeAndGroup)
		// r.GET("/request/course/printstory", request.GetRequestAndPrintStoryAndCourse)

		r.GET("/activity", activity.GetActivity)
		r.POST("/activity", activity.CreateActivity)
		r.PATCH("/editActivity/:id", activity.UpdateActivity)
		r.DELETE("/activity/:id", activity.DeleteActivitys)
		r.GET("/activity/:id",activity.GetActivityById) 
		r.POST("/api/enroll", activity.EnrollActivity)
		r.POST("/api/unenroll", activity.UnenrollActivity)

		// //Test ดู Pdf
		// r.GET("/print2", func(ctx *gin.Context) {
		// 	billPDF, err := pdf.GeneratePDF2() 
		// 	if err != nil {
		// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 		return
		// 	}
		// 	ctx.Data(http.StatusOK, "application/pdf", billPDF.Bytes())
		// })

		r.GET("/print2", func(c *gin.Context) {
			inputName := ""
			inputStudentID := ""
			Faculty := ""
			Major := ""
			EducationSector := ""
			datestring := ""
			SpecifyReason := ""
			inputPhoneNumber := ""
			Date := time.Now()

			pdfData, err := pdf.GeneratePDF2(inputName, inputStudentID,  Faculty, Major, EducationSector, datestring, SpecifyReason, inputPhoneNumber, Date )
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.Data(http.StatusOK, "application/pdf", pdfData.Bytes())
		})


		r.PATCH("/pdf2", pdf.PatchPDF2)
		r.POST("/CreatePrintStory2", pdf.CreatePrintStory2)


	}

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "API RUNNING... PORT: %s", PORT)
	})

	// Run the server
	r.Run("localhost:" + PORT)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
 