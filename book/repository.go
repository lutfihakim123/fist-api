package book

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(book Book) (Book, error)
}

type repository struct {
	db *gorm.DB
}

// make connection from main to get DB
func NewRepository(db *gorm.DB) *repository  {
	return &repository{db}
}


// find all data
func (r *repository) FindAll()([]Book, error)  {
	var books []Book
	err := r.db.Find(&books).Error
	return books, err
}

// find data by id 
func (r *repository) FindById(ID int)(Book, error)  {
	var book Book
	err := r.db.Find(&book, ID).Error
	return book, err
}

// create data 
func (r *repository) Create(book Book)(Book, error)  {
	err := r.db.Create(&book).Error
	return book, err
}





