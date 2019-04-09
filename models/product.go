package models

import (
	"database/sql"
	"errors"
)

// Product represents product table
type Product struct {
	id    int     `json:"id"`
	name  string  `json:"name"`
	price float64 `json:"price"`
}

func (product *Product) getProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getProducts(db *sql.DB, start, count int) ([]Product, error) {
	rows, err := db.Query(
		"SELECT id, name, price FROM products LIMIT $1 OFFSET $2",
		count,
		start,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
}
