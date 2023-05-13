package repository

import (
	"database/sql"
	"web-service/app/config"
	"web-service/app/model"
	"web-service/app/util"
)

func GetAllProducts() []model.Product {
	var sqlProducts *sql.Rows = executeQuery("SELECT * FROM products order by id asc")

	products := []model.Product{}

	for sqlProducts.Next() {
		p := readObjectFromDatabase(sqlProducts)
		products = append(products, p)
	}
	return products
}

func GetProductById(productId string) model.Product {
	sqlProducts := executeQuery("SELECT * FROM products WHERE id=$1", productId)

	var product model.Product
	for sqlProducts.Next() {
		product = readObjectFromDatabase(sqlProducts)
	}
	return product
}

func CreateNewProduct(name string, description string, price float64, amount int) {
	db := config.DatabaseConnector()
	_, err := db.Query(
		"INSERT INTO products(name, description, price, amount) VALUES ($1, $2, $3, $4)",
		name, description, price, amount,
	)

	util.ErrorHandler(err)
	defer db.Close()
}

func UpdateProduct(id int, name string, description string, price float64, amount int) {
	db := config.DatabaseConnector()
	_, err := db.Query(
		"UPDATE products SET name=$1, description=$2, price=$3, amount=$4 WHERE id=$5",
		name, description, price, amount, id,
	)

	util.ErrorHandler(err)
	defer db.Close()
}

func DeleteProduct(productId string) {
	executeQuery("DELETE FROM products WHERE id=$1", productId)
}

func executeQuery(query string, parameters ...string) *sql.Rows {
	db := config.DatabaseConnector()
	var sqlResult *sql.Rows
	var err error

	if parameters == nil {
		sqlResult, err = db.Query(query)
	} else {
		sqlResult, err = db.Query(query, parameters[0])
	}

	util.ErrorHandler(err)
	defer db.Close()

	return sqlResult
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
