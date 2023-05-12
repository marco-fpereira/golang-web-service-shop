package repository

import (
	"database/sql"
	"reflect"
	"web-service/app/config"
	"web-service/app/model"
	"web-service/app/util"
)

type ProductTypes interface {
	model.Product | string
}

func GetAllProducts() []model.Product {
	var sqlProducts *sql.Rows = executeQuery("SELECT * FROM products")

	products := []model.Product{}

	for sqlProducts.Next() {
		p := readObjectFromDatabase(sqlProducts)
		products = append(products, p)
	}
	return products
}

func GetProductById(productId string) model.Product {
	sqlProducts := executeQueryWithParameters("SELECT * FROM products WHERE id=$1", productId)

	var product model.Product
	for sqlProducts.Next() {
		product = readObjectFromDatabase(sqlProducts)
	}
	return product
}

func CreateNewProduct(name string, description string, price float64, amount int) {
	executeQueryWithParameters(
		"INSERT INTO products(name, description, price, amount) VALUES ($1, $2, $3, $4)",
		model.Product{Name: name, Description: description, Price: price, Amount: amount},
	)
}

func DeleteProduct(productId string) {
	executeQueryWithParameters("DELETE FROM products WHERE id=$1", productId)
}

func executeQuery(query string) *sql.Rows {
	db := config.DatabaseConnector()

	sqlResult, err := db.Query(query)
	util.ErrorHandler(err)
	defer db.Close()

	return sqlResult
}

func executeQueryWithParameters(requestQuery string, parameters interface{}) *sql.Rows {
	var result *sql.Rows
	var err error

	db := config.DatabaseConnector()

	switch reflect.TypeOf(parameters).String() {
	case "string":
		id, _ := parameters.(string)
		result, err = db.Query(requestQuery, id)
	case "model.Product":
		products, _ := parameters.(model.Product)
		result, err = db.Query(requestQuery, products.Name, products.Description, products.Price, products.Amount)
	}

	util.ErrorHandler(err)
	defer db.Close()

	return result
}

func readObjectFromDatabase(sqlProducts *sql.Rows) model.Product {
	var name, description string
	var price float64
	var amount, id int
	err := sqlProducts.Scan(&id, &name, &description, &price, &amount)
	util.ErrorHandler(err)

	p := model.Product{}
	p.Id = id
	p.Name = name
	p.Description = description
	p.Price = price
	p.Amount = amount

	return p
}
