package model

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"

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

func (book *Books) CreateBook(db *gorm.DB) error {
	err := db.Model(Books{}).Create(&book).Error
	if err != nil {
		return err
	}

	return nil
}

func (book *Books) GetBookById(db *gorm.DB) (Books, error) {
	res := Books{}

	err := db.Model(Books{}).Where("id = ?", book.Model.ID).Take(&res).Error
	if err != nil {
		return Books{}, err
	}

	return res, nil
}

func (book *Books) GetAllBooks(db *gorm.DB) ([]Books, error) {
	var books []Books

	err := db.Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (book *Books) UpdateOneByID(db *gorm.DB) error {
	err := db.
		Model(Books{}).
		Select("isbn", "penulis", "tahun", "judul", "gambar", "stok").
		Where("id = ?", book.Model.ID).
		Updates(map[string]any{
			"isbn":    book.ISBN,
			"penulis": book.Penulis,
			"tahun":   book.Tahun,
			"judul":   book.Judul,
			"gambar":  book.Gambar,
			"stok":    book.Stok,
		}).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (book *Books) DeleteByID(db *gorm.DB) error {
	err := db.
		Model(Books{}).
		Where("id = ?", book.Model.ID).
		Delete(&book).
		Error

	if err != nil {
		return err
	}

	return nil
}

// InsertCsvFromFile import data from csv file to database
func InsertCsvFromFile(db *gorm.DB, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	reader.LazyQuotes = true

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if len(line) < 6 {
			return errors.New("invalid CSV format: insufficient fields")
		}

		stok, err := strconv.ParseUint(line[5], 10, 32)
		if err != nil {
			return err
		}

		year, err := strconv.ParseUint(line[2], 10, 32)
		if err != nil {
			return err
		}

		book := Books{
			ISBN:    line[0],
			Penulis: line[1],
			Tahun:   uint(year),
			Judul:   line[3],
			Gambar:  line[4],
			Stok:    uint(stok),
		}

		err = db.Create(&book).Error
		if err != nil {
			return err
		}
	}

	return nil
}

// func InsertCsvFromFile(db *gorm.DB, filePath string) error {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	reader := csv.NewReader(file)
// 	reader.Comma = ';'
// 	reader.LazyQuotes = true

// 	for {
// 		line, err := reader.Read()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			return err
// 		}

// 		stok, err := strconv.ParseUint(line[5], 10, 32)
// 		if err != nil {
// 			return err
// 		}

// 		year, err := strconv.ParseUint(line[2], 10, 32)
// 		if err != nil {
// 			return err
// 		}

// 		book := Books{
// 			ISBN:    line[0],
// 			Penulis: line[1],
// 			Tahun:   uint(year),
// 			Judul:   line[3],
// 			Gambar:  line[4],
// 			Stok:    uint(stok),
// 		}

// 		err = db.Create(&book).Error
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
