package unit

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"github.com/sut67/team09/entity"
)

func TestClassroomValidation(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("RoomNumber is required and must be exactly 3 digits", func(t *testing.T) {
		classroom := entity.Classroom{
			RoomNumber:          "",
			Capacity:            30,
			Status:              "Available",
			Floor:               2,
			EndTimeClassroom:    time.Now(),
			BuildingID:          1,
			ConditionID:         1,
		}
		ok, err := govalidator.ValidateStruct(classroom)
		t.Logf("RoomNumber error: %v", err)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("RoomNumber is required"))
	})

	t.Run("Capacity must be between 1 and 2000", func(t *testing.T) {
		classroom := entity.Classroom{
			RoomNumber:          "101",
			Capacity:            2500,
			Status:              "Available",
			Floor:               2,
			EndTimeClassroom:    time.Now(),
			BuildingID:          1,
			ConditionID:         1,
		}
		ok, err := govalidator.ValidateStruct(classroom)
		t.Logf("Capacity range error: %v", err)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("Capacity must be between 1 and 2000"))
	})

	t.Run("Floor must be between 1 and 99", func(t *testing.T) {
		classroom := entity.Classroom{
			RoomNumber:          "101",
			Capacity:            30,
			Status:              "Available",
			Floor:               100,
			EndTimeClassroom:    time.Now(),
			BuildingID:          1,
			ConditionID:         1,
		}
		ok, err := govalidator.ValidateStruct(classroom)
		t.Logf("Floor range error: %v", err)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("Floor must be between 1 and 99"))
	})

	t.Run("EndTimeClassroom is required", func(t *testing.T) {
		classroom := entity.Classroom{
			RoomNumber:          "101",
			Capacity:            30,
			Status:              "Available",
			Floor:               2,
			EndTimeClassroom:    time.Time{}, // Empty time
			BuildingID:          1,
			ConditionID:         1,
		}
		ok, err := govalidator.ValidateStruct(classroom)
		t.Logf("EndTimeClassroom error: %v", err)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("EndTimeClassroom is required"))
	})

	t.Run("BuildingID is required", func(t *testing.T) {
		classroom := entity.Classroom{
			RoomNumber:          "101",
			Capacity:            30,
			Status:              "Available",
			Floor:               2,
			EndTimeClassroom:    time.Now(),
			BuildingID:          0,
			ConditionID:         1,
		}
		ok, err := govalidator.ValidateStruct(classroom)
		t.Logf("BuildingID error: %v", err)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("BuildingID is required"))
	})

	t.Run("ConditionID is required", func(t *testing.T) {
		classroom := entity.Classroom{
			RoomNumber:          "101",
			Capacity:            30,
			Status:              "Available",
			Floor:               2,
			EndTimeClassroom:    time.Now(),
			BuildingID:          1,
			ConditionID:         0,
		}
		ok, err := govalidator.ValidateStruct(classroom)
		t.Logf("ConditionID error: %v", err)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("ConditionID is required"))
	})
}