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
		defer subjectMenu(db)
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
