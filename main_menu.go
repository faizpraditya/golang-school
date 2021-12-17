package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func MainMenu(db *DBConn) {
	fmt.Println("Enigma School")
	fmt.Println("1. Data Mahasiswa")
	fmt.Println("2. Subject/Mata Kuliah")
	fmt.Println("3. Mahasiswa pilih Mata Kuliah")
	fmt.Println("4. Tampilkan data Mahasiswa dengan list Mata Kuliah")
	fmt.Println("5. Exit")
	fmt.Print("Pilih Menu: ")
	scanner.Scan()
	menu, _ := strconv.Atoi(scanner.Text())
	mainMenuController(menu, db)
}

func mainMenuController(menu int, db *DBConn) {
	switch menu {
	case 1:
		defer studentMenu(db)
		return
	case 2:
		defer subjectMenu(db)
		return
	case 3:
		defer chooseSubject(db)
		return
	case 4:
		defer studentSubjects(db)
		return
	case 5:
		fmt.Println("Exit from program")
		os.Exit(1)
	default:
		defer MainMenu(db)
		fmt.Println("Wrong input")
		return
	}
}

func chooseSubject(db *DBConn) {
	defer MainMenu(db)
	fmt.Println("Mahasiswa memilih Mata Kuliah")
	fmt.Print("Masukkan NIM Mahasiswa  : ")
	scanner.Scan()
	nim := scanner.Text()
	fmt.Print("Masukkan ID Mata Kuliah : ")
	scanner.Scan()
	idmk := scanner.Text()

	studentRepo := NewStudentRepo(db)
	studentRepo.OpenSubjectForExistingStudent(nim, idmk)
}

func studentSubjects(db *DBConn) {
	defer MainMenu(db)
	fmt.Println("Tampilkan data Mahasiswa dengan list Mata Kuliah")
	fmt.Print("Masukkan NIM Mahasiswa  : ")
	scanner.Scan()
	nim := scanner.Text()

	studentRepo := NewStudentRepo(db)
	showStudentSubject, _ := studentRepo.StudentSubjectList(nim)

	fmt.Println("Data Mahasiswa dengan Mata Kuliah yang diambil:")
	fmt.Println("NIM           : ", showStudentSubject.NIM)
	fmt.Println("Nama Lengkap  : ", showStudentSubject.Fullname)
	fmt.Println("Alamat        : ", showStudentSubject.Address)
	fmt.Println("Tanggal Lahir : ", showStudentSubject.BirthDate)
	fmt.Println("Mata Kuliah   : ")
	for i, mk := range showStudentSubject.Subjects {
		fmt.Println("(", i+1, ")")
		fmt.Println("Nama MK       : ", mk.Name)
		fmt.Println("Kredit (SKS)  : ", mk.Credits)
	}
}
