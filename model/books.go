package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Books struct {
	Model
	ID uint `gorm:"primarykey" json:"id"`
	// CreatedAt time.Time      `json:"created_at"`
	// UpdatedAt time.Time      `json:"updated_at"`
	// DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ISBN    string `gorm:"not null" json:"isbn"`
	Penulis string `gorm:"not null" json:"penulis"`
	Tahun   uint   `gorm:"not null" json:"tahun"`
	Judul   string `gorm:"not null" json:"judul"`
	Gambar  string `gorm:"not null" json:"gambar"`
	Stok    uint   `gorm:"not null" json:"stok"`
}

func (book *Books) ImportCsv(db *gorm.DB) error {
	err := db.
		Model(Books{}).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"isbn", "penulis", "tahun", "judul", "gambar", "stok"}),
		}).Create(&book).
		Error
	if err != nil {
		return err
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
