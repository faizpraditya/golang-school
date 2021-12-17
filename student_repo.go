package main

import (
	"fmt"
)

type StudentRepo struct {
	BaseRepo
}

func (sr *StudentRepo) Insert(newStudent Student) (Student, error) {
	// Cek validasi manual
	fmt.Println("insert")
	result := sr.conn.Db.Create(&newStudent)
	// Cek validasi manual
	fmt.Println(newStudent.ID)
	return newStudent, sr.HandleError(result)
}

type DetailStudent struct {
	NIM       string
	Fullname  string
	Address   string
	BirthDate string
}

func (sr *StudentRepo) FindByNIM(nim string) (DetailStudent, error) {
	var student DetailStudent
	result := sr.conn.Db.Model(&Student{}).Where("nim = ?", nim).Scan(&student)
	return student, sr.HandleError(result)
}

func (sr *StudentRepo) DeleteByNIM(student Student) error {
	result := sr.conn.Db.Delete(&student, "nim = ?", student.NIM)
	return sr.HandleError(result)
}

func (sr *StudentRepo) UpdateByNIM(updateCustomerInfo Student) (Student, error) {
	result := sr.conn.Db.Model(&updateCustomerInfo).Where("nim = ?", updateCustomerInfo.NIM).Updates(updateCustomerInfo)
	err := sr.HandleError(result)
	if err != nil {
		return Student{}, err
	}
	return updateCustomerInfo, nil
}

func (sr *StudentRepo) UpdateFirstName(updateCustomerInfo Student) (Student, error) {
	result := sr.conn.Db.Model(&updateCustomerInfo).Select("first_name").Updates(updateCustomerInfo)
	err := sr.HandleError(result)
	if err != nil {
		return Student{}, err
	}
	return updateCustomerInfo, nil
}

func (sr *StudentRepo) OpenSubjectForExistingStudent(studentWithProduct Student) error {
	result := sr.conn.Db.Model(&studentWithProduct).Updates(studentWithProduct)
	return sr.HandleError(result)
}

func NewStudentRepo(conn *DBConn) *StudentRepo {
	return &StudentRepo{BaseRepo{conn: conn}}
}
