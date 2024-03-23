package model

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Books struct {
	Model
	ISBN    string `gorm:"not null" json:"isbn"`
	Penulis string `gorm:"not null" json:"penulis"`
	Tahun   uint   `gorm:"not null" json:"tahun"`
	Judul   string `gorm:"not null" json:"judul"`
	Gambar  string `gorm:"not null" json:"gambar"`
	Stok    uint   `gorm:"not null" json:"stok"`
}

func (book *Books) UploadGambar(file io.Reader) error {
	imageName := GenerateUniqueImageName()

	err := SaveImage(file, imageName)
	if err != nil {
		return err
	}

	book.Gambar = imageName

	return nil
}

func SaveImage(file io.Reader, imageName string) error {
	f, err := os.Create("../gambar" + imageName)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		return err
	}

	return nil
}

func GenerateUniqueImageName() string {
	timestamp := fmt.Sprintf("%d", time.Now().UnixNano())
	randomNum := rand.Intn(1000)
	return timestamp + "_" + strconv.Itoa(randomNum) + ".jpeg"
}

func (book *Books) Create(db *gorm.DB) error {
	err := db.Create(&book).Error
	if err != nil {
		return err
	}
	return nil
}
