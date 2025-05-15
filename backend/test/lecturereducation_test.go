package unit

import (
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"github.com/sut67/team09/entity"
	"testing"
	"time"
)

func TestLecturerEducation(t *testing.T){
	g := NewGomegaWithT(t)

	var lec0, educate0, degreeID0 *uint

	lec := uint(1)
	educate := uint(1)
	degreeID := uint(1)
	postionID := uint(1)
	genderID := uint(1)
	majorID := uint(1)

	lec101 := entity.Lecturer{
		LecturerCode: "L1234567",
		FirstName: "Somsak",
		LastName: "Sri",
		BirthDay: time.Now().AddDate(-30, 0, 0),
		Email: "L1234567@gmail.com",
		Password:"L1234567",
		NationalID: "1234567890123",
		Phone: "0987654321",
		Profile: "test",
		DateEmployed: time.Now().AddDate(-30, 0, 0),
		PositionID: &postionID,
		GenderID: &genderID,
		MajorID: &majorID,
	}

	education1 := entity.Education{
		Instution: "Test",
	}

	degree1 := entity.Degree{
		Degree: "Test",
	}

	t.Run(`lecturer_education is valid`, func(t *testing.T) {
		lecedu := entity.LecturerEducation{
			LecturerID: &lec,
			Lecturer: lec101,
			EducationID: &educate,
			Education: education1,
			DegreeID: &degreeID,
			Degree: degree1,
			Certificate: "Test",
		}

		ok, err := govalidator.ValidateStruct(lecedu)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})

	t.Run(`lecturer_id is required`, func(t *testing.T) {
		lecedu := entity.LecturerEducation{
			LecturerID: lec0, //ผิดตรงนี้
			Lecturer: lec101,
			EducationID: &educate,
			Education: education1,
			DegreeID: &degreeID,
			Degree: degree1,
			Certificate: "Test",
		}

		ok, err := govalidator.ValidateStruct(lecedu)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("LecturerID is required"))
	})

	t.Run(`education_id is required`, func(t *testing.T) {
		lecedu := entity.LecturerEducation{
			LecturerID: &lec, 
			Lecturer: lec101,
			EducationID: educate0, //ผิดตรงนี้
			Education: education1,
			DegreeID: &degreeID,
			Degree: degree1,
			Certificate: "Test",
		}

		ok, err := govalidator.ValidateStruct(lecedu)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("EducationID is required"))
	})

	t.Run(`degree_id is required`, func(t *testing.T) {
		lecedu := entity.LecturerEducation{
			LecturerID: &lec, 
			Lecturer: lec101,
			EducationID: &educate, 
			Education: education1,
			DegreeID: degreeID0, //ผิดตรงนี้
			Degree: degree1,
			Certificate: "Test",
		}

		ok, err := govalidator.ValidateStruct(lecedu)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("DegreeID is required"))
	})

	t.Run(`certificate is required`, func(t *testing.T) {
		lecedu := entity.LecturerEducation{
			LecturerID: &lec, 
			Lecturer: lec101,
			EducationID: &educate, 
			Education: education1,
			DegreeID: &degreeID, 
			Degree: degree1,
			Certificate: "", //ผิดตรงนี้
		}

		ok, err := govalidator.ValidateStruct(lecedu)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Certificate is required"))
	})

}