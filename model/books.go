package model

import (
	"encoding/csv"
	"fmt"
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

func ImportDataFromCSV(filePath string, db *gorm.DB) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error reading CSV: %v", err)
		}

		tahunStr := record[2]
		tahun, err := strconv.ParseUint(tahunStr, 10, 64)
		if err != nil {
			return err
		}

		stok, err := strconv.ParseUint(tahunStr, 10, 64)
		if err != nil {
			return fmt.Errorf("error parsing stok: %v", err)
		}

		book := &Books{
			ISBN:    record[0],
			Penulis: record[1],
			Tahun:   uint(tahun),
			Judul:   record[3],
			Gambar:  record[4],
			Stok:    uint(stok),
		}

		var existingBook Books
		result := db.First(&existingBook, "isbn = ?", book.ISBN)
		if result.RowsAffected > 0 {
			book.ID = existingBook.ID
			db.Save(book)
		} else {
			db.Create(book)
		}
	}

	return nil
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
