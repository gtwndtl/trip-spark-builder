package entity

import (
	"time"

	"gorm.io/gorm"
)

type Accommodation struct {
	gorm.Model
	Name          string    
	Lat            float32    
	Lon             float32              
	Time_open time.Time
	Time_close time.Time
	Total_people int
	Price float32
	Review int
	City             string             
	Street	string   
}
