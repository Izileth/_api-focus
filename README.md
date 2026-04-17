# API Focus 🚀 - Guia de Integração Frontend

API backend em Go para processamento de pagamentos via Stripe (Pix, Boleto e Cartão) e gestão de serviços.

## 🔌 Endpoints de Pagamento

A API utiliza versionamento. Para novos desenvolvimentos, utilize a **v1**. A URL base é `http://localhost:8080` (ou o domínio de produção).

---

### 1. Criar Link de Pagamento (Hosted Checkout)
**URL:** `POST /api/v1/payments/create-checkout-session`

Ideal para quando você deseja redirecionar o usuário para uma página segura hospedada pelo Stripe que já contém todos os métodos de pagamento (Cartão, Pix, Boleto).

- **Request Body:**
```json
{
  "amount": 5000,
  "currency": "brl",
  "success_url": "https://moudusfocus.online/sucesso",
  "cancel_url": "https://moudusfocus.online/cancelado"
}
```

- **Resposta:**
```json
{
  "id": "cs_test_...",
  "url": "https://checkout.stripe.com/c/pay/..."
}
```
*Ação: O frontend deve redirecionar o usuário para a `url` retornada.*

---

### 2. Criar Intenção de Pagamento (Embedded)
**URL:** `POST /api/v1/payments/create-intent`

Utilizado para integrações customizadas onde o pagamento acontece dentro do seu site (usando Stripe Elements).

#### A. Pix (Obrigatório: Email, Nome, CPF e Endereço)
- **Request Body:**
```json
{
  "amount": 1000,
  "currency": "brl",
  "payment_method": "pix",
  "email": "cliente@email.com",
  "name": "Nome do Cliente",
  "tax_id": "123.456.789-00",
  "line1": "Rua Exemplo, 123",
  "city": "São Paulo",
  "state": "SP",
  "postal_code": "01234-567"
}
```

#### B. Boleto (Obrigatório: Email, Nome, CPF e Endereço)
- **Request Body:**
```json
{
  "amount": 10000,
  "currency": "brl",
  "payment_method": "boleto",
  "email": "cliente@email.com",
  "name": "Nome do Cliente",
  "tax_id": "123.456.789-00",
  "line1": "Rua Exemplo, 123",
  "city": "São Paulo",
  "state": "SP",
  "postal_code": "01234-567"
}
```

#### C. Cartão de Crédito
- **Request Body:**
```json
{
  "amount": 5000,
  "currency": "brl",
  "payment_method": "card"
}
```

---

### 📥 Resposta da API (Intent Mapping)

| Campo | Tipo | Descrição |
| :--- | :--- | :--- |
| `id` | `string` | Identificador do `PaymentIntent` (ex: `pi_...`). |
| `clientSecret` | `string` | Usado pelo Stripe Elements para confirmar o pagamento. |
| `status` | `string` | Status atual (ex: `requires_action`). |
| `nextAction` | `object` | Contém dados para Pix (QR Code) ou Boleto (URL). |

---

## ✨ Instruções Importantes para o Frontend

### 1. Dados de Endereço (CRÍTICO)
Para **Pix** e **Boleto**, o Stripe agora exige os detalhes de cobrança (`billing_details[address]`). Certifique-se de coletar e enviar:
- `line1`: Logradouro e número.
- `postal_code`: CEP formatado ou apenas números.
- `city` e `state`: Cidade e UF.

### 2. Moeda
- O **Pix** aceita exclusivamente a moeda `brl`. Tentativas com outras moedas retornarão erro 400.

### 3. Tratamento de Erros
A API retorna erros estruturados. Exemplo:
```json
{
  "error": "Email, Nome, TaxID e Endereço (Rua e CEP) são obrigatórios para Pix"
}
```

---

## 🔍 Saúde do Sistema
**URL:** `GET /health`
```json
{
  "api": "UP",
  "database": "UP",
  "message": "API Focus systems check"
}
```

---

## 🛠️ Configuração de CORS
A API está configurada para aceitar requisições de:
- `http://localhost:3000`
- `https://moudusfocus.online`
- `http://moudusfocus.online`

## 📝 Notas
- **Valores:** Sempre em **centavos** (R$ 1,00 = `100`).
- **Webhooks:** O status final do pagamento deve ser monitorado via webhook para maior segurança.
