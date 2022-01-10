package main

import (
	"github.com/Ad3bay0c/golang-testing/server"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)
const PORT = ":8080"

func main () {
	router := gin.Default()
	setUpRouter(router)

	s := &http.Server{
		Handler: router,
		Addr: PORT,
	}
	log.Printf("Server started at http://localhost%v", PORT)
	log.Fatal(s.ListenAndServe())
}

func setUpRouter(r *gin.Engine) {
	s := server.NewServer()

	r.POST("/", s.GetMostUsedWords)
}
