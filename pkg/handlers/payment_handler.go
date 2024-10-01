package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Jpdsbarbosa/noxpay-client/pkg/models"
	"github.com/Jpdsbarbosa/noxpay-client/pkg/noxpay"
)

// CreatePaymentHandler é o handler para criar um novo pagamento
func CreatePaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var paymentRequest models.PaymentRequest
	err := json.NewDecoder(r.Body).Decode(&paymentRequest)
	if err != nil {
		http.Error(w, "Erro ao decodificar o pedido", http.StatusBadRequest)
		return
	}

	client := noxpay.NewClient()
	paymentResponse, err := client.CreatePayment(paymentRequest)
	if err != nil {
		http.Error(w, "Erro ao criar pagamento", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paymentResponse)
}

// ConsultPaymentHandler é o handler para consultar os detalhes de um pagamento
func ConsultPaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Extrai o txID da URL
	txID := strings.TrimPrefix(r.URL.Path, "/payment/")
	if txID == "" {
		http.Error(w, "ID do pagamento não fornecido", http.StatusBadRequest)
		return
	}

	client := noxpay.NewClient()
	paymentResponse, err := client.ConsultPayment(txID)
	if err != nil {
		http.Error(w, "Erro ao consultar pagamento", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paymentResponse)
}

// CreateCreditCardPaymentHandler é o handler para criar um novo pagamento por cartão de crédito
func CreateCreditCardPaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var ccPaymentRequest models.CreditCardPaymentRequest
	err := json.NewDecoder(r.Body).Decode(&ccPaymentRequest)
	if err != nil {
		http.Error(w, "Erro ao decodificar o pedido", http.StatusBadRequest)
		return
	}

	client := noxpay.NewClient()
	paymentResponse, err := client.CreateCreditCardPayment(ccPaymentRequest)
	if err != nil {
		http.Error(w, "Erro ao criar pagamento por cartão de crédito", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paymentResponse)
}
