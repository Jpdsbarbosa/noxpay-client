# NoxPay API Client

Este projeto implementa um cliente para a API da NoxPay, que permite a criação de pagamentos por Pix e cartão de crédito, além da consulta de pagamentos e obtenção de dados da conta. O servidor foi feito para rodar localmente e expõe vários endpoints para interagir com a API.

## Requisitos

- Go 1.20+
- Uma conta na NoxPay com uma API Key válida

## Instalação

1. Clone o repositório:

```bash
git clone https://github.com/Jpdsbarbosa/noxpay-client.git
cd noxpay-client
```

2. Instale as dependências:

```bash
go mod tidy
```

3. Execute o servidor localmente:

```bash
go run main.go
```

O servidor será iniciado na porta `8080`.

## Endpoints

Abaixo estão os endpoints disponíveis e como realizar as requisições usando `curl`.

### 1. Criar Pagamento por Pix

**Endpoint**: `/payment/pix`  
**Método**: `POST`

**Exemplo de requisição**:

```bash
curl -X POST http://localhost:8080/payment/pix \
  -H "Content-Type: application/json" \
  -d '{
    "amount": 1000,
    "currency": "BRL",
    "description": "Pagamento teste via Pix"
  }'
```

### 2. Consultar Pagamento

**Endpoint**: `/payment/{payment_id}`  
**Método**: `GET`

**Exemplo de requisição**:

```bash
curl -X GET http://localhost:8080/payment/{payment_id}
```

Substitua `{payment_id}` pelo ID do pagamento que deseja consultar.

### 3. Criar Pagamento por Cartão de Crédito

**Endpoint**: `/payment/creditcard`  
**Método**: `POST`

**Exemplo de requisição**:

```bash
curl -X POST http://localhost:8080/payment/creditcard \
  -H "Content-Type: application/json" \
  -d '{
    "amount": 1000,
    "currency": "BRL",
    "description": "Pagamento teste via Cartão de Crédito",
    "card_number": "4111111111111111",
    "expiration_date": "12/25",
    "cvv": "123"
  }'
```

### 4. Obter Dados da Conta

**Endpoint**: `/account-data`  
**Método**: `GET`

**Exemplo de requisição**:

```bash
curl -X GET http://localhost:8080/account-data
```

## Estrutura do Projeto

- `main.go`: Contém a inicialização do servidor e o roteamento dos endpoints.
- `pkg/handlers`: Contém os manipuladores de endpoints que lidam com as requisições HTTP.

## Logs

Os logs do servidor são exibidos no console. Exemplo de log ao iniciar o servidor:

```
Servidor iniciado na porta 8080
```

Se houver algum erro ao iniciar o servidor, ele será exibido no console.

## Considerações Finais

Este projeto foi desenvolvido como parte de um teste técnico. O servidor roda localmente e os exemplos de requisições são baseados no ambiente local.
```