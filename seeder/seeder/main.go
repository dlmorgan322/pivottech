package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Inventory struct {
	db *sql.DB
}

const file string = "products.db"

type Products struct {
	Products []Product
}

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

const products string = `
  CREATE TABLE IF NOT EXISTS products (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT,
  price REAL,
  );`

func NewProducts() (*Inventory, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if _, err := db.Exec(products); err != nil {
		return nil, err
	}

	return &Inventory{
		db: db,
	}, nil

}

func (p *Product) Unmarshal(data []byte) error {
	var product Product
	if err := json.Unmarshal(data, &p); err != nil {
		return err
	}
	*p = Product(product)
	return nil

}

func main() {
	err := os.Remove(file)
	if err != nil {
		fmt.Println(err)
	}

	data, err := os.Open("products.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File opened")

	defer data.Close()

	byteResult, err := ioutil.ReadFile("products.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer data.Close()

	decoder := json.NewDecoder(data)

	var product Product
	json.Unmarshal(byteResult, &product)
	for {
		err = decoder.Decode(&product)
		if err != nil {
			fmt.Print(err)
		}

		json.Unmarshal(byteResult, &product)
	}

}
