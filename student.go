package main

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID        string `gorm:"column:id;size:100;primaryKey"`
	NIM       string `gorm:"column:nim;size:10;not null"`
	Fullname  string `gorm:"column:fullname;size:50;not null"`
	Address   string `gorm:"column:address;not null"`
	BirthDate time.Time
	Subjects  []Subject `gorm:"many2many:tr_student_subjects"`
	gorm.Model
}
