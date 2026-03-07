package main

import (
	"fmt"
	"log"

	"github.com/ncostamagna/go-simple-project/internal/title"
	"github.com/ncostamagna/go-simple-project/transport/httpapi"
	"github.com/ncostamagna/go-simple-project/adapter/memorydb"
)

func main() {

	repo := memorydb.New()

	srv := title.NewService(repo)

	endpoints := httpapi.MakePostsEndpoints(srv)
	
	apiServer := httpapi.New(endpoints)

	errs := make(chan error, 1)

	go func() {
		url := fmt.Sprintf("127.0.0.1:8085")
		log.Println("Listening", "url", url)
		errs <- apiServer.Run(url)
	}()

	<-errs
	log.Println("end program", errs)
}
