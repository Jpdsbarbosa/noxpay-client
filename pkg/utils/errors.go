package utils

import (
	"fmt"
	"io"
	"net/http"
)

// HandleAPIError trata os erros de API, lendo o corpo da resposta e retornando uma mensagem de erro mais informativa.
func HandleAPIError(resp *http.Response) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// Retorna um erro que inclui tanto o código de status quanto o fato de que o corpo da resposta não pôde ser lido.
		return fmt.Errorf("erro da API: código de status %d, falha ao ler o corpo da resposta: %v", resp.StatusCode, err)
	}

	// Se o corpo foi lido com sucesso, inclui o corpo na mensagem de erro.
	return fmt.Errorf("erro da API: código de status %d, corpo: %s", resp.StatusCode, string(body))
}
