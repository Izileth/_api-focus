package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func Connect() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("⚠️ DATABASE_URL não configurada no .env")
	}

	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("❌ Erro ao conectar ao banco de dados: %v", err)
	}

	// Ping para validar a conexão
	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatalf("❌ Falha no ping do banco de dados: %v", err)
	}

	log.Println("✅ Conexão com o banco de dados estabelecida com sucesso!")
	DB = conn
}
