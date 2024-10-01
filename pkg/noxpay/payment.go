package noxpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Jpdsbarbosa/noxpay-client/pkg/models"
	"github.com/Jpdsbarbosa/noxpay-client/pkg/utils"
)

// CreatePayment: cria um novo pagamento (pix)
func (c *Client) CreatePayment(pr models.PaymentRequest) (*models.PaymentResponse, error) {
	payload, err := json.Marshal(pr)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/payment", BaseURL), bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	resp, err := c.DoRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, utils.HandleAPIError(resp)
	}

	var PaymentResp models.PaymentResponse
	err = json.NewDecoder(resp.Body).Decode(&PaymentResp)
	if err != nil {
		return nil, err
	}

	return &PaymentResp, nil
}

// ConsultPayment: consulta os detalhes de um pagamento
func (c *Client) ConsultPayment(txID string) (*models.PaymentResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/payment/%s", BaseURL, txID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.DoRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, utils.HandleAPIError(resp)
	}

	var paymentResp models.PaymentResponse
	err = json.NewDecoder(resp.Body).Decode(&paymentResp)
	if err != nil {
		return nil, err
	}

	return &paymentResp, nil
}

// CreateCreditCardPayment: cria um novo pagamento (cartão de crédito)
func (c *Client) CreateCreditCardPayment(ccpr models.CreditCardPaymentRequest) (*models.PaymentResponse, error) {
	payload, err := json.Marshal(ccpr)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/creditcard", BaseURL), bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	resp, err := c.DoRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, utils.HandleAPIError(resp)
	}

	var paymentResp models.PaymentResponse
	err = json.NewDecoder(resp.Body).Decode(&paymentResp)
	if err != nil {
		return nil, err
	}

	return &paymentResp, nil
}
