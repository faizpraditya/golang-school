package main

type StudentRepo struct {
	BaseRepo
}

func (sr *StudentRepo) Insert(newStudent Student) (Student, error) {
	result := sr.conn.Db.Create(&newStudent)
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

func (sr *StudentRepo) FindByNIMFull(nim string) (Student, error) {
	var student Student
	result := sr.conn.Db.Model(&Student{}).Where("nim = ?", nim).Scan(&student)
	return student, sr.HandleError(result)
}

func (sr *StudentRepo) OpenSubjectForExistingStudent(nim string, idmk string) error {
	dataStudent, _ := sr.FindByNIMFull(nim)
	studentWithSubject := Student{ID: dataStudent.ID, Subjects: []Subject{{ID: idmk}}}
	result := sr.conn.Db.Model(&studentWithSubject).Updates(studentWithSubject)
	return sr.HandleError(result)
}

func (sr *StudentRepo) StudentSubjectList(nim string) (Student, error) {
	var student Student
	result := sr.conn.Db.Preload("Subjects").Where("nim = ?", nim).Find(&student)
	return student, sr.HandleError(result)
}

func NewStudentRepo(conn *DBConn) *StudentRepo {
	return &StudentRepo{BaseRepo{conn: conn}}
}
