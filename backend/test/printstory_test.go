package unit

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"github.com/sut67/team09/entity"
)

func TestPrintStoryCode(t *testing.T) {

	g := NewGomegaWithT(t)


	t.Run(`PrintStoryCode is required`, func(t *testing.T) {
		pdf := entity.PrintStory{
			PrintStoryCode: "", // ผิดตรงนี้ ส่งแบบไม่มีข้อมูล code
			DocumentPath: "test",
			CreateAt: time.Now(),
			RequestID: 1,  

		}

		ok, err := govalidator.ValidateStruct(pdf)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		// g.Expect(err.Error()).To(Equal("PrintStoryCode is required"))
		g.Expect(err.Error()).To(ContainSubstring("PrintStoryCode is required"))

	})

	t.Run(`PrintStoryCode pattern is invalid`, func(t *testing.T) {
		pdf := entity.PrintStory{

			PrintStoryCode: "R9073423781111000", //pattern ผิด 
			DocumentPath: "test",
			CreateAt: time.Now(),
			RequestID: 1,  

		}

		ok, err := govalidator.ValidateStruct(pdf)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(ContainSubstring("PrintStoryCode is invalid"))
		
	})

	t.Run(`PrintStoryCode is invalid`, func(t *testing.T) {
		pdf := entity.PrintStory{

			PrintStoryCode: "R907342378", //กรณีถูก
			DocumentPath: "test",
			CreateAt: time.Now(),
			RequestID: 1,  

		}

		ok, _ := govalidator.ValidateStruct(pdf)

		g.Expect(ok).To(BeTrue())


	})
}


func TestDocumentPath(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run(`Document_Path is required`, func(t *testing.T) {
		pdf := entity.PrintStory{

			PrintStoryCode: "R907342378",
			DocumentPath: "",  // ผิดตรงนี้ ส่งแบบไม่มีข้อมูล file
			CreateAt: time.Now(),
			RequestID: 1,  

		}

		ok, err := govalidator.ValidateStruct(pdf)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(ContainSubstring("DocumentPath is required"))

	})

	t.Run(`Document_Path is correct`, func(t *testing.T) {
		pdf := entity.PrintStory{

			PrintStoryCode: "R907342378", 
			DocumentPath: "test",  //กรณีถูก
			CreateAt: time.Now(),
			RequestID: 2,  

		}

		ok, err := govalidator.ValidateStruct(pdf)

		if (err != nil) {
			print("err",err.Error())
		}

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})
}

