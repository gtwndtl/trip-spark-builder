package unit

import (
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"github.com/sut67/team09/entity"
	"testing"
	"time"
)

func TestLecturerThesis(t *testing.T) {
	g := NewGomegaWithT(t)

	var lec0, thesis0, role0 *uint

	lec := uint(1)
	thsID := uint(1)
	roleID := uint(1)
	postionID := uint(1)
	genderID := uint(1)
	majorID := uint(1)

	lec101 := entity.Lecturer{
		LecturerCode: "L1234567", //กรอกข้อมูลถูก
		FirstName:    "Somsak",
		LastName:     "Sri",
		BirthDay:     time.Now().AddDate(-30, 0, 0),
		Email:        "L1234567@gmail.com",
		Password:     "L1234567",
		NationalID: "1234567890123",
		Phone:        "0987654321",
		Profile:      "test",
		DateEmployed: time.Now().AddDate(-30, 0, 0),
		PositionID:   &postionID,
		GenderID:     &genderID,
		MajorID:      &majorID,
	}
	thesisID := entity.Thesis{
		Title: "TestThesis",
		PublicationDate: time.Now().AddDate(-30, 0, 0),
		URL: "Test.com",
	}

	roleofthesisID := entity.RoleOfThesis{
		RoleOfThesisName: "Advisor",
	}

	t.Run(`lecturer_thesis is valid`, func(t *testing.T) {
		lecthe0 := entity.LecturerThesis{
			LecturerID:     &lec,
			Lecturer: lec101,
			ThesisID:       &thsID,
			Thesis: thesisID,
			RoleOfThesisID: &roleID,
			RoleOfThesis: roleofthesisID,
		}

		ok, err := govalidator.ValidateStruct(lecthe0)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})

	t.Run(`lecturer_id is required`, func(t *testing.T) {
		lecthe0 := entity.LecturerThesis{
			LecturerID:     lec0, //ผิดตรงนี้
			Lecturer: lec101,
			ThesisID:       &thsID,
			Thesis: thesisID,
			RoleOfThesisID: &roleID,
			RoleOfThesis: roleofthesisID,
		}

		ok, err := govalidator.ValidateStruct(lecthe0)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("LecturerID is required"))
	})

	t.Run(`thesis_id is required`, func(t *testing.T) {
		lecthe0 := entity.LecturerThesis{
			LecturerID:     &lec,
			Lecturer: lec101,
			ThesisID:       thesis0, //ผิดตรงนี้
			Thesis: thesisID,
			RoleOfThesisID: &roleID,
			RoleOfThesis: roleofthesisID,
		}

		ok, err := govalidator.ValidateStruct(lecthe0)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("ThesisID is required"))
	})

	t.Run(`role_of_thesis_id is required`, func(t *testing.T) {
		lecthe0 := entity.LecturerThesis{
			LecturerID:     &lec,
			Lecturer: lec101,
			ThesisID:       &thsID,
			Thesis: thesisID,
			RoleOfThesisID: role0, //ผิดตรงนี้
			RoleOfThesis: roleofthesisID,
		}

		ok, err := govalidator.ValidateStruct(lecthe0)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("RoleOfThesisID is required"))
	})
}
