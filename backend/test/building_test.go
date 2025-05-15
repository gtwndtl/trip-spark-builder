package unit

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"github.com/sut67/team09/entity"
)

func TestBuildingValidation(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("BuildingCode is required", func(t *testing.T) {
		building := entity.Building{
			TotalFloors:     1,
			Location:        "SUT",
			BuildingName:    "Main Building",
			BuildingCode:    "",
			BuildingPicture: "Picture.jpg",
			ConditionID:     1,
		}
		ok, err := govalidator.ValidateStruct(building)
		t.Logf("BuildingCode error: %v", err)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("BuildingCode is required"))
	})

	t.Run("Location is required", func(t *testing.T) {
		building := entity.Building{
			TotalFloors:     1,
			Location:        "",
			BuildingName:    "Main Building",
			BuildingCode:    "B01",
			BuildingPicture: "Picture.jpg",
			ConditionID:     1,
		}
		ok, err := govalidator.ValidateStruct(building)
		t.Logf("Location error: %v", err)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("Location is required"))
	})

	t.Run("TotalFloors is required and must be between 1 and 99", func(t *testing.T) {
		building := entity.Building{
			TotalFloors:     -1, // Invalid value
			Location:        "SUT",
			BuildingName:    "Main Building",
			BuildingCode:    "B01",
			BuildingPicture: "Picture.jpg",
			ConditionID:     1,
		}
		ok, err := govalidator.ValidateStruct(building)
		t.Logf("TotalFloors error: %v", err)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("TotalFloors must be between 1 and 99"))
	})

	t.Run("BuildingName is required", func(t *testing.T) {
		building := entity.Building{
			TotalFloors:     1,
			Location:        "SUT",
			BuildingName:    "",
			BuildingCode:    "B01",
			BuildingPicture: "Picture.jpg",
			ConditionID:     1,
		}
		ok, err := govalidator.ValidateStruct(building)
		t.Logf("BuildingName error: %v", err)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("BuildingName is required"))
	})

	t.Run("BuildingPicture is required", func(t *testing.T) {
		building := entity.Building{
			TotalFloors:     1,
			Location:        "SUT",
			BuildingName:    "Main Building",
			BuildingCode:    "B01",
			BuildingPicture: "",
			ConditionID:     1,
		}
		ok, err := govalidator.ValidateStruct(building)
		t.Logf("BuildingPicture error: %v", err)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("BuildingPicture is required"))
	})

	t.Run("ConditionID is required", func(t *testing.T) {
		building := entity.Building{
			TotalFloors:     1,
			Location:        "SUT",
			BuildingName:    "Main Building",
			BuildingCode:    "B01",
			BuildingPicture: "Picture.jpg",
			ConditionID:     0, // Invalid condition ID
		}
		ok, err := govalidator.ValidateStruct(building)
		t.Logf("ConditionID error: %v", err)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("ConditionID is required"))
	})
}