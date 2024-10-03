package noxpay

import (
	"net/http"
	"time"
)

// Client representa a estrutura para o cliente HTTP que interage com a API da NoxPay.
type Client struct {
	HTTPClient *http.Client // Cliente HTTP para fazer as requisições.
	BaseURL    string       // BaseURL é a URL base da API da NoxPay.
	APIKey     string       // APIKey é a chave de API para autenticação nas requisições.
}

// NewClient cria e retorna uma nova instância do Cliente com a URL base da API e a chave de API configuradas.
// Recebe a URL base da API e a chave de API como argumentos.
func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		HTTPClient: &http.Client{Timeout: 15 * time.Second},
		BaseURL:    baseURL,
		APIKey:     apiKey,
	}
}

// DoRequest executa a requisição HTTP preparada, injetando os headers necessários, e retorna a resposta.
// Adiciona automaticamente o Content-Type e a chave de API a cada requisição.
func (c *Client) DoRequest(req *http.Request) (*http.Response, error) {
	req.Header.Set("Content-Type", "application/json") // Define o Content-Type das requisições como JSON.
	req.Header.Set("api-key", c.APIKey)                // Adiciona a chave de API ao header da requisição.

	return c.HTTPClient.Do(req) // Executa a requisição e retorna a resposta.
}
