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
	return nil, errors.New("Not implemented")
}
