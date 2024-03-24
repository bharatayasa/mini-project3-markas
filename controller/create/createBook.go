package create

import (
	"github.com/bharatayasa/mini-project3-markas/model"
	"gorm.io/gorm"
)

func CreateNewBook(book *model.Books, db *gorm.DB) error {
	err := book.CreateBook(db)
	if err != nil {
		return err
	}

	return nil
}
