package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/G-Villarinho/parallelizing-golang-ops/models"
	"go.mongodb.org/mongo-driver/mongo"
)

const amount = 10000000

func SeedMongoDB(collection *mongo.Collection) {
	ctx := context.Background()
	collection.Drop(ctx)

	var students []interface{}
	for i := 0; i < amount; i++ {
		student := models.Student{
			Name:         fmt.Sprintf("Student %d", i),
			Email:        fmt.Sprintf("student%d@example.com", i),
			Age:          18 + (i % 42),
			RegisteredAt: time.Now().AddDate(0, 0, -i),
		}
		students = append(students, student)
	}

	_, err := collection.InsertMany(ctx, students)
	if err != nil {
		log.Fatalf("Erro ao inserir no MongoDB: %v", err)
	}

	log.Printf("Inseridos %d estudantes no MongoDB\n", amount)
}
