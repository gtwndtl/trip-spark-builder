package unit

import (
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"github.com/sut67/team09/entity"
	"testing"
)

func TestEducation(t *testing.T){
	g := NewGomegaWithT(t)


	t.Run(`education is valid`, func(t *testing.T) {
		educate := entity.Education{
			Instution: "Test",
		}

		ok, err := govalidator.ValidateStruct(educate)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})

	t.Run(`education is required`, func(t *testing.T) {
		educate := entity.Education{
			Instution: "", //ผิดตรงนี้
		}

		ok, err := govalidator.ValidateStruct(educate)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Instution is required"))
	})
}