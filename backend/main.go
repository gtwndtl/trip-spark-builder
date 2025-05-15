package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gtwndtl/trip-spark-builder/controllers"
	"github.com/gtwndtl/trip-spark-builder/entity"
)

func main() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("เชื่อมต่อฐานข้อมูลล้มเหลว")
	}

	// สร้างตารางในฐานข้อมูล
	err = db.AutoMigrate(
		&entity.User{},
		&entity.Accommodation{},
		&entity.Condition{},
		&entity.Landmark{},
		&entity.Restaurant{},
		&entity.Shortestpath{},
		&entity.Trips{},
	)
	if err != nil {
		panic("AutoMigrate ล้มเหลว: " + err.Error())
	}

	r := gin.Default()

	// สร้าง Controller สำหรับแต่ละ entity (สมมติว่าคุณสร้าง controller ตามนี้)
	userCtrl := controllers.NewUserController(db)
	accommodationCtrl := controllers.NewAccommodationController(db)
	conditionCtrl := controllers.NewConditionController(db)
	landmarkCtrl := controllers.NewLandmarkController(db)
	restaurantCtrl := controllers.NewRestaurantController(db)
	shortestpathCtrl := controllers.NewShortestpathController(db)
	tripsCtrl := controllers.NewTripsController(db)

	// Route User
	r.POST("/users", userCtrl.CreateUser)
	r.GET("/users", userCtrl.GetAll)
	r.GET("/users/:id", userCtrl.GetByID)
	r.PUT("/users/:id", userCtrl.Update)
	r.DELETE("/users/:id", userCtrl.Delete)

	// Route Accommodation
	r.POST("/accommodations", accommodationCtrl.CreateAccommodation)
	r.GET("/accommodations", accommodationCtrl.GetAll)
	r.GET("/accommodations/:id", accommodationCtrl.GetByID)
	r.PUT("/accommodations/:id", accommodationCtrl.Update)
	r.DELETE("/accommodations/:id", accommodationCtrl.Delete)

	// Route Condition
	r.POST("/conditions", conditionCtrl.CreateCondition)
	r.GET("/conditions", conditionCtrl.GetAll)
	r.GET("/conditions/:id", conditionCtrl.GetByID)
	r.PUT("/conditions/:id", conditionCtrl.Update)
	r.DELETE("/conditions/:id", conditionCtrl.Delete)

	// Route Landmark
	r.POST("/landmarks", landmarkCtrl.CreateLandmark)
	r.GET("/landmarks", landmarkCtrl.GetAll)
	r.GET("/landmarks/:id", landmarkCtrl.GetByID)
	r.PUT("/landmarks/:id", landmarkCtrl.Update)
	r.DELETE("/landmarks/:id", landmarkCtrl.Delete)

	// Route Restaurant
	r.POST("/restaurants", restaurantCtrl.CreateRestaurant)
	r.GET("/restaurants", restaurantCtrl.GetAll)
	r.GET("/restaurants/:id", restaurantCtrl.GetByID)
	r.PUT("/restaurants/:id", restaurantCtrl.Update)
	r.DELETE("/restaurants/:id", restaurantCtrl.Delete)

	// Route Shortestpath
	r.POST("/shortestpaths", shortestpathCtrl.CreateShortestpath)
	r.GET("/shortestpaths", shortestpathCtrl.GetAll)
	r.GET("/shortestpaths/:id", shortestpathCtrl.GetByID)
	r.PUT("/shortestpaths/:id", shortestpathCtrl.Update)
	r.DELETE("/shortestpaths/:id", shortestpathCtrl.Delete)

	// Route Trips
	r.POST("/trips", tripsCtrl.CreateTrip)
	r.GET("/trips", tripsCtrl.GetAll)
	r.GET("/trips/:id", tripsCtrl.GetByID)
	r.PUT("/trips/:id", tripsCtrl.Update)
	r.DELETE("/trips/:id", tripsCtrl.Delete)

	r.Run() // รันเซิร์ฟเวอร์ที่พอร์ต 8080
}
