package noxpay

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Jpdsbarbosa/noxpay-client/pkg/models"
	"github.com/Jpdsbarbosa/noxpay-client/pkg/utils"
)

// GetAccountData obtém os dados da conta do usuário.
func (c *Client) GetAccountData() (*models.AccountData, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/account", BaseURL), nil)
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

	var accountData models.AccountData
	err = json.NewDecoder(resp.Body).Decode(&accountData)
	if err != nil {
		return nil, err
	}

	return &accountData, nil
}
