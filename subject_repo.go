package main

type SubjectRepo struct {
	BaseRepo
}

func (sur *SubjectRepo) Insert(newSubject Subject) error {
	result := sur.conn.Db.Create(&newSubject)
	return sur.HandleError(result)
}

type DetailSubject struct {
	ID      string
	Name    string
	Credits string
}

func (sur *SubjectRepo) FindByID(id string) (DetailSubject, error) {
	var subject DetailSubject
	result := sur.conn.Db.Model(&Subject{}).Where("id = ?", id).Scan(&subject)
	return subject, sur.HandleError(result)
}

func (sur *SubjectRepo) DeleteByID(subject Subject) error {
	result := sur.conn.Db.Delete(&subject, "id = ?", subject.ID)
	return sur.HandleError(result)
}

func (sur *SubjectRepo) UpdateByID(updateCustomerInfo Subject) (Subject, error) {
	result := sur.conn.Db.Model(&updateCustomerInfo).Where("id = ?", updateCustomerInfo.ID).Updates(updateCustomerInfo)
	err := sur.HandleError(result)
	if err != nil {
		return Subject{}, err
	}
	return updateCustomerInfo, nil
}

func NewSubjectRepo(conn *DBConn) *SubjectRepo {
	return &SubjectRepo{BaseRepo{conn: conn}}
}
