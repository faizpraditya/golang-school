package main

import "gorm.io/gorm"

type Subject struct {
	ID       string    `gorm:"column:id;size:100;primaryKey"`
	Name     string    `gorm:"column:name;size:50"`
	Credits  int       `gorm:"column:credits;"`
	Students []Student `gorm:"many2many:tr_student_subjects"`
	gorm.Model
}
