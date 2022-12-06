package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Inventory struct {
	db *sql.DB
}

const file string = "products.db"

func newDB() {
	db, err := os.Create(file)
	if err != nil {
		fmt.Print("Failed to creat a new database")
		return
	}
	defer db.Close()
}

const productTable string = `
  CREATE TABLE IF NOT EXISTS products (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "name" TEXT,
  "price" REAL, 
  );`

const jsonFile string = "products.json"

type ProductStruct struct {
	Product []Product
}

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	err := os.Remove(file)
	if err != nil {
		fmt.Println("Failed to delete existing database")
	}

	data, err := os.Open(jsonFile)
	if err != nil {
		fmt.Println("Failed to open json file")
		return
	}

	fmt.Println("File opened")

	defer data.Close()

	byteResult, err := os.ReadFile(jsonFile)
	if err != nil {
		fmt.Println("Failed to read json file")
		return
	}

	list := ProductStruct{}
	if err := json.Unmarshal(byteResult, &list.Product); err != nil {
		fmt.Printf("Failed to unmarshal json file")
		return
	}

	db, _ := sql.Open("sqlite3", file)

	newDB()
	createTable(db)

	var newItem []Product

	for i := 0; i < 5; i++ {
		if i < 5 {
			for _, item := range list.Product {
				newItem = append(newItem, item)
				addProduct(db, newItem)
				return
			}
		}
	}
	viewDB(db)
}

func createTable(db *sql.DB) {
	if _, err := db.Prepare(productTable); err != nil {
		fmt.Println("Failed to create table in database")
		return
	}

	fmt.Print("Table created")
}

func addProduct(db *sql.DB, newItem []Product) {
	insert := `INSERT INTO products(id, name, price) VALUES (null,?,?)`
	q, err := db.Prepare(insert)
	if err != nil {
		fmt.Print(("Failed to prepare new item"))
		return
	}

	_, err = q.Exec(newItem)
	if err != nil {
		fmt.Print("Failed to add a new item")
		return
	}
}

func viewDB(db *sql.DB) {
	records, err := db.Query("SELECT * FROM products")
	if err != nil {
		fmt.Print("Failed to Query the database")
	}

	defer records.Close()

	for records.Next() {
		var id int
		var name string
		var price int
		records.Scan(&id, &name, &price)
		fmt.Printf("Product: %d %s %d", id, name, price)
	}
}
