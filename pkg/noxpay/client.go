package noxpay

import (
	"net/http"
	"time"
)

// Client: estrutura para o cliente HTTP
type Client struct {
	HTTPClient *http.Client
}

// NewClient: cria e retorna uma nova instância do Cliente
func NewClient() *Client {
	return &Client{
		HTTPClient: &http.Client{Timeout: 15 * time.Second},
	}
}

// DoRequest: executa a requisição HTTP preparada e retorna a resposta
func (c *Client) DoRequest(req *http.Request) (*http.Response, error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", APIKey)
	return c.HTTPClient.Do(req)
}
