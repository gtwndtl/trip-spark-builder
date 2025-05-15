package unit

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"github.com/sut67/team09/entity"
)

func TestEnrollmentValidation(t *testing.T) {
	g := NewGomegaWithT(t)

	// ทดสอบกรณีที่ไม่มี Status
	t.Run("missing Status", func(t *testing.T) {
		studentID := uint(1)
		courseID := uint(1)
		enrollment := entity.Enrollment{
			Status:    false,   // ให้ค่า Status เป็น false แทนที่จะไม่มีค่า
			StudentID: &studentID,
			CourseID:  &courseID,
		}

		// ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(enrollment)

		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

	})

	// ทดสอบกรณีที่ไม่มี StudentID
	t.Run("missing StudentID", func(t *testing.T) {
		courseID := uint(1)
		enrollment := entity.Enrollment{
			Status:    true,
			StudentID: nil, // ไม่มีการตั้ง StudentID
			CourseID:  &courseID,
		}

		// ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(enrollment)

		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

	})

	// ทดสอบกรณีที่ไม่มี CourseID
	t.Run("missing CourseID", func(t *testing.T) {
		studentID := uint(1)
		enrollment := entity.Enrollment{
			Status:    true,
			StudentID: &studentID,
			CourseID:  nil, // ไม่มีการตั้ง CourseID
		}

		// ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(enrollment)

		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

	})
}
