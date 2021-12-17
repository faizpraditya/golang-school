package main

func main() {
	// NewDbConn()
	db := NewDbConn()
	defer db.Close()

	// Migrate table
	db.Migration(Student{}, Subject{})

	MainMenu(db)
}
