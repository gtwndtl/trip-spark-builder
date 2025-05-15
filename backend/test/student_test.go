package unit

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"github.com/sut67/team09/entity"
)

func TestStudentCode(t *testing.T){
	g := NewGomegaWithT(t)

	gender := uint(1)
	major := uint(1)
	status := uint(1)
	semester := uint(1)

	t.Run(`create student is success`, func(t *testing.T){
		student := entity.Student{
			StudentCode: "B1234567",
			FirstName: "Som",
			LastName: "Sak",
			Email: "StudentTest@smail.com",
			Password: "123456",
			NationalID: "1234567890123",
			Phone: "0987654321",
			Profile: "Test",
			BirthDay: time.Now(),
			YearOfStudy: 1,
			GenderID: &gender,
			MajorID: &major,
			StatusStaffID: &status,
			SemesterID: semester,
		}

		ok, err := govalidator.ValidateStruct(student)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})

	t.Run(`student_code is required`, func(t *testing.T) {
		student := entity.Student{
			StudentCode: "", //ผิดตรงนี้
			FirstName: "Som",
			LastName: "Sak",
			Email: "StudentTest@smail.com",
			Password: "123456",
			NationalID: "1234567890123",
			Phone: "0987654321",
			Profile: "Test",
			BirthDay: time.Now(),
			YearOfStudy: 1,
			GenderID: &gender,
			MajorID: &major,
			StatusStaffID: &status,
			SemesterID: uint(1),
		}

		ok, err := govalidator.ValidateStruct(student)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("StudentCode is required"))
	})

	t.Run(`student_code is invalid`, func(t *testing.T) {
		student := entity.Student{
			StudentCode: "A12345", //ผิดตรงนี้
			FirstName: "Som",
			LastName: "Sak",
			Email: "StudentTest@smail.com",
			Password: "123456",
			NationalID: "1234567890123",
			Phone: "0987654321",
			Profile: "Test",
			BirthDay: time.Now(),
			YearOfStudy: 1,
			GenderID: &gender,
			MajorID: &major,
			StatusStaffID: &status,
			SemesterID: uint(1),
		}

		ok, err := govalidator.ValidateStruct(student)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("StudentCode is invalid"))
	})
}

func TestFirstName(t *testing.T){
	g := NewGomegaWithT(t)

	gender := uint(1)
	major := uint(1)
	status := uint(1)

	t.Run(`first_name is required`, func(t *testing.T){
		student := entity.Student{
			StudentCode: "B1234567",
			FirstName: "", //ผิดตรงนี้
			LastName: "Sak",
			Email: "StudentTest@smail.com",
			Password: "123456",
			NationalID: "1234567890123",
			Phone: "0987654321",
			Profile: "Test",
			BirthDay: time.Now(),
			YearOfStudy: 1,
			GenderID: &gender,
			MajorID: &major,
			StatusStaffID: &status,
			SemesterID: uint(1),
		}

		ok, err := govalidator.ValidateStruct(student)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("FirstName is required"))
	})
}

func TestLastName(t *testing.T){
	g := NewGomegaWithT(t)

	gender := uint(1)
	major := uint(1)
	status := uint(1)

	t.Run(`last_name is required`, func(t *testing.T){
		student := entity.Student{
			StudentCode: "B1234567",
			FirstName: "Som", 
			LastName: "", //ผิดตรงนี้
			Email: "StudentTest@smail.com",
			Password: "123456",
			NationalID: "1234567890123",
			Phone: "0987654321",
			Profile: "Test",
			BirthDay: time.Now(),
			YearOfStudy: 1,
			GenderID: &gender,
			MajorID: &major,
			StatusStaffID: &status,
			SemesterID: uint(1),
		}

		ok, err := govalidator.ValidateStruct(student)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("LastName is required"))
	})
}

func TestEmailPasswordNationIDStudent(t *testing.T){
	g := NewGomegaWithT(t)

	gender := uint(1)
	major := uint(1)
	status := uint(1)

	t.Run(`email is required`, func(t *testing.T){
		student := entity.Student{
			StudentCode: "B1234567",
			FirstName: "Som", 
			LastName: "Sak", 
			Email: "", //ผิดตรงนี้
			Password: "123456",
			NationalID: "1234567890123",
			Phone: "0987654321",
			Profile: "Test",
			BirthDay: time.Now(),
			YearOfStudy: 1,
			GenderID: &gender,
			MajorID: &major,
			StatusStaffID: &status,
			SemesterID: uint(1),
		}

		ok, err := govalidator.ValidateStruct(student)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Email is required"))
	})

	t.Run(`email is invalid`, func(t *testing.T) {
		student := entity.Student{
			StudentCode: "B1234567",
			FirstName: "Som", 
			LastName: "Sak", 
			Email: "Sut.com", //ผิดตรงนี้
			Password: "123456",
			NationalID: "1234567890123",
			Phone: "0987654321",
			Profile: "Test",
			BirthDay: time.Now(),
			YearOfStudy: 1,
			GenderID: &gender,
			MajorID: &major,
			StatusStaffID: &status,
			SemesterID: uint(1),
		}

		ok, err := govalidator.ValidateStruct(student)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Email is invalid"))
	})

	t.Run(`email is not matches pattern`, func(t *testing.T) {
		student := entity.Student{
			StudentCode: "B1234567",
			FirstName: "Som", 
			LastName: "Sak", 
			Email: "Test@ut.com", //ผิดตรงนี้
			Password: "123456",
			NationalID: "1234567890123",
			Phone: "0987654321",
			Profile: "Test",
			BirthDay: time.Now(),
			YearOfStudy: 1,
			GenderID: &gender,
			MajorID: &major,
			StatusStaffID: &status,
			SemesterID: uint(1),
		}

		ok, err := govalidator.ValidateStruct(student)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Email must end with @smail.com"))
	})

	t.Run(`password is required`, func(t *testing.T) {
		student := entity.Student{
			StudentCode: "B1234567",
			FirstName: "Som", 
			LastName: "Sak", 
			Email: "Sut@smail.com", 
			Password: "", //ผิดตรงนี้
			NationalID: "1234567890123",
			Phone: "0987654321",
			Profile: "Test",
			BirthDay: time.Now(),
			YearOfStudy: 1,
			GenderID: &gender,
			MajorID: &major,
			StatusStaffID: &status,
			SemesterID: uint(1),
		}

		ok, err := govalidator.ValidateStruct(student)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Password is required"))
	})

	t.Run(`NationalID is required`, func(t *testing.T) {
		student := entity.Student{
			StudentCode: "B1234567",
			FirstName: "Som", 
			LastName: "Sak", 
			Email: "Sut@smail.com", 
			Password: "12345",
			NationalID: "", //ผิดตรงนี้
			Phone: "0987654321",
			Profile: "Test",
			BirthDay: time.Now(),
			YearOfStudy: 1,
			GenderID: &gender,
			MajorID: &major,
			StatusStaffID: &status,
			SemesterID: uint(1),
		}

		ok, err := govalidator.ValidateStruct(student)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("NationalID is required"))
	})

	t.Run(`NationalID is invalid`, func(t *testing.T) {
		student := entity.Student{
			StudentCode: "B1234567",
			FirstName: "Som", 
			LastName: "Sak", 
			Email: "Sut@smail.com", 
			Password: "12345",
			NationalID: "1243532", //ผิดตรงนี้
			Phone: "0987654321",
			Profile: "Test",
			BirthDay: time.Now(),
			YearOfStudy: 1,
			GenderID: &gender,
			MajorID: &major,
			StatusStaffID: &status,
			SemesterID: uint(1),
		}

		ok, err := govalidator.ValidateStruct(student)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("NationalID must be exactly 13 digits"))
	})
}

func TestPhone(t *testing.T){
	g := NewGomegaWithT(t)

	gender := uint(1)
	major := uint(1)
	status := uint(1)

	t.Run(`phone is required`, func(t *testing.T){
		student := entity.Student{
			StudentCode: "B1234567",
			FirstName: "Som", 
			LastName: "Sak", 
			Email: "Sut@smail.com", 
			Password: "123456", 
			NationalID: "1234567890123",
			Phone: "", //ผิดตรงนี้
			Profile: "Test",
			BirthDay: time.Now(),
			YearOfStudy: 1,
			GenderID: &gender,
			MajorID: &major,
			StatusStaffID: &status,
			SemesterID: uint(1),
		}

		ok, err := govalidator.ValidateStruct(student)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Phone is required"))
	})

	t.Run(`phone is invalid`, func(t *testing.T){
		student := entity.Student{
			StudentCode: "B1234567",
			FirstName: "Som", 
			LastName: "Sak", 
			Email: "Sut@smail.com", 
			Password: "123456", 
			NationalID: "1234567890123",
			Phone: "123456789", //ผิดตรงนี้
			Profile: "Test",
			BirthDay: time.Now(),
			YearOfStudy: 1,
			GenderID: &gender,
			MajorID: &major,
			StatusStaffID: &status,
			SemesterID: uint(1),
		}

		ok, err := govalidator.ValidateStruct(student)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Phone must start with 0 and be 10 digits"))
	})
}

func TestProfile(t *testing.T){
	g := NewGomegaWithT(t)

	gender := uint(1)
	major := uint(1)
	status := uint(1)

	t.Run(`profile is required`, func(t *testing.T){
		student := entity.Student{
			StudentCode: "B1234567",
			FirstName: "Som", 
			LastName: "Sak", 
			Email: "Sut@smail.com", 
			Password: "123456", 
			NationalID: "1234567890123",
			Phone: "0987654321",
			Profile: "", //ผิดตรงนี้
			BirthDay: time.Now(),
			YearOfStudy: 1,
			GenderID: &gender,
			MajorID: &major,
			StatusStaffID: &status,
			SemesterID: uint(1),
		}

		ok, err := govalidator.ValidateStruct(student)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Profile is required"))
	})
}

func TestBirthDay(t *testing.T){
	g := NewGomegaWithT(t)

	gender := uint(1)
	major := uint(1)
	status := uint(1)

	var birthday time.Time

	t.Run(`birthday is required`, func(t *testing.T) {
		student := entity.Student{
			StudentCode: "B1234567",
			FirstName: "Som", 
			LastName: "Sak", 
			Email: "Sut@smail.com", 
			Password: "123456", 
			NationalID: "1234567890123",
			Phone: "0987654321",
			Profile: "Test",
			BirthDay: birthday, //ผิดตรงนี้
			YearOfStudy: 1,
			GenderID: &gender,
			MajorID: &major,
			StatusStaffID: &status,
			SemesterID: uint(1),
		}

		ok, err := govalidator.ValidateStruct(student)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("BirthDay is required"))
	})
}

func TestYearOfStudy(t *testing.T) {
	g := NewGomegaWithT(t)

	var year int
	gender := uint(1)
	major := uint(1)
	status := uint(1)

	t.Run(`year_of_study is required`, func(t *testing.T) {
		student := entity.Student{
			StudentCode: "B1234567",
			FirstName: "Som", 
			LastName: "Sak", 
			Email: "Sut@smail.com", 
			Password: "123456", 
			NationalID: "1234567890123",
			Phone: "0987654321",
			Profile: "Test",
			BirthDay: time.Now(), 
			YearOfStudy: year,//ผิดตรงนี้
			GenderID: &gender,
			MajorID: &major,
			StatusStaffID: &status,
			SemesterID: uint(1),
		}

		ok, err := govalidator.ValidateStruct(student)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("YearOfStudy is required"))
	})
}

func TestReferenceID(t *testing.T) {
	g := NewGomegaWithT(t)

	var g0, m0, s0 uint

	gender := uint(1)
	major := uint(1)
	status := uint(1)

	t.Run(`gender_id is required`, func(t *testing.T) {
		student := entity.Student{
			StudentCode: "B1234567",
			FirstName: "Som", 
			LastName: "Sak", 
			Email: "Sut@smail.com", 
			Password: "123456", 
			NationalID: "1234567890123",
			Phone: "0987654321",
			Profile: "Test",
			BirthDay: time.Now(), 
			YearOfStudy: 1,
			GenderID: &g0, //ผิดตรงนี้
			MajorID: &major,
			StatusStaffID: &status,
			SemesterID: uint(1),
		}

		ok, err := govalidator.ValidateStruct(student)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("GenderID is required"))
	})

	t.Run(`major_id is required`, func(t *testing.T) {
		student := entity.Student{
			StudentCode: "B1234567",
			FirstName: "Som", 
			LastName: "Sak", 
			Email: "Sut@smail.com", 
			Password: "123456", 
			NationalID: "1234567890123",
			Phone: "0987654321",
			Profile: "Test",
			BirthDay: time.Now(), 
			YearOfStudy: 1,
			GenderID: &gender, 
			MajorID: &m0, //ผิดตรงนี้
			StatusStaffID: &status,
			SemesterID: uint(1),
		}

		ok, err := govalidator.ValidateStruct(student)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("MajorID is required"))
	})

	t.Run(`status_staff_id is required`, func(t *testing.T) {
		student := entity.Student{
			StudentCode: "B1234567",
			FirstName: "Som", 
			LastName: "Sak", 
			Email: "Sut@smail.com", 
			Password: "123456", 
			NationalID: "1234567890123",
			Phone: "0987654321",
			Profile: "Test",
			BirthDay: time.Now(), 
			YearOfStudy: 1,
			GenderID: &gender, 
			MajorID: &major,
			StatusStaffID: &s0, //ผิดตรงนี้
			SemesterID: uint(1),
		}

		ok, err := govalidator.ValidateStruct(student)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("StatusStaffID is required"))
	})

	t.Run(`semester_id is required`, func(t *testing.T) {
		student := entity.Student{
			StudentCode: "B1234567",
			FirstName: "Som", 
			LastName: "Sak", 
			Email: "Sut@smail.com", 
			Password: "123456", 
			NationalID: "1234567890123",
			Phone: "0987654321",
			Profile: "Test",
			BirthDay: time.Now(), 
			YearOfStudy: 1,
			GenderID: &gender, 
			MajorID: &major,
			StatusStaffID: &status,
			SemesterID: s0, //ผิดตรงนี้
		}

		ok, err := govalidator.ValidateStruct(student)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("SemesterID is required"))
	})
}