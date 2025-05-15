package unit

import (
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"github.com/sut67/team09/entity"
	"testing"
	"time"
)

func TestThesis(t *testing.T){
	g := NewGomegaWithT(t)

	var publicdate time.Time

	t.Run(`thesis is valid`, func(t *testing.T) {
		thesis := entity.Thesis{
			Title: "Test",
			PublicationDate: time.Now().AddDate(-30, 0, 0),
			URL: "test.com",
		}

		ok, err := govalidator.ValidateStruct(thesis)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})

	t.Run(`title is required`, func(t *testing.T) {
		thesis := entity.Thesis{
			Title: "", //ผิดตรงนี้
			PublicationDate: time.Now().AddDate(-30, 0, 0),
			URL: "test.com",
		}

		ok, err := govalidator.ValidateStruct(thesis)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Title is required"))
	})

	t.Run(`publication_date is required`, func(t *testing.T) {
		thesis := entity.Thesis{
			Title: "Test", 
			PublicationDate: publicdate, //ผิดตรงนี้
			URL: "test.com",
		}

		ok, err := govalidator.ValidateStruct(thesis)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("PublicationDate is required"))
	})

	t.Run(`title is required`, func(t *testing.T) {
		thesis := entity.Thesis{
			Title: "Test", 
			PublicationDate: time.Now().AddDate(-30, 0, 0),
			URL: "", //ผิดตรงนี้
		}

		ok, err := govalidator.ValidateStruct(thesis)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("URL is required"))
	})
}