package routes

import (
	"net/http"

	"github.com/carlos-moreno/loja/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
