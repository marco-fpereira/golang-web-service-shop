package controller

import (
	"html/template"
	"net/http"
	"strconv"
	"web-service/app/repository"
	"web-service/app/util"

	_ "github.com/lib/pq"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := repository.GetAllProducts()

	templates.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		util.ErrorHandler(err)
		amount, err := strconv.Atoi(r.FormValue("amount"))
		util.ErrorHandler(err)

		repository.CreateNewProduct(name, description, price, amount)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	repository.DeleteProduct(productId)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := repository.GetProductById(productId)

	templates.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, err := strconv.Atoi(r.FormValue("id"))
		util.ErrorHandler(err)

		name := r.FormValue("name")

		description := r.FormValue("description")

		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		util.ErrorHandler(err)

		amount, err := strconv.Atoi(r.FormValue("amount"))
		util.ErrorHandler(err)

		repository.UpdateProduct(id, name, description, price, amount)
	}
	http.Redirect(w, r, "/", 301)
}
