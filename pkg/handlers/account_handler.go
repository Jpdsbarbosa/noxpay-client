package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Jpdsbarbosa/noxpay-client/pkg/noxpay"
)

// GetAccountDataHandler é responsável por lidar com as requisições GET para obter os dados da conta do usuário.
// Este handler verifica se o método da requisição é GET e, em caso afirmativo, utiliza o cliente da NoxPay
// para obter os dados da conta. Em caso de sucesso, os dados da conta são retornados como JSON.
// Se o método da requisição não for GET, retorna um erro 405 (Method Not Allowed).
// Em caso de erro na obtenção dos dados da conta, retorna um erro 500 (Internal Server Error) e registra o erro.
func GetAccountDataHandler(w http.ResponseWriter, r *http.Request) {
	// Verifica se o método da requisição é GET.
	if r.Method != http.MethodGet {
		msg := "Método não permitido"
		log.Printf("GetAccountDataHandler: %s", msg)
		http.Error(w, msg, http.StatusMethodNotAllowed)
		return
	}

	client := noxpay.NewClient(noxpay.BaseURL, noxpay.APIKey)
	accountData, err := client.GetAccountData()
	if err != nil {
		// Log do erro para facilitar a depuração.
		log.Printf("Erro ao obter dados da conta: %v", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// Utiliza json.NewEncoder(w).Encode ao invés de json.Marshal para escrever diretamente na resposta.
	// Isso evita a necessidade de lidar com um potencial erro de json.Marshal e é mais eficiente.
	if err := json.NewEncoder(w).Encode(accountData); err != nil {
		// Log do erro ao tentar codificar a resposta JSON.
		log.Printf("Erro ao codificar dados da conta para JSON: %v", err)
		http.Error(w, "Erro ao processar dados da conta", http.StatusInternalServerError)
	}
}
