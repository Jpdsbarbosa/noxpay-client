package noxpay

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Jpdsbarbosa/noxpay-client/pkg/models"
	"github.com/Jpdsbarbosa/noxpay-client/pkg/utils"
)

// GetAccountData realiza uma solicitação GET para a API da NoxPay para obter os dados da conta do usuário.
// Retorna uma instância de AccountData contendo informações como nome e saldo da conta,
// ou um erro caso a solicitação não seja bem-sucedida ou a resposta não possa ser processada.
func (c *Client) GetAccountData() (*models.AccountData, error) {
	// Constrói a requisição HTTP GET usando a URL base concatenada com o endpoint específico.
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/account", BaseURL), nil)
	if err != nil {
		// Retorna o erro caso não seja possível construir a requisição.
		return nil, fmt.Errorf("falha ao construir a requisição: %w", err)
	}

	// Realiza a requisição HTTP através do método DoRequest do cliente.
	resp, err := c.DoRequest(req)
	if err != nil {
		// Retorna o erro caso a requisição falhe.
		return nil, fmt.Errorf("falha na requisição para obter dados da conta: %w", err)
	}
	defer resp.Body.Close() // Garante o fechamento do corpo da resposta.

	// Verifica o código de status da resposta. Se não for 200 (OK), trata o erro.
	if resp.StatusCode != http.StatusOK {
		return nil, utils.HandleAPIError(resp)
	}

	var accountData models.AccountData
	// Decodifica o corpo da resposta para a struct AccountData.
	err = json.NewDecoder(resp.Body).Decode(&accountData)
	if err != nil {
		// Retorna o erro caso a decodificação falhe.
		return nil, fmt.Errorf("falha ao decodificar os dados da conta: %w", err)
	}

	// Retorna os dados da conta e nil para o erro caso tudo tenha sido bem-sucedido.
	return &accountData, nil
}
