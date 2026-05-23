package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Worker de Go iniciado...")
	for {
		fmt.Println("Revisando tareas en la base de datos...")
		time.Sleep(10 * time.Second)
	}
}