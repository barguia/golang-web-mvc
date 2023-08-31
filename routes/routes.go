package routes

import (
	"golang-web/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/cadastro", controllers.Create)
	http.HandleFunc("/cadastra-produto", controllers.CadastraProduto)
}
