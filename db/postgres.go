package db

import (
	"context"
	"log"

	"github.com/G-Villarinho/parallelizing-golang-ops/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectPostgres() *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(config.PostgresConfig())
	if err != nil {
		log.Fatalf("Erro ao configurar a conexão com o PostgreSQL: %v", err)
	}

	conn, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Erro ao conectar no PostgreSQL: %v", err)
	}

	return conn
}

func CreateStudentsTable(conn *pgxpool.Pool) {
	query := `
		CREATE TABLE IF NOT EXISTS students (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			age INT NOT NULL,
			registered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`
	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("Erro ao criar tabela no PostgreSQL: %v", err)
	}
	log.Println("Tabela 'students' criada com sucesso!")
}
