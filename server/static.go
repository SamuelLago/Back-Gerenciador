package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("../../Front-Gerenciador/public"))
	http.Handle("/", fs)

	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
