# API Focus 🚀

API backend desenvolvida em Go para processamento de pagamentos e gestão de serviços, integrada com Stripe e PostgreSQL (Supabase).

## 🛠️ Tecnologias

- **Linguagem:** [Go 1.26+](https://go.dev/)
- **Framework Web:** [Gin Gonic](https://gin-gonic.com/)
- **Banco de Dados:** PostgreSQL via [pgx](https://github.com/jackc/pgx)
- **Pagamentos:** [Stripe SDK](https://github.com/stripe/stripe-go)
- **Ambiente:** [godotenv](https://github.com/joho/godotenv) e [Air](https://github.com/air-verse/air) (Live Reload)

## 📁 Estrutura do Projeto

```text
├── cmd/api/          # Ponto de entrada da aplicação
├── internal/
│   ├── config/       # Carregamento de variáveis de ambiente
│   ├── database/     # Conexão com o banco de dados
│   ├── handlers/     # Controladores (Lógica das rotas)
│   ├── stripe/       # Inicialização do cliente Stripe
│   └── ...           # Repositories e Services (em desenvolvimento)
└── .env              # Variáveis sensíveis (não commitado)
```

## 🚀 Como Iniciar

### 1. Pré-requisitos
- Go instalado.
- Conta no [Stripe](https://stripe.com/) (para chaves de teste).
- Banco de dados PostgreSQL (Configurado no Supabase por padrão).

### 2. Configuração do Ambiente
Crie ou edite o arquivo `.env` na raiz do projeto:
```env
PORT=8080
DATABASE_URL=sua_url_do_postgres
STRIPE_SECRET_KEY=sua_chave_secreta_do_stripe
STRIPE_WEBHOOK_SECRET=sua_chave_de_webhook_do_stripe
```

### 3. Executando a API
Para rodar em modo de desenvolvimento com live reload:
```bash
air
```
Ou manualmente:
```bash
go run cmd/api/main.go
```

## 🔌 Endpoints de Pagamento

### Criar Intenção de Pagamento
Cria um `PaymentIntent` no Stripe e retorna o `clientSecret` para o frontend.
- **URL:** `POST /payments/create-intent`
- **Body:**
```json
{
  "amount": 1000,
  "currency": "brl"
}
```

### Webhook do Stripe
Endpoint para receber notificações assíncronas do Stripe.
- **URL:** `POST /payments/webhook`
- **Nota:** Requer o [Stripe CLI](https://stripe.com/docs/stripe-cli) para testes locais de webhook:
  ```bash
  stripe listen --forward-to localhost:8080/payments/webhook
  ```

## 📝 Licença
Este projeto está sob a licença MIT.
