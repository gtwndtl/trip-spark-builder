package unit

import (
	"fmt"
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"github.com/sut67/team09/entity"
)

func TestContactValidation(t *testing.T) {
    g := NewGomegaWithT(t)

    t.Run("Contact is required", func(t *testing.T) {
        report := entity.Report{
            Note:        "Valid Note",
            Contact:     "", // Empty contact
            Status:      "approve",
            Photo:       "",
            DateApprove: time.Now(),
            DateReport:  time.Now(),
            UsersID:     1,
            AdminID:     1,
            DormitoryID: 1,
            BooksID:     1,
            RoomID:      1,
        }

        ok, err := govalidator.ValidateStruct(report)

        // Debug actual error message if needed
        t.Logf("Validation error: %v", err)

        // Adjust case sensitivity and ensure validation fails as expected
        g.Expect(ok).To(BeFalse()) 
        g.Expect(err).To(HaveOccurred())
        g.Expect(err.Error()).To(ContainSubstring("Contact is required")) // Use lowercase 'contact'
    })

    t.Run("Contact length is invalid", func(t *testing.T) {
        report := entity.Report{
            Note:        "Valid Note",
            Contact:     "09672113166", // Invalid length
            Status:      "approve",
            Photo:       "",
            DateApprove: time.Now(),
            DateReport:  time.Now(),
            UsersID:     1,
            AdminID:     1,
            DormitoryID: 1,
            BooksID:     1,
            RoomID:      1,
        }

        ok, err := govalidator.ValidateStruct(report)

        // Debug actual error message if needed
        t.Logf("Validation error: %v", err)

        // Ensure validation fails and matches the exact error message
        g.Expect(ok).To(BeFalse()) 
        g.Expect(err).To(HaveOccurred())
        g.Expect(err.Error()).To(Equal(fmt.Sprintf("contact: %s does not validate as stringlength(10|10)", report.Contact))) 
    })
}

func TestNoteValidation(t *testing.T) {
    g := NewGomegaWithT(t)

    t.Run("Note is required", func(t *testing.T) {
        report := entity.Report{
            Note:        "",
            Contact:     "0967211316", // Empty contact
            Status:      "approve",
            Photo:       "",
            DateApprove: time.Now(),
            DateReport:  time.Now(),
            UsersID:     1,
            AdminID:     1,
            DormitoryID: 1,
            BooksID:     1,
            RoomID:      1,
        }

        ok, err := govalidator.ValidateStruct(report)

        // Debug actual error message if needed
        t.Logf("Validation error: %v", err)

        // Adjust case sensitivity and ensure validation fails as expected
        g.Expect(ok).To(BeFalse()) 
        g.Expect(err).To(HaveOccurred())
        g.Expect(err.Error()).To(ContainSubstring("Note is required")) // Use lowercase 'contact'
    })
}

func TestStatusValidation(t *testing.T) {
    g := NewGomegaWithT(t)

    t.Run("Status is required", func(t *testing.T) {
        report := entity.Report{
            Note:        "HEY",
            Contact:     "0967211316", // Contact is valid
            Status:      "",
            Photo:       "",
            DateApprove: time.Now(),
            DateReport:  time.Now(),
            UsersID:     1,
            AdminID:     1,
            DormitoryID: 1,
            BooksID:     1,
            RoomID:      1,
        }

        ok, err := govalidator.ValidateStruct(report)

        t.Logf("Validation error: %v", err)

        g.Expect(ok).To(BeFalse()) 
        g.Expect(err).To(HaveOccurred())
        g.Expect(err.Error()).To(ContainSubstring("Status is required"))
    })

    t.Run("Status is invalid", func(t *testing.T) {
        report := entity.Report{
            Note:        "HEY",
            Contact:     "0967211316", // Contact is valid
            Status:      "pending", // Invalid status
            Photo:       "",
            DateApprove: time.Now(),
            DateReport:  time.Now(),
            UsersID:     1,
            AdminID:     1,
            DormitoryID: 1,
            BooksID:     1,
            RoomID:      1,
        }

        ok, err := govalidator.ValidateStruct(report)

        t.Logf("Validation error: %v", err)

        g.Expect(ok).To(BeFalse()) 
        g.Expect(err).To(HaveOccurred())
        g.Expect(err.Error()).To(ContainSubstring("Status must be either 'approve' or 'not approve'"))
    })

    t.Run("Status is valid", func(t *testing.T) {
        report := entity.Report{
            Note:        "HEY",
            Contact:     "0967211316", // Contact is valid
            Status:      "approve", // Valid status
            Photo:       "",
            DateApprove: time.Now(),
            DateReport:  time.Now(),
            UsersID:     1,
            AdminID:     1,
            DormitoryID: 1,
            BooksID:     1,
            RoomID:      1,
        }

        ok, err := govalidator.ValidateStruct(report)

        t.Logf("Validation error: %v", err)

        g.Expect(ok).To(BeTrue()) 
        g.Expect(err).To(BeNil())
    })
}

func TestUsersIDValidation(t *testing.T) {
    g := NewGomegaWithT(t)

    t.Run("Users is required", func(t *testing.T) {
        report := entity.Report{
            Note:        "Hey",
            Contact:     "0967211316", // Valid contact
            Status:      "approve",
            Photo:       "",
            DateApprove: time.Now(),
            DateReport:  time.Now(),
            // UsersID is missing or invalid
            AdminID:     1,
            DormitoryID: 1,
            BooksID:     1,
            RoomID:      1,
        }

        ok, err := govalidator.ValidateStruct(report)

        // Debug actual error message if needed
        t.Logf("Validation error: %v", err)

        // Ensure validation fails for missing UsersID
        g.Expect(ok).To(BeFalse())
        g.Expect(err).To(HaveOccurred())
        g.Expect(err.Error()).To(ContainSubstring("Users is required"))
    })
}

func TestAdminIDValidation(t *testing.T) {
    g := NewGomegaWithT(t)

    t.Run("Admin is required", func(t *testing.T) {
        report := entity.Report{
            Note:        "Hey",
            Contact:     "0967211316", // Valid contact
            Status:      "approve",
            Photo:       "",
            DateApprove: time.Now(),
            DateReport:  time.Now(),
            UsersID:     1,
            DormitoryID: 1,
            BooksID:     1,
            RoomID:      1,
        }

        ok, err := govalidator.ValidateStruct(report)

        // Debug actual error message if needed
        t.Logf("Validation error: %v", err)

        // Ensure validation fails for missing UsersID
        g.Expect(ok).To(BeFalse())
        g.Expect(err).To(HaveOccurred())
        g.Expect(err.Error()).To(ContainSubstring("Admin is required"))
    })
}

func TestDormitoryIDValidation(t *testing.T) {
    g := NewGomegaWithT(t)

    t.Run("Dormitory is required", func(t *testing.T) {
        report := entity.Report{
            Note:        "Hey",
            Contact:     "0967211316", // Valid contact
            Status:      "approve",
            Photo:       "",
            DateApprove: time.Now(),
            DateReport:  time.Now(),
            UsersID:     1,
			AdminID:     1,
            BooksID:     1,
            RoomID:      1,
        }

        ok, err := govalidator.ValidateStruct(report)

        // Debug actual error message if needed
        t.Logf("Validation error: %v", err)

        // Ensure validation fails for missing UsersID
        g.Expect(ok).To(BeFalse())
        g.Expect(err).To(HaveOccurred())
        g.Expect(err.Error()).To(ContainSubstring("Dormitory is required"))
    })
}

func TestBooksIDValidation(t *testing.T) {
    g := NewGomegaWithT(t)

    t.Run("Books is required", func(t *testing.T) {
        report := entity.Report{
            Note:        "Hey",
            Contact:     "0967211316", // Valid contact
            Status:      "approve",
            Photo:       "",
            DateApprove: time.Now(),
            DateReport:  time.Now(),
            UsersID:     1,
			AdminID:     1,
			DormitoryID: 1,
            RoomID:      1,
        }

        ok, err := govalidator.ValidateStruct(report)

        // Debug actual error message if needed
        t.Logf("Validation error: %v", err)

        // Ensure validation fails for missing UsersID
        g.Expect(ok).To(BeFalse())
        g.Expect(err).To(HaveOccurred())
        g.Expect(err.Error()).To(ContainSubstring("Books is required"))
    })
}

func TestRoomIDValidation(t *testing.T) {
    g := NewGomegaWithT(t)

    t.Run("Room is required", func(t *testing.T) {
        report := entity.Report{
            Note:        "Hey",
            Contact:     "0967211316", // Valid contact
            Status:      "approve",
            Photo:       "",
            DateApprove: time.Now(),
            DateReport:  time.Now(),
            UsersID:     1,
			AdminID:     1,
			DormitoryID: 1,
			BooksID:     1,
        }

        ok, err := govalidator.ValidateStruct(report)

        // Debug actual error message if needed
        t.Logf("Validation error: %v", err)

        // Ensure validation fails for missing UsersID
        g.Expect(ok).To(BeFalse())
        g.Expect(err).To(HaveOccurred())
        g.Expect(err.Error()).To(ContainSubstring("Room is required"))
    })
}


func TestValid(t *testing.T) {
    g := NewGomegaWithT(t)

    t.Run(`valid`, func(t *testing.T) {
        report := entity.Report{
            Note:        "Hey",
            Contact:     "0967211316", // Valid contact
            Status:      "approve",    // Valid status
            Photo:       "",
            DateApprove: time.Now(),
            DateReport:  time.Now(),
            UsersID:     1,
            AdminID:     1,
            DormitoryID: 1,
            BooksID:     1,
            RoomID:      1,
        }

        // ตรวจสอบด้วย govalidator
        ok, err := govalidator.ValidateStruct(report)

        // ok ต้องเป็นค่า true แปลว่าการตรวจสอบผ่าน
        g.Expect(ok).To(BeTrue())
        // err ต้องเป็นค่า nil แปลว่าการตรวจสอบไม่พบข้อผิดพลาด
        g.Expect(err).To(BeNil())
    })
}
