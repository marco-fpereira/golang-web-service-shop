package repository

import (
	"web-service/app/config"
	"web-service/app/model"
	"web-service/app/util"
)

func GetAllProducts() []model.Product {
	db := config.DatabaseConnector()

	sqlProducts, err := db.Query("SELECT * FROM products")
	util.ErrorHandler(err)

	p := model.Product{}
	products := []model.Product{}

	for sqlProducts.Next() {
		var name, description string
		var price float64
		var amount, id int
		err = sqlProducts.Scan(&id, &name, &description, &price, &amount)
		util.ErrorHandler(err)

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Amount = amount

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func CreateNewProduct(name string, description string, price float64, amount int) {
	db := config.DatabaseConnector()
	query, err := db.Prepare("INSERT INTO products(name, description, price, amount) VALUES ($1, $2, $3, $4)")
	util.ErrorHandler(err)

	query.Exec(name, description, price, amount)

	defer db.Close()
}

func DeleteProduct(productId string) {
	db := config.DatabaseConnector()

	query, err := db.Prepare("DELETE FROM products WHERE id=$1")
	util.ErrorHandler(err)

	query.Exec(productId)
	defer db.Close()
}
