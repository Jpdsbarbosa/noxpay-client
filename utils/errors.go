package utils

import (
	"fmt"
	"net/http"
)

// HandleError analisa a resposta de uma requisição HTTP e retorna um erro personalizado.
func HandleError(resp *http.Response) error {
	switch resp.StatusCode {
	case 400:
		return fmt.Errorf("erro de validação de parâmetros") // Erro 400.
	case 401:
		return fmt.Errorf("erro de autenticação") // Erro 401.
	default:
		return fmt.Errorf("erro de comunicação com a API: %d", resp.StatusCode) // Outros erros.
	}
}
