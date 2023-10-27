// main.go
package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"api/controllers"
	"api/services"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	classService := services.NewClassService()
	classController := controllers.NewClassController(classService)

	r.Post("/classes", classController.CreateClass)

	// Obtém o endereço IP atribuído ao servidor
	serverAddr, err := determineServerAddress()
	if err != nil {
		fmt.Println("Erro ao determinar o endereço IP do servidor:", err)
		os.Exit(1)
	}
	fmt.Println("Servidor iniciado em:", serverAddr)

	http.ListenAndServe(serverAddr, r)
}

func determineServerAddress() (string, error) {
	// Obtenha o endereço IP local da máquina
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.String(), nil
}
