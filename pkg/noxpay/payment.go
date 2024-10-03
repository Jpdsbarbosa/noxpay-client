package noxpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Jpdsbarbosa/noxpay-client/pkg/models"
	"github.com/Jpdsbarbosa/noxpay-client/pkg/utils"
)

// CreatePayment cria um novo pagamento utilizando o método PIX.
// Recebe um PaymentRequest e retorna um PaymentResponse ou um erro.
func (c *Client) CreatePayment(pr models.PaymentRequest) (*models.PaymentResponse, error) {
	payload, err := json.Marshal(pr)
	if err != nil {
		return nil, fmt.Errorf("falha ao serializar o payload do pagamento: %w", err)
	}

	// Utiliza a BaseURL do cliente para construir a URL do endpoint de pagamento.
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/payment", c.BaseURL), bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("falha ao criar a requisição de pagamento: %w", err)
	}

	resp, err := c.DoRequest(req)
	if err != nil {
		return nil, fmt.Errorf("falha ao realizar a requisição de pagamento: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, utils.HandleAPIError(resp)
	}

	var paymentResp models.PaymentResponse
	err = json.NewDecoder(resp.Body).Decode(&paymentResp)
	if err != nil {
		return nil, fmt.Errorf("falha ao decodificar a resposta do pagamento: %w", err)
	}

	return &paymentResp, nil
}

// ConsultPayment consulta os detalhes de um pagamento existente pelo txID.
// Retorna um PaymentResponse com os detalhes do pagamento ou um erro.
func (c *Client) ConsultPayment(txID string) (*models.PaymentResponse, error) {
	// Utiliza a BaseURL do cliente para construir a URL do endpoint de consulta de pagamento.
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/payment/%s", c.BaseURL, txID), nil)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar a requisição de consulta de pagamento: %w", err)
	}

	resp, err := c.DoRequest(req)
	if err != nil {
		return nil, fmt.Errorf("falha ao realizar a requisição de consulta de pagamento: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, utils.HandleAPIError(resp)
	}

	var paymentResp models.PaymentResponse
	err = json.NewDecoder(resp.Body).Decode(&paymentResp)
	if err != nil {
		return nil, fmt.Errorf("falha ao decodificar a resposta da consulta de pagamento: %w", err)
	}

	return &paymentResp, nil
}
