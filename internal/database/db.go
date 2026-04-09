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
		log.Fatal("âš ï¸  DATABASE_URL nÃ£o configurada no .env")
	}

	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("â Œ Erro ao conectar ao banco de dados: %v", err)
	}

	// Ping para validar a conexÃ£o
	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatalf("â Œ Falha no ping do banco de dados: %v", err)
	}

	log.Println("âœ… ConexÃ£o com o banco de dados estabelecida com sucesso!")
	DB = conn
}