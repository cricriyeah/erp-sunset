package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Microservicio de PDFs activo...")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Servidor iniciado en el puerto 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}