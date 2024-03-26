package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bharatayasa/mini-project3-markas/config"
	"github.com/bharatayasa/mini-project3-markas/controller/create"
	uploadfile "github.com/bharatayasa/mini-project3-markas/controller/uploadFile"
	"github.com/bharatayasa/mini-project3-markas/model"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("env not found")
	}
}

func TambahBuku(db *gorm.DB) {
	book := model.Books{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan ISBN buku: ")
	book.ISBN, _ = reader.ReadString('\n')
	book.ISBN = strings.TrimSpace(book.ISBN)

	fmt.Print("Masukkan Penulis buku: ")
	book.Penulis, _ = reader.ReadString('\n')
	book.Penulis = strings.TrimSpace(book.Penulis)

	fmt.Print("Masukan Tahun terbit buku: ")
	_, err := fmt.Scanln(&book.Tahun)
	if err != nil {
		log.Fatalf("Error parsing input: %v", err)
	}

	fmt.Print("Masukkan Judul buku: ")
	book.Judul, _ = reader.ReadString('\n')
	book.Judul = strings.TrimSpace(book.Judul)

	fmt.Print("Masukkan Gambar buku: ")
	fileName, err := uploadfile.UploadFile()
	if err != nil {
		log.Fatalf("Error uploading file: %v", err)
	}
	book.Gambar = fileName

	fmt.Print("Masukkan Jumlah Stok buku: ")
	_, err = fmt.Scanln(&book.Stok)
	if err != nil {
		log.Fatalf("Error parsing input: %v", err)
	}

	err = create.CreateNewBook(&book, db)
	if err != nil {
		log.Fatalf("Error adding book: %v", err)
	}

	fmt.Println("Buku berhasil ditambahkan!")
}

func TampilBuku(db *gorm.DB) {
	var books []model.Books
	db.Find(&books)

	if len(books) == 0 {
		fmt.Println("Tidak ada buku yang tersedia")
		return
	}

	fmt.Println("Daftar buku:")
	fmt.Println("==================")
	for _, book := range books {
		fmt.Printf("ID: %d\n", book.ID)
		fmt.Printf("ISBN: %s\n", book.ISBN)
		fmt.Printf("Penulis: %s\n", book.Penulis)
		fmt.Printf("Tahun Terbit: %d\n", book.Tahun)
		fmt.Printf("Judul: %s\n", book.Judul)
		fmt.Printf("Gambar: %s\n", book.Gambar)
		fmt.Printf("Stok: %d\n\n", book.Stok)
	}
}

func EditBook(db *gorm.DB) {
	var book model.Books

	fmt.Print("Masukkan ID buku yang ingin diedit: ")
	_, err := fmt.Scanln(&book.Model.ID)
	if err != nil {
		log.Fatalf("Error parsing input: %v", err)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan ISBN baru: ")
	book.ISBN, _ = reader.ReadString('\n')
	book.ISBN = strings.TrimSpace(book.ISBN)

	fmt.Print("Masukkan penulis baru: ")
	book.Penulis, _ = reader.ReadString('\n')
	book.Penulis = strings.TrimSpace(book.Penulis)

	fmt.Print("Masukan Tahun terbit baru: ")
	_, err = fmt.Scanln(&book.Tahun)
	if err != nil {
		log.Fatalf("Error parsing input: %v", err)
	}

	fmt.Print("Masukkan Judul baru: ")
	book.Judul, _ = reader.ReadString('\n')
	book.Judul = strings.TrimSpace(book.Judul)

	fmt.Print("Masukkan Gambar baru: ")
	fileName, err := uploadfile.UploadFile()
	if err != nil {
		log.Fatalf("Error uploading file: %v", err)
	}
	book.Gambar = fileName

	fmt.Print("Masukkan Jumlah Stok baru: ")
	_, err = fmt.Scanln(&book.Stok)
	if err != nil {
		log.Fatalf("Error parsing input: %v", err)
	}

	err = book.UpdateOneByID(db)
	if err != nil {
		log.Fatalf("Error editing book: %v", err)
	}

	fmt.Println("Buku berhasil diedit!")
}

func HapusBuku() {
	var book model.Books

	fmt.Print("Masukkan ID buku yang ingin dihapus: ")
	_, err := fmt.Scanln(&book.Model.ID)
	if err != nil {
		log.Fatalf("Error parsing input: %v", err)
	}

	db, err := config.OpenDb()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	err = book.DeleteByID(db)
	if err != nil {
		log.Fatalf("Error deleting book: %v", err)
	}

	fmt.Println("Buku berhasil dihapus!")
}

func PrintOneBook() {
	fmt.Println("print one book")
}

func PrintAllBooks() {
	fmt.Println("print all book")
}

func ImportCsv(db *gorm.DB) {
	fmt.Println("insert csv file")
}

func main() {
	Init()
	db, err := config.OpenDb()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	for {
		fmt.Print("---------------------------------------------")
		fmt.Println("\nSistem manajemen buku")
		fmt.Println("---------------------------------------------")
		fmt.Println("1. Tambah Buku")
		fmt.Println("2. Tampil Buku")
		fmt.Println("3. Edit Buku")
		fmt.Println("4. Hapus Buku")
		fmt.Println("5. Print Buku")
		fmt.Println("6. import CSV file")
		fmt.Println("7. Keluar")
		fmt.Println("---------------------------------------------")

		var choice int
		fmt.Print("Pilih menu opsi: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			TambahBuku(db)
		case 2:
			TampilBuku(db)
		case 3:
			EditBook(db)
		case 4:
			HapusBuku()
		case 5:
			PrintOneBook()
		case 6:
			ImportCsv(db)
		case 7:
			fmt.Println("Terima kasih telah menggunakan program ini.")
			os.Exit(0)
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
