package main

import (
	"log"
	"net/http"

	"github.com/Jpdsbarbosa/noxpay-client/pkg/handlers"
)

func main() {

	http.HandleFunc("/payment", handlers.CreatePaymentHandler)       // Para criar um pagamento por pix
	http.HandleFunc("/payment/", handlers.ConsultPaymentHandler)     // Para consultar um pagamento
	http.HandleFunc("/account-data", handlers.GetAccountDataHandler) // Para obter dados da conta

	log.Println("Servidor iniciado na porta 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Erro ao iniciar o servidor: ", err)
	}
}
