package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func New() {
	mux := http.NewServeMux() //Multiplexizador de requisicao ==> Aceitar varias requsicoes ao mesmo tempo (Exemplo ingressos) // NewServeMux ==> Configurar servidor

	mux.HandleFunc("/health", greet) // Caminho // funcao executada

	mux.HandleFunc("/add", register)

	http.ListenAndServe(":5555", mux) // ListenAndServe ==> Criar localhost
}

func greet(w http.ResponseWriter, r *http.Request) { // w resposta para usuario // requisicao do usuario
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() // Fecha o corpo (evitando vazamento de memoria,e automaticamente)

	type Person struct {
		Name    string `json:"Name"`
		Surname string `json:"Surname"`
		Age     int    `json:"Age"`
	}

	var person Person

	body, err := io.ReadAll(r.Body) //Le o corpo da requisicao (Formulario)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &person) //Pega o que recebe no corpo da requisiscao(bytes) e passa para o tipo struct
	if err != nil {
		http.Error(w, "Erro ao analisar JSON", http.StatusBadRequest)
		return
	}

	fmt.Println("Corpo da requisição:", string(body))
	fmt.Fprintf(w, "Nome: %s %s,Idade: %d", person.Name, person.Surname, person.Age)

}
