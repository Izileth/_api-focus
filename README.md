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

### 📥 Resposta da API (Mapeamento)

Todas as requisições bem-sucedidas para criação de intenção retornam o seguinte objeto:

| Campo | Tipo | Descrição |
| :--- | :--- | :--- |
| `id` | `string` | Identificador único do `PaymentIntent` (ex: `pi_...`). |
| `clientSecret` | `string` | Chave secreta usada pelo Stripe Elements para confirmar o pagamento. |
| `status` | `string` | Status atual (ex: `requires_payment_method`, `requires_action`). |
| `nextAction` | `object` | (Opcional) Contém dados para Pix (QR Code) ou Boleto (URL/Linha). |

#### Exemplo de Resposta (Pix/Boleto)
```json
{
  "id": "pi_3O...",
  "clientSecret": "pi_3O..._secret_...",
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

---

## ✨ Instruções para o Frontend (UX/Sucesso)

Para garantir uma boa experiência ao usuário, siga estas diretrizes após receber a resposta da API:

### 1. Mensagens de Sucesso e Feedback
- **Cartão de Crédito:** Se o status for `succeeded`, exiba uma tela de "Pagamento Confirmado" imediatamente.
- **Pix:** Exiba o QR Code de forma clara, o botão "Copiar Código Pix" e informe que a confirmação é automática (geralmente em segundos).
- **Boleto:** Ofereça o botão para "Baixar PDF" e mostre a linha digitável com um botão de cópia. Avise que o prazo de compensação é de até 3 dias úteis.

### 2. Estados de Pagamento
Recomendamos tratar os status retornados pelo Stripe:
- `requires_payment_method`: O pagamento falhou ou ainda não foi iniciado.
- `requires_action`: O usuário precisa completar um passo (escanear Pix, pagar Boleto ou 3D Secure).
- `processing`: O pagamento está sendo processado (comum em boletos e cartões com análise de fraude).
- `succeeded`: Dinheiro garantido! Libere o acesso ou produto.

### 3. Tratamento de Erros
Em caso de erro (Status 400 ou 500), a API retornará:
```json
{
  "error": "Descrição detalhada do erro para o desenvolvedor"
}
```
*Dica: No frontend, exiba uma mensagem amigável como "Ops! Algo deu errado com seu pagamento. Tente novamente ou use outro método."*

---

## 🔍 Saúde do Sistema
Você pode verificar o status da API e do banco de dados:
**URL:** `GET /health`
```json
{
  "api": "UP",
  "database": "UP",
  "message": "API Focus systems check"
}
```

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
