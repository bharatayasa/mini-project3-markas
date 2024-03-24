package main

import (
	"fmt"
	"log"
	"os"

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
	fmt.Print("ISBN: ")
	fmt.Scanln(&book.ISBN)
	fmt.Print("Penulis: ")
	fmt.Scanln(&book.Penulis)
	fmt.Print("Tahun: ")
	fmt.Scanln(&book.Tahun)
	fmt.Print("Judul: ")
	fmt.Scanln(&book.Judul)
	fmt.Print("Gambar: ")
	fmt.Scanln(&book.Gambar)
	fmt.Print("Stok: ")
	fmt.Scanln(&book.Stok)

	err := CreateNewBook(&book, db)
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
		fmt.Println("6. Keluar")
		fmt.Println("---------------------------------------------")

		var choice int
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			TambahBuku(db)
		case 6:
			fmt.Println("Terima kasih telah menggunakan program ini.")
			os.Exit(0)
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
