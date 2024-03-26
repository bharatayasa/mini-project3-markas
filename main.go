package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

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

	fmt.Print("Drag & drop file gambar disini: ")
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

	fmt.Print("Drag & drop file gambar disini: ")
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

func ImportCsv() {
	var direktori string

	fmt.Println("Import Data Buku dari File CSV")

	fmt.Print("drag & drop file .csv disini: ")
	_, err := fmt.Scanln(&direktori)
	if err != nil {
		fmt.Println("Terjadi Error : ", err)
	}

	file, err := openFile(direktori)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	csvChan, err := loadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	jmlGoroutine := 5

	var bookChanTemp []<-chan model.Books

	for i := 0; i < jmlGoroutine; i++ {
		bookChanTemp = append(bookChanTemp, processConverStruct(csvChan))
	}

	mergeCh := appendBooks(bookChanTemp...)

	var books []model.Books

	for ch := range mergeCh {
		books = append(books, ch)
	}

	ch := make(chan model.Books)
	wg := sync.WaitGroup{}

	for i := 0; i < jmlGoroutine; i++ {
		wg.Add(1)
		go simpanImportBuku(ch, &wg)
	}

	for _, book := range books {
		ch <- book
	}

	close(ch)
	wg.Wait()

	fmt.Println("Import Data Selesai!")
}

func openFile(path string) (*os.File, error) {
	return os.Open(path)
}

func loadFile(file *os.File) (<-chan []string, error) {
	bookChan := make(chan []string)
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return bookChan, err
	}

	go func() {
		for i, book := range records {
			if i == 0 {
				continue
			}
			bookChan <- book
		}

		close(bookChan)
	}()

	return bookChan, nil
}

func processConverStruct(csvChan <-chan []string) <-chan model.Books {
	booksChan := make(chan model.Books)

	go func() {
		for book := range csvChan {
			id, err := strconv.ParseUint(book[0], 10, 64)
			if err != nil {
				fmt.Println(err)
			}

			tahun, err := strconv.ParseUint(book[3], 10, 64)
			if err != nil {
				fmt.Println(err)
			}

			stok, err := strconv.ParseUint(book[6], 10, 64)
			if err != nil {
				fmt.Println(err)
			}

			booksChan <- model.Books{
				ID:      uint(id),
				ISBN:    book[1],
				Penulis: book[2],
				Tahun:   uint(tahun),
				Judul:   book[4],
				Gambar:  book[5],
				Stok:    uint(stok),
			}
		}

		close(booksChan)
	}()

	return booksChan
}

func appendBooks(bookChanMany ...<-chan model.Books) <-chan model.Books {
	wg := sync.WaitGroup{}

	mergedChan := make(chan model.Books)

	wg.Add(len(bookChanMany))
	for _, ch := range bookChanMany {
		go func(ch <-chan model.Books) {
			for books := range ch {
				mergedChan <- books
			}
			wg.Done()
		}(ch)
	}

	go func() {
		wg.Wait()
		close(mergedChan)
	}()

	return mergedChan
}

func simpanImportBuku(ch <-chan model.Books, wg *sync.WaitGroup) {
	for model := range ch {
		err := model.ImportCsv(config.Mysql.DB)
		if err != nil {
			fmt.Println(err)
		}
	}

	wg.Done()
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
		fmt.Println("5. import CSV file")
		fmt.Println("6. Keluar")
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
			ImportCsv()
		case 6:
			fmt.Println("Terima kasih telah menggunakan program ini.")
			os.Exit(0)
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
