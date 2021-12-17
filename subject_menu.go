package main

import (
	"fmt"
	"log"
	"strconv"
)

func subjectMenu(db *DBConn) {
	fmt.Println("Data Subject/Mata Kuliah")
	fmt.Println("1. Detail matakuliah")
	fmt.Println("2. Simpan data baru matakuliah")
	fmt.Println("3. Hapus data matakuliah")
	fmt.Println("4. Ubah data matakuliah")
	fmt.Println("5. Kembali ke menu utama")
	fmt.Print("Pilih Menu: ")
	scanner.Scan()
	menu, _ := strconv.Atoi(scanner.Text())
	subjectMenuController(menu, db)
}

func subjectMenuController(menu int, db *DBConn) {
	switch menu {
	case 1:
		defer MainMenu(db)
		featureDetailSubject(db)
		return
	case 2:
		defer MainMenu(db)
		featureCreateSubject(db)
		return
	case 3:
		defer MainMenu(db)
		featureDeleteSubject(db)
		return
	case 4:
		defer MainMenu(db)
		featureUpdateSubject(db)
		return
	case 5:
		defer MainMenu(db)
		return
	default:
		defer subjectMenu(db)
		fmt.Println("Wrong input")
		return
	}
}

func featureCreateSubject(db *DBConn) {
	var newSubject Subject
	fmt.Println("Simpan Data Baru Matakuliah")
	fmt.Print("ID               : ")
	scanner.Scan()
	newSubject.ID = scanner.Text()
	fmt.Print("Nama Matakuliah  : ")
	scanner.Scan()
	newSubject.Name = scanner.Text()
	fmt.Print("Kredit (SKS)     : ")
	scanner.Scan()
	newSubject.Credits, _ = strconv.Atoi(scanner.Text())

	SubjectRepo := NewSubjectRepo(db)
	SubjectRepo.Insert(newSubject)
	log.Println("Data Berhasil Ditambahkan")
}

func featureDetailSubject(db *DBConn) {
	fmt.Println("Detail Matakuliah")
	fmt.Print("Masukkan ID Matakuliah : ")
	scanner.Scan()
	subjectRepo := NewSubjectRepo(db)
	subject, _ := subjectRepo.FindByID(scanner.Text())
	fmt.Println("ID Matakuliah   : ", subject.ID)
	fmt.Println("Nama Matakuliah : ", subject.Name)
	fmt.Println("Kredit (SKS)    : ", subject.Credits)
}

func featureDeleteSubject(db *DBConn) {
	var deleteSubject Subject
	fmt.Println("Hapus Data Matakuliah")
	fmt.Print("Masukkan ID: ")
	scanner.Scan()
	deleteSubject.ID = scanner.Text()
	subjectRepo := NewSubjectRepo(db)
	subjectRepo.DeleteByID(deleteSubject)
	log.Println("Data Berhasil Dihapus")
}

func featureUpdateSubject(db *DBConn) {
	var updateSubject Subject
	fmt.Println("Ubah Data Matakuliah")
	fmt.Print("Masukkan ID dari data yang akan diubah: ")
	scanner.Scan()
	updateSubject.ID = scanner.Text()
	fmt.Println("Masukkan data yang akan diubah")
	fmt.Print("Nama Matakuliah : ")
	scanner.Scan()
	updateSubject.Name = scanner.Text()
	fmt.Print("Kredit (SKS)    : ")
	scanner.Scan()
	updateSubject.Credits, _ = strconv.Atoi(scanner.Text())

	subjectRepo := NewSubjectRepo(db)
	subjectRepo.UpdateByID(updateSubject)
	log.Println("Data Berhasil Diubah")
}
