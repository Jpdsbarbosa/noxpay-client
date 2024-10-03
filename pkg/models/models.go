package models

// PaymentRequest representa a estrutura de dados para uma requisição de pagamento.
// Inclui informações sobre o método de pagamento, código único do pagamento, valor,
// e opcionalmente, uma URL de webhook, nome do cliente e documento do cliente.
type PaymentRequest struct {
	Method         string  `json:"method"`                    // Método de pagamento (ex: PIX).
	Code           string  `json:"code"`                      // Código único identificador do pagamento.
	Amount         float64 `json:"amount"`                    // Valor do pagamento.
	WebhookURL     string  `json:"webhook_url,omitempty"`     // URL de webhook para notificações (opcional).
	ClientName     string  `json:"client_name,omitempty"`     // Nome do cliente (opcional).
	ClientDocument string  `json:"client_document,omitempty"` // Documento do cliente (CPF/CNPJ) (opcional).
}

// PaymentResponse representa a estrutura de dados para a resposta de um pagamento.
// Contém detalhes sobre o método, status, código, identificador de transação (TxID),
// valor, e opcionalmente, QR Code e informações para cópia e colagem.
type PaymentResponse struct {
	Method       string  `json:"method"`                 // Método de pagamento utilizado.
	Status       string  `json:"status"`                 // Status atual do pagamento.
	Code         string  `json:"code"`                   // Código único identificador do pagamento.
	TxID         string  `json:"txid"`                   // Identificador único da transação.
	Amount       float64 `json:"amount"`                 // Valor do pagamento.
	QRCode       string  `json:"qrcode,omitempty"`       // URL do QR Code para pagamento (opcional).
	QRCodeBase64 string  `json:"qrcodebase64,omitempty"` // QR Code em formato base64 (opcional).
	CopyPaste    string  `json:"copypaste,omitempty"`    // String para cópia e colagem (opcional).
}

// AccountData representa a estrutura de dados para informações da conta.
// Inclui o nome da conta e o saldo atual.
type AccountData struct {
	Name    string  `json:"name"`    // Nome da conta.
	Balance float64 `json:"balance"` // Saldo atual da conta.
}
