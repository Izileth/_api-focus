# Guia de Testes e Simulação

Para garantir que a integração está funcionando, utilize as ferramentas e dados de teste do Stripe.

## 💳 Cartões de Teste

| Bandeira | Número | CVV | Expiração |
| :--- | :--- | :--- | :--- |
| Sucesso (Visa) | `4242 4242 4242 4242` | `123` | Qualquer data futura |
| Recusado | `4000 0000 0000 0002` | `123` | Qualquer data futura |

## 💠 Simulação de Pix e Boleto

Ao usar o `clientSecret` no Stripe Elements em modo de teste:
*   **Pix:** O Stripe exibirá um botão "Simular Pagamento com Sucesso". Clique nele para disparar o webhook de sucesso.
*   **Boleto:** O Stripe gerará um PDF de teste. No dashboard do Stripe, você pode forçar o pagamento desse boleto de teste.

## 🛠️ Testando Webhooks Localmente (Stripe CLI)

Para que o backend receba os avisos do Stripe na sua máquina local, você deve usar o **Stripe CLI**:

1.  **Instale o Stripe CLI:** `scoop install stripe` ou baixe o binário.
2.  **Login:** `stripe login`
3.  **Forwarding:**
    ```bash
    stripe listen --forward-to localhost:8080/api/v1/payments/webhook
    ```
4.  **Copie a Secret:** O comando acima vai gerar uma chave começando com `whsec_...`. Coloque-a no seu `.env` como `STRIPE_WEBHOOK_SECRET`.

## 🧪 Validando a API

Você pode usar o cURL para testar a criação de intenções rapidamente:

```bash
curl -X POST http://localhost:8080/api/v1/payments/create-intent \
     -H "Content-Type: application/json" \
     -d '{"amount": 2000, "currency": "brl", "email": "teste@exemplo.com"}'
```
