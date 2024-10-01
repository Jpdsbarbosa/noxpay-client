package models

// PaymentRequest representa a estrutura de dados para uma requisição de pagamento.
type PaymentRequest struct {
	Method         string  `json:"method"`
	Code           string  `json:"code"`
	Amount         float64 `json:"amount"`
	WebhookURL     string  `json:"webhook_url,omitempty"`
	ClientName     string  `json:"client_name,omitempty"`
	ClientDocument string  `json:"client_document,omitempty"`
}

// PaymentResponse representa a estrutura de dados para a resposta de um pagamento.
type PaymentResponse struct {
	Method       string  `json:"method"`
	Status       string  `json:"status"`
	Code         string  `json:"code"`
	TxID         string  `json:"txid"`
	Amount       float64 `json:"amount"`
	QRCode       string  `json:"qrcode,omitempty"`
	QRCodeBase64 string  `json:"qrcodebase64,omitempty"`
	CopyPaste    string  `json:"copypaste,omitempty"`
}

// AccountData representa a estrutura de dados para informações da conta.
type AccountData struct {
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

// CreditCardPaymentRequest representa a estrutura de dados para uma requisição de pagamento por cartão de crédito.
type CreditCardPaymentRequest struct {
	Code                 string  `json:"code"`
	Amount               float64 `json:"amount"`
	Email                string  `json:"email"`
	Name                 string  `json:"name"`
	CpfCnpj              string  `json:"cpf_cnpj"`
	ExpiredUrl           string  `json:"expired_url"`
	ReturnUrl            string  `json:"return_url"`
	MaxInstallmentsValue int     `json:"max_installments_value"`
	SoftDescriptorLight  string  `json:"soft_descriptor_light"`
}
