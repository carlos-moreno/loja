package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39, Quantidade: 10},
		{"Tênis", "Confortável", 99, 10},
		{"Fone", "Muito bom", 149.90, 5},
		{"Sapato", "Muito confortável", 189.90, 8},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}
