package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	// "sync"

	"github.com/brijesh025/students-api/internal/config"
	"github.com/joho/godotenv"
)
func main(){
	//load config
	e := godotenv.Load();
	if(e!=nil){
		log.Println(".env file was not loaded")
	}
	cfg := config.MustLoad();
	// database setup
	// setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to students api"))
	})
	// setup server
	server := http.Server{
		Addr: cfg.HTTPServer.Address,
		Handler: router,
	}
	fmt.Printf("Server started and Ready to work!! Check %s", cfg.HTTPServer.Address)
	// var wg sync.WaitGroup
	// wg.Add(1);
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM);
	var err error
	go func(/*wg *sync.WaitGroup*/){
		// defer wg.Done()
		err = server.ListenAndServe()
		if (err != nil) {
			log.Fatal("Failed to start server")
		}
	}(/*&wg*/)
	// wg.Wait()
	<-done
	fmt.Println("this is after locking the server listening funciton")
}