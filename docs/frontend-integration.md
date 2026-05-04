# Guia de Integração Frontend - API Focus

Este guia detalha como o frontend deve interagir com a API Focus para processar pagamentos de forma segura e eficiente.

## 📌 Visão Geral

A API utiliza o **Stripe** como processador. Existem dois fluxos principais:

1.  **Hosted Checkout:** O usuário é redirecionado para uma página segura do Stripe.
2.  **Embedded Payment (Intents):** O pagamento ocorre dentro da sua aplicação usando **Stripe Elements**.

---

## 💰 1. Fluxo de Intenção (Embedded)

Ideal para manter o usuário no seu site.

### Criar Intenção
**Endpoint:** `POST /api/v1/payments/create-intent`

**Payload:**
```json
{
  "amount": 5000,
  "currency": "brl",
  "email": "cliente@email.com"
}
```

**Resposta de Sucesso (200 OK):**
```json
{
  "clientSecret": "pi_3P..._secret_...",
  "id": "pi_3P...",
  "status": "requires_payment_method"
}
```

**Ação do Frontend:**
Use o `clientSecret` para inicializar o `Stripe Elements`. O Stripe decidirá quais métodos mostrar (Pix, Boleto, Cartão) com base na moeda e localização.

---

## 🔗 2. Fluxo de Checkout (Hosted)

Ideal para uma integração rápida e simples.

### Criar Sessão
**Endpoint:** `POST /api/v1/payments/create-checkout-session`

**Payload:**
```json
{
  "amount": 10000,
  "currency": "brl",
  "success_url": "https://seu-site.com/sucesso?session_id={CHECKOUT_SESSION_ID}",
  "cancel_url": "https://seu-site.com/cancelado"
}
```

**Resposta de Sucesso (200 OK):**
```json
{
  "id": "cs_test_...",
  "url": "https://checkout.stripe.com/c/pay/..."
}
```

**Ação do Frontend:**
Redirecione o usuário para a `url` retornada.

---

## 📋 3. Requisitos de Dados por Método

Embora a API aceite payloads simplificados, o Stripe exige dados específicos no momento da confirmação (via Frontend SDK):

### Pix & Boleto (Obrigatórios para conformidade no Brasil)
*   **Nome Completo**
*   **E-mail**
*   **CPF/CNPJ (Tax ID)**
*   **Endereço:** Rua, Número, Cidade, Estado (UF) e CEP.

> **Nota:** A API Focus está configurada para habilitar `Automatic Payment Methods`. Certifique-se de coletar os dados acima no seu formulário antes de chamar `confirmPayment` no Stripe SDK.

---

## ⚠️ 4. Tratamento de Erros

A API utiliza códigos HTTP padrão:

*   `400 Bad Request`: Payload malformado ou dados faltando.
*   `500 Internal Server Error`: Falha na comunicação com Stripe ou Banco de Dados.

**Estrutura de Erro:**
```json
{
  "error": "Descrição detalhada do erro para o desenvolvedor"
}
```

---

## 💡 Dicas Importantes

1.  **Centavos:** O campo `amount` é sempre um inteiro em centavos. `R$ 10,50` deve ser enviado como `1050`.
2.  **Moeda:** Atualmente apenas `brl` é suportado para métodos locais (Pix/Boleto).
3.  **Ambiente:** Utilize chaves `pk_test_...` no frontend para testar com os dados de teste do Stripe.
