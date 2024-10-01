# NoxPay Client

Esta é uma biblioteca Go para integração com a API de pagamentos da NoxPay, fornecendo operações básicas como criar pagamentos, consultar o status de pagamentos e obter dados da conta. O projeto inclui exemplos prontos para rodar e suporte para diferentes métodos de pagamento, como Pix e cartão de crédito.

## Índice

- [Instalação](#instalação)
- [Configuração](#configuração)
- [Uso](#uso)
  - [Criar Pagamento (Pix)](#criar-pagamento-pix)
  - [Criar Pagamento (Cartão de Crédito)](#criar-pagamento-cartão-de-crédito)
  - [Consultar Pagamento](#consultar-pagamento)
  - [Obter Dados da Conta](#obter-dados-da-conta)
- [Executar o Servidor](#executar-o-servidor)
- [Testes](#testes)
- [Contribuição](#contribuição)
- [Licença](#licença)

## Instalação

Você pode usar a biblioteca diretamente em seu projeto Go, mas os exemplos abaixo mostram como realizar requisições diretamente com `cURL` para interagir com a API NoxPay.

## Configuração

Antes de utilizar a API NoxPay, você precisará da sua `API Key` e da `Base URL`. No exemplo abaixo, usamos a URL padrão da API NoxPay:

- **Base URL**: `https://api2.noxpay.io`
- **API Key**: `sua-api-key-aqui`

## Uso

### Criar Pagamento (Pix)

Para criar um pagamento via Pix, faça uma requisição `POST` à API com os dados do pagamento. O exemplo abaixo cria um pagamento no valor de 100.50 com o método "PIX".

```bash
curl -X POST https://api2.noxpay.io/payment/pix \
-H "Authorization: Bearer sua-api-key-aqui" \
-H "Content-Type: application/json" \
-d '{
  "method": "PIX",
  "code": "123456",
  "amount": 100.50
}'
```

### Criar Pagamento (Cartão de Crédito)

Para criar um pagamento via cartão de crédito, envie uma requisição `POST` com as informações do cartão e do pagamento:

```bash
curl -X POST https://api2.noxpay.io//payment/creditcard \
-H "Authorization: Bearer sua-api-key-aqui" \
-H "Content-Type: application/json" \
-d '{
  "method": "CREDIT_CARD",
  "code": "123456",
  "amount": 250.00,
  "email": "cliente@email.com",
  "name": "Cliente Teste",
  "cpf_cnpj": "123.456.789-00",
  "max_installments_value": 12,
  "soft_descriptor_light": "CompraTest"
}'
```

### Consultar Pagamento

Para consultar o status de um pagamento já criado, basta fazer uma requisição `GET` à API, passando o ID da transação:

```bash
curl -X GET https://api2.noxpay.io/payment/{id-do-pagamento} \
-H "Authorization: Bearer sua-api-key-aqui"
```

### Obter Dados da Conta

Para obter os dados da sua conta, você pode usar o seguinte comando `cURL`:

```bash
curl -X GET https://api2.noxpay.io/aaccount-data \
-H "Authorization: Bearer sua-api-key-aqui"
```

## Executar o Servidor

Caso esteja utilizando esta biblioteca como base para um serviço, certifique-se de que as variáveis de ambiente para `API Key` e `Base URL` estão configuradas corretamente no seu ambiente de desenvolvimento.

## Testes

Para rodar os testes da biblioteca, execute o seguinte comando:

```bash
go test ./...
```

Os testes utilizam mocks para simular as respostas da API.

## Contribuição

Contribuições são bem-vindas! Por favor, abra um *pull request* ou envie sugestões na seção de *issues*.

## Licença

Este projeto está licenciado sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---
