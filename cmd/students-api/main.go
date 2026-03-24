package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/brijesh025/students-api/internal/config"
)
func main(){
	//load config
	cfg := config.MustLoad();
	// database setup
	// setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to sstudents api"))
	})
	// setup server
	server := http.Server{
		Addr: cfg.HTTPServer.Address,
		Handler: router,
	}
	err := server.ListenAndServe()
	if (err != nil) {
		log.Fatal("Failed to start server")
	}
	fmt.Println("Server started and Ready to work!!")
}