package model_test

import (
	"fmt"
	"testing"

	"github.com/bharatayasa/mini-project3-markas/config"
	"github.com/bharatayasa/mini-project3-markas/model"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("error, env not found")
	}

	config.OpenDb()
}

func TestCreateBook(test *testing.T) {
	Init()

	bookData := model.Books{
		ISBN:    "1231231234",
		Penulis: "Bharata",
		Tahun:   2021,
		Judul:   "hahah",
		Gambar:  "hahah",
		Stok:    2000,
	}

	err := bookData.CreateBook(config.Mysql.DB)
	assert.Nil(test, err)
}

func TestGetBookById(test *testing.T) {
	Init()

	bookData := model.Books{
		Model: model.Model{
			ID: 7,
		},
	}

	data, err := bookData.GetBookById(config.Mysql.DB)
	assert.Nil(test, err)
	assert.NotNil(test, data)

	fmt.Println(data)
}

func TestGetAllBooks(test *testing.T) {
	Init()

	bookData := model.Books{}

	data, err := bookData.GetAllBooks(config.Mysql.DB)
	assert.Nil(test, err)
	assert.NotNil(test, data)

	fmt.Println(data)
}

func TestUpdateBookByID(test *testing.T) {
	Init()

	bookData := model.Books{
		Model: model.Model{
			ID: 4,
		},
		ISBN:    "1231231234",
		Penulis: "Bharata ubah",
		Tahun:   2024,
		Judul:   "heheeh ubah",
		Gambar:  "hoho ubah",
		Stok:    2000,
	}

	err := bookData.UpdateOneByID(config.Mysql.DB)
	assert.Nil(test, err)
}

func TestDeleteByID(test *testing.T) {
	Init()

	bookData := model.Books{
		Model: model.Model{
			ID: 1,
		},
	}

	err := bookData.DeleteByID(config.Mysql.DB)
	assert.Nil(test, err)
}

func TestInsertCsv(test *testing.T) {
	Init()

	err := model.InsertCsvFromFile(config.Mysql.DB, "/Users/bharata/Desktop/markas/miniProject/mini-project3-markas/sample_books.csv")
	assert.Nil(test, err)
}
