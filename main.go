package main

import (
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
	"github.com/AllanKley/GO_WEB_EXPLORATION/Database"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var templates = template.Must(template.ParseGlob("./templates/*.html"))

func main() {

	dbm := &Database.DatabaseManager{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		DbName:   "GO_WEB",
		Password: "root",
		Ssl:      "disabled",
	}

	db := dbm.Connect()

	defer db.Close()

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{"product_1", "description_1", 100, 10},
	}

	templates.ExecuteTemplate(w, "Index", products)
}
