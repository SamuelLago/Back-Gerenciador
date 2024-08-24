package server

import (
	"fmt"
	"net/http"
	"time"
)

func New() {
	mux := http.NewServeMux() //Multiplexizador de requisicao ==> Aceitar varias requsicoes ao mesmo tempo (Exemplo ingressos) // NewServeMux ==> Configurar servidor

	mux.HandleFunc("/health", greet) // Caminho // funcao executada

	http.ListenAndServe(":5555", mux) // ListenAndServe ==> Criar localhost
}

func greet(w http.ResponseWriter, r *http.Request) { // w resposta para usuario // requisicao do usuario
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}
