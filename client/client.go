package client

import (
	"net/http"
	"time"

	"github.com/Jpdsbarbosa/noxpay-client/config"
)

// Client: estutura que mantém o cliente HTTP e configurações para fazer requisições
type Client struct {
	HTTPClient *http.Client // Cliente HTTP para fazer requisições
}

// New: cria e retorna uma nova instância do cliente configurada
func New() *Client {
	return &Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second}, // timeout de 10 segundos para as requisições
	}
}

// DoRequest: executa uma requisição HTTP preparada e adiciona o header da APIKey
func (c *Client) DoRequest(req *http.Request) (*http.Response, error) {
	req.Header.Set("api-key", config.APIKey)
	return c.HTTPClient.Do(req) // Executa a requisição e retorna a resposta
}
