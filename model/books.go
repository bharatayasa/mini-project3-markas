package model

import (
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
