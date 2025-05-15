package unit

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"github.com/sut67/team09/entity"
)

func TestActivity(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run(`ActivityName is required`, func(t *testing.T) {
		act := entity.Activity{
			ActivityName: "", // ผิดตรงนี้
			Description:   "สัมมนาเพื่อเรียนรู้เทคโนโลยีปัญญาประดิษฐ์",
			ActivityDate:  time.Date(2025, time.January, 25, 0, 0, 0, 0, time.UTC), 
			StartTime:     time.Date(1, time.January, 1, 10, 0, 0, 0, time.UTC),
			EndTime:       time.Date(1, time.January, 1, 17, 0, 0, 0, time.UTC), 
			Location:      "ห้องสัมมนา A",
			Organizer:     "ทีมงาน A",
			MaxParticipants: 100,
			ActivityPic:   "", 
			AdminID:       1,
			DormitoryID:   1,
			StatusActivityID: 1,
		}


		ok, err := govalidator.ValidateStruct(act)

		
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		
		g.Expect(err.Error()).To(ContainSubstring("ActivityName is required"))
	})


	t.Run(`Description is required`, func(t *testing.T) {
		act := entity.Activity{
			ActivityName: "งานสัมมนาเทคโนโลยี AI",
			Description:   "", // ผิดตรงนี้
			ActivityDate:  time.Date(2025, time.January, 25, 0, 0, 0, 0, time.UTC), 
			StartTime:     time.Date(1, time.January, 1, 10, 0, 0, 0, time.UTC),
			EndTime:       time.Date(1, time.January, 1, 17, 0, 0, 0, time.UTC), 
			Location:      "ห้องสัมมนา A",
			Organizer:     "ทีมงาน A",
			MaxParticipants: 100,
			ActivityPic:   "", 
			AdminID:       1,
			DormitoryID:   1,
			StatusActivityID: 1,
		}


		ok, err := govalidator.ValidateStruct(act)

		
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		
		g.Expect(err.Error()).To(ContainSubstring("Description is required"))
	})

	t.Run(`Location is required`, func(t *testing.T) {
		act := entity.Activity{
			ActivityName: "งานสัมมนาเทคโนโลยี AI",
			Description:   "สัมมนาเพื่อเรียนรู้เทคโนโลยีปัญญาประดิษฐ์", 
			ActivityDate:  time.Date(2025, time.January, 25, 0, 0, 0, 0, time.UTC), 
			StartTime:     time.Date(1, time.January, 1, 10, 0, 0, 0, time.UTC),
			EndTime:       time.Date(1, time.January, 1, 17, 0, 0, 0, time.UTC), 
			Location:      "",// ผิดตรงนี้
			Organizer:     "ทีมงาน A",
			MaxParticipants: 100,
			ActivityPic:   "", 
			AdminID:       1,
			DormitoryID:   1,
			StatusActivityID: 1,
		}


		ok, err := govalidator.ValidateStruct(act)

		
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		
		g.Expect(err.Error()).To(ContainSubstring("Location is required"))
	})

	t.Run(`Organizer is required`, func(t *testing.T) {
		act := entity.Activity{
			ActivityName: "งานสัมมนาเทคโนโลยี AI",
			Description:   "สัมมนาเพื่อเรียนรู้เทคโนโลยีปัญญาประดิษฐ์", 
			ActivityDate:  time.Date(2025, time.January, 25, 0, 0, 0, 0, time.UTC), 
			StartTime:     time.Date(1, time.January, 1, 10, 0, 0, 0, time.UTC),
			EndTime:       time.Date(1, time.January, 1, 17, 0, 0, 0, time.UTC), 
			Location:      "ห้องสัมมนา A",
			Organizer:     "", // ผิดตรงนี้
			MaxParticipants: 100,
			ActivityPic:   "", 
			AdminID:       1,
			DormitoryID:   1,
			StatusActivityID: 1,
		}


		ok, err := govalidator.ValidateStruct(act)

		
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		
		g.Expect(err.Error()).To(ContainSubstring("Organizer is required"))
	})



	t.Run(`MaxParticipants is required`, func(t *testing.T) {
		act := entity.Activity{
			ActivityName: "งานสัมมนาเทคโนโลยี AI",
			Description:   "สัมมนาเพื่อเรียนรู้เทคโนโลยีปัญญาประดิษฐ์", 
			ActivityDate:  time.Date(2025, time.January, 25, 0, 0, 0, 0, time.UTC), 
			StartTime:     time.Date(1, time.January, 1, 10, 0, 0, 0, time.UTC),
			EndTime:       time.Date(1, time.January, 1, 17, 0, 0, 0, time.UTC), 
			Location:      "",
			Organizer:     "ทีมงาน A",
			MaxParticipants: 0,// ผิดตรงนี้
			ActivityPic:   "", 
			AdminID:       1,
			DormitoryID:   1,
			StatusActivityID: 1,
		}


		ok, err := govalidator.ValidateStruct(act)

		
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		
		g.Expect(err.Error()).To(ContainSubstring("MaxParticipants is required"))
	})

}
