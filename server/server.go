package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// New configura o servidor e as rotas
func New() {
	mux := http.NewServeMux() // Multiplexador de requisições

	mux.HandleFunc("/health", greet) // Configura a rota "/health"
	mux.HandleFunc("/add", register) // Configura a rota "/add"

	fmt.Println("Server is running on http://localhost:5555")
	http.ListenAndServe(":5555", mux) // Inicia o servidor na porta 5555
}

// greet responde com uma mensagem simples
func greet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Permite requisições de qualquer origem
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

// register processa e responde a requisições POST com JSON
func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Permite requisições de qualquer origem
	defer r.Body.Close()

	type Person struct {
		Name    string `json:"Name"`
		Surname string `json:"Surname"`
		Age     int    `json:"Age"`
	}

	var person Person

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &person)
	if err != nil {
		http.Error(w, "Erro ao analisar JSON", http.StatusBadRequest)
		return
	}

	fmt.Println("Corpo da requisição:", string(body))
	fmt.Fprintf(w, "Nome: %s %s, Idade: %d", person.Name, person.Surname, person.Age)
}
