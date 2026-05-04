# Payloads e Schemas de Dados

Este documento serve como referência técnica para os objetos trafegados entre Frontend e Backend.

## 📥 Payloads de Entrada (Requests)

### PaymentRequest
Usado em `/create-intent`.

| Campo | Tipo | Obrigatório | Descrição |
| :--- | :--- | :--- | :--- |
| `amount` | `int64` | Sim | Valor em centavos (Ex: 1000 = R$ 10,00) |
| `currency` | `string` | Sim | Código da moeda (Ex: "brl") |
| `email` | `string` | Sim | E-mail do cliente para recibo |

### CheckoutRequest
Usado em `/create-checkout-session`.

| Campo | Tipo | Obrigatório | Descrição |
| :--- | :--- | :--- | :--- |
| `amount` | `int64` | Sim | Valor em centavos |
| `currency` | `string` | Sim | Código da moeda |
| `success_url` | `string` | Sim | URL de retorno após sucesso |
| `cancel_url` | `string` | Sim | URL de retorno após cancelamento |

---

## 📤 Objetos de Saída (Responses)

### Intent Response
```json
{
  "clientSecret": "string",
  "id": "string",
  "status": "string",
  "nextAction": "object | null"
}
```
*   `nextAction`: Presente se o pagamento exigir uma ação imediata (ex: exibir QR Code de Pix gerado pelo backend, embora o recomendado seja via SDK).

### Checkout Response
```json
{
  "id": "string",
  "url": "string"
}
```

---

## 🔐 Segurança e Headers

*   **Content-Type:** `application/json`
*   **CORS:** As requisições devem partir de origens autorizadas (ver `internal/middleware/cors.go`).
*   **Versionamento:** Sempre inclua a versão na rota: `/api/v1/...`.
