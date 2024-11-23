package db

import (
	"context"
	"log"

	"github.com/ClickHouse/clickhouse-go/v2"
)

func ConnectClickHouse() clickhouse.Conn {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"localhost:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
		Debug: false,
	})
	if err != nil {
		log.Fatalf("Erro ao conectar no ClickHouse: %v", err)
	}
	return conn
}

func CreateTableClickHouse(conn clickhouse.Conn) {
	ctx := context.Background()

	query := `
	CREATE TABLE IF NOT EXISTS students (
		id UInt32,
		name String,
		email String,
		age UInt8,
		registered_at DateTime
	) ENGINE = MergeTree()
	ORDER BY (id)`

	if err := conn.Exec(ctx, query); err != nil {
		log.Fatalf("Erro ao criar tabela no ClickHouse: %v", err)
	}

	log.Println("Tabela 'students' criada com sucesso no ClickHouse.")
}
