package unit

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"github.com/sut67/team09/entity"
)

func TestCourseValidation(t *testing.T) {
	g := NewGomegaWithT(t)

	// กำหนดค่า ID ที่จำเป็น
	lecturerID := uint(1)

	// ทดสอบกรณีที่ CourseCode เป็นค่าว่าง
	t.Run("invalid course code", func(t *testing.T) {
		invalidCourse := entity.Course{
			CourseCode: "", // ค่าว่าง
			CourseName: "Computer Science",
			Credit:     3,
			Group:      "1",
			MaxSeat:    50,
			Status:     true,
			LecturerID: &lecturerID,
		}
		ok, err := govalidator.ValidateStruct(invalidCourse)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil())
	})

	// ทดสอบกรณีที่ CourseName เป็นค่าว่าง
	t.Run("invalid course name", func(t *testing.T) {
		invalidCourse := entity.Course{
			CourseCode: "CS101",
			CourseName: "", // ค่าว่าง
			Credit:     3,
			Group:      "1",
			MaxSeat:    50,
			Status:     true,
			LecturerID: &lecturerID,
		}
		ok, err := govalidator.ValidateStruct(invalidCourse)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil()) // คาดหวังข้อผิดพลาด 1 ตัว (CourseName)
	})

	// ทดสอบกรณีที่ Credit เป็นค่าน้อยกว่า 1
	t.Run("invalid credit", func(t *testing.T) {
		invalidCourse := entity.Course{
			CourseCode: "CS101",
			CourseName: "Computer Science",
			Credit:     0, 
			Group:      "1",
			MaxSeat:    50,
			Status:     true,
			LecturerID: &lecturerID,
		}
		ok, err := govalidator.ValidateStruct(invalidCourse)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil()) // คาดหวังข้อผิดพลาด 1 ตัว (Credit)
	})

	// ทดสอบกรณีที่ MaxSeat เป็นค่าติดลบ
	t.Run("invalid max seat", func(t *testing.T) {
		invalidCourse := entity.Course{
			CourseCode: "CS101",
			CourseName: "Computer Science",
			Credit:     3,
			Group:      "1",
			MaxSeat:    0,
			Status:     true,
			LecturerID: &lecturerID,
		}
		ok, err := govalidator.ValidateStruct(invalidCourse)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil()) // คาดหวังข้อผิดพลาด 1 ตัว (MaxSeat)
	})

	// ทดสอบกรณีที่ Status เป็นค่า null หรือไม่กำหนด
	t.Run("invalid status", func(t *testing.T) {
		invalidCourse := entity.Course{
			CourseCode: "CS101",
			CourseName: "Computer Science",
			Credit:     3,
			Group:      "1",
			MaxSeat:    50,
			Status:     false, // สถานะเป็น false
			LecturerID: &lecturerID,
		}
		ok, err := govalidator.ValidateStruct(invalidCourse)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil()) // ไม่คาดหวังข้อผิดพลาดเนื่องจาก Status ไม่จำเป็นต้องเป็น true เสมอไป
	})

	// ทดสอบกรณีที่ LecturerID เป็น nil
	t.Run("invalid lecturer id", func(t *testing.T) {
		invalidCourse := entity.Course{
			CourseCode: "CS101",
			CourseName: "Computer Science",
			Credit:     3,
			Group:      "1",
			MaxSeat:    50,
			Status:     true,
			LecturerID: nil, // ไม่กำหนด LecturerID
		}
		ok, err := govalidator.ValidateStruct(invalidCourse)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil()) // คาดหวังข้อผิดพลาด 1 ตัว (LecturerID)
	})
}
