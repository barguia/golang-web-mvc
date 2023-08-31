package controllers

import (
	"fmt"
	"golang-web/models"
	"html/template"
	"net/http"
)

var app_template = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.GetAllProdutos()
	app_template.ExecuteTemplate(w, "Index", produtos)
}

func Create(w http.ResponseWriter, r *http.Request) {
	app_template.ExecuteTemplate(w, "Create", nil)
}

func CadastraProduto(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Oi mundo")
}
