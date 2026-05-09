# Modificações Necessárias no Front-end (Fix 400 Error)

Este documento detalha as alterações necessárias no front-end para que as requisições de pagamento parem de retornar erro `400 Bad Request`.

## 1. Formatação do Valor (`amount`)

O Back-end utiliza o Stripe, que exige que valores monetários sejam enviados em **centavos** como números inteiros.

*   **Como está:** `amount: 195.8` (Decimal/Float)
*   **Como deve ser:** `amount: 19580` (Inteiro)

**Exemplo de correção no TypeScript/JS:**
```typescript
const amountInCents = Math.round(valorEmReais * 100);
```

## 2. Mapeamento de Campos (Payload)

O Back-end espera nomes de campos específicos na struct de `PaymentRequest`.

| Campo no Front-end atual | Campo esperado pelo Back-end |
| :--- | :--- |
| `customerEmail` | `email` |
| `amount` (float) | `amount` (int64) |

**Payload Correto para `/api/v1/payments/create-intent`:**
```json
{
  "amount": 19580,
  "currency": "brl",
  "email": "usuario@exemplo.com"
}
```

## 3. Integração com Supabase (Interactions)

O erro `400` na URL `wnotjxleeltjysdlplkc.supabase.co/rest/v1/interactions` ocorre porque o objeto enviado contém tipos de dados incompatíveis com o schema do banco de dados (ex: tentando salvar float em coluna integer).

**Recomendação:**
Certifique-se de que a função que salva a interação no Supabase use o valor já convertido em centavos ou que a coluna no Supabase aceite decimais (numeric/float). O ideal é manter a consistência com o Stripe e usar **centavos (inteiro)** em todo o fluxo.

## 4. Checklist de Verificação

1. [ ] O campo `amount` é um número inteiro sem casas decimais?
2. [ ] O campo de e-mail foi renomeado de `customerEmail` para `email`?
3. [ ] A `currency` está sempre em letras minúsculas (ex: `"brl"`)?
4. [ ] Os dados extras (city, state, taxId) estão sendo enviados apenas se o endpoint os suportar ou se forem para o Supabase separadamente? (O endpoint de `create-intent` atual ignora campos extras, mas falha se os obrigatórios estiverem errados).
