package model

import (
	"time"
)

type Faculty struct {
	ID int `db:"id"`
	Name string `db:"name"`
}

type Group struct {
	ID int `db:"id"`
	Name string `db:"name"`
	FacultyID int `db:"faculty_id"`
}

type Student struct {
	ID int `db:"id"`
	Name string `db:"name"`
	Gender string `db:"gender"`
	BirthDate time.Time `db:"birth_date"`
	FacultyID int `db:"faculty_id"`
	GroupID int `db:"group_id"`
}
	
type Schedule struct {
	ID int `db:"id"`
	Subject string `db:"subject"`
	GroupID int `db:"group_id"`
	FacultyID int `db:"faculty_id"`
	TimeSlot time.Time `db:"time_slot"`
}