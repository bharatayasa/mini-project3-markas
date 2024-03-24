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

	println("id data baru: ", bookData.ID)
}
