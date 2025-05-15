package config

import (
	"fmt"
	"time"

	"github.com/sut67/team09/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func ConnectionDB() {
	database, err := gorm.Open(sqlite.Open("SE_University_Team09.db?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connected database")
	db = database
}

func SetupDatabase() {
	db.AutoMigrate(
		//ระบบจัดการข้อมูลอาจารย์
		&entity.Accommodation{},
		&entity.Condition{},
		&entity.Landmark{},
		&entity.Restaurant{},
		&entity.Shortestpath{},
		&entity.Trips{},
		&entity.User{},
	)

	//hashedPassword, _ := HashPassword("123456")

	// for _, Dormitory := range Dormitory {
	// 	db.FirstOrCreate(&Dormitory)
	// }

	//  db.FirstOrCreate(Dormitory, &entity.Dormitory)
	// 	{

	// 	// Price: 5,

	//  })
	Accommodation := &entity.Accommodation{
		Lat:          13.75919,
		Lon:          100.53657,
		Time_open:    time.Date(2023, 10, 1, 8, 0, 0, 0, time.UTC),
		Time_close:   time.Date(2023, 10, 1, 8, 0, 0, 0, time.UTC),
		Total_people: 5,
		Price:        10000.0,
		Review:       5,
		City:         "ราชเทวี",
		Street:       "45 ซอยศรีอยุธยา 12",
	}

	db.FirstOrCreate(Accommodation, &entity.Accommodation{
		Name: "โรงแรมทรูสยามบางกอก",
	})

	Landmark := &entity.Landmark{
		Lat:          13.75389,
		Lon:          100.50841,
		City:         "ป้อมปราบศัตรูพ่าย",
		Street:       "344 ถนนจักรพรรดิพงษ์",
		Time_open:    time.Date(2023, 10, 1, 8, 0, 0, 0, time.UTC),
		Time_close:   time.Date(2023, 10, 1, 8, 0, 0, 0, time.UTC),
		Total_people: 15,
		Price:        500.0,
		Review:       3,
	}

	db.FirstOrCreate(Landmark, &entity.Landmark{
		Name: "วัดสระเกศ ราชวรมหาวิหาร (วัดภูเขาทอง)",
	})

	Restaurant := &entity.Restaurant{
		Lat:          13.84104,
		Lon:          100.61726,
		Time_open:    time.Date(2023, 10, 1, 8, 0, 0, 0, time.UTC),
		Time_close:   time.Date(2023, 10, 1, 8, 0, 0, 0, time.UTC),
		Total_people: 10,
		Price:        300.0,
		Review:       4,
		City:         "ลาดพร้าว",
		Street:       "จรเข้บัว",
	}

	db.FirstOrCreate(Restaurant, &entity.Restaurant{
		Name: "วรชาติ",
	})

}
