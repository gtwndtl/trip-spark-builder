package unit

import (
	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"github.com/sut67/team09/entity"

)

func TestDormitory(t *testing.T) {
	g := NewGomegaWithT(t)
	

	t.Run(`DormName is required`, func(t *testing.T) {
		dorm := entity.Dormitory{
			DormName:         "", // Invalid
			DormDescription: "A comfortable dormitory",
			DormType:        "Single",
			DormPic:         "example.jpg",
			Price:           10000,
		}

		ok, err := govalidator.ValidateStruct(dorm)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("DormName is required"))
	})

	t.Run(`DormDescription is required`, func(t *testing.T) {
		dorm := entity.Dormitory{
			DormName:         "Green Dormitory",
			DormDescription: "", // Invalid
			DormType:        "Single",
			DormPic:         "example.jpg",
			Price:           10000,
		}

		ok, err := govalidator.ValidateStruct(dorm)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("DormDescription is required"))
	})

	t.Run(`DormType is required`, func(t *testing.T) {
		dorm := entity.Dormitory{
			DormName:         "Green Dormitory",
			DormDescription: "A comfortable dormitory",
			DormType:        "", // Invalid
			DormPic:         "example.jpg",
			Price:           10000,
		}

		ok, err := govalidator.ValidateStruct(dorm)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("DormType is required"))
	})

	t.Run(`Dormitory is valid`, func(t *testing.T) {
		dorm := entity.Dormitory{
			DormName:         "Green Dormitory",
			DormDescription: "A comfortable dormitory",
			DormType:        "Single",
			DormPic:         "example.jpg",
			Price:           10000,
		}

		ok, err := govalidator.ValidateStruct(dorm)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})

	
	



	
}
