package main

import (
	"github.com/Ad3bay0c/golang-testing/handler"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const PORT = ":8080"

func main() {
	router := gin.Default()
	s := handler.NewServer()
	s.SetUpRouter(router)

	server := &http.Server{
		Handler: router,
		Addr:    PORT,
	}
	log.Printf("Server started at http://localhost%v", PORT)
	log.Fatal(server.ListenAndServe())
}


