package unit

import (
	//"fmt"
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"github.com/sut67/team09/entity"
)

func TestPaymentUserID(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`users_id is required`, func(t *testing.T){
		payment := entity.Payment{
			UsersID:       0,
			DormitoryID:   1,
			RoomID: 1,
            Wages:         200,
            WagesStudent:  320,
            StatusDor:     "approve",
            StatusStudent: "approve",
            TermStudent:   2565,
            YearStudent:   1,
            Credit:        16,
            Date:          time.Now(),
		}

		ok, err := govalidator.ValidateStruct(payment)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("UsersID is required"))
	})

}

func TestPaymentDorID(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`DormitoryID is required`, func(t *testing.T){
		payment := entity.Payment{
			UsersID:       1,
			DormitoryID:   0,
			RoomID: 1,
            Wages:         200,
            WagesStudent:  320,
            StatusDor:     "approve",
            StatusStudent: "approve",
            TermStudent:   2565,
            YearStudent:   1,
            Credit:        16,
            Date:          time.Now(),
		}

		ok, err := govalidator.ValidateStruct(payment)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("DormitoryID is required"))
	})
}

func TestPaymentWages(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`Wages is required`, func(t *testing.T){
		payment := entity.Payment{
			UsersID:       1,
			DormitoryID:   1,
			RoomID: 1,
            Wages:         0,
            WagesStudent:  320,
            StatusDor:     "approve",
            StatusStudent: "approve",
            TermStudent:   2565,
            YearStudent:   1,
            Credit:        16,
            Date:          time.Now(),
		}

		ok, err := govalidator.ValidateStruct(payment)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Wages is required"))
	})

}

func TestPaymentWagesStudent(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`WagesStudent is required`, func(t *testing.T){
		payment := entity.Payment{
			UsersID:       1,
			DormitoryID:   1,
			RoomID: 1,
            Wages:         320,
            WagesStudent:  0,
            StatusDor:     "approve",
            StatusStudent: "approve",
            TermStudent:   2565,
            YearStudent:   1,
            Credit:        16,
            Date:          time.Now(),
		}

		ok, err := govalidator.ValidateStruct(payment)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("WagesStudent is required"))
	})
}

func TestPaymentStatusDor(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run("StatusDor is required", func(t *testing.T){
		payment := entity.Payment{
			UsersID:       1,
			DormitoryID:   1,
			RoomID: 1,
            Wages:         320,
            WagesStudent:  1,
            StatusDor:     "",
            StatusStudent: "approve",
            TermStudent:   2565,
            YearStudent:   1,
            Credit:        16,
            Date:          time.Now(),
		}

		ok, err := govalidator.ValidateStruct(payment)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("StatusDor is required"))

		t.Run("StatusDor must be 'approve' or 'notapprove'", func(t *testing.T){
			payment := entity.Payment{
				UsersID:       1,
				DormitoryID:   1,
				RoomID: 1,
            	Wages:         320,
            	WagesStudent:  1,
            	StatusDor:     "Approve",
            	StatusStudent: "approve",
            	TermStudent:   2565,
            	YearStudent:   1,
           	 	Credit:        16,
            	Date:          time.Now(),
			}

			ok, err := govalidator.ValidateStruct(payment)

			g.Expect(ok).NotTo(BeTrue())

			g.Expect(err).NotTo(BeNil())

			g.Expect(err.Error()).To(Equal("StatusDor must be 'approve' or 'notapprove'"))
		})
	})
}

func TestPaymentStatusStudent(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run("StatusDor is required", func(t *testing.T){
		payment := entity.Payment{
			UsersID:       1,
			DormitoryID:   1,
			RoomID: 1,
            Wages:         320,
            WagesStudent:  1,
            StatusDor:     "approve",
            StatusStudent: "",
            TermStudent:   2565,
            YearStudent:   1,
            Credit:        16,
            Date:          time.Now(),
		}

		ok, err := govalidator.ValidateStruct(payment)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("StatusStudent is required"))

		t.Run("StatusStudent must be 'approve' or 'notapprove'", func(t *testing.T){
			payment := entity.Payment{
				UsersID:       1,
				DormitoryID:   1,
				RoomID: 1,
            	Wages:         320,
            	WagesStudent:  1,
            	StatusDor:     "approve",
            	StatusStudent: "Approve",
            	TermStudent:   2565,
            	YearStudent:   1,
           	 	Credit:        16,
            	Date:          time.Now(),
			}

			ok, err := govalidator.ValidateStruct(payment)

			g.Expect(ok).NotTo(BeTrue())

			g.Expect(err).NotTo(BeNil())

			g.Expect(err.Error()).To(Equal("StatusStudent must be 'approve' or 'notapprove'"))
		})
	})
}

func TestPaymentTermStudent(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("TermStudent is required", func(t *testing.T){
		payment := entity.Payment{
			UsersID:       1,
			DormitoryID:   1,
			RoomID: 1,
            Wages:         320,
            WagesStudent:  1,
            StatusDor:     "approve",
            StatusStudent: "approve",
            TermStudent:   0,
            YearStudent:   1,
           	Credit:        16,
            Date:          time.Now(),
		}

		ok, err := govalidator.ValidateStruct(payment)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("TermStudent is required"))
	})
}

func TestPaymentYearStudent(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run("YearStudent is required", func(t *testing.T){
		payment := entity.Payment{
			UsersID:       1,
			DormitoryID:   1,
			RoomID: 1,
            Wages:         320,
            WagesStudent:  1,
            StatusDor:     "approve",
            StatusStudent: "approve",
            TermStudent:   2565,
            YearStudent:   0,
           	Credit:        16,
            Date:          time.Now(),
		}

		ok, err := govalidator.ValidateStruct(payment)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("YearStudent is required"))
	})
}

func TestPaymentCredit(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run("Credit is required", func(t *testing.T){
		payment := entity.Payment{
			UsersID:       1,
			DormitoryID:   1,
			RoomID: 1,
            Wages:         320,
            WagesStudent:  1,
            StatusDor:     "approve",
            StatusStudent: "approve",
            TermStudent:   2565,
            YearStudent:   1,
           	Credit:        0,
            Date:          time.Now(),
		}

		ok, err := govalidator.ValidateStruct(payment)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Credit is required"))
	})
}

func TestPaymentValid(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`valid`, func(t *testing.T){
		payment := entity.Payment{
			UsersID:       1,
			DormitoryID:   1,
			RoomID: 1,
        	Wages:         200,
        	WagesStudent:  320,
        	StatusDor:     "approve",
        	StatusStudent: "approve",
        	TermStudent:   2565,
        	YearStudent:   1,
        	Credit:        16,
        	Date:          time.Now(),
		}

		ok, err := govalidator.ValidateStruct(payment)

		g.Expect(ok).To(BeTrue())

		g.Expect(err).To(BeNil())

	})
}