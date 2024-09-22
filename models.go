package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func getProducts(db *sql.DB) ([]product, error) {
	query := "Select * from products"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	products := []product{}
	for rows.Next() {
		var p product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Quantity)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (p *product) getProduct(db *sql.DB) error {
	query := fmt.Sprintf("Select * from products Where id=%d", p.ID)
	log.Println("Query", query)
	row := db.QueryRow(query)
	log.Println(row)
	err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Quantity)

	if err != nil {
		return err
	}
	return nil
}

func (p *product) createProduct(db *sql.DB) error {
	query := fmt.Sprintf("Insert Into products(name,quantity,price) values('%v','%v','%v')", p.Name, p.Quantity, p.Price)
	result, err := db.Exec(query)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = int(id)
	return nil
}

func (p *product) updateProduct(db *sql.DB) error {
	query := fmt.Sprintf("update products set name='%v', quantity=%v, price=%v where id=%v", p.Name, p.Quantity, p.Price, p.ID)
	result, err := db.Exec(query)
	if err != nil {
		return err
	}
	rowsModified, err := result.RowsAffected()

	if rowsModified == 0 {
		return errors.New("no such row exist")
	}
	log.Println("Rows Affected", rowsModified)

	return err
}

func (p *product) deleteProduct(db *sql.DB) error {
	query := fmt.Sprintf("DELETE from products where id=%v", p.ID)
	res, err := db.Exec(query)
	if err != nil {
		return err
	}
	rowsModified, err := res.RowsAffected()

	if rowsModified == 0 {
		return errors.New("no such product exist")
	}
	return err
}
