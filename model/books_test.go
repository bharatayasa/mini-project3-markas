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

func TestAll(test *testing.T) {
	Init()
	TestCreateBook(test)
	TestGetBookById(test)
	TestGetAllBooks(test)
	TestUpdateBookByID(test)
	TestDeleteByID(test)
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
		ISBN:    "1231231234",
		Penulis: "Bharata",
		Tahun:   2021,
		Judul:   "hahah",
		Gambar:  "hahah",
		Stok:    2000,
	}

	err := bookData.CreateBook(config.Mysql.DB)
	assert.Nil(test, err)

	GetBookById := model.Books{
		Model: model.Model{
			ID: bookData.ID,
		},
	}

	data, err := GetBookById.GetBookById(config.Mysql.DB)
	assert.Nil(test, err)
	assert.NotNil(test, data)

	fmt.Println(data)
}

func TestGetAllBooks(test *testing.T) {
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

	bookData = model.Books{}

	data, err := bookData.GetAllBooks(config.Mysql.DB)
	assert.Nil(test, err)
	assert.NotNil(test, data)

	fmt.Println(data)
}

func TestUpdateBookByID(test *testing.T) {
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

	bookDataUpdate := model.Books{
		Model: model.Model{
			ID: bookData.ID,
		},
		ISBN:    "1231231234",
		Penulis: "Bharata ubah",
		Tahun:   2024,
		Judul:   "heheeh ubah",
		Gambar:  "hoho ubah",
		Stok:    2000,
	}

	err = bookDataUpdate.UpdateOneByID(config.Mysql.DB)
	assert.Nil(test, err)
}

func TestDeleteByID(test *testing.T) {
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

	DeleteByID := model.Books{
		Model: model.Model{
			ID: bookData.ID,
		},
	}

	err = DeleteByID.DeleteByID(config.Mysql.DB)
	assert.Nil(test, err)
}
