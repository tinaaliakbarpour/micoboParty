package entity

import (
	"time"

	"gorm.io/gorm"
)

//Event stores information about each event in micoboParty
type Event struct {
	gorm.Model
	Name       string    `json:"name" ,gorm:"name, size:128,notnull"`
	LaunchDate time.Time `json:"launch_date" ,gorm:"launch_date"`
}
