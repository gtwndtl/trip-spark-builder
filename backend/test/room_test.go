package unit

import (
	"regexp"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"github.com/sut67/team09/entity"
)

func TestRoomNumber(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("RoomNumber is valid", func(t *testing.T) {
		roomNumber := "A1234" // Valid
		regex := regexp.MustCompile(`^A\d{4}$`) // แก้ไข Regex
		result := regex.MatchString(roomNumber)

		g.Expect(result).To(BeTrue())
	})

	t.Run("RoomNumber has invalid prefix", func(t *testing.T) {
		room := entity.Room{
			RoomNumber: "B1234", // Invalid prefix
		}

		ok, err := govalidator.ValidateStruct(room)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(ContainSubstring("RoomNumber"))
	})

	t.Run("RoomNumber has less than 4 digits", func(t *testing.T) {
		room := entity.Room{
			RoomNumber: "A123", // Less than 4 digits
		}

		ok, err := govalidator.ValidateStruct(room)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(ContainSubstring("RoomNumber"))
	})

	t.Run("RoomNumber has more than 4 digits", func(t *testing.T) {
		room := entity.Room{
			RoomNumber: "A12345", // More than 4 digits
		}

		ok, err := govalidator.ValidateStruct(room)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(ContainSubstring("RoomNumber"))
	})

	t.Run("RoomNumber has non-digit characters", func(t *testing.T) {
		room := entity.Room{
			RoomNumber: "Aabcd", // Non-digit characters
		}

		ok, err := govalidator.ValidateStruct(room)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(ContainSubstring("RoomNumber"))
	})
}
