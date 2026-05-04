# Ciclo de Vida de Pagamento e Webhooks

Diferente de uma API comum, o resultado de um pagamento nem sempre é imediato. Este documento explica como lidar com a assincronia, especialmente para Pix e Boleto.

## 🔄 O Fluxo de Confirmação

1.  **Frontend:** Chama `/create-intent`.
2.  **Backend:** Retorna `clientSecret`.
3.  **Frontend:** Usa o SDK do Stripe para confirmar o pagamento.
4.  **Stripe:** Processa a transação.
    *   **Cartão:** Geralmente imediato.
    *   **Pix/Boleto:** O Stripe aguarda o pagamento do usuário.
5.  **Stripe Webhook:** O Stripe avisa o **Backend** da Focus que o dinheiro caiu.
6.  **Backend:** Atualiza o banco de dados e libera o serviço/produto.

## ⚓ Webhooks Suportados

O backend da Focus escuta os seguintes eventos do Stripe:

| Evento | Descrição | Ação Recomendada no Frontend |
| :--- | :--- | :--- |
| `payment_intent.succeeded` | Pagamento confirmado com sucesso. | Mostrar tela de sucesso/liberação. |
| `payment_intent.payment_failed` | Falha no processamento (ex: cartão recusado). | Pedir ao usuário outro método de pagamento. |
| `checkout.session.completed` | Sessão de checkout finalizada com sucesso. | Redirecionar para a `success_url`. |

## ⚠️ A Regra de Ouro do Frontend

**Nunca libere um produto baseando-se apenas na resposta do SDK no frontend.** 

O frontend pode ser fechado ou a conexão cair antes da resposta. A única fonte da verdade é o banco de dados da API, que é atualizado via Webhook. 

**Estratégia recomendada:**
1.  Após a confirmação no frontend, mostre uma tela de "Processando".
2.  Faça um polling (consultar de 3 em 3 segundos) em um endpoint de status do pedido (em desenvolvimento) ou use WebSockets para avisar o usuário quando o Webhook chegar.
