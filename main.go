package main

import (
	"fmt"
	"kamus_api/server"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {
	routes := chi.NewRouter()
	r := server.NewRoute(routes)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "9090"
	}
	fmt.Println("server started at localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
