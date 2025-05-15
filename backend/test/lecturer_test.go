package unit

import (
	"time"
	"testing"
	"github.com/sut67/team09/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestLecturerCode(t *testing.T){
	g := NewGomegaWithT(t)

	postionID := uint(1)
	genderID := uint(1)
	majorID := uint(3)

	t.Run(`lecturer_code is required`, func(t *testing.T) {
		lec0 := entity.Lecturer{
			LecturerCode: "", //ไม่ได้กรอก
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

		ok, err := govalidator.ValidateStruct(lec0)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("LecturerCode is required"))
	})

	t.Run(`lecturer_code pattern is invalid`, func(t *testing.T) {
		lec1 := entity.Lecturer{
			LecturerCode: "D1234567", //pattern ผิด
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

		ok, err := govalidator.ValidateStruct(lec1)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("LecturerCode is invalid"))
	})

	t.Run(`lecturer_code is valid`, func(t *testing.T) {
		lec101 := entity.Lecturer{
			LecturerCode: "L1234567", //กรอกข้อมูลถูก
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

		ok, err := govalidator.ValidateStruct(lec101)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})
}

func TestFirstLastName(t *testing.T){

	g := NewGomegaWithT(t)

	postionID := uint(1)
	genderID := uint(1)
	majorID := uint(3)

	t.Run(`first_name is required`, func(t *testing.T) {
		lec2 := entity.Lecturer{
			LecturerCode: "L1234567",
			FirstName: "", //ผิดตรงนี้
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

		ok, err := govalidator.ValidateStruct(lec2)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("FirstName is required"))
	})


	t.Run(`last_name is required`, func(t *testing.T) {
		lec3 := entity.Lecturer{
			LecturerCode: "L1234567",
			FirstName: "Somsak",
			LastName: "",  //ผิดตรงนี้
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

		ok, err := govalidator.ValidateStruct(lec3)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("LastName is required"))
	})

}

func TestTime(t *testing.T){
	g := NewGomegaWithT(t)

	var birth_day, date_employed time.Time

	postionID := uint(1)
	genderID := uint(1)
	majorID := uint(3)

	t.Run(`birth_day is required`, func(t *testing.T) {
		lec101 := entity.Lecturer{
			LecturerCode: "L1234567", 
			FirstName: "Somsak",
			LastName: "Sri",
			BirthDay: birth_day,//ผิดตรงนี้
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

		ok, err := govalidator.ValidateStruct(lec101)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("BirthDay is required"))
	})

	t.Run(`date_employed is required`, func(t *testing.T) {
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
			DateEmployed: date_employed, //ผิดตรงนี้
			PositionID: &postionID,
			GenderID: &genderID,
			MajorID: &majorID,
		}

		ok, err := govalidator.ValidateStruct(lec101)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("DateEmployed is required"))
	})
}

func TestEmailPasswordNationalID(t *testing.T){
	g := NewGomegaWithT(t)

	postionID := uint(1)
	genderID := uint(1)
	majorID := uint(3)

	t.Run(`email is required`, func(t *testing.T) {
		email_lec := entity.Lecturer{
			LecturerCode: "L1234567",
			FirstName: "Somsak",
			LastName: "Sri",
			BirthDay: time.Now().AddDate(-30, 0, 0),
			Email: "", //ผิดตรงนี้
			Password:"L1234567",
			NationalID: "1234567890123",
			Phone: "0987654321",
			Profile: "test",
			DateEmployed: time.Now().AddDate(-30, 0, 0), 
			PositionID: &postionID,
			GenderID: &genderID,
			MajorID: &majorID,
		}

		ok, err := govalidator.ValidateStruct(email_lec)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Email is required"))
	})

	t.Run(`email is invalid`, func(t *testing.T) {
		email_lec := entity.Lecturer{
			LecturerCode: "L1234567",
			FirstName: "Somsak",
			LastName: "Sri",
			BirthDay: time.Now().AddDate(-30, 0, 0),
			Email: "l1245@gmail", //ผิดตรงนี้
			Password:"L1234567",
			NationalID: "1234567890123",
			Phone: "0987654321",
			Profile: "test",
			DateEmployed: time.Now().AddDate(-30, 0, 0), 
			PositionID: &postionID,
			GenderID: &genderID,
			MajorID: &majorID,
		}

		ok, err := govalidator.ValidateStruct(email_lec)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Email is invalid"))
	})

	t.Run(`email is not matches pattern`, func(t *testing.T) {
		email_lec := entity.Lecturer{
			LecturerCode: "L1234567",
			FirstName: "Somsak",
			LastName: "Sri",
			BirthDay: time.Now().AddDate(-30, 0, 0),
			Email: "l1245@sut.com", //ผิดตรงนี้
			Password:"L1234567",
			NationalID: "1234567890123",
			Phone: "0987654321",
			Profile: "test",
			DateEmployed: time.Now().AddDate(-30, 0, 0), 
			PositionID: &postionID,
			GenderID: &genderID,
			MajorID: &majorID,
		}

		ok, err := govalidator.ValidateStruct(email_lec)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Email must end with @gmail.com"))
	})

	t.Run(`password is required`, func(t *testing.T) {
		email_lec := entity.Lecturer{
			LecturerCode: "L1234567",
			FirstName: "Somsak",
			LastName: "Sri",
			BirthDay: time.Now().AddDate(-30, 0, 0),
			Email: "l12345678@gmail.com",
			Password:"",//ผิดตรงนี้
			NationalID: "1234567890123",
			Phone: "0987654321",
			Profile: "test",
			DateEmployed: time.Now().AddDate(-30, 0, 0), 
			PositionID: &postionID,
			GenderID: &genderID,
			MajorID: &majorID,
		}

		ok, err := govalidator.ValidateStruct(email_lec)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Password is required"))
	})

	t.Run(`NationalID is required`, func(t *testing.T) {
		email_lec := entity.Lecturer{
			LecturerCode: "L1234567",
			FirstName: "Somsak",
			LastName: "Sri",
			BirthDay: time.Now().AddDate(-30, 0, 0),
			Email: "l12345678@gmail.com",
			Password:"1234567890123",
			NationalID: "",//ผิดตรงนี้
			Phone: "0987654321",
			Profile: "test",
			DateEmployed: time.Now().AddDate(-30, 0, 0), 
			PositionID: &postionID,
			GenderID: &genderID,
			MajorID: &majorID,
		}

		ok, err := govalidator.ValidateStruct(email_lec)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("NationalID is required"))
	})

	t.Run(`NationalID is invalid pattern`, func(t *testing.T) {
		email_lec := entity.Lecturer{
			LecturerCode: "L1234567",
			FirstName: "Somsak",
			LastName: "Sri",
			BirthDay: time.Now().AddDate(-30, 0, 0),
			Email: "l12345678@gmail.com",
			Password:"1234567890123",
			NationalID: "123",//ผิดตรงนี้
			Phone: "0987654321",
			Profile: "test",
			DateEmployed: time.Now().AddDate(-30, 0, 0), 
			PositionID: &postionID,
			GenderID: &genderID,
			MajorID: &majorID,
		}

		ok, err := govalidator.ValidateStruct(email_lec)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("NationalID must be exactly 13 digits"))
	})
}

func TestPositionID(t *testing.T){
	g := NewGomegaWithT(t)

	var positionID *uint
	genderID := uint(1)
	majorID := uint(3)

	t.Run(`position_id is required`, func(t *testing.T) {
		position := entity.Lecturer{
			LecturerCode: "L1234567",
			FirstName: "Somsak",
			LastName: "Sri",
			BirthDay: time.Now().AddDate(-30, 0, 0),
			Email: "l12345678@gmail.com",
			Password:"123456789",
			NationalID: "1234567890123",
			Phone: "0987654321",
			Profile: "test",
			DateEmployed: time.Now().AddDate(-30, 0, 0), 
			PositionID: positionID, //ผิดตรงนี้(ไม่ได้กรอก)
			GenderID: &genderID, 
			MajorID: &majorID, 
		}

		ok, err := govalidator.ValidateStruct(position)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("PositionID is required"))
	})

}

func TestGenderID(t *testing.T){
	g := NewGomegaWithT(t)

	var genderID *uint
	positionID := uint(1)
	majorID := uint(3)

	t.Run(`gender_id is required`, func(t *testing.T) {
		gender := entity.Lecturer{
			LecturerCode: "L1234567",
			FirstName: "Somsak",
			LastName: "Sri",
			BirthDay: time.Now().AddDate(-30, 0, 0),
			Email: "l12345678@gmail.com",
			Password:"123456789",
			NationalID: "1234567890123",
			Phone: "0987654321",
			Profile: "test",
			DateEmployed: time.Now().AddDate(-30, 0, 0), 
			PositionID: &positionID,
			GenderID: genderID, //ผิดตรงนี้(ไม่ได้กรอก)
			MajorID: &majorID, 
		}

		ok, err := govalidator.ValidateStruct(gender)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("GenderID is required"))
	})

}

func TestMajorID(t *testing.T){
	g := NewGomegaWithT(t)

	var majorID *uint
	positionID := uint(1)
	genderID := uint(1)

	t.Run(`major_id is required`, func(t *testing.T) {
		major := entity.Lecturer{
			LecturerCode: "L1234567",
			FirstName: "Somsak",
			LastName: "Sri",
			BirthDay: time.Now().AddDate(-30, 0, 0),
			Email: "l12345678@gmail.com",
			Password:"123456789",
			NationalID: "1234567890123",
			Phone: "0987654321",
			Profile: "test",
			DateEmployed: time.Now().AddDate(-30, 0, 0), 
			PositionID: &positionID,
			GenderID: &genderID, 
			MajorID: majorID, //ผิดตรงนี้(ไม่ได้กรอก)
		}

		ok, err := govalidator.ValidateStruct(major)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("MajorID is required"))
	})

}

func TestPhoneProfile(t *testing.T){
	g := NewGomegaWithT(t)

	majorID := uint(1)
	positionID := uint(1)
	genderID := uint(1)

	t.Run(`Phone is required`, func(t *testing.T) {
		lec := entity.Lecturer{
			LecturerCode: "L1234567",
			FirstName: "Somsak",
			LastName: "Sri",
			BirthDay: time.Now().AddDate(-30, 0, 0),
			Email: "l12345678@gmail.com",
			Password:"123456789",
			NationalID: "1234567890123",
			Phone: "", //ผิดตรงนี้
			Profile: "test",
			DateEmployed: time.Now().AddDate(-30, 0, 0), 
			PositionID: &positionID,
			GenderID: &genderID, 
			MajorID: &majorID, 
		}

		ok, err := govalidator.ValidateStruct(lec)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Phone is required"))
	})

	t.Run(`Phone is invalid`, func(t *testing.T) {
		lec := entity.Lecturer{
			LecturerCode: "L1234567",
			FirstName: "Somsak",
			LastName: "Sri",
			BirthDay: time.Now().AddDate(-30, 0, 0),
			Email: "l12345678@gmail.com",
			Password:"123456789",
			NationalID: "1234567890123",
			Phone: "3124", //ผิดตรงนี้
			Profile: "test",
			DateEmployed: time.Now().AddDate(-30, 0, 0), 
			PositionID: &positionID,
			GenderID: &genderID, 
			MajorID: &majorID, 
		}

		ok, err := govalidator.ValidateStruct(lec)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Phone must start with 0 and be 10 digits"))
	})

	t.Run(`Profile is required`, func(t *testing.T) {
		lec := entity.Lecturer{
			LecturerCode: "L1234567",
			FirstName: "Somsak",
			LastName: "Sri",
			BirthDay: time.Now().AddDate(-30, 0, 0),
			Email: "l12345678@gmail.com",
			Password:"123456789",
			NationalID: "1234567890123",
			Phone: "0987654321", 
			Profile: "",//ผิดตรงนี้
			DateEmployed: time.Now().AddDate(-30, 0, 0), 
			PositionID: &positionID,
			GenderID: &genderID, 
			MajorID: &majorID, 
		}

		ok, err := govalidator.ValidateStruct(lec)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Profile is required"))
	})
}
