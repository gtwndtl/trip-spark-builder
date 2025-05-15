package unit

import (
	"testing"
	"time"

	. "github.com/onsi/gomega"
	"github.com/sut67/team09/entity"
)

func TestScheduleValidation(t *testing.T) {
	g := NewGomegaWithT(t)

	// ทดสอบกรณีที่ StartTime หลังจาก EndTime
	t.Run("StartTime after EndTime", func(t *testing.T) {
		schedule := entity.Schedule{
			StartTime: time.Now().Add(2 * time.Hour),
			EndTime:   time.Now(),
		}

		// ตรวจสอบ custom validation
		err := schedule.Validate()

		// ตรวจสอบข้อผิดพลาดว่าเกิดจาก EndTime ที่มาก่อน StartTime
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("EndTime must be after StartTime"))
	})
}
