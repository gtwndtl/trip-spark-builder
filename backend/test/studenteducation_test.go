package unit

import (
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"github.com/sut67/team09/entity"
	"testing"
	"time"
)

func TestStudentEducation(t *testing.T) {
	g := NewGomegaWithT(t)

	gender := uint(1)
	major := uint(1)
	status := uint(1)
	educate := uint(1)
	degree := uint(1)

	var s0, e0, d0 uint
	studentID := uint(1)

	education1 := entity.Education{
		Instution: "Test",
	}
	student := entity.Student{
		StudentCode:   "B1234567",
		FirstName:     "Som",
		LastName:      "Sak",
		Email:         "StudentTest@smail.com",
		Password:      "123456",
		NationalID: "1234567890123",
		Phone:         "0987654321",
		Profile:       "Test",
		BirthDay:      time.Now(),
		YearOfStudy:   1,
		GenderID:      &gender,
		MajorID:       &major,
		StatusStaffID: &status,
		SemesterID:    1,
	}

	t.Run(`create student_education is success`, func(t *testing.T) {
		se := entity.StudentEducation{
			StudentID: &studentID,
			Student: student,
			EducationID: &educate,
			Education: education1,
			DegreeID: &degree,
			Certificate: "Test Certificate",
		}

		ok, err := govalidator.ValidateStruct(se)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})

	t.Run(`student_id is required`, func(t *testing.T) {
		se := entity.StudentEducation{
			StudentID: &s0, //ผิดตรงนี้
			Student: student,
			EducationID: &educate,
			Education: education1,
			DegreeID: &degree,
			Certificate: "Test Certificate",
		}

		ok, err := govalidator.ValidateStruct(se)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("StudentID is required"))
	})

	t.Run(`education_id is required`, func(t *testing.T) {
		se := entity.StudentEducation{
			StudentID: &studentID,
			Student: student,
			EducationID: &e0, //ผิดตรงนี้
			Education: education1,
			DegreeID: &degree,
			Certificate: "Test Certificate",
		}

		ok, err := govalidator.ValidateStruct(se)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("EducationID is required"))
	})

	t.Run(`degree_id is required`, func(t *testing.T) {
		se := entity.StudentEducation{
			StudentID: &studentID,
			Student: student,
			EducationID: &educate, 
			Education: education1,
			DegreeID: &d0, //ผิดตรงนี้
			Certificate: "Test Certificate",
		}

		ok, err := govalidator.ValidateStruct(se)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("DegreeID is required"))
	})

	t.Run(`certificate is required`, func(t *testing.T) {
		se := entity.StudentEducation{
			StudentID: &studentID,
			Student: student,
			EducationID: &educate, 
			Education: education1,
			DegreeID: &degree, 
			Certificate: "",//ผิดตรงนี้
		}

		ok, err := govalidator.ValidateStruct(se)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Certificate is required"))
	})
}
