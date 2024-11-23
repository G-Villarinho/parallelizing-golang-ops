package main

import (
	"context"
	"log"

	"github.com/G-Villarinho/parallelizing-golang-ops/db"
	"github.com/G-Villarinho/parallelizing-golang-ops/services"
)

func main() {
	mongoClient, mongoCollection := db.ConnectMongoDB()
	defer mongoClient.Disconnect(context.Background())

	pgConn := db.ConnectPostgres()
	defer pgConn.Close()

	services.SeedMongoDB(mongoCollection)

	db.CreateStudentsTable(pgConn)

	services.TransferData(mongoCollection, pgConn, 1000)

	log.Println("Processo finalizado com sucesso!")
}
