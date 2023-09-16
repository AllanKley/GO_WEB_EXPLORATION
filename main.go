package main

import (
	"net/http"
	"text/template"

	"github.com/AllanKley/GO_WEB_EXPLORATION/Database"
	_ "github.com/lib/pq"
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
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err.Error())
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{"product_1", "description_1", 100, 10},
	}

	err := templates.ExecuteTemplate(w, "Index", products)
	if err != nil {
		panic(err.Error())
	}
}
