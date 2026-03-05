package main

import (
	"fmt"
	"log"

	"github.com/ncostamagna/go-simple-project/service"
	"github.com/ncostamagna/go-simple-project/transport/httpapi"
)

func main() {

	srv := service.New()

	endpoints := httpapi.MakePostsEndpoints(srv)
	
	apiServer := httpapi.New(endpoints)

	errs := make(chan error, 1)

	go func() {
		url := fmt.Sprintf("127.0.0.1:8085")
		log.Println("Listening", "url", url)
		errs <- apiServer.Run(url)
	}()

	<-errs
	log.Println("end program")
}
