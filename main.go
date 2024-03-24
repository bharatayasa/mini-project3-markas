package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/bharatayasa/mini-project3-markas/config"
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
	fileName, err := UploadFile()
	if err != nil {
		log.Fatalf("Error uploading file: %v", err)
	}
	book.Gambar = fileName

	fmt.Print("Masukkan Jumlah Stok buku: ")
	_, err = fmt.Scanln(&book.Stok)
	if err != nil {
		log.Fatalf("Error parsing input: %v", err)
	}

	err = CreateNewBook(&book, db)
	if err != nil {
		log.Fatalf("Error adding book: %v", err)
	}

	fmt.Println("Buku berhasil ditambahkan!")
}

func CreateNewBook(book *model.Books, db *gorm.DB) error {
	err := book.CreateBook(db)
	if err != nil {
		return err
	}

	return nil
}

func UploadFile() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan path file gambar: ")
	path, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	path = strings.TrimSpace(path)

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return "", err
	}

	fileBytes := make([]byte, fileInfo.Size())
	_, err = file.Read(fileBytes)
	if err != nil {
		return "", err
	}

	fileType := http.DetectContentType(fileBytes)
	if fileType != "image/jpeg" && fileType != "image/png" {
		return "", fmt.Errorf("file format not supported")
	}

	fileName := filepath.Base(file.Name())

	dst, err := os.Create("./images/" + fileName)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return fileName, nil
}

func TampilBuku() {
	fmt.Println("Tampil buku")
}

func EditBook() {
	fmt.Println("edit buku")
}

func HapusBuku() {
	fmt.Println("hapus buku")
}

func PrintOneBook() {
	fmt.Println("print one book")
}

func PrintAllBooks() {
	fmt.Println("print all book")
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
		fmt.Println("6. Keluar")
		fmt.Println("---------------------------------------------")

		var choice int
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			TambahBuku(db)
		case 2:
			TampilBuku()
		case 3:
			EditBook()
		case 4:
			HapusBuku()
		case 5:
			PrintOneBook()
		case 6:
			fmt.Println("Terima kasih telah menggunakan program ini.")
			os.Exit(0)
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
