package service

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/Jpdsbarbosa/noxpay-client/client"
	"github.com/Jpdsbarbosa/noxpay-client/config"
	"github.com/Jpdsbarbosa/noxpay-client/models"
)

// CreatePayment: requisição para criar pagamento
func CreatePayment(c *client.Client, pr models.PaymentRequest) (*models.PaymentResponse, error) {
	payloadBuf, err := json.Marshal(pr) // converte a estrutura PaymentRequest para JSON
	if err != nil {
		return nil, err
	}

	// prepara a requisição HTTP para o endpoint de criação de pagamento
	req, err := http.NewRequest("POST", config.BaseURL+"/payment", bytes.NewBuffer(payloadBuf))
	if err != nil {
		return nil, err
	}

	resp, err := c.DoRequest(req) // Executa a requisição usando o cliente HTTP
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // garante o fechamento do corpo de resposta

	var paymentResp models.PaymentResponse // variavel para armazenar resposta
	// decodificação do corpo da response para a estrutura PaymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&paymentResp); err != nil {
		return nil, err
	}
	return &paymentResp, nil // Retorna a resposta decodificada

}
