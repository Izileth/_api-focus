# API Focus 🚀 - Guia de Integração Frontend

API backend em Go para processamento de pagamentos via Stripe (Pix, Boleto e Cartão) e gestão de serviços.

## 🔌 Endpoints de Pagamento

A API utiliza versionamento. Para novos desenvolvimentos, utilize a **v1**.

### Criar Intenção de Pagamento
**URL:** `POST /api/v1/payments/create-intent`

Cria um `PaymentIntent` no Stripe e retorna o `clientSecret` e ações necessárias para concluir o pagamento.

#### 1. Cartão de Crédito
Apenas o valor e a moeda são obrigatórios para a intenção, mas o Stripe Elements cuidará do restante no front.
- **Request Body:**
```json
{
  "amount": 5000,
  "currency": "brl",
  "payment_method": "card"
}
```

#### 2. Pix (Obrigatório: Email e Nome)
Retorna o `nextAction` com os dados para gerar o QR Code.
- **Request Body:**
```json
{
  "amount": 1000,
  "currency": "brl",
  "payment_method": "pix",
  "email": "cliente@email.com",
  "name": "Nome do Cliente"
}
```

#### 3. Boleto (Obrigatório: Email, Nome e TaxID/CPF)
Retorna o `nextAction` com o link para o PDF e a linha digitável.
- **Request Body:**
```json
{
  "amount": 10000,
  "currency": "brl",
  "payment_method": "boleto",
  "email": "cliente@email.com",
  "name": "Nome do Cliente",
  "tax_id": "123.456.789-00"
}
```

---

### 📥 Resposta da API (Exemplo Pix/Boleto)
Ao criar uma intenção para Pix ou Boleto, a API retornará o campo `nextAction`. Você deve usar esses dados para exibir o pagamento ao usuário.

```json
{
  "clientSecret": "pi_3O...",
  "id": "pi_3O...",
  "status": "requires_action",
  "nextAction": {
    "type": "display_pix_qr_code",
    "display_pix_qr_code": {
      "data": "00020101021226870014...", 
      "image_url_png": "https://...",
      "expires_at": 1712950000
    }
  }
}
```
*No caso de **Boleto**, o `type` será `verify_with_microdeposits` ou similar, contendo a URL do PDF.*

---

## 🛠️ Como Testar Localmente (Dev)

### 1. Iniciar a API
```bash
air
```

### 2. Configurar Webhooks (Opcional, mas Recomendado)
Para receber confirmações de pagamento em tempo real no seu ambiente local, use o Stripe CLI:
```bash
stripe listen --forward-to localhost:8080/api/v1/payments/webhook
```

### 3. Variáveis de Ambiente Necessárias (.env)
```env
PORT=8080
STRIPE_SECRET_KEY=sk_test_...
STRIPE_WEBHOOK_SECRET=whsec_...
```

## 📝 Notas de Implementação
- **Valores:** O campo `amount` é sempre em **centavos** (ex: R$ 10,00 = `1000`).
- **Segurança:** Nunca armazene a `STRIPE_SECRET_KEY` no frontend. Utilize apenas a `Publishable Key` no cliente.
