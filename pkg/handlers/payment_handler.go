package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/Jpdsbarbosa/noxpay-client/pkg/models"
	"github.com/Jpdsbarbosa/noxpay-client/pkg/noxpay"
)

// CreatePaymentHandler lida com a criação de novos pagamentos.
// Aceita apenas métodos POST e espera um corpo de requisição JSON com os dados do pagamento.
func CreatePaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("CreatePaymentHandler: método não permitido %v", r.Method)
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var paymentRequest models.PaymentRequest
	err := json.NewDecoder(r.Body).Decode(&paymentRequest)
	if err != nil {
		log.Printf("CreatePaymentHandler: erro ao decodificar o pedido %v", err)
		http.Error(w, "Erro ao decodificar o pedido", http.StatusBadRequest)
		return
	}

	client := noxpay.NewClient(noxpay.BaseURL, noxpay.APIKey)
	paymentResponse, err := client.CreatePayment(paymentRequest)
	if err != nil {
		log.Printf("CreatePaymentHandler: erro ao criar pagamento %v", err)
		http.Error(w, "Erro ao criar pagamento", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(paymentResponse); err != nil {
		log.Printf("CreatePaymentHandler: erro ao codificar resposta JSON %v", err)
		http.Error(w, "Erro ao processar resposta", http.StatusInternalServerError)
	}
}

// ConsultPaymentHandler lida com a consulta de detalhes de pagamentos existentes.
// Aceita apenas métodos GET e espera um txID como parte da URL.
func ConsultPaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Printf("ConsultPaymentHandler: método não permitido %v", r.Method)
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	txID := strings.TrimPrefix(r.URL.Path, "/payment/")
	if txID == "" {
		log.Printf("ConsultPaymentHandler: ID do pagamento não fornecido")
		http.Error(w, "ID do pagamento não fornecido", http.StatusBadRequest)
		return
	}

	client := noxpay.NewClient(noxpay.BaseURL, noxpay.APIKey)
	paymentResponse, err := client.ConsultPayment(txID)
	if err != nil {
		log.Printf("ConsultPaymentHandler: erro ao consultar pagamento %v", err)
		http.Error(w, "Erro ao consultar pagamento", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(paymentResponse); err != nil {
		log.Printf("ConsultPaymentHandler: erro ao codificar resposta JSON %v", err)
		http.Error(w, "Erro ao processar resposta", http.StatusInternalServerError)
	}
}
