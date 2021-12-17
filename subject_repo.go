package main

type SubjectRepo struct {
	BaseRepo
}

func (cr *SubjectRepo) Insert(newSubject Subject) error {
	result := cr.conn.Db.Create(&newSubject)
	return cr.HandleError(result)
}

func newSubjectRepo(conn *DBConn) *SubjectRepo {
	return &SubjectRepo{
		BaseRepo{
			conn: conn,
		},
	}
}
