package unit

import (
	"testing"
	"time"

	// "time"

	. "github.com/onsi/gomega"
	"github.com/sut67/team09/controller/pdf"
)

func TestGeneratePDF(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run(`InputChek`, func(t *testing.T) {
		//กรณีถูก
		var InputName = "Sirion"
		var InputStudentID = "B6505066"
		var Degree = "1"
		var FacultyName = "1"
		var MajorName = "1"
		var Details = "เพิ่มวิชา"
		var CourseCode = "Eng23 3032"
 		var CourseTitle = "Software Engineering"
		var Group = "1"
		var OldGroup = "1"
		var NewGroup= "1"
		var SpecifyReason= "ต้องการเพิ่มวิชาเนื่องจากมีวิชาตัวต่อ"
		var InputPhoneNumber = "0651018312"
		var Date = time.Now()



		pdf, err := pdf.GeneratePDF(
			InputName,
			InputStudentID,
			Degree,
			FacultyName,
			MajorName,
			Details,
			CourseCode,
			CourseTitle,
			Group,
			OldGroup,
			NewGroup,
			SpecifyReason,
			InputPhoneNumber,
			Date,
		)


		
		if (err != nil) {
			print("err",err.Error())
		}
	

		g.Expect(pdf.Len() > 0).To(BeTrue())
		g.Expect(err).To(BeNil())

	})

	t.Run(`InputName Check`, func(t *testing.T) {
		var InputName = "" //ผิดตรงนี้ ส่งค่าว่างไป
		var InputStudentID = "B6505066" 
		var Degree = "1"
		var FacultyName = "1"
		var MajorName = "1"
		var Details = "เพิ่มวิชา"
		var CourseCode = "Eng23 3032"
 		var CourseTitle = "Software Engineering"
		var Group = "1"
		var OldGroup = "1"
		var NewGroup= "1"
		var SpecifyReason= "ต้องการเพิ่มวิชาเนื่องจากมีวิชาตัวต่อ"
		var InputPhoneNumber = "0651018312"
		var Date = time.Now()



		_, err := pdf.GeneratePDF(
			InputName,
			InputStudentID,
			Degree,
			FacultyName,
			MajorName,
			Details,
			CourseCode,
			CourseTitle,
			Group,
			OldGroup,
			NewGroup,
			SpecifyReason,
			InputPhoneNumber,
			Date,
		)


		
		if (err != nil) {
			print("err",err.Error())
		}

		print("error check",err)
	
		g.Expect(err.Error()).To(Equal("InputName is required"))


	})

	t.Run(`InputStudentID Check`, func(t *testing.T) {
		var InputName = "Sirion" 
		var InputStudentID = "" //ผิดตรงนี้ ส่งค่าว่างไป
		var Degree = "1"
		var FacultyName = "1"
		var MajorName = "1"
		var Details = "เพิ่มวิชา"
		var CourseCode = "Eng23 3032"
 		var CourseTitle = "Software Engineering"
		var Group = "1"
		var OldGroup = "1"
		var NewGroup= "1"
		var SpecifyReason= "ต้องการเพิ่มวิชาเนื่องจากมีวิชาตัวต่อ"
		var InputPhoneNumber = "0651018312"
		var Date = time.Now()



		_, err := pdf.GeneratePDF(
			InputName,
			InputStudentID,
			Degree,
			FacultyName,
			MajorName,
			Details,
			CourseCode,
			CourseTitle,
			Group,
			OldGroup,
			NewGroup,
			SpecifyReason,
			InputPhoneNumber,
			Date,
		)


		
		if (err != nil) {
			print("err",err.Error())
		}

		print("error check",err)
	
		g.Expect(err.Error()).To(Equal("InputStudentID is required"))


	})


	t.Run(`Degree Check`, func(t *testing.T) {
		var InputName = "Sirion" 
		var InputStudentID = "B6505066" 
		var Degree = "" //ผิดตรงนี้ ส่งค่าว่างไป
		var FacultyName = "1"
		var MajorName = "1"
		var Details = "เพิ่มวิชา"
		var CourseCode = "Eng23 3032"
 		var CourseTitle = "Software Engineering"
		var Group = "1"
		var OldGroup = "1"
		var NewGroup= "1"
		var SpecifyReason= "ต้องการเพิ่มวิชาเนื่องจากมีวิชาตัวต่อ"
		var InputPhoneNumber = "0651018312"
		var Date = time.Now()



		_, err := pdf.GeneratePDF(
			InputName,
			InputStudentID,
			Degree,
			FacultyName,
			MajorName,
			Details,
			CourseCode,
			CourseTitle,
			Group,
			OldGroup,
			NewGroup,
			SpecifyReason,
			InputPhoneNumber,
			Date,
		)


		
		if (err != nil) {
			print("err",err.Error())
		}

		print("error check",err)
	
		g.Expect(err.Error()).To(Equal("Degree is required"))


	})


	t.Run(`FacultyName Check`, func(t *testing.T) {
		var InputName = "Sirion" 
		var InputStudentID = "B6505066" 
		var Degree = "1" 
		var FacultyName = "" //ผิดตรงนี้ ส่งค่าว่างไป
		var MajorName = "1"
		var Details = "เพิ่มวิชา"
		var CourseCode = "Eng23 3032"
 		var CourseTitle = "Software Engineering"
		var Group = "1"
		var OldGroup = "1"
		var NewGroup= "1"
		var SpecifyReason= "ต้องการเพิ่มวิชาเนื่องจากมีวิชาตัวต่อ"
		var InputPhoneNumber = "0651018312"
		var Date = time.Now()



		_, err := pdf.GeneratePDF(
			InputName,
			InputStudentID,
			Degree,
			FacultyName,
			MajorName,
			Details,
			CourseCode,
			CourseTitle,
			Group,
			OldGroup,
			NewGroup,
			SpecifyReason,
			InputPhoneNumber,
			Date,
		)


		
		if (err != nil) {
			print("err",err.Error())
		}

		print("error check",err)
	
		g.Expect(err.Error()).To(Equal("Faculty is required"))


	})

	t.Run(`MajorName Check`, func(t *testing.T) {
		var InputName = "Sirion" 
		var InputStudentID = "B6505066" 
		var Degree = "1" 
		var FacultyName = "1" 
		var MajorName = "" //ผิดตรงนี้ ส่งค่าว่างไป
		var Details = "เพิ่มวิชา"
		var CourseCode = "Eng23 3032"
 		var CourseTitle = "Software Engineering"
		var Group = "1"
		var OldGroup = "1"
		var NewGroup= "1"
		var SpecifyReason= "ต้องการเพิ่มวิชาเนื่องจากมีวิชาตัวต่อ"
		var InputPhoneNumber = "0651018312"
		var Date = time.Now()



		_, err := pdf.GeneratePDF(
			InputName,
			InputStudentID,
			Degree,
			FacultyName,
			MajorName,
			Details,
			CourseCode,
			CourseTitle,
			Group,
			OldGroup,
			NewGroup,
			SpecifyReason,
			InputPhoneNumber,
			Date,
		)


		
		if (err != nil) {
			print("err",err.Error())
		}

		print("error check",err)
	
		g.Expect(err.Error()).To(Equal("Major is required"))


	})


	// t.Run(`Details Check`, func(t *testing.T) {
	// 	var InputName = "Sirion" 
	// 	var InputStudentID = "B6505066" 
	// 	var Degree = "1" 
	// 	var FacultyName = "1" 
	// 	var MajorName = "1" 
	// 	var Details = "" //ผิดตรงนี้ ส่งค่าว่างไป
	// 	var CourseCode = "Eng23 3032"
 	// 	var CourseTitle = "Software Engineering"
	// 	var Group = "1"
	// 	var OldGroup = "1"
	// 	var NewGroup= "1"
	// 	var SpecifyReason= "ต้องการเพิ่มวิชาเนื่องจากมีวิชาตัวต่อ"
	// 	var InputPhoneNumber = "0651018312"
	// 	var Date = time.Now()



	// 	_, err := pdf.GeneratePDF(
	// 		InputName,
	// 		InputStudentID,
	// 		Degree,
	// 		FacultyName,
	// 		MajorName,
	// 		Details,
	// 		CourseCode,
	// 		CourseTitle,
	// 		Group,
	// 		OldGroup,
	// 		NewGroup,
	// 		SpecifyReason,
	// 		InputPhoneNumber,
	// 		Date,
	// 	)


		
	// 	if (err != nil) {
	// 		print("err",err.Error())
	// 	}

	// 	print("error check",err)
	
	// 	g.Expect(err.Error()).To(Equal("Details is required"))


	// })


	t.Run(`CourseCode Check`, func(t *testing.T) {
		var InputName = "Sirion" 
		var InputStudentID = "B6505066" 
		var Degree = "1" 
		var FacultyName = "1" 
		var MajorName = "1" 
		var Details = "เพิ่มวิชา" 
		var CourseCode = "" //ผิดตรงนี้ ส่งค่าว่างไป
 		var CourseTitle = "Software Engineering"
		var Group = "1"
		var OldGroup = "1"
		var NewGroup= "1"
		var SpecifyReason= "ต้องการเพิ่มวิชาเนื่องจากมีวิชาตัวต่อ"
		var InputPhoneNumber = "0651018312"
		var Date = time.Now()



		_, err := pdf.GeneratePDF(
			InputName,
			InputStudentID,
			Degree,
			FacultyName,
			MajorName,
			Details,
			CourseCode,
			CourseTitle,
			Group,
			OldGroup,
			NewGroup,
			SpecifyReason,
			InputPhoneNumber,
			Date,
		)


		
		if (err != nil) {
			print("err",err.Error())
		}

		print("error check",err)
	
		g.Expect(err.Error()).To(Equal("CourseCode is required"))


	})


	t.Run(`CourseTitle Check`, func(t *testing.T) {
		var InputName = "Sirion" 
		var InputStudentID = "B6505066" 
		var Degree = "1" 
		var FacultyName = "1" 
		var MajorName = "1" 
		var Details = "เพิ่มวิชา" 
		var CourseCode = "Eng23 3032" 
 		var CourseTitle = "" //ผิดตรงนี้ ส่งค่าว่างไป
		var Group = "1"
		var OldGroup = "1"
		var NewGroup= "1"
		var SpecifyReason= "ต้องการเพิ่มวิชาเนื่องจากมีวิชาตัวต่อ"
		var InputPhoneNumber = "0651018312"
		var Date = time.Now()



		_, err := pdf.GeneratePDF(
			InputName,
			InputStudentID,
			Degree,
			FacultyName,
			MajorName,
			Details,
			CourseCode,
			CourseTitle,
			Group,
			OldGroup,
			NewGroup,
			SpecifyReason,
			InputPhoneNumber,
			Date,
		)


		
		if (err != nil) {
			print("err",err.Error())
		}

		print("error check",err)
	
		g.Expect(err.Error()).To(Equal("CourseTitle is required"))


	})


	t.Run(`Group Check`, func(t *testing.T) {
		var InputName = "Sirion" 
		var InputStudentID = "B6505066" 
		var Degree = "1" 
		var FacultyName = "1" 
		var MajorName = "1" 
		var Details = "เพิ่มวิชา" 
		var CourseCode = "Eng23 3032" 
 		var CourseTitle = "Software Engineering" 
		var Group = "" //ผิดตรงนี้ ส่งค่าว่างไป
		var OldGroup = "1"
		var NewGroup= "1"
		var SpecifyReason= "ต้องการเพิ่มวิชาเนื่องจากมีวิชาตัวต่อ"
		var InputPhoneNumber = "0651018312"
		var Date = time.Now()



		_, err := pdf.GeneratePDF(
			InputName,
			InputStudentID,
			Degree,
			FacultyName,
			MajorName,
			Details,
			CourseCode,
			CourseTitle,
			Group,
			OldGroup,
			NewGroup,
			SpecifyReason,
			InputPhoneNumber,
			Date,
		)


		
		if (err != nil) {
			print("err",err.Error())
		}

		print("error check",err)
	
		g.Expect(err.Error()).To(Equal("Group is required"))


	})


	// t.Run(`OldGroup Check`, func(t *testing.T) {
	// 	var InputName = "Sirion" 
	// 	var InputStudentID = "B6505066" 
	// 	var Degree = "1" 
	// 	var FacultyName = "1" 
	// 	var MajorName = "1" 
	// 	var Details = "เพิ่มวิชา" 
	// 	var CourseCode = "Eng23 3032" 
 	// 	var CourseTitle = "Software Engineering" 
	// 	var Group = "1" 
	// 	var OldGroup = "" //ผิดตรงนี้ ส่งค่าว่างไป
	// 	var NewGroup= "1"
	// 	var SpecifyReason= "ต้องการเพิ่มวิชาเนื่องจากมีวิชาตัวต่อ"
	// 	var InputPhoneNumber = "0651018312"
	// 	var Date = time.Now()



	// 	_, err := pdf.GeneratePDF(
	// 		InputName,
	// 		InputStudentID,
	// 		Degree,
	// 		FacultyName,
	// 		MajorName,
	// 		Details,
	// 		CourseCode,
	// 		CourseTitle,
	// 		Group,
	// 		OldGroup,
	// 		NewGroup,
	// 		SpecifyReason,
	// 		InputPhoneNumber,
	// 		Date,
	// 	)


		
	// 	if (err != nil) {
	// 		print("err",err.Error())
	// 	}

	// 	print("error check",err)
	
	// 	g.Expect(err.Error()).To(Equal("OldGroup is required"))


	// })


	// t.Run(`NewGroup Check`, func(t *testing.T) {
	// 	var InputName = "Sirion" 
	// 	var InputStudentID = "B6505066" 
	// 	var Degree = "1" 
	// 	var FacultyName = "1" 
	// 	var MajorName = "1" 
	// 	var Details = "เพิ่มวิชา" 
	// 	var CourseCode = "Eng23 3032" 
 	// 	var CourseTitle = "Software Engineering" 
	// 	var Group = "1" 
	// 	var OldGroup = "1" 
	// 	var NewGroup= "" //ผิดตรงนี้ ส่งค่าว่างไป
	// 	var SpecifyReason= "ต้องการเพิ่มวิชาเนื่องจากมีวิชาตัวต่อ"
	// 	var InputPhoneNumber = "0651018312"
	// 	var Date = time.Now()



	// 	_, err := pdf.GeneratePDF(
	// 		InputName,
	// 		InputStudentID,
	// 		Degree,
	// 		FacultyName,
	// 		MajorName,
	// 		Details,
	// 		CourseCode,
	// 		CourseTitle,
	// 		Group,
	// 		OldGroup,
	// 		NewGroup,
	// 		SpecifyReason,
	// 		InputPhoneNumber,
	// 		Date,
	// 	)


		
	// 	if (err != nil) {
	// 		print("err",err.Error())
	// 	}

	// 	print("error check",err)
	
	// 	g.Expect(err.Error()).To(Equal("NewGroup is required"))


	// })


	// t.Run(`SpecifyReason Check`, func(t *testing.T) {
	// 	var InputName = "Sirion" 
	// 	var InputStudentID = "B6505066" 
	// 	var Degree = "1" 
	// 	var FacultyName = "1" 
	// 	var MajorName = "1" 
	// 	var Details = "เพิ่มวิชา" 
	// 	var CourseCode = "Eng23 3032" 
 	// 	var CourseTitle = "Software Engineering" 
	// 	var Group = "1" 
	// 	var OldGroup = "1" 
	// 	var NewGroup= "1" 
	// 	var SpecifyReason= "" //ผิดตรงนี้ ส่งค่าว่างไป
	// 	var InputPhoneNumber = "0651018312"
	// 	var Date = time.Now()



	// 	_, err := pdf.GeneratePDF(
	// 		InputName,
	// 		InputStudentID,
	// 		Degree,
	// 		FacultyName,
	// 		MajorName,
	// 		Details,
	// 		CourseCode,
	// 		CourseTitle,
	// 		Group,
	// 		OldGroup,
	// 		NewGroup,
	// 		SpecifyReason,
	// 		InputPhoneNumber,
	// 		Date,
	// 	)


		
	// 	if (err != nil) {
	// 		print("err",err.Error())
	// 	}

	// 	print("error check",err)
	
	// 	g.Expect(err.Error()).To(Equal("SpecifyReason is required"))


	// })


	t.Run(`InputPhoneNumber Check`, func(t *testing.T) {
		var InputName = "Sirion" 
		var InputStudentID = "B6505066" 
		var Degree = "1" 
		var FacultyName = "1" 
		var MajorName = "1" 
		var Details = "เพิ่มวิชา" 
		var CourseCode = "Eng23 3032" 
 		var CourseTitle = "Software Engineering" 
		var Group = "1" 
		var OldGroup = "1" 
		var NewGroup= "1" 
		var SpecifyReason= "ต้องการเพิ่มวิชาเนื่องจากมีวิชาตัวต่อ" 
		var InputPhoneNumber = "" //ผิดตรงนี้ ส่งค่าว่างไป
		var Date = time.Now()



		_, err := pdf.GeneratePDF(
			InputName,
			InputStudentID,
			Degree,
			FacultyName,
			MajorName,
			Details,
			CourseCode,
			CourseTitle,
			Group,
			OldGroup,
			NewGroup,
			SpecifyReason,
			InputPhoneNumber,
			Date,
		)


		
		if (err != nil) {
			print("err",err.Error())
		}

		print("error check",err)
	
		g.Expect(err.Error()).To(Equal("InputPhoneNumber is required"))


	})



	// t.Run(`InputStudentID Check pattern`, func(t *testing.T) {
	// 	var InputName = "Sirion" 
	// 	var InputStudentID = "B6505066666" //ผิดตรงนี้ pattern ผิด
	// 	var Degree = "1" 
	// 	var FacultyName = "1" 
	// 	var MajorName = "1" 
	// 	var Details = "เพิ่มวิชา" 
	// 	var CourseCode = "Eng23 3032" 
 	// 	var CourseTitle = "Software Engineering" 
	// 	var Group = "1" 
	// 	var OldGroup = "1" 
	// 	var NewGroup= "1" 
	// 	var SpecifyReason= "ต้องการเพิ่มวิชาเนื่องจากมีวิชาตัวต่อ" 
	// 	var InputPhoneNumber = "0651018312" 
	// 	var Date = time.Now()



	// 	_, err := pdf.GeneratePDF(
	// 		InputName,
	// 		InputStudentID,
	// 		Degree,
	// 		FacultyName,
	// 		MajorName,
	// 		Details,
	// 		CourseCode,
	// 		CourseTitle,
	// 		Group,
	// 		OldGroup,
	// 		NewGroup,
	// 		SpecifyReason,
	// 		InputPhoneNumber,
	// 		Date,
	// 	)


		
	// 	if (err != nil) {
	// 		print("err",err.Error())
	// 	}

	// 	print("error check",err)
	
	// 	g.Expect(err.Error()).To(ContainSubstring("InputStudentID is invalid"))


	// })


}

