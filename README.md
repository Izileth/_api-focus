# API Focus 🚀

API robusta em Go para processamento de pagamentos via Stripe e gestão de serviços. Desenvolvida com foco em performance, escalabilidade e facilidade de integração.

---

## 🛠️ Stack Tecnológica

- **Linguagem:** [Go (Golang)](https://golang.org/) 1.26+
- **Framework Web:** [Gin Gonic](https://gin-gonic.com/)
- **Banco de Dados:** [PostgreSQL](https://www.postgresql.org/) (via pgx)
- **Integração de Pagamentos:** [Stripe SDK](https://stripe.com/docs/api)
- **Gestão de Configuração:** Godotenv
- **Hot Reload:** [Air](https://github.com/cosmtrek/air)

---

## 📂 Arquitetura do Projeto

O projeto segue uma estrutura modular para facilitar a manutenção e evolução:

```text
.
├── cmd/
│   └── api/                # Ponto de entrada da aplicação (main.go)
├── internal/
│   ├── config/             # Carregamento de variáveis de ambiente
│   ├── database/           # Conexão e configuração do banco de dados
│   ├── handlers/           # Lógica de controle dos endpoints
│   ├── middleware/         # Middlewares (CORS, Versionamento, etc.)
│   ├── repositories/       # Abstração de acesso aos dados (Em desenvolvimento)
│   ├── services/           # Regras de negócio (Em desenvolvimento)
│   └── stripe/             # Cliente e integração direta com a API do Stripe
├── tmp/                    # Binários e logs temporários (Air)
├── .air.toml               # Configuração do live reloading
└── go.mod                  # Dependências do projeto
```

---

## 🚀 Como Começar

### Pré-requisitos
- Go 1.26 ou superior instalado.
- Instância do PostgreSQL rodando.
- Conta no Stripe com chaves de API.

### Instalação

1. Clone o repositório:
   ```bash
   git clone <repo-url>
   cd api-focus
   ```

2. Instale as dependências:
   ```bash
   go mod tidy
   ```

3. Configure as variáveis de ambiente:
   Crie um arquivo `.env` na raiz do projeto com as seguintes chaves:
   ```env
   PORT=8080
   DATABASE_URL=postgres://user:password@localhost:5432/dbname
   STRIPE_SECRET_KEY=sk_test_...
   STRIPE_WEBHOOK_SECRET=whsec_...
   ```

### Executando a API

**Modo Desenvolvimento (com Air):**
```bash
air
```

**Modo Produção:**
```bash
go run cmd/api/main.go
```

---

## 🔌 Referência da API

### URL Base
`http://localhost:8080/api/{version}`

### Versionamento
A API suporta versionamento via prefixo na URL: `/v1`, `/v2`, `/v3`. Atualmente, a **v1** é a versão estável recomendada.

---

### 1. Pagamentos

#### A. Criar Intenção de Pagamento (Embedded)
Inicia um processo de pagamento onde o Stripe decide dinamicamente os métodos disponíveis (Cartão, Pix, Boleto) com base na configuração do dashboard.

- **Endpoint:** `POST /api/v1/payments/create-intent`
- **Request Body:**
```json
{
  "amount": 5000,
  "currency": "brl",
  "email": "cliente@email.com"
}
```
- **Resposta:** Retorna o `clientSecret` para ser usado no frontend com o Stripe Elements.

#### B. Criar Sessão de Checkout (Hosted)
Redireciona o usuário para uma página hospedada pelo Stripe.

- **Endpoint:** `POST /api/v1/payments/create-checkout-session`
- **Request Body:**
```json
{
  "amount": 10000,
  "currency": "brl",
  "success_url": "https://seu-site.com/sucesso",
  "cancel_url": "https://seu-site.com/cancelado"
}
```

#### C. Webhook de Notificação
Endpoint para receber atualizações de status do Stripe de forma assíncrona.

- **Endpoint:** `POST /api/v1/payments/webhook`
- **Segurança:** Revalida a assinatura usando `STRIPE_WEBHOOK_SECRET`.

---

### 2. Sistema & Saúde

- **Health Check:** `GET /health` ou `GET /api/v1/health`
  - Verifica status da API e conexão com o Banco de Dados.
- **Info:** `GET /`
  - Retorna informações básicas e versões disponíveis.

---

## 🛡️ CORS & Segurança

A API está configurada para aceitar requisições das seguintes origens por padrão:
- `http://localhost:3000`
- `https://moudusfocus.online`

Para adicionar novas origens, altere o middleware em `internal/middleware/cors.go`.

---

## 📝 Notas de Implementação

- **Valores:** Todos os valores financeiros são tratados em **centavos** (ex: R$ 10,00 = `1000`).
- **Tratamento de Erros:** A API retorna códigos HTTP semânticos e mensagens claras em caso de falha.
- **Banco de Dados:** Utiliza o driver de alta performance `pgx` para interações com PostgreSQL.

---

## 📖 Documentação Adicional

Para detalhes profundos sobre a integração, consulte nossos guias especializados:

*   **[Guia de Integração Frontend](./docs/frontend-integration.md)**: Fluxos de UX e requisitos.
*   **[Referência de Payloads](./docs/payloads.md)**: Schemas de JSON e tipos de dados.
*   **[Webhooks e Ciclo de Vida](./docs/webhooks.md)**: Como lidar com pagamentos assíncronos.
*   **[Guia de Testes](./docs/testing.md)**: Dados de teste e como usar o Stripe CLI.

---

Desenvolvido com ❤️ por API Focus Team.
