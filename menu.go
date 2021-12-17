package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func MainMenu(db *DBConn) {
	fmt.Println("Enigma School")
	fmt.Println("1. Data Mahasiswa")
	fmt.Println("2. Subject/Mata Kuliah")
	fmt.Println("3. Mahasiswa pilih Mata Kuliah")
	fmt.Println("4. Tampilkan data Mahasiswa dengan list Mata Kuliah yang diambil berdasarkan nim")
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
	case 3:
	case 4:
	case 5:
		fmt.Println("Exit from program")
		os.Exit(1)
	default:
		defer MainMenu(db)
		fmt.Println("Wrong input")
		return
	}
}

func studentMenu(db *DBConn) {
	fmt.Println("Data Mahasiswa")
	fmt.Println("1. Detail mahasiswa")
	fmt.Println("2. Simpan data baru mahasiswa")
	fmt.Println("3. Hapus data mahasiswa")
	fmt.Println("4. Ubah data mahasiswa")
	fmt.Println("5. Kembali ke menu utama")
	fmt.Print("Pilih Menu: ")
	scanner.Scan()
	menu, _ := strconv.Atoi(scanner.Text())
	studentMenuController(menu, db)
}

func studentMenuController(menu int, db *DBConn) {
	switch menu {
	case 1:
		defer MainMenu(db)
		featureDetailStudent(db)
		return
	case 2:
		defer MainMenu(db)
		featureCreateStudent(db)
		return
	case 3:
		defer MainMenu(db)
		featureDeleteStudent(db)
		return
	case 4:
		defer MainMenu(db)
		featureUpdateStudent(db)
		return
	case 5:
		defer MainMenu(db)
		return
	default:
		defer studentMenu(db)
		fmt.Println("Wrong input")
		return
	}
}

func featureCreateStudent(db *DBConn) {
	var newStudent Student
	fmt.Println("Simpan Data Baru Mahasiswa")
	fmt.Print("ID           : ")
	scanner.Scan()
	newStudent.ID = scanner.Text()
	fmt.Print("NIM          : ")
	scanner.Scan()
	newStudent.NIM = scanner.Text()
	fmt.Print("Nama Lengkap : ")
	scanner.Scan()
	newStudent.Fullname = scanner.Text()
	fmt.Print("Alamat       : ")
	scanner.Scan()
	newStudent.Address = scanner.Text()
	fmt.Println("Tanggal Lahir")
	fmt.Print("(YYYY-MM-DD) : ")
	scanner.Scan()
	newStudent.BirthDate = StringToDate(scanner.Text())

	studentRepo := NewStudentRepo(db)
	studentRepo.Insert(newStudent)
	log.Println("Data Berhasil Ditambahkan")
}

func featureDetailStudent(db *DBConn) {
	fmt.Println("Detail Mahasiswa")
	fmt.Print("Masukkan NIM: ")
	scanner.Scan()
	studentRepo := NewStudentRepo(db)
	student, _ := studentRepo.FindByNIM(scanner.Text())
	fmt.Println("NIM           : ", student.NIM)
	fmt.Println("Nama Lengkap  : ", student.Fullname)
	fmt.Println("Alamat        : ", student.Address)
	fmt.Println("Tanggal Lahir : ", student.BirthDate)
}

func featureDeleteStudent(db *DBConn) {
	var deleteStudent Student
	fmt.Println("Hapus Data Mahasiswa")
	fmt.Print("Masukkan NIM: ")
	scanner.Scan()
	deleteStudent.NIM = scanner.Text()
	studentRepo := NewStudentRepo(db)
	studentRepo.DeleteByNIM(deleteStudent)
	log.Println("Data Berhasil Dihapus")
}

func featureUpdateStudent(db *DBConn) {
	var updateStudent Student
	fmt.Println("Ubah Data Mahasiswa")
	fmt.Print("Masukkan NIM dari data yang akan diubah: ")
	scanner.Scan()
	updateStudent.NIM = scanner.Text()
	fmt.Println("Masukkan data yang akan diubah")
	fmt.Print("Nama Lengkap : ")
	scanner.Scan()
	updateStudent.Fullname = scanner.Text()
	fmt.Print("Alamat       : ")
	scanner.Scan()
	updateStudent.Address = scanner.Text()

	studentRepo := NewStudentRepo(db)
	studentRepo.UpdateByNIM(updateStudent)
}
