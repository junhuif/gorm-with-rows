package main

import (
	"errors"
	"fmt"
	"os"

	// Using postgres sql driver
	_ "github.com/lib/pq"

	"github.com/jinzhu/gorm"
)

var (
	// DB returns a gorm.DB interface, it is used to access to database
	DB *gorm.DB
)

type Product struct {
	gorm.Model
	Title       string `sql:"not null"`
	Description string `sql:"not null;size:2000"`
}

func init() {
	initDB()
	migrate()
}

func initDB() {
	var err error
	var db gorm.DB

	dbParams := os.Getenv("DB_PARAMS")
	if dbParams == "" {
		panic(errors.New("DB_PARAMS environment variable not set"))
	}

	db, err = gorm.Open("postgres", fmt.Sprintf(dbParams))
	if err == nil {
		DB = &db
	} else {
		panic(err)
	}
}

func migrate() {
	DB.DropTableIfExists(&Product{})
	DB.AutoMigrate(&Product{})
}

func loadProductsWithRows() (products []Product, err error) {
	products = []Product{}
	rows, err := DB.Model(&Product{}).Select("id, title, description, created_at, updated_at, deleted_at").Rows()
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		product := Product{}
		rows.Scan(&product.ID, &product.Title, &product.Description, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt)
		products = append(products, product)
	}
	return
}

func loadProductsWithFind() (products []Product, err error) {
	products = []Product{}

	err = DB.Find(&products).Error
	if err != nil {
		return
	}
	return
}
