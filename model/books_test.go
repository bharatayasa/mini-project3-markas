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
		fmt.Println(".env not found")
	}

	config.OpenDb()
}

func TestCreateSuccess(test *testing.T) {
	Init()

	bookData := model.Books{
		ISBN:    "Kanginkauh221",
		Penulis: "Bharata",
		Tahun:   2021,
		Judul:   "aku menyukai orang yang bahkah dia tidak tahu namaku",
		Gambar:  "../UpGambar/Mangrove4.jpeg",
		Stok:    2000,
	}

	err := bookData.Create(config.Mysql.DB)
	assert.Nil(test, err)
}
