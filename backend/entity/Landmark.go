package entity

import (
	"time"

	"gorm.io/gorm"
)

type Landmark struct {
	gorm.Model
	Name          string    
	Lat            float32    
	Lon             float32   
	City             string             
	Street	string   
	Time_open time.Time
	Time_close time.Time
	Total_people int
	Price float32
	Review int
}


