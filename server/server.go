package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Adiciona cabeçalhos CORS
func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Definição da estrutura para tarefas
type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Função para lidar com a rota "/health"
func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

// Função para lidar com a rota "/add-task"
func addTask(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var task Task
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &task)
	if err != nil {
		http.Error(w, "Erro ao analisar JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("Tarefa recebida: %+v\n", task)

	// Aqui você pode adicionar o código para salvar a tarefa no banco de dados

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Tarefa adicionada com sucesso: %+v", task)
}

// Função para inicializar o servidor
func New() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", greet)
	mux.HandleFunc("/add-task", addTask)

	// Adiciona o middleware CORS
	http.ListenAndServe(":5555", withCORS(mux))
}
