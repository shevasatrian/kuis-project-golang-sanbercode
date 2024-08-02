package models

import (
	"database/sql"
	"time"
)

type Category struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
}

func GetAllCategories() ([]Category, error) {
	rows, err := db.Query("SELECT id, name, created_at, created_by, modified_at, modified_by FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := []Category{}
	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func GetCategoryByID(id int) (Category, error) {
	var category Category
	err := db.QueryRow("SELECT id, name, created_at, created_by, modified_at, modified_by FROM categories WHERE id = $1", id).Scan(
		&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return category, err
		}
		return category, err
	}
	return category, nil
}

func CreateCategory(category Category) (Category, error) {
	err := db.QueryRow(
		"INSERT INTO categories (name, created_by) VALUES ($1, $2) RETURNING id, name, created_at, created_by, modified_at, modified_by",
		category.Name, category.CreatedBy).Scan(
		&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy)
	if err != nil {
		return category, err
	}
	return category, nil
}

func DeleteCategoryByID(id int) error {
	result, err := db.Exec("DELETE FROM categories WHERE id = $1", id)
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

func GetBooksByCategoryID(id int) ([]Book, error) {
	rows, err := db.Query("SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by FROM books WHERE category_id = $1", id)
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
