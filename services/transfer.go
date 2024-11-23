package services

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/G-Villarinho/parallelizing-golang-ops/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.mongodb.org/mongo-driver/mongo"
)

func TransferData(mongoCollection *mongo.Collection, pgConn *pgxpool.Pool, batchSize int) {
	startTime := time.Now()

	ctx := context.Background()

	totalDocs, err := mongoCollection.CountDocuments(ctx, map[string]interface{}{})
	if err != nil {
		log.Fatalf("Erro ao contar documentos no MongoDB: %v", err)
	}
	log.Printf("Total de documentos no MongoDB: %d", totalDocs)

	batchChannel := make(chan []models.Student, 100)

	var wg sync.WaitGroup

	numWorkers := 200
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for batch := range batchChannel {
				insertBatch(pgConn, batch)
			}
		}()
	}

	cursor, err := mongoCollection.Find(ctx, map[string]interface{}{})
	if err != nil {
		log.Fatalf("Erro ao ler do MongoDB: %v", err)
	}
	defer cursor.Close(ctx)

	var batch []models.Student
	for cursor.Next(ctx) {
		var student models.Student
		if err := cursor.Decode(&student); err != nil {
			log.Fatalf("Erro ao decodificar documento do MongoDB: %v", err)
		}

		batch = append(batch, student)
		if len(batch) == batchSize {
			batchChannel <- batch
			batch = nil
		}
	}

	if len(batch) > 0 {
		batchChannel <- batch
	}

	close(batchChannel)
	wg.Wait()

	elapsedTime := time.Since(startTime)
	fmt.Printf("Transferência de dados concluída em %v segundos\n", elapsedTime.Seconds())
}

func insertBatch(conn *pgxpool.Pool, batch []models.Student) {
	if len(batch) == 0 {
		return
	}

	query := `INSERT INTO students (name, email, age, registered_at) VALUES `
	args := []interface{}{}
	argCounter := 1

	for _, student := range batch {
		query += fmt.Sprintf("($%d, $%d, $%d, $%d),", argCounter, argCounter+1, argCounter+2, argCounter+3)
		args = append(args, student.Name, student.Email, student.Age, student.RegisteredAt)
		argCounter += 4
	}

	query = query[:len(query)-1]

	_, err := conn.Exec(context.Background(), query, args...)
	if err != nil {
		log.Fatalf("Erro ao inserir no PostgreSQL: %v", err)
	}
}
