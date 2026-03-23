package main

import (
	"fmt"
	"log"

	"github.com/ncostamagna/go-simple-project/adapter/memorydb"
	"github.com/ncostamagna/go-simple-project/transport/httpapi"
	"github.com/ncostamagna/go-simple-project/adapter/postgres"
	"github.com/ncostamagna/go-simple-project/bootstrap"
	"github.com/ncostamagna/go-simple-project/internal/title"
)

func main() {

	var repoPostgresdb postgres.Repository
	var repoMemorydb memorydb.Repository

	db, err := bootstrap.InitPostgres()
	if err != nil {
		log.Fatalf("failed to connect to Postgres: %v", err)
	}

	repoPostgresdb = postgres.NewRepository(db)
	repoMemorydb = memorydb.NewRepository()

	srv := title.NewService(repoPostgresdb, repoMemorydb)

	endpoints := httpapi.MakePostsEndpoints(srv)
	apiServer := httpapi.New(endpoints)

	errs := make(chan error, 1)

	go func() {
		url := fmt.Sprintf("0.0.0.0:8085")
		log.Println("Listening", "url", url)
		errs <- apiServer.Run(url)
	}()

	fatalErr := <-errs
	log.Println("Program ended:", fatalErr)
}
