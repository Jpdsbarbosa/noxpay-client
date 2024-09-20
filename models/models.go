package models

type PaymentRequest struct {
	Method         string  `json:"method"`                    // método de pagamento
	Code           string  `json:"code"`                      // código único para o pagamento
	Amount         float64 `json:"amount"`                    // valor do pagamento
	Webhook        string  `json:"webhook_url,omitempty"`     // URL do webhook para notificações (opcional)
	ClientName     string  `json:"client_name,omitempty"`     // Nome do cliente (opcional)
	ClientDocument string  `json:"client_document,omitempty"` // Documento do cliente (cpf/cnpj, opcional)
}

// PaymentResponse: resposta da API ao criar pagamento
type PaymentResponse struct {
	Method       string  `json:"method"`                 // método de pagamento
	Status       string  `json:"status"`                 // statys do pagamento
	Code         string  `json:"code"`                   // código do pagamento
	TxID         string  `json:"txid"`                   // ID da transação
	Amount       float64 `json:"amount"`                 // valor do pagamento
	QRCode       string  `json:"qrcode,omitempty"`       // Qr Code para pagamento (se aplicável)
	QRCodeBase64 string  `json:"qrcodebase64,omitempty"` // Qr Code em Base64 (se aplicável)
	CopyPaste    string  `json:"copypaste,omitempty"`    // String para copiar e colar (se aplicável)
}

// AcountData: dados da conta obtidos da API
type AccountData struct {
	Name    string  `json:"name"`    // Nome da conta
	Balance float64 `json:"balance"` // saldo disponível na conta
}
