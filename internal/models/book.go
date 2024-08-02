package models

import (
	"database/sql"
	"time"
)

type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	ReleaseYear int       `json:"release_year"`
	Price       int       `json:"price"`
	TotalPage   int       `json:"total_page"`
	Thickness   string    `json:"thickness"`
	CategoryID  int       `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
	ModifiedAt  time.Time `json:"modified_at"`
	ModifiedBy  string    `json:"modified_by"`
}

func GetAllBooks() ([]Book, error) {
	rows, err := db.Query("SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []Book{}
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func GetBookByID(id int) (Book, error) {
	var book Book
	err := db.QueryRow("SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by FROM books WHERE id = $1", id).Scan(
		&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return book, err
		}
		return book, err
	}
	return book, nil
}

func CreateBook(book Book) (Book, error) {
	err := db.QueryRow(
		"INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, category_id, created_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by",
		book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID, book.CreatedBy).Scan(
		&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy)
	if err != nil {
		return book, err
	}
	return book, nil
}

func DeleteBookByID(id int) error {
	result, err := db.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
