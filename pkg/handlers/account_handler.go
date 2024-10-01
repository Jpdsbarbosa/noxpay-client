package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Jpdsbarbosa/noxpay-client/pkg/noxpay"
)

// GetAccountDataHandler é o handler para obter os dados da conta
func GetAccountDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	client := noxpay.NewClient()
	accountData, err := client.GetAccountData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accountData)
}
