package entity

import (
	"time"

	"gorm.io/gorm"
)

//Employee stores information about the employee
type Employee struct {
	gorm.Model
	EventID   int64     `json:"event_id"`
	FirstName string    `json:"first_name" ,gorm:"size:32"`
	LastName  string    `json:"last_name" ,gorm:"size:64"`
	Birthday  time.Time `json:"birth_day"`
	Gender    Gender    `json:"gender" ,gorm:"size:10"`
}

type Gender string

const (
	MALE   Gender = "male"
	FEMALE Gender = "female"
)

// @dev : this is considered that we have these 2 types of gender and also the first
//and last name should be seperated
